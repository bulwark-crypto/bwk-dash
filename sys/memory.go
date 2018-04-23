package sys

import "github.com/pbnjay/memory"

// GetMemorySize will return the total bytes in RAM
// for the running system.
func GetMemorySize() (size float64) {
	size = float64(memory.TotalMemory())
	return
}
