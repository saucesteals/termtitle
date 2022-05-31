//go:build windows

package termtitle

import (
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	libkernel32          = windows.NewLazySystemDLL("kernel32.dll")
	procSetConsoleTitleW = libkernel32.NewProc("SetConsoleTitleW")
	procGetConsoleTitleW = libkernel32.NewProc("GetConsoleTitleW")
)

var (
	maxTitleBuffer = 256
	titlePool      = sync.Pool{
		New: func() interface{} {
			return make([]uint16, maxTitleBuffer)
		},
	}
)

func setTitle(title string) error {
	var ptr *uint16

	if title == "" {
		ptr = new(uint16)
	} else {
		var err error
		ptr, err = windows.UTF16PtrFromString(title)

		if err != nil {
			return err
		}
	}

	_, _, err := procSetConsoleTitleW.Call(
		uintptr(unsafe.Pointer(ptr)),
	)

	if err != windows.ERROR_SUCCESS {
		return err
	}

	return nil
}

func getTitle() (string, error) {
	buff := titlePool.Get().([]uint16)
	defer titlePool.Put(buff)

	ptr := &buff[0]

	_, _, err := procGetConsoleTitleW.Call(
		uintptr(unsafe.Pointer(ptr)), uintptr(maxTitleBuffer),
	)

	if err != windows.ERROR_SUCCESS {
		return "", err
	}

	return windows.UTF16PtrToString(ptr), nil
}
