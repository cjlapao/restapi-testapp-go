package bootstrap

import (
	"os"
)

func Start() {
	if !FileExists(TempFolder()) {
		os.MkdirAll(TempFolder(), os.ModePerm)
	}
}
