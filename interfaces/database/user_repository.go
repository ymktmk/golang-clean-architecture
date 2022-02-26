// User関連のSQL実行のコードを書く
package database

import (
	"Golang-CleanArchitecture/domain"
	"fmt"
)

type UserRepository struct {
    SqlHandler
}

func (repo *UserRepository) Store(user domain.User) (int, error) {
    result, err := repo.Execute(
        "INSERT INTO users (name) VALUES (?)", user.Name,
    )
    if err != nil {
        fmt.Println(err)
    }
    id64, err := result.LastInsertId()
    if err != nil {
        fmt.Println(err)
    }
    id := int(id64)
    return id, err
}

func (repo *UserRepository) FindById(identifier int) (domain.User, error) {
    row, err := repo.Query("SELECT id, name FROM users WHERE id = ?", identifier)
    defer row.Close()
    if err != nil {
        fmt.Println(err)
    }
    var user domain.User
    var id int
    var name string
    row.Next()
    if err = row.Scan(&id, &name); err != nil {
        fmt.Println(err)
    }
    user.ID = id
    user.Name = name
    return user, err
}
