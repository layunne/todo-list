package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/layunne/todo-list/backend/errors"
	"github.com/layunne/todo-list/backend/models"
	"github.com/layunne/todo-list/backend/repositories"
	"net/http"
)

func NewProjectsService(
	projectsRepository repositories.ProjectsRepository,
	taskRepository repositories.TasksRepository,
) ProjectsService {
	return &projectsService{projectsRepository: projectsRepository, taskRepository: taskRepository}
}

type ProjectsService interface {
	Get(userId string, id string) *models.ProjectDTO
	GetAll(userId string) []*models.ProjectDTO
	Create(createProject *models.CreateProject) (*models.ProjectDTO, *errors.Error)
	Update(updateProject *models.UpdateProject) (*models.ProjectDTO, *errors.Error)
	Delete(userId string, projectId string) *errors.Error
}

type projectsService struct {
	projectsRepository repositories.ProjectsRepository
	taskRepository repositories.TasksRepository
}

func (s *projectsService) Get(userId string, id string) *models.ProjectDTO {

	project := s.projectsRepository.Get(id)

	if project == nil {
		return nil
	}

	if project.UserId != userId {
		return nil
	}

	tasks := s.taskRepository.GetByProjectId(project.Id)

	return project.ToDTO(tasks)
}

func (s *projectsService) GetAll(userId string) []*models.ProjectDTO {

	projects := s.projectsRepository.GetByUserId(userId)

	projectsDTO := make([]*models.ProjectDTO, 0)

	for _, project := range projects {
		tasks := s.taskRepository.GetByProjectId(project.Id)

		p := project.ToDTO(tasks)

		projectsDTO = append(projectsDTO, p)
	}

	return projectsDTO
}

func (s *projectsService) Create(createProject *models.CreateProject) (*models.ProjectDTO, *errors.Error) {

	if len(createProject.Name) < 4 {
		return nil, errors.New(http.StatusBadRequest, "name needs to be greater 3")
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("uuid error: %v", err.Error()))
	}

	project := &models.Project{
		Id:     id.String(),
		UserId: createProject.UserId,
		Name:   createProject.Name,
	}

	s.projectsRepository.Save(project)

	return project.ToDTO(make([]*models.Task, 0)), nil
}

func (s *projectsService) Update(updateProject *models.UpdateProject) (*models.ProjectDTO, *errors.Error) {

	project := s.projectsRepository.Get(updateProject.Id)

	if project == nil {
		return nil, errors.New(http.StatusBadRequest, "project not fount for id: " + updateProject.Id)
	}

	if project.UserId != updateProject.UserId {
		return nil, errors.New(http.StatusUnauthorized, "project does not belong to user")
	}

	if len(updateProject.Name) < 4 {
		return nil, errors.New(http.StatusBadRequest, "name needs to be greater 3")
	}

	project.Name = updateProject.Name

	s.projectsRepository.Save(project)

	tasks := s.taskRepository.GetByProjectId(project.Id)

	return project.ToDTO(tasks), nil
}

func (s *projectsService) Delete(userId string, projectId string) *errors.Error {
	project := s.projectsRepository.Get(projectId)

	if project == nil {
		return errors.New(http.StatusBadRequest, "project not fount for id: " + projectId)
	}

	if project.UserId != userId {
		return errors.New(http.StatusUnauthorized, "project does not belong to user")
	}

	return s.projectsRepository.Delete(projectId)
}
