package util

import (
	"os"
)

func Build() {
	os.Chdir("workspace/react-webpack")
	Run("pnpm i")
}
