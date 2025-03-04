package task

import (
	"encoding/json"
	"log"
	"task-cli-go/internal/storage"
	"time"
)

var (
	DONE       = "done"
	TODO       = "todo"
	InPROGRESS = "in-progress"
)

type Service interface {
	CreateTask(task TaskDTO) (bool, *TaskDTO)
	DeleteTask(taskID string) bool
	UpdateTask()
	GetAllTasks() []TaskDTO
}

type Task struct {
	storage storage.Repository
}

func CreateNewTask(storage storage.Repository) *Task {
	return &Task{
		storage: storage,
	}
}

func (t *Task) CreateTask(task TaskDTO) (bool, *TaskDTO) {
	data, err := t.storage.GetAll()

	if err != nil {
		log.Println(ErrUnableToGetStorageData)

		return false, nil
	}

	now := time.Now().String()
	id := data.Total + 1

	newTask := TaskDTO{
		Id:          &id,
		Description: task.Description,
		Status:      TODO,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	byteData, err := json.Marshal(newTask)

	if err != nil {
		log.Println(ErrUnableToCreateNewTask)

		return false, nil
	}

	savedTaskString, err := t.storage.InsertOrUpdate(byteData)

	if err != nil || savedTaskString == nil {
		log.Println(ErrUnableToCreateNewTask)

		return false, nil
	}
	var savedTask TaskDTO

	err = json.Unmarshal([]byte(*savedTaskString), &savedTask)

	if err != nil {
		log.Println(ErrUnableToCreateNewTask)

		return false, nil
	}

	return true, &savedTask
}

func (t *Task) DeleteTask(taskID string) bool {
	deleted, err := t.storage.Delete(taskID)

	if err != nil || !deleted {
		log.Println(ErrUnableToDeleteTask)

		return false
	}

	return true
}

func (t *Task) UpdateTask() {}

func (t *Task) GetAllTasks() []TaskDTO {
	var list []TaskDTO
	data, err := t.storage.GetAll()

	if err != nil {
		log.Println("negerai")
	}

	for _, value := range data.Records {
		var task TaskDTO

		err := json.Unmarshal([]byte(value), &task)

		if err != nil {
			log.Println(err)
		}

		list = append(list, task)
	}

	return list
}

func (t *Task) GetTaskByStatus(status string) {}
