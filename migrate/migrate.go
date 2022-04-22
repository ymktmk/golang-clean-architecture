package main

import (
	"os"

	// v1
	"github.com/joho/godotenv"
	"github.com/ymktmk/golang-clean-architecture/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	dsn := os.Getenv(("MYSQL_USER")) +":"+os.Getenv(("MYSQL_PASSWORD")) +"@tcp("+ os.Getenv(("MYSQL_HOST")) +":" +os.Getenv(("MYSQL_PORT"))+ ")/"+ os.Getenv(("MYSQL_DATABASE")) +"?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// defer func() {
	// 	if err := conn.Close(); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// conn.LogMode(true)
	// if err := conn.DB().Ping(); err != nil {
	// 	panic(err)
	// }
	// drop & migration
	// conn.DropTable(
	// 	&domain.User{},
	// 	&domain.Todo{},
	// )
	// conn.AutoMigrate(&domain.Company{})
	conn.AutoMigrate(
		&domain.User{},
		// &domain.Todo{},
	)
	// .AddForeignKey("company_id", "campanies(id)", "RESTRICT", "RESTRICT")
}