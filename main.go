package main

import (
	"github.com/ymktmk/golang-clean-architecture/app/infrastructure"
)

func main() {
	e := infrastructure.NewRouter()
	e.Logger.Fatal(e.Start(":9000"))
}
