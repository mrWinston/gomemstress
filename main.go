package main

import (
	"flag"
	"fmt"
	"math"
	"runtime"
	"time"
)

var s [][]byte

func main() {
	fmt.Println("Starting memalloc")
  
  sleepDuration := flag.Int("sleep", 250, "Time in ms to wait between allocations")
  chunkSize := flag.Int("size", 100, "Chunk size in MiB to allocate per cycle")
  maxMem := flag.Int("max", math.MaxInt, "Maximum memory that should be consumed")
  flag.Parse()
	var m runtime.MemStats
  for x := 0; x < math.MaxInt; x++ {
    if x * *chunkSize >= *maxMem {
      fmt.Printf("Allocated enough, from now on i just wait\n")

    } else {
      p := make([]byte, 1024*1024 * *chunkSize)
      p[0] = 12
      for i := 1; i < len(p); i *= 2 {
        copy(p[i:], p[:i])
      }
      s = append(s, p)
    }
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
		fmt.Printf("\tTotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
    fmt.Printf("----")
    time.Sleep(time.Millisecond * time.Duration(*sleepDuration))
	}

}
