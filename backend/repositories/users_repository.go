package repositories

import (
	"github.com/layunne/todo-list/backend/errors"
	"github.com/layunne/todo-list/backend/infrastructure/mongo"
	"github.com/layunne/todo-list/backend/models"
	"net/http"
)

type UsersRepository interface {
	Get(id string) *models.User
	GetByUsername(username string) *models.User
	Save(user *models.User) *errors.Error
	Update(user *models.User) *errors.Error
	Delete(id string) *errors.Error
}

func NewUsersRepository(mongoClient mongo.Client) UsersRepository  {
	return &usersRepositoryMongo{
		mongoClient: mongoClient,
		collection: "users",
	}
}

type usersRepositoryMongo struct {
	mongoClient mongo.Client
	collection string
}

func (r *usersRepositoryMongo) Get(id string) *models.User {

	res := &models.User{}

	err := r.mongoClient.GetById(r.collection, id, res)

	if err != nil {
		return nil
	}

	return res
}

func (r *usersRepositoryMongo) GetByUsername(username string) *models.User {

	res := &models.User{}

	err := r.mongoClient.FindOne(r.collection, "username", username, res)

	if err != nil {
		return nil
	}

	return res
}

func (r *usersRepositoryMongo) Save(user *models.User) *errors.Error {

	savedUser := r.GetByUsername(user.Username)

	if savedUser != nil {
		return errors.New(http.StatusNotAcceptable, "username already exists for: " + user.Username)
	}

	err := r.mongoClient.InsertOrUpdate(r.collection, user, user.Id)

	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *usersRepositoryMongo) Update(user *models.User) *errors.Error {
	savedUser := r.Get(user.Id)

	if savedUser == nil {
		return errors.New(http.StatusNotAcceptable, "user not found for id: " + user.Id)
	}

	err := r.mongoClient.InsertOrUpdate(r.collection, user, user.Id)

	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (r *usersRepositoryMongo) Delete(id string) *errors.Error {

	err := r.mongoClient.Remove(r.collection, id)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err.Error())
	}

	return nil
}

