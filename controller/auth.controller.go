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

func (c *AuthController) Signup(ctx *fiber.Ctx) error {
	var signupDto dto.SignupDto
	if err := ctx.BodyParser(&signupDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	result := c.authService.Signup(signupDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *AuthController) Signin(ctx *fiber.Ctx) error {
	var authDto dto.AuthDto
	if err := ctx.BodyParser(&authDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	result := c.authService.Signin(authDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *AuthController) Refresh(ctx *fiber.Ctx) error {
	refreshToken := ctx.Get("Authorization")
	result := c.authService.RefreshToken(refreshToken)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *AuthController) GetProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	result := c.authService.GetProfile(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *AuthController) EditProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	var editProfileDto dto.EditProfileDto
	if err := ctx.BodyParser(&editProfileDto); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "잘못된 입력입니다"})
	}

	result := c.authService.EditProfile(userId, editProfileDto)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	result := c.authService.Logout(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}

func (c *AuthController) DeleteAccount(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(uint)
	result := c.authService.DeleteAccount(userId)
	return ctx.Status(result.Status).JSON(result.Data)
}