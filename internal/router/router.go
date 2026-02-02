package router

import (
	"CliPorto/internal/controller"
	"CliPorto/internal/repository"
	"CliPorto/internal/service"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	// Showcase routes
	showcaseRepo := repository.NewShowcaseRepository(db)
	showcaseService := service.NewShowcaseService(showcaseRepo)
	showcaseController := controller.NewShowcaseController(showcaseService)

	// Profile routes
	profileRepo := repository.NewProfileRepository(db)
	profileService := service.NewProfileService(profileRepo)
	profileController := controller.NewProfileController(profileService)

	api := r.Group("/api")
	{
		api.GET("/showcase", showcaseController.GetAll)
		api.GET("/showcase/:id", showcaseController.GetByID)
		api.GET("/profile", profileController.GetAll)
	}
}
