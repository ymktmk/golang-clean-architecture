package output

type Todo struct {
	ID     uint   `json:"id"`
	Name   string `json:"Name"`
	UserID int    `json:"user_id"`
}

type Todos []Todo
