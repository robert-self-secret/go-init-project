package main

import (
	"fmt"

	"github.com/robert-self-secret/go-init-project.git/internal/config"
	"github.com/robert-self-secret/go-init-project.git/internal/connection"
)

func main() {
	cnf := config.GetConfig()
	connection.GetDBOpenConenction(cnf)
	connection.GetRedisConnection(cnf)

	fmt.Println("Hello world")
}
