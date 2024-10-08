package controller

import (
	"swai/dto"
	"swai/service"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Signup godoc
// @Summary 회원가입
// @Description 새로운 사용자를 등록합니다.
// @Tags Auth
// @Accept json
// @Produce json
// @Param signupDto body dto.SignupDto true "회원가입 정보"
// @Success 201 {object} map[string]interface{} "회원가입이 완료되었습니다"
// @Failure 400 {object} map[string]interface{} "잘못된 입력입니다"
// @Failure 500 {object} map[string]interface{} "회원가입에 실패했습니다"
// @Router /auth/signup [post]
func (c *AuthController) Signup(ctx *fiber.Ctx) error {
	var signupDto dto.SignupDto
	if err := ctx.BodyParser(&signupDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	result := c.authService.Signup(signupDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

// Signin godoc
// @Summary 로그인
// @Description 사용자가 로그인합니다.
// @Tags Auth
// @Accept json
// @Produce json
// @Param authDto body dto.AuthDto true "로그인 정보"
// @Success 200 {object} map[string]interface{} "로그인이 성공적으로 되었습니다"
// @Failure 400 {object} map[string]interface{} "잘못된 입력입니다"
// @Failure 401 {object} map[string]interface{} "이메일 또는 비밀번호가 일치하지 않습니다"
// @Failure 500 {object} map[string]interface{} "토큰 생성에 실패했습니다"
// @Router /auth/signin [post]
func (c *AuthController) Signin(ctx *fiber.Ctx) error {
	var authDto dto.AuthDto
	if err := ctx.BodyParser(&authDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	result := c.authService.Signin(authDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

// Refresh godoc
// @Summary 토큰 갱신
// @Description Refresh 토큰을 사용하여 새로운 Access 토큰을 발급받습니다.
// @Tags Auth
// @Produce json
// @Param Authorization header string true "Refresh 토큰"
// @Success 200 {object} map[string]interface{} "새로운 토큰이 발급되었습니다"
// @Failure 400 {object} map[string]interface{} "Refresh 토큰이 필요합니다"
// @Failure 401 {object} map[string]interface{} "유효하지 않은 Refresh 토큰입니다"
// @Failure 500 {object} map[string]interface{} "토큰 생성에 실패했습니다"
// @Router /auth/refresh [post]
func (c *AuthController) Refresh(ctx *fiber.Ctx) error {
	refreshToken := ctx.Get("Authorization")
	result := c.authService.RefreshToken(refreshToken)
	return ctx.Status(result.Status).JSON(result.Data)
}

// GetProfile godoc
// @Summary 프로필 조회
// @Description 사용자의 프로필 정보를 조회합니다.
// @Tags Auth
// @Produce json
// @Success 200 {object} dto.ProfileDto "프로필 정보"
// @Failure 404 {object} map[string]interface{} "프로필을 찾을 수 없습니다"
// @Router /auth/profile [get]
func (c *AuthController) GetProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	result := c.authService.GetProfile(userId)
	
	if result.Status != fiber.StatusOK {
		return ctx.Status(result.Status).JSON(result.Data)
	}

	return ctx.Status(result.Status).JSON(result.Data)
}

// EditProfile godoc
// @Summary 프로필 수정
// @Description 사용자의 프로필 정보를 수정합니다.
// @Tags Auth
// @Accept json
// @Produce json
// @Param editProfileDto body dto.EditProfileDto true "프로필 수정 정보"
// @Success 200 {object} entity.User "수정된 프로필 정보"
// @Failure 400 {object} map[string]interface{} "잘못된 입력입니다"
// @Failure 404 {object} map[string]interface{} "사용자를 찾을 수 없습니다"
// @Failure 500 {object} map[string]interface{} "프로필 수정에 실패했습니다"
// @Router /auth/profile [put]
func (c *AuthController) EditProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	var editProfileDto dto.EditProfileDto
	if err := ctx.BodyParser(&editProfileDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	result := c.authService.EditProfile(userId, editProfileDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

// Logout godoc
// @Summary 로그아웃
// @Description 사용자가 로그아웃합니다.
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]interface{} "로그아웃되었습니다"
// @Failure 500 {object} map[string]interface{} "로그아웃에 실패했습니다"
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	result := c.authService.Logout(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}

// DeleteAccount godoc
// @Summary 계정 삭제
// @Description 사용자의 계정을 삭제합니다.
// @Tags Auth
// @Produce json
// @Success 200 {object} map[string]interface{} "계정이 삭제되었습니다"
// @Failure 500 {object} map[string]interface{} "계정 삭제에 실패했습니다"
// @Router /auth/delete [delete]
func (c *AuthController) DeleteAccount(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	result := c.authService.DeleteAccount(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}