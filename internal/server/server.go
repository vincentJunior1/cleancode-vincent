package server

import (
	"fmt"

	router "github.com/vincentJunior1/cleancode-vincent/internal/router"
)

func Init() {
	router.Init()
	fmt.Println("init server")
}
