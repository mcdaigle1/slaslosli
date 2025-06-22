package fileutils

import (
	"os"
)

func GetWorkingDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return wd, nil
}
