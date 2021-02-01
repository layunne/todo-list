package repositories

import (
	"github.com/layunne/todo-list/backend/errors"
	"github.com/layunne/todo-list/backend/infrastructure/mongo"
	"github.com/layunne/todo-list/backend/models"
	"net/http"
)

type TasksRepository interface {
	Get(id string) *models.Task
	GetByProjectId(projectId string) []*models.Task
	Save(task *models.Task) *errors.Error
	Delete(id string) *errors.Error
}

func NewTasksRepository(mongoClient mongo.Client) TasksRepository  {
	return &tasksRepositoryMongo{
		mongoClient: mongoClient,
		collection: "tasks",
	}
}

type tasksRepositoryMongo struct {
	mongoClient mongo.Client
	collection string
}

func (r *tasksRepositoryMongo) Get(id string) *models.Task {

	res := &models.Task{}

	err := r.mongoClient.GetById(r.collection, id, res)

	if err != nil {
		return nil
	}

	return res
}

func (r *tasksRepositoryMongo) GetByProjectId(projectId string) []*models.Task {

	res := &([]*models.Task{})

	err := r.mongoClient.Find(r.collection, "projectId", projectId, res)

	if err != nil {
		return nil
	}

	return *res
}

func (r *tasksRepositoryMongo) Save(task *models.Task) *errors.Error {

	err := r.mongoClient.InsertOrUpdate(r.collection, task, task.Id)

	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *tasksRepositoryMongo) Delete(id string) *errors.Error {

	err := r.mongoClient.Remove(r.collection, id)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

