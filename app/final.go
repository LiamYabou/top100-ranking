package app

import (
	"github.com/LiamYabou/top100-ranking/variable"
)

func Finalize() {
	if variable.Env == "development" {
		file.Close()
	}
	DBPool.Close()
}
