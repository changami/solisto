package main

import (
	"context"
	"solisto/entities"
	"solisto/repositories"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GetAllTasks returns a list of all tasks
func (a *App) GetAllTasks() []entities.Task {
	taskRepository := repositories.NewTaskRepository()
	return taskRepository.GetAll()
}
