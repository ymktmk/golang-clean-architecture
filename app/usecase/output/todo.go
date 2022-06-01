package output

type Todo struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id"`
}

type Todos []Todo
