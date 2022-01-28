package util

import (
	"fmt"
	"os"
)

func Clone(url string) {
	os.Chdir("workspace")
	cmd := fmt.Sprintf("git clone %s", url)
	Run(cmd)
}
