package util

import (
	"os"
)

var BasePath string

func InitBasePath() {
	if BasePath == "" {
		pwd, _ := os.Getwd()
		BasePath = pwd
	}
}
