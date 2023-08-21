package main

import (
	"go/controller"
	"go/db"
	"go/repository"
	"go/router"
	"go/usecase"
)

// サーバー立ちあげ
func main() {
	// データベースに接続
	db := db.NewDB()

	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)

	e := router.Setup(taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
