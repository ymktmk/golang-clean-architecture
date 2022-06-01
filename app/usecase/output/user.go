package output

type User struct {
	ID     uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type Users []User