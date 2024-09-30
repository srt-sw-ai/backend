package controller

import (
	"fmt"
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

func (c *AuthController) Signup(ctx *fiber.Ctx) error {
	var signupDto dto.SignupDto
	if err := ctx.BodyParser(&signupDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	err := c.authService.Signup(signupDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "회원가입에 실패했습니다"})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "회원가입이 완료되었습니다"})
}

func (c *AuthController) Signin(ctx *fiber.Ctx) error {
	var authDto dto.AuthDto
	if err := ctx.BodyParser(&authDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	accessToken, refreshToken, err := c.authService.Signin(authDto)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "이메일 또는 비밀번호가 일치하지 않습니다"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "로그인이 성공적으로 되었습니다",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (c *AuthController) Refresh(ctx *fiber.Ctx) error {
	refreshToken := ctx.Get("Authorization")
	if refreshToken == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Refresh 토큰이 필요합니다"})
	}

	accessToken, newRefreshToken, err := c.authService.RefreshToken(refreshToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "유효하지 않은 Refresh 토큰입니다"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"accessToken":  accessToken,
		"refreshToken": newRefreshToken,
	})
}

func (c *AuthController) GetProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	fmt.Println(userId)
	profile, err := c.authService.GetProfile(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "프로필을 찾을 수 없습니다"})
	}

	return ctx.Status(fiber.StatusOK).JSON(profile)
}

func (c *AuthController) EditProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	var editProfileDto dto.EditProfileDto
	if err := ctx.BodyParser(&editProfileDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	updatedProfile, err := c.authService.EditProfile(userId, editProfileDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "프로필 수정에 실패했습니다"})
	}

	return ctx.Status(fiber.StatusOK).JSON(updatedProfile)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	err := c.authService.Logout(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "로그아웃에 실패했습니다"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "로그아웃되었습니다"})
}

func (c *AuthController) DeleteAccount(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	err := c.authService.DeleteAccount(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "계정 삭제에 실패했습니다"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "계정이 삭제되었습니다"})
}
