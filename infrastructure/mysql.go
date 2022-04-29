package infrastructure

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ymktmk/golang-clean-architecture/interfaces/database"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() database.SqlHandler {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	dsn := os.Getenv(("MYSQL_USER")) + ":" + os.Getenv(("MYSQL_PASSWORD")) + "@tcp(" + os.Getenv(("MYSQL_HOST")) + ":" + os.Getenv(("MYSQL_PORT")) + ")/" + os.Getenv(("MYSQL_DATABASE")) + "?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func (handler *SqlHandler) Joins(query string, args ...interface{}) *gorm.DB {
	return handler.Conn.Joins(query, args...)
}

func (handler *SqlHandler) Model(value interface{}) *gorm.DB {
	return handler.Conn.Model(value)
}

func (handler *SqlHandler) Preload(query string, args ...interface{}) *gorm.DB {
	return handler.Conn.Preload(query, args...)
}

func (handler *SqlHandler) Table(query string, args ...interface{}) *gorm.DB {
	return handler.Conn.Table(query, args...)
}
