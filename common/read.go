package common

import (
	"os"
	"path"
)

func Read(year string, filename string) (string, error) {
	pathTemplate := path.Join("./data", year, filename+".txt")
	file, err := os.ReadFile(pathTemplate)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
