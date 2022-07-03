package utils

import (
	"fmt"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func searchup(dir string, filename string) string {
	if dir == "/" || dir == "" {
		return ""
	}

	if _, err := os.Stat(path.Join(dir, filename)); err == nil {
		return path.Join(dir, filename)
	}

	return searchup(path.Dir(dir), filename)
}

func SetEnviroment(env string) error {
	directory, err := os.Getwd()
	if err != nil {
		return err
	}

	filename := fmt.Sprintf(".env%s", env)
	filepath := searchup(directory, filename)

	if filepath == "" {
		return fmt.Errorf("could not find dot env file: %s", filename)
	}

	return godotenv.Load(filepath)
}
