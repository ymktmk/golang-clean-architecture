package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	dsn := os.Getenv(("MYSQL_USER")) + ":" + os.Getenv(("MYSQL_PASSWORD")) + "@tcp(" + os.Getenv(("MYSQL_HOST")) + ":" + os.Getenv(("MYSQL_PORT")) + ")/" + os.Getenv(("MYSQL_DATABASE")) + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	// logger
	db.Logger = db.Logger.LogMode(logger.Info)
	// drop & migration
	err = db.Migrator().DropTable(
		&domain.User{},
		&domain.Todo{},
	)
	err = db.Migrator().CreateTable(
		domain.User{},
		&domain.Todo{},
	)
	// .AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
