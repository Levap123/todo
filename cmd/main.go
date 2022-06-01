package main

import (
	"github.com/Levap123/toDo.git"
	"github.com/Levap123/toDo.git/package/handler"
	"log"
)

func main() {
	handler := new(handler.Handler)
	server := new(todo.Server)
	err := server.Run("4444", handler.InitRoutes())
	if err != nil {
		log.Fatalln(err)
	}
}
