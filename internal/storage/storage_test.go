package storage_test

import (
	"main/internal/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenCallSaveWithMemoryStructThenSuccess(t *testing.T) {
	storage := storage.NewMemory()
	result, error := storage.Save(1)

	assert.Nil(t, error)
	assert.EqualValues(t, true, result)
}

func TestWhenCallSaveWithFileStructThenSuccess(t *testing.T) {
	storage := storage.NewFile()
	result, error := storage.Save(1)

	assert.Nil(t, error)
	assert.EqualValues(t, true, result)
}
