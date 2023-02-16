package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const projectDirName = "go-gin-boilerplate"

func loadEnv() error {
	var (
		err      error
		envFile  string
		env      = os.Getenv("GO_ENV")
		cwd, _   = os.Getwd()
		rootPath = cwd
	)

	if env == "" {
		envFile = ".env"
	} else {
		envFile = ".env." + env
	}

	f := buildFilePath(rootPath, envFile)

	if testFile(f) {
		log.Println("Using environment from", f)
		err = godotenv.Load(f)
	} else {
		// Fallback to .env in case GO_ENV is set but doesn't exist.
		f = buildFilePath(rootPath, ".env")
		log.Println("Using environment from", f)
		err = godotenv.Load(f)
		return err
	}

	return err
}

func testFile(p string) bool {
	info, err := os.Stat(p)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func buildFilePath(rootPath string, filename string) string {
	return fmt.Sprintf("%s/%s", rootPath, filename)
}
