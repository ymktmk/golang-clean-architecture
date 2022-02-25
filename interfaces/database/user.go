package database

import "Golang-CleanArchitecture/domain"


type UserRepository struct {
    SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
    result, err := repo.Execute(
        "INSERT INTO users (name) VALUES (?)", u.Name,
    )
    if err != nil {
        return
    }
    id64, err := result.LastInsertId()
    if err != nil {
        return
    }
    id = int(id64)
    return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
    row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
    defer row.Close()
    if err != nil {
        return
    }
    var id int
    var name string
    row.Next()
    if err = row.Scan(&id, &name); err != nil {
        return
    }
    user.ID = id
	user.Name = name
    return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
    rows, err := repo.Query("SELECT id, first_name, last_name FROM users")
    defer rows.Close()
    if err != nil {
        return
    }
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            continue
        }
        user := domain.User{
            ID:        id,
            Name: name,
        }
        users = append(users, user)
    }
    return
}