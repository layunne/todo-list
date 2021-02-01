package services

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/layunne/todo-list/backend/errors"
	"github.com/layunne/todo-list/backend/models"
	"github.com/layunne/todo-list/backend/repositories"
	"net/http"
	"time"
)

func NewTasksService(
	taskRepository repositories.TasksRepository,
	projectsRepository repositories.ProjectsRepository,
) TasksService {
	return &tasksService{taskRepository: taskRepository, projectsRepository: projectsRepository}
}

type TasksService interface {
	Get(userId string, id string) *models.Task
	GetAll(userId string, projectId string) []*models.Task
	Create(userId string, createTask *models.CreateTask) (*models.Task, *errors.Error)
	Update(userId string, updateTask *models.UpdateTask) (*models.Task, *errors.Error)
	ChangeStatus(userId string, updateTask *models.UpdateTask) (*models.Task, *errors.Error)
	Delete(userId string, taskId string)
}

type tasksService struct {
	taskRepository     repositories.TasksRepository
	projectsRepository repositories.ProjectsRepository
}

func (s *tasksService) Get(userId string, id string) *models.Task {

	task := s.taskRepository.Get(id)

	if task == nil {
		return nil
	}

	project := s.projectsRepository.Get(task.ProjectId)

	if project == nil || project.UserId != userId {
		return nil
	}

	return task
}

func (s *tasksService) GetAll(userId string, projectId string) []*models.Task {
	project := s.projectsRepository.Get(projectId)
	if project == nil || project.UserId != userId {
		return make([]*models.Task, 0)
	}
	return s.taskRepository.GetByProjectId(projectId)
}

func (s *tasksService) Create(userId string, createTask *models.CreateTask) (*models.Task, *errors.Error) {

	project := s.projectsRepository.Get(createTask.ProjectId)

	if project == nil {
		return nil, errors.New(http.StatusBadRequest, "project not found for id: "+createTask.ProjectId)
	}

	if project.UserId != userId {
		return nil, errors.New(http.StatusBadRequest, "project not found for id: "+createTask.ProjectId)
	}

	if createTask.Description == "" {
		return nil, errors.New(http.StatusBadRequest, "description cannot be empty")
	}

	id, err := uuid.NewUUID()

	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("uuid error: %v", err.Error()))
	}

	task := &models.Task{
		Id:          id.String(),
		ProjectId:   project.Id,
		Description: createTask.Description,
		ToFinishAt:  createTask.ToFinishAt,
		CreatedAt:   time.Now().Unix(),
	}

	err2 := s.taskRepository.Save(task)

	if err2 != nil {
		return nil, err2
	}

	return task, nil
}

func (s *tasksService) Update(userId string, updateTask *models.UpdateTask) (*models.Task, *errors.Error) {

	task := s.taskRepository.Get(updateTask.Id)

	if task == nil {
		return nil, errors.New(http.StatusNotAcceptable, "task not found for id: " + updateTask.Id)
	}

	project := s.projectsRepository.Get(task.ProjectId)

	if project == nil || project.UserId != userId {
		return nil, errors.New(http.StatusBadRequest, "project not found for id: "+task.ProjectId)
	}

	if updateTask.Description == "" {
		return nil, errors.New(http.StatusBadRequest, "description cannot be empty")
	}

	task.Description = updateTask.Description
	task.ToFinishAt = updateTask.ToFinishAt

	err := s.taskRepository.Save(task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *tasksService) ChangeStatus(userId string, updateTask *models.UpdateTask) (*models.Task, *errors.Error) {

	task := s.taskRepository.Get(updateTask.Id)

	if task == nil {
		return nil, errors.New(http.StatusNotAcceptable, "task not found for id: " + updateTask.Id)
	}

	project := s.projectsRepository.Get(task.ProjectId)

	if project == nil || project.UserId != userId {
		return nil, errors.New(http.StatusBadRequest, "project not found for id: "+task.ProjectId)
	}

	if task.Status == updateTask.Status {
		return task, nil
	}

	task.Status = updateTask.Status

	if task.Status == true {
		task.FinishedAt = time.Now().Unix()
	} else {
		task.FinishedAt = 0
	}

	err := s.taskRepository.Save(task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *tasksService) Delete(userId string, taskId string) {

	task := s.taskRepository.Get(taskId)

	if task == nil {
		return
	}

	project := s.projectsRepository.Get(task.ProjectId)

	if project == nil || project.UserId != userId {
		return
	}

	s.taskRepository.Delete(taskId)
}
