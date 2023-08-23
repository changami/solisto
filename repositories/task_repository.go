package repositories

import (
	"solisto/entities"
)

type TaskRepository struct {
	Database *Database
}

func NewTaskRepository(db *Database) *TaskRepository {
	err := db.AutoMigrate(&entities.Task{})
	if err != nil {
		panic("failed to migrate Task table")
	}

	return &TaskRepository{Database: db}
}

func (r *TaskRepository) GetAll() []entities.Task {
	var results []entities.Task
	r.Database.Find(&results)

	return results
}

func (r *TaskRepository) Add(task entities.Task) error {
	return r.Database.Create(&task).Error
}
