package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"swai/dto"
	"swai/entity"
	"time"
)

type AuthService struct {
	db        *gorm.DB
	jwtSecret string
}

func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{db: db, jwtSecret: jwtSecret}
}

func (s *AuthService) Signup(signupDto dto.SignupDto) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signupDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		Email:    signupDto.Email,
		Password: string(hashedPassword),
		NickName: signupDto.NickName,
	}

	result := s.db.Create(&user)
	return result.Error
}

func (s *AuthService) Signin(authDto dto.AuthDto) (string, string, error) {
	var user entity.User
	result := s.db.Where("email = ?", authDto.Email).First(&user)
	if result.Error != nil {
		return "", "", errors.New("이메일 또는 비밀번호가 일치하지 않습니다")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(authDto.Password))
	if err != nil {
		return "", "", errors.New("이메일 또는 비밀번호가 일치하지 않습니다")
	}

	accessToken, refreshToken, err := s.generateTokens(user.ID)
	if err != nil {
		return "", "", err
	}

	err = s.updateHashedRefreshToken(user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return "", "", errors.New("유효하지 않은 Refresh 토큰입니다")
	}

	userId := uint(claims["userId"].(float64))

	var user entity.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return "", "", errors.New("사용자를 찾을 수 없습니다")
	}

	hashedRefreshToken := sha256.Sum256([]byte(refreshToken))
	if user.HashedRefreshToken != hex.EncodeToString(hashedRefreshToken[:]) {
		return "", "", errors.New("유효하지 않은 Refresh 토큰입니다")
	}

	newAccessToken, newRefreshToken, err := s.generateTokens(userId)
	if err != nil {
		return "", "", err
	}

	err = s.updateHashedRefreshToken(userId, newRefreshToken)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *AuthService) GetProfile(userId uint) (*entity.User, error) {
	var user entity.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) EditProfile(userId uint, editProfileDto dto.EditProfileDto) (*entity.User, error) {
	var user entity.User
	if err := s.db.First(&user, userId).Error; err != nil {
		return nil, err
	}

	user.NickName = editProfileDto.NickName
	user.Email = editProfileDto.Email

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Logout(userId uint) error {
	return s.db.Model(&entity.User{}).Where("id = ?", userId).Update("hashed_refresh_token", nil).Error
}

func (s *AuthService) DeleteAccount(userId uint) error {
	return s.db.Delete(&entity.User{}, userId).Error
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
