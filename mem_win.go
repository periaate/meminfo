//go:build windows
// +build windows

package meminfo

import (
	"syscall"
	"unsafe"
)

type MEMORYSTATUSEX struct {
	dwLength                uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func Get() (*MemoryInfo, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	GlobalMemoryStatusEx := kernel32.NewProc("GlobalMemoryStatusEx")

	var memStat MEMORYSTATUSEX
	memStat.dwLength = uint32(unsafe.Sizeof(memStat))
	_, _, err := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memStat)))
	if err != syscall.Errno(0) {
		return nil, err
	}

	return &MemoryInfo{
		Load:      memStat.dwMemoryLoad,
		PhysTotal: memStat.ullTotalPhys / 1024,
		PhysAvail: memStat.ullAvailPhys / 1024,
		PhysUsed:  (memStat.ullTotalPhys - memStat.ullAvailPhys) / 1024,
		PageTotal: memStat.ullTotalPageFile / 1024,
		PageAvail: memStat.ullAvailPageFile / 1024,
		PageUsed:  (memStat.ullTotalPageFile - memStat.ullAvailPageFile) / 1024,
		VirtTotal: memStat.ullTotalVirtual / 1024,
		VirtAvail: memStat.ullAvailVirtual / 1024,
		VirtUsed:  (memStat.ullTotalVirtual - memStat.ullAvailVirtual) / 1024,
	}, nil
}
