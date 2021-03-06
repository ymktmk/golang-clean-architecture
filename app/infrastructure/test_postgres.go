package infrastructure

import (
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ymktmk/golang-clean-architecture/app/interfaces/database"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	mockDB.Logger = mockDB.Logger.LogMode(logger.Info)
	return mockDB, mock, err
}

func SqlMockHandler(conn *gorm.DB) database.SqlHandler {
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}
