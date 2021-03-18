// +build linux

package godmi

import (
	"os"
	"syscall"
)

func getMem(base uint32, length uint32) (mem []byte, err error) {
	file, err := os.Open("/dev/mem")
	if err != nil {
		return
	}
	defer file.Close()
	fd := file.Fd()
	mmoffset := base % uint32(os.Getpagesize())
	mm, err := syscall.Mmap(int(fd), int64(base-mmoffset), int(mmoffset+length), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		return
	}
	mem = make([]byte, len(mm))
	copy(mem, mm)
	err = syscall.Munmap(mm)
	if err != nil {
		return
	}
	return
}
