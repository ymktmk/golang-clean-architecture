// User関連のSQL実行のコードを書く
package database

import (
	"Golang-CleanArchitecture/domain"
	"fmt"
)

// ここで再定義して使う
type UserRepository struct {
    // type SqlHandler interface {
    //     Execute(string, ...interface{}) (Result, error)
    //     Query(string, ...interface{}) (Row, error)
    // }
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


func (repo *UserRepository) FindAll() (domain.Users, error) {
    rows, err := repo.Query("SELECT id, name FROM users")
    defer rows.Close()
    if err != nil {
        fmt.Println(err)
    }

    var users domain.Users
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            continue
        }
        user := domain.User{
            ID: id,
            Name: name,
        }
        users = append(users, user)
    }
    
    return users, err
}
