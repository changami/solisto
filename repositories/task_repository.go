package repositories

import (
	"gorm.io/gorm"
	"solisto/entities"
)

type ITaskRepository interface {
	GetAll() []entities.Task
	Add(entities.Task) (entities.Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func init() {
	err := db.AutoMigrate(&entities.Task{})
	if err != nil {
		panic("failed to migrate Task table")
	}
}

func NewTaskRepository() ITaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) GetAll() []entities.Task {
	var results []entities.Task
	r.db.Find(&results)

	return results
}

func (r *TaskRepository) Add(task entities.Task) (entities.Task, error) {
	err := r.db.Create(&task).Error
	if err != nil {
		panic("failed to add task")
	}

	return task, err
}
