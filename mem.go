package meminfo

// MemoryInfo holds data about system memory usage.
type MemoryInfo struct {
	Load      uint32 // 0-100 percentage
	PhysTotal uint64 // in KiB
	PhysAvail uint64 // in KiB
	PhysUsed  uint64 // in KiB
	PageTotal uint64 // in KiB
	PageAvail uint64 // in KiB
	PageUsed  uint64 // in KiB
	VirtTotal uint64 // in KiB
	VirtAvail uint64 // in KiB
	VirtUsed  uint64 // in KiB
}
