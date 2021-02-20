package main

import (
	"os"
	"path/filepath"
)

func main() {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, todoFilename)
		_, err = os.Stat(filename)
		if err == nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todoFilename)
	}

}