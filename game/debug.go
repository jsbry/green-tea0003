package game

import (
	"fmt"
	"runtime"
)

var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}
