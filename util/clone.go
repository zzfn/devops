package util

import (
	"fmt"
	"os"
)

func Clone(url string, repo string) {
	basePath := BasePath + "/workspace"
	dir, err := os.Stat(basePath)
	if err != nil || !dir.IsDir() {
		os.Chdir(BasePath)
		os.Mkdir("workspace", os.ModePerm)
	}
	os.Chdir(basePath)
	if Exists(basePath + "/" + repo) {
		Run("git pull")
	} else {
		cmd := fmt.Sprintf("git clone %s %s", url, repo)
		Run(cmd)
	}
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
