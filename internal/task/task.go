package task

import (
	"encoding/json"
	"fmt"
	"task-cli-go/internal/storage"
	"time"
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
		Description: task.Description,
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

func (t *Task) updateTask() {

}

func (t *Task) DeleteTask(taskID string) {
	t.storage.Delete(taskID)
}
