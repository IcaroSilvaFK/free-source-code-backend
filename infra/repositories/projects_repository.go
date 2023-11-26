package repositories

import (
	"errors"
	"fmt"

	"github.com/IcaroSilvaFK/free-code-source-back/infra/database"
	"github.com/IcaroSilvaFK/free-code-source-back/infra/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectsRepository struct {
	db *mongo.Database
}

type ProjectsRepositoryInterface interface {
	Create(p *models.ProjectModel) error
	FindAll() (*[]models.ProjectModel, error)
	FindById(id string) (*models.ProjectModel, error)
	Delete(id string) error
}

func NewProjectsRepository(
	db *mongo.Database,
) ProjectsRepositoryInterface {

	return &ProjectsRepository{
		db,
	}
}

func (pr *ProjectsRepository) Create(p *models.ProjectModel) error {

	// err := pr.db.Model(&models.UserModel{}).Create(p).Error

	_, err := pr.db.Collection("projects").InsertOne(database.DB_CONTEXT, p)

	return err
}

func (pr *ProjectsRepository) FindAll() (*[]models.ProjectModel, error) {

	var projects []models.ProjectModel

	query, err := pr.db.Collection("projects").Find(database.DB_CONTEXT, bson.M{})

	fmt.Println(err)

	if !errors.Is(err, nil) {
		return nil, err
	}

	defer query.Close(database.DB_CONTEXT)

	err = query.All(database.DB_CONTEXT, &projects)

	fmt.Println(projects)

	return &projects, err
}

func (pr *ProjectsRepository) FindById(id string) (*models.ProjectModel, error) {

	var project models.ProjectModel

	filter := bson.M{"id": id}

	query := pr.db.Collection("projects").FindOne(database.DB_CONTEXT, filter)

	if !errors.Is(query.Err(), nil) {
		return nil, query.Err()
	}

	err := query.Decode(&project)

	return &project, err
}

func (pr *ProjectsRepository) Delete(id string) error {

	_, err := pr.db.Collection("projects").DeleteOne(database.DB_CONTEXT, bson.M{"id": id})

	return err
}
