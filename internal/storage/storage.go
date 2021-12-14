package storage

import (
	"fmt"
	"io/ioutil"
)

type Memory struct {
	memory float64
}

type File struct {
	name string
}

func NewMemory() *Memory {
	return &Memory{memory: 0}
}

func NewFile() *File {
	return &File{name: "go_file.txt"}
}

func (m *Memory) Save(result float64) (bool, error) {
	m.memory = result
	return true, nil
}

func (f *File) Save(result float64) (bool, error) {
    b := []byte(fmt.Sprintf("El resultado es: %f", result))
    err := ioutil.WriteFile(f.name, b, 0644)
	fmt.Println("ERROOOOR", err)
    if err != nil {
        return false, err
    }
	return true, nil
}