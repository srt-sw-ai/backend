package main

import (
	"log"

	"github.com/gofiber/contrib/swagger"

	"swai/config"
	"swai/controller"
	_ "swai/docs"
	"swai/middleware"
	"swai/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("구성을 로드하지 못했습니다. %v", err)
	}

	db, err := config.InitDB(&cfg)
	if err != nil {
		log.Fatalf("데이터베이스 초기화 실패: %v", err)
	}

	var ConfigDefault = swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "swagger",
		Title:    "Swagger API Documention",
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(swagger.New(ConfigDefault))

	authService := service.NewAuthService(db, cfg.JWTSecret)
	authController := controller.NewAuthController(authService)

	mapService := service.NewMapService(db)
	mapController := controller.NewMapController(mapService)

	reportsService := service.NewReportsService(db)
	reportsController := controller.NewReportsController(reportsService, mapService)

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

	mapGroup := app.Group("/map")
	mapGroup.Use(middleware.JWTMiddleware(cfg.JWTSecret))
	mapGroup.Post("/", mapController.CreateMarker)
	mapGroup.Get("/", mapController.FindAllMarker)
	mapGroup.Get("/:markerId", mapController.FindMarker)

	log.Fatal(app.Listen(":8080"))
}
