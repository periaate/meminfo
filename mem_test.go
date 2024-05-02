package meminfo

import (
	"fmt"
	"testing"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func TestGet(t *testing.T) {
	mi, err := Get()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Total load\t%v%%\n", mi.Load)
	fmt.Printf("PageTotal\t%v\n", HumanizeBytes(1, mi.PageTotal, 2, true))
	fmt.Printf("PageAvail\t%v\n", HumanizeBytes(1, mi.PageAvail, 2, true))
	fmt.Printf("PageUsed\t%v\n", HumanizeBytes(1, mi.PageUsed, 2, true))
	fmt.Printf("PhysTotal\t%v\n", HumanizeBytes(1, mi.PhysTotal, 2, true))
	fmt.Printf("PhysAvail\t%v\n", HumanizeBytes(1, mi.PhysAvail, 2, true))
	fmt.Printf("PhysUsed\t%v\n", HumanizeBytes(1, mi.PhysUsed, 2, true))
	fmt.Printf("VirtTotal\t%v\n", HumanizeBytes(1, mi.VirtTotal, 2, true))
	fmt.Printf("VirtAvail\t%v\n", HumanizeBytes(1, mi.VirtAvail, 2, true))
	fmt.Printf("VirtUsed\t%v\n", HumanizeBytes(1, mi.VirtUsed, 2, true))
}

// HumanizeBytes converts an integer byte value into a human-readable string.
// base of 0 means the input is in bytes, of 1 in kB, ...
func HumanizeBytes[T Integer](base int, val T, decimals int, asKiB bool) string {
	var unit float64 = 1000 // Use 1000 as base for KB, MB, GB...
	var suffixes = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	if asKiB {
		unit = 1024 // Use 1024 as base for KiB, MiB, GiB...
		suffixes = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	}
	suffixes = suffixes[base:]

	if val == 0 {
		return fmt.Sprintf("%.*f %s", decimals, 0.0, suffixes[0])
	}

	negative := val < 0
	val = Abs(val)

	size := float64(val)
	i := 0
	for size >= unit && i < len(suffixes)-1 {
		size /= unit
		i++
	}

	if negative {
		size = -size
	}

	return fmt.Sprintf("%.*f %s", decimals, size, suffixes[i])
}

func Abs[N Integer](x N) (zero N) {
	if x < zero {
		return -x
	}
	return x
}
