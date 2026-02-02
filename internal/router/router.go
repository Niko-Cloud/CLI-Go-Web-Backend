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

	// Contact routes
	contactRepo := repository.NewContactRepository(db)
	contactService := service.NewContactService(contactRepo)
	contactController := controller.NewContactController(contactService)

	// Skill routes
	skillRepo := repository.NewSkillRepository(db)
	skillService := service.NewSkillService(skillRepo)
	skillController := controller.NewSkillController(skillService)

	// Education routes
	educationRepo := repository.NewEducationRepository(db)
	educationService := service.NewEducationService(educationRepo)
	educationController := controller.NewEducationController(educationService)

	// Work Experience routes
	workRepo := repository.NewWorkExperienceRepository(db)
	workService := service.NewWorkExperienceService(workRepo)
	workController := controller.NewWorkExperienceController(workService)

	api := r.Group("/api")
	{
		api.GET("/showcase", showcaseController.GetAll)
		api.GET("/showcase/:id", showcaseController.GetByID)
		api.GET("/profile", profileController.GetAll)
		api.GET("/contact", contactController.GetContactInfo)
		api.GET("/skill", skillController.GetAllSkills)
		api.GET("/education", educationController.GetAll)
		api.GET("/work-experience", workController.GetAll)
		api.GET("/work-experience/:id", workController.GetByID)
	}
}
