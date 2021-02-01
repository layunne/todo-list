package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/layunne/todo-list/backend/errors"
	"github.com/layunne/todo-list/backend/models"
	"github.com/layunne/todo-list/backend/repositories"
	"net/http"
)

func NewUsersService(
	usersRepository repositories.UsersRepository,
	authService AuthService,
	encryptionService EncryptionService,
) UsersService {
	return &usersService{usersRepository: usersRepository, authService: authService, encryptionService: encryptionService}
}

type UsersService interface {
	Get(token string) (*models.UserDTO, *errors.Error)
	Create(user *models.CreateUser) (*models.UserDTO, *errors.Error)
	Update(token string, user *models.UpdateUser) (*models.UserDTO, *errors.Error)
	Delete(token string) *errors.Error
	Login(login *models.UserLogin) (*models.UserDTO, *errors.Error)
}

type usersService struct {
	usersRepository repositories.UsersRepository
	authService     AuthService
	encryptionService EncryptionService
}

func (s *usersService) Get(token string) (*models.UserDTO, *errors.Error) {

	id, err := s.authService.GetId(token)

	if err != nil {
		return nil, err
	}

	return s.usersRepository.Get(id).ToDTO(token), nil
}

func (s *usersService) Create(createUser *models.CreateUser) (*models.UserDTO, *errors.Error) {

	if len(createUser.Name) < 4 {
		return nil, errors.New(http.StatusBadRequest, "name needs to be greater 3")
	}

	if len(createUser.Username) < 4 {
		return nil, errors.New(http.StatusBadRequest, "username needs to be greater 3")
	}

	if len(createUser.Password) < 6 {
		return nil, errors.New(http.StatusBadRequest, "password needs to be greater 5")
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("uuid error: %v", err.Error()))
	}

	user := &models.User{
		Id:       id.String(),
		Name:     createUser.Name,
		Username: createUser.Username,
		Password: s.encryptionService.GetEncryption(createUser.Password),
	}

	s.usersRepository.Save(user)

	newToken := s.authService.GetToken(user.Id)

	return user.ToDTO(newToken), nil
}

func (s *usersService) Update(token string, updateUser *models.UpdateUser) (*models.UserDTO, *errors.Error) {

	id, err := s.authService.GetId(token)

	if err != nil {
		return nil, err
	}

	user := s.usersRepository.Get(id)

	if user == nil {
		return nil, errors.New(http.StatusNotFound, "user not found for id: "+id)
	}

	if len(user.Name) < 4 {
		return nil, errors.New(http.StatusBadRequest, "name needs to be greater 3")
	}
	if updateUser.OldPassword != "" && len(updateUser.Password) < 6 {
		return nil, errors.New(http.StatusBadRequest, "password needs to be greater 5")
	}

	user.Name = updateUser.Name
	user.Username = updateUser.Username

	if updateUser.Password != "" {
		if !s.encryptionService.Check(user.Password, updateUser.OldPassword) {
			return nil, errors.New(http.StatusUnauthorized, "invalid old password")
		}
		user.Password = s.encryptionService.GetEncryption(updateUser.Password)
	}

	s.usersRepository.Save(user)

	newToken := s.authService.GetToken(user.Id)

	return user.ToDTO(newToken), nil
}

func (s *usersService) Delete(token string) *errors.Error {

	id, err := s.authService.GetId(token)

	if err != nil {
		return err
	}

	s.usersRepository.Delete(id)

	return nil
}

func (s *usersService) Login(login *models.UserLogin) (*models.UserDTO, *errors.Error) {

	user := s.usersRepository.GetByUsername(login.Username)

	if user == nil {
		return nil, errors.New(http.StatusUnauthorized, "username not found")
	}

	if !s.encryptionService.Check(user.Password, login.Password) {
		return nil, errors.New(http.StatusUnauthorized, "invalid password")
	}

	newToken := s.authService.GetToken(user.Id)

	return user.ToDTO(newToken), nil
}
