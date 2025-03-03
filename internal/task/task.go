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
		Status:      TODO,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	data, err := json.Marshal(newTask)

	log.Println(task.Id)
	if err != nil {
		panic(ErrUnableToCreateNewTask)
	}

	t.storage.InsertOrUpdate(data)
}

func (t *Task) DeleteTask(taskID int) {
	t.storage.Delete(taskID)

	log.Println("Task has been deleted")
}

func (t *Task) UpdateTask() {

}

func (t *Task) GetAllTasks() []TaskDTO {
	var list []TaskDTO
	data, err := t.storage.GetAll()

	if err != nil {
		log.Println("negerai")
	}

	for _, value := range data.Records {
		var task TaskDTO

		err = json.Unmarshal([]byte(value), &task)

		if err != nil {
			log.Println(err)
		}

		list = append(list, task)
	}

	return list
}

func (t *Task) GetTaskByStatus(status string) {}
