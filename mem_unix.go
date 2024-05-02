//go:build unix
// +build unix

package meminfo

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// parseMemInfo helps parse lines from /proc/meminfo and extract memory values.
func parseMemInfo(line string) (string, uint64, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return "", 0, errors.New("invalid line format")
	}
	value, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return "", 0, err
	}
	return fields[0], value, nil
}

// Get fetches the current memory usage data from /proc/meminfo.
func Get() (*MemoryInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalMem := uint64(0)
	freeMem := uint64(0)
	bufferMem := uint64(0)
	cachedMem := uint64(0)
	swapTotal := uint64(0)
	swapFree := uint64(0)

	for scanner.Scan() {
		key, value, err := parseMemInfo(scanner.Text())
		if err != nil {
			continue // Skip lines with parsing errors
		}

		switch key {
		case "MemTotal:":
			totalMem = value
		case "MemFree:":
			freeMem = value
		case "Buffers:":
			bufferMem = value
		case "Cached:":
			cachedMem = value
		case "SwapTotal:":
			swapTotal = value
		case "SwapFree:":
			swapFree = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &MemoryInfo{
		Load:      uint32(100 - (freeMem*100)/totalMem),
		PhysTotal: totalMem,
		PhysAvail: freeMem + bufferMem + cachedMem,
		PhysUsed:  totalMem - freeMem - bufferMem - cachedMem,
		PageTotal: swapTotal,
		PageAvail: swapFree,
		PageUsed:  swapTotal - swapFree,
		VirtTotal: totalMem + swapTotal,
		VirtAvail: freeMem + bufferMem + cachedMem + swapFree,
		VirtUsed:  totalMem - freeMem - bufferMem - cachedMem + swapTotal - swapFree,
	}, nil
}
