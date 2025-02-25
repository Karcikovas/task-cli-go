package task

type TaskDTO struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	CreatedAt   *string `json:"createdAt"`
	UpdatedAt   *string `json:"updatedAt"`
}
