package output

type CreateTodo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type UpdateTodo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type GetTodo struct {
	TodoID int `json:"todo_id"`
}

type GetAllTodos struct {
	UserID int `json:"user_id"`
}
