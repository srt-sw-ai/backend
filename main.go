package main

import (
	"github.com/gofiber/contrib/swagger"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"swai/config"
	"swai/controller"
	_ "swai/docs"
	"swai/middleware"
	"swai/service"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := config.InitDB(&cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	var ConfigDefault = swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Docs",
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(swagger.New(ConfigDefault))

	authService := service.NewAuthService(db, cfg.JWTSecret)
	authController := controller.NewAuthController(authService)

	reportsService := service.NewReportsService(db)
	reportsController := controller.NewReportsController(reportsService)

	api := app.Group("/auth")
	api.Post("/signup", authController.Signup)
	api.Post("/signin", authController.Signin)
	api.Get("/refresh", authController.Refresh)

	api.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	api.Get("/me", authController.GetProfile)
	api.Patch("/me", authController.EditProfile)
	api.Post("/logout", authController.Logout)
	api.Delete("/me", authController.DeleteAccount)

	reports := app.Group("/reports")
	reports.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	reports.Post("/", reportsController.CreateReport)
	reports.Get("/", reportsController.FindAllReports)
	reports.Get("/by-user", reportsController.FindReportByUserId)
	reports.Get("/:reportId", reportsController.FindReport)

	log.Fatal(app.Listen(":8080"))
}
