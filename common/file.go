package common

import (
	"os"
)

func ReadFile(filename string) string {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(file)
}
