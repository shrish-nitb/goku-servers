package servernativeapp

import (
	"servernative/pkg/httpserver/httpserverapp"
	"servernative/pkg/httpserver/httpserverapp/middleware"
	"servernative/pkg/todo"
)

func Run() {
	todos := todo.NewTodos()

	todoHandler := todos.CreateHandlers()

	config := httpserverapp.Config{
		Addr: "127.0.0.1:8000",
	}

	masterHandle := httpserverapp.New()

	//-> Logger <-> Parser <-> todoHandler
	masterHandle.Use(middleware.Logger())
	masterHandle.Use(middleware.BodyParser())
	masterHandle.Pass(todoHandler)

	httpserverapp.Run(&config, masterHandle)
}
