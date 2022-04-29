package utils_test

import (
	"fmt"
	"testing"

	"github.com/ymktmk/golang-clean-architecture/app/utils"
)

func TestCreateToken(t *testing.T) {
	token := utils.CreateToken(2)
	fmt.Println(token)
}
