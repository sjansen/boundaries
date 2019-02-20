package c

import (
	"path/filepath"
	"runtime"
)

func Value() string {
	if _, path, _, ok := runtime.Caller(0); ok {
		basename := filepath.Base(path)
		ext := filepath.Ext(basename)
		return basename[:len(basename)-len(ext)]
	} else {
		return "huh..."
	}
}
