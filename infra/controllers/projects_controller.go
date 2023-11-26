package controllers

import (
	"errors"
	"net/http"

	"github.com/IcaroSilvaFK/free-code-source-back/infra/controllers/views"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/models"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/services"
	"github.com/labstack/echo/v4"
)

type ProjectsController struct {
	repo services.ProjectServiceInterface
}

type ProjectsControllerInterface interface {
	Create(ctx echo.Context) error
	FindAll(ctx echo.Context) error
	FindById(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

func NewProjectsController(
	repo services.ProjectServiceInterface,
) ProjectsControllerInterface {
	return &ProjectsController{
		repo,
	}
}

func (pc *ProjectsController) Create(ctx echo.Context) error {

	p := new(views.ProjectInput)

	if err := ctx.Bind(p); !errors.Is(err, nil) {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(p); !errors.Is(err, nil) {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	project := models.ProjectModel{
		ProjectName:        p.Title,
		ProjectDescription: p.Description,
		LinkToSocialMedia:  p.LinkToSocialMedia,
		ProjectType:        p.ProjectType,
		Tecs:               p.Tecs,
		User: models.UserProject{
			Email:         p.User.Email,
			AvatarUrl:     p.User.AvatarUrl,
			LinkToProfile: p.User.LinkToProfile,
			Username:      p.User.Username,
		},
	}

	err := pc.repo.CreateProject(project)

	if !errors.Is(err, nil) {

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, echo.Map{
		"message": "Project created successfully",
	})
}

func (pc *ProjectsController) FindAll(ctx echo.Context) error {

	projects, err := pc.repo.FindAllProjects()

	if !errors.Is(err, nil) {

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, projects)
}

func (pc *ProjectsController) FindById(ctx echo.Context) error {

	id := ctx.Param("id")

	if id == "" {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Id is required",
		})
	}

	project, err := pc.repo.FindProjectById(id)

	if !errors.Is(err, nil) {

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, project)
}

func (pc *ProjectsController) Delete(ctx echo.Context) error {

	id := ctx.Param("id")

	if id == "" {

		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Id is required",
		})
	}

	err := pc.repo.DeleteProject(id)

	if !errors.Is(err, nil) {

		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusNoContent, nil)
}
