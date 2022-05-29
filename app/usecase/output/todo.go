package output

type Todo struct {
	ID     uint   `json:"id"`
	Name   string `json:"Name"`
	UserID uint   `json:"user_id"`
}

type Todos []Todo
