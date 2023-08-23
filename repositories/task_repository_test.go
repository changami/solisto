package repositories

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"solisto/entities"
	"testing"
)

func TestTaskRepository_Add(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	task := entities.Task{Name: "Adding Test"}
	err := repo.Add(task)
	assert.NoError(t, err)

	var actual entities.Task
	db.First(&actual)

	assert.NotNil(t, actual.ID)
	assert.Equal(t, "Adding Test", actual.Name)
	assert.NotNil(t, actual.CreatedAt)
	assert.NotNil(t, actual.UpdatedAt)
	assert.False(t, actual.DeletedAt.Valid)
}

func TestTaskRepository_GetAll(t *testing.T) {
	db := setupTestDB(t)
	repo := NewTaskRepository(db)

	db.Create(&entities.Task{Name: "GetAll Test 1"})
	db.Create(&entities.Task{Name: "GetAll Test 2"})

	targets := repo.GetAll()
	assert.Equal(t, 2, len(targets))
}

func setupTestDB(t *testing.T) *Database {
	testDB, err := gorm.Open(sqlite.Open("file::memory:" /* +"?cache=shared" */), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	return &Database{DB: testDB}
}
