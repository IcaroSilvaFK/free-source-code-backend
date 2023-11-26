package services

import (
	"time"

	"github.com/IcaroSilvaFK/free-code-source-back/infra/models"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/repositories"
	"github.com/google/uuid"
)

type ProjectService struct {
	repo repositories.ProjectsRepositoryInterface
}

type ProjectServiceInterface interface {
	CreateProject(models.ProjectModel) error
	FindAllProjects() (*[]models.ProjectModel, error)
	FindProjectById(id string) (*models.ProjectModel, error)
	DeleteProject(id string) error
}

func NewProjectService(repo repositories.ProjectsRepositoryInterface) ProjectServiceInterface {

	return &ProjectService{
		repo: repo,
	}
}

func (ps *ProjectService) CreateProject(p models.ProjectModel) error {

	project := models.ProjectModel{
		ID:                 uuid.NewString(),
		ProjectName:        p.ProjectName,
		ProjectDescription: p.ProjectDescription,
		LinkToSocialMedia:  p.LinkToSocialMedia,
		CreatedAt:          time.Now(),
		ProjectType:        p.ProjectType,
		Stack:              p.Stack,
		Tecs:               p.Tecs,
		User: models.UserProject{
			Email:         p.User.Email,
			AvatarUrl:     p.User.AvatarUrl,
			LinkToProfile: p.User.LinkToProfile,
			Username:      p.User.Username,
		},
	}

	return ps.repo.Create(&project)
}
func (ps *ProjectService) FindAllProjects() (*[]models.ProjectModel, error) {

	return ps.repo.FindAll()
}

func (ps *ProjectService) FindProjectById(id string) (*models.ProjectModel, error) {

	project, err := ps.repo.FindById(id)

	return project, err
}

func (ps *ProjectService) DeleteProject(id string) error {

	return ps.repo.Delete(id)
}
