// +build !linux

package godmi

import (
	"fmt"
	"runtime"
)

func getMem(base uint32, length uint32) (mem []byte, err error) {
	panic(fmt.Errorf("not support os: %v", runtime.GOOS))
}