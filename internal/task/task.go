package task

import (
	"encoding/json"
	"strconv"
	"task-cli-go/internal/dto"
	"task-cli-go/internal/logger"
	"task-cli-go/internal/storage"
	"time"
)

var (
	DONE       = "done"
	TODO       = "todo"
	InPROGRESS = "in-progress"
)

type Service interface {
	CreateTask(task dto.TaskDTO) (bool, *dto.TaskDTO)
	DeleteTask(taskID string) bool
	UpdateTask(task dto.UpdateTaskDTO) bool
	GetAllTasks() []dto.TaskDTO
	FilterByStatus(status string) []dto.TaskDTO
}

type Task struct {
	storage storage.Repository
	logger  logger.Service
}

func CreateNewTask(storage storage.Repository, logger logger.Service) Service {
	return &Task{
		storage: storage,
		logger:  logger,
	}
}

func (t *Task) CreateTask(task dto.TaskDTO) (bool, *dto.TaskDTO) {
	data, err := t.storage.GetAll()

	if err != nil {
		t.logger.LogError(ErrUnableToGetStorageData.Error())

		return false, nil
	}

	now := time.Now().Format(time.RFC3339)
	id := t.storage.GenerateID(data)

	newTask := dto.TaskDTO{
		ID:          &id,
		Description: task.Description,
		Status:      TODO,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	byteData, err := json.Marshal(newTask)

	if err != nil {
		t.logger.LogError(ErrUnableToCreateNewTask.Error())

		return false, nil
	}

	savedTaskString, err := t.storage.Upsert(strconv.Itoa(id), byteData)

	if err != nil || savedTaskString == nil {
		t.logger.LogError(ErrUnableToCreateNewTask.Error())

		return false, nil
	}
	var savedTask dto.TaskDTO

	err = json.Unmarshal([]byte(*savedTaskString), &savedTask)

	if err != nil {
		t.logger.LogError(ErrUnableToCreateNewTask.Error())

		return false, nil
	}

	return true, &savedTask
}

func (t *Task) DeleteTask(taskID string) bool {
	deleted, err := t.storage.Delete(taskID)

	if err != nil || !deleted {
		t.logger.LogError(ErrUnableToDeleteTask.Error())

		return false
	}

	return true
}

func (t *Task) GetAllTasks() []dto.TaskDTO {
	var list []dto.TaskDTO
	data, err := t.storage.GetAll()

	if err != nil {
		t.logger.LogError(ErrUnableToGetAllTask.Error())

		return list
	}

	for _, value := range data.Records {
		var task dto.TaskDTO

		err = json.Unmarshal([]byte(value), &task)

		if err != nil {
			t.logger.LogError(ErrUnableToGetAllTask.Error())

			return []dto.TaskDTO{}
		}

		list = append(list, task)
	}

	return list
}

func (t *Task) FilterByStatus(status string) []dto.TaskDTO {
	var list []dto.TaskDTO
	data, err := t.storage.GetAll()

	if err != nil {
		t.logger.LogError(ErrUnableToGetAllTask.Error())

		return list
	}

	for _, value := range data.Records {
		var task dto.TaskDTO

		err = json.Unmarshal([]byte(value), &task)

		if err != nil {
			t.logger.LogError(ErrUnableToGetAllTask.Error())

			return []dto.TaskDTO{}
		}

		if task.Status == status {
			list = append(list, task)
		}
	}

	return list
}

func (t *Task) UpdateTask(updateDto dto.UpdateTaskDTO) bool {
	taskID := strconv.Itoa(updateDto.ID)
	dbTask, err := t.storage.GetOneBy(taskID)

	if err != nil {
		t.logger.LogError(err.Error())

		return false
	}

	var task dto.TaskDTO

	err = json.Unmarshal([]byte(*dbTask), &task)

	if err != nil {
		t.logger.LogError(err.Error())

		return false
	}

	timeNow := time.Now().Format(time.RFC3339)

	updateTask := dto.TaskDTO{
		ID:          task.ID,
		Description: task.Description,
		Status:      task.Status,
		UpdatedAt:   &timeNow,
		CreatedAt:   task.CreatedAt,
	}

	if updateDto.Description != nil {
		updateTask.Description = *updateDto.Description
	}

	if updateDto.Status != nil {
		updateTask.Status = *updateDto.Status
	}

	byteData, err := json.Marshal(updateTask)

	if err != nil {
		t.logger.LogError(err.Error())

		return false
	}

	_, err = t.storage.Upsert(taskID, byteData)

	if err != nil {
		t.logger.LogError(err.Error())

		return false
	}

	return true
}
