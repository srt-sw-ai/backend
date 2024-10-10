package service

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"swai/common"
	"swai/dto"
	"swai/entity"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{db: db, jwtSecret: jwtSecret}
}

func (s *AuthService) Signup(signupDto dto.SignupDto) common.ServiceResult {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "비밀번호 해싱에 실패했습니다"}}
	}

	user := entity.User{
		Email:    signupDto.Email,
		Password: string(hashedPassword),
		NickName: signupDto.NickName,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "회원가입에 실패했습니다"}}
	}

	return common.ServiceResult{Status: fiber.StatusCreated, Data: fiber.Map{"message": "회원가입이 완료되었습니다"}}
}

func (s *AuthService) Signin(authDto dto.AuthDto) common.ServiceResult {
	var user entity.User
	if err := s.db.Where("email = ?", authDto.Email).First(&user).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusUnauthorized, Data: fiber.Map{"error": "이메일 또는 비밀번호가 일치하지 않습니다"}}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authDto.Password)); err != nil {
		return common.ServiceResult{Status: fiber.StatusUnauthorized, Data: fiber.Map{"error": "이메일 또는 비밀번호가 일치하지 않습니다"}}
	}

	accessToken, refreshToken, err := s.generateTokens(user.ID)
	if err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "토큰 생성에 실패했습니다"}}
	}

	if err := s.updateHashedRefreshToken(user.ID, refreshToken); err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "Refresh 토큰 업데이트에 실패했습니다"}}
	}

	return common.ServiceResult{Status: fiber.StatusOK, Data: fiber.Map{
		"message":      "로그인이 성공적으로 되었습니다",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}}
}

func (s *AuthService) RefreshToken(refreshToken string) common.ServiceResult {
	if refreshToken == "" {
		return common.ServiceResult{Status: fiber.StatusBadRequest, Data: fiber.Map{"error": "Refresh 토큰이 필요합니다"}}
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return common.ServiceResult{Status: fiber.StatusUnauthorized, Data: fiber.Map{"error": "유효하지 않은 Refresh 토큰입니다"}}
	}

	userId := uint(claims["userId"].(float64))

	var user entity.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusNotFound, Data: fiber.Map{"error": "사용자를 찾을 수 없습니다"}}
	}

	hashedRefreshToken := sha256.Sum256([]byte(refreshToken))
	if user.HashedRefreshToken != hex.EncodeToString(hashedRefreshToken[:]) {
		return common.ServiceResult{Status: fiber.StatusUnauthorized, Data: fiber.Map{"error": "유효하지 않은 Refresh 토큰입니다"}}
	}

	newAccessToken, newRefreshToken, err := s.generateTokens(userId)
	if err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "토큰 생성에 실패했습니다"}}
	}

	if err := s.updateHashedRefreshToken(userId, newRefreshToken); err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "Refresh 토큰 업데이트에 실패했습니다"}}
	}

	return common.ServiceResult{Status: fiber.StatusOK, Data: fiber.Map{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	}}
}

func (s *AuthService) GetProfile(userId uint) common.ServiceResult {
	var user entity.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusNotFound, Data: fiber.Map{"error": "프로필을 찾을 수 없습니다"}}
	}

	profileDto := dto.ProfileDto{
		ID:       user.ID,
		Email:    user.Email,
		NickName: user.NickName,
		ImageUri: user.ImageUri,
		Gender:   user.Gender,
		Birthday: user.Birthday,
		Phone:    user.Phone,
		EmergencyPhone: user.EmergencyPhone,
		Address:  user.Address,
		Allergys: user.Allergys,
		UnderlyingDiseases: user.UnderlyingDiseases,
		Medicines: user.Medicines,
		BloodType: user.BloodType,
		Weight: user.Weight,
		Height: user.Height,
	}

	return common.ServiceResult{Status: fiber.StatusOK, Data: profileDto}
}

func (s *AuthService) EditProfile(userId uint, editProfileDto dto.EditProfileDto) common.ServiceResult {
	var user entity.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusNotFound, Data: fiber.Map{"error": "사용자를 찾을 수 없습니다"}}
	}

	user.NickName = editProfileDto.NickName
	user.Email = editProfileDto.Email
	user.Gender = editProfileDto.Gender
	user.Birthday = editProfileDto.Birthday
	user.Phone = editProfileDto.Phone
	user.EmergencyPhone = editProfileDto.EmergencyPhone
	user.Address = editProfileDto.Address
	user.Allergys = editProfileDto.Allergys
	user.UnderlyingDiseases = editProfileDto.UnderlyingDiseases
	user.Medicines = editProfileDto.Medicines
	user.BloodType = editProfileDto.BloodType
	user.Weight = editProfileDto.Weight
	user.Height = editProfileDto.Height

	if err := s.db.Save(&user).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "프로필 수정에 실패했습니다"}}
	}

	return common.ServiceResult{Status: fiber.StatusOK, Data: user}
}

func (s *AuthService) Logout(userId uint) common.ServiceResult {
	if err := s.db.Model(&entity.User{}).Where("id = ?", userId).Update("hashed_refresh_token", nil).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "로그아웃에 실패했습니다"}}
	}

	return common.ServiceResult{Status: fiber.StatusOK, Data: fiber.Map{"message": "로그아웃되었습니다"}}
}

func (s *AuthService) DeleteAccount(userId uint) common.ServiceResult {
	if err := s.db.Delete(&entity.User{}, userId).Error; err != nil {
		return common.ServiceResult{Status: fiber.StatusInternalServerError, Data: fiber.Map{"error": "계정 삭제에 실패했습니다"}}
	}

	return common.ServiceResult{Status: fiber.StatusOK, Data: fiber.Map{"message": "계정이 삭제되었습니다"}}
}

func (s *AuthService) generateTokens(userId uint) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *AuthService) updateHashedRefreshToken(userID uint, refreshToken string) error {
	hashedRefreshToken := sha256.Sum256([]byte(refreshToken))
	result := s.db.Model(&entity.User{}).Where("id = ?", userID).Update("hashed_refresh_token", hex.EncodeToString(hashedRefreshToken[:]))
	return result.Error
}