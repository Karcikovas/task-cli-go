package task

import (
	"encoding/json"
	"fmt"
	"log"
	"task-cli-go/internal/storage"
	"time"
)

var (
	DONE       = "done"
	TODO       = "todo"
	InPROGRESS = "in-progress"
)

type Task struct {
	storage *storage.Storage
}

func CreateNewTask(storage *storage.Storage) *Task {
	return &Task{
		storage: storage,
	}
}

func (t *Task) CreateTask(task TaskDTO) {
	now := time.Now().String()

	newTask := TaskDTO{
		Id:          task.Id,
		Description: TODO,
		Status:      task.Status,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	data, err := json.Marshal(newTask)

	if err != nil {
		panic(ErrUnableToCreateNewTask)
	}

	t.storage.Save(data, fmt.Sprintf(`task-%s-data`, task.Id))
}

func (t *Task) DeleteTask(taskID string) {
	t.storage.Delete(fmt.Sprintf(`task-%s-data`, taskID))

	log.Println("Task has been deleted")
}

func (t *Task) UpdateTask() {

}

func (t *Task) GetAllTasks() []TaskDTO {
	var list []TaskDTO
	data, err := t.storage.GetAllStorageData()

	if err != nil {
		log.Println("negerai")
	}

	for _, value := range data {
		log.Println(value)
	}

	return list
}

func (t *Task) GetTaskByStatus(status string) {}
