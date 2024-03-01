package godb

import (
	"os"
)

func createDB(name string, data_path string) {
	// Create a new file
	file, err := os.Create(data_path + name + ".db")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
