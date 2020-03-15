package internal

import (
	"path"
	"runtime"
)

// FuncName returns the current function name
func FuncName(back int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(back+2, pc)
	return path.Base(runtime.FuncForPC(pc[0]).Name())
}
