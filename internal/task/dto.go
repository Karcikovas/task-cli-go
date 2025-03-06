package task

type TaskDTO struct {
	Id          *int    `json:"id"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
}

type UpdateTaskDTO struct {
	ID          int
	Description *string
	Status      *string
	UpdatedAt   *string
}
