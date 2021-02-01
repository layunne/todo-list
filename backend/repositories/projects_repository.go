package repositories

import (
	"github.com/layunne/todo-list/backend/errors"
	"github.com/layunne/todo-list/backend/infrastructure/mongo"
	"github.com/layunne/todo-list/backend/models"
	"net/http"
)

type ProjectsRepository interface {
	Get(id string) *models.Project
	GetByUserId(userId string) []*models.Project
	Save(project *models.Project) *errors.Error
	Update(project *models.Project) *errors.Error
	Delete(id string) *errors.Error
}

func NewProjectsRepository(mongoClient mongo.Client) ProjectsRepository  {
	return &projectsRepositoryMongo{
		mongoClient: mongoClient,
		collection: "projects",
	}
}

type projectsRepositoryMongo struct {
	mongoClient mongo.Client
	collection string
}

func (r *projectsRepositoryMongo) Get(id string) *models.Project {

	res := &models.Project{}

	err := r.mongoClient.GetById(r.collection, id, res)

	if err != nil {
		return nil
	}

	return res
}

func (r *projectsRepositoryMongo) GetByUserId(userId string) []*models.Project {

	res := &([]*models.Project{})

	err := r.mongoClient.Find(r.collection, "userId", userId, res)

	if err != nil {
		return nil
	}

	return *res
}

func (r *projectsRepositoryMongo) Save(project *models.Project) *errors.Error {

	err := r.mongoClient.InsertOrUpdate(r.collection, project, project.Id)

	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *projectsRepositoryMongo) Update(project *models.Project) *errors.Error {
	savedUser := r.Get(project.Id)

	if savedUser == nil {
		return errors.New(http.StatusNotAcceptable, "project not found for id: " + project.Id)
	}

	err := r.mongoClient.InsertOrUpdate(r.collection, project, project.Id)

	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *projectsRepositoryMongo) Delete(id string) *errors.Error {

	err := r.mongoClient.Remove(r.collection, id)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

