package util

import (
	"os"
)

func Build() {
	os.Chdir(BasePath + "/workspace/react-webpack")
	Run("pnpm i")
	Run("pnpm run build:prod")
}
