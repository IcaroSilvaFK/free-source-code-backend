package routes

import (
	"github.com/IcaroSilvaFK/free-code-source-back/infra/controllers"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/database"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/repositories"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/services"
	"github.com/labstack/echo/v4"
)

func NewAppRoutes(r *echo.Group) {

	loginController := controllers.NewLoginController()

	db := database.NewDBConnection()
	projectRepo := repositories.NewProjectsRepository(db)
	projectSvc := services.NewProjectService(projectRepo)
	projectController := controllers.NewProjectsController(projectSvc)

	r.GET("/login/:code", loginController.Login)
	r.POST("/projects", projectController.Create)
	r.GET("/projects", projectController.FindAll)
	r.GET("/projects/:id", projectController.FindById)
	r.DELETE("/projects/:id", projectController.Delete)

}
