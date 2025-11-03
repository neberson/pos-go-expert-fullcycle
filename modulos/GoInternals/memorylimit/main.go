package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	//Definindo um limite de memória de 10 MB
	debug.SetMemoryLimit(10 * 1024 * 1024)

	// Função que alocal memória para demonstrar o funcionamento do GC
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// Loop para alocar memória repetidamente
	for i := 0; i < 10; i++ {
		_ = allocateMemory(20 * 1024 * 1024) // Alocando 20 MB
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
		fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
		fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
		fmt.Printf("Lookups = %v\n", m.Lookups)
		fmt.Printf("Mallocs = %v\n", m.Mallocs)
		fmt.Printf("Frees = %v\n", m.Frees)
		fmt.Printf("HeapAlloc = %v\n", m.HeapAlloc/1024/1024)
		fmt.Printf("HeapSys = %v MiB\n", m.HeapSys/1024/1024)
		fmt.Printf("HeapIdle = %v MiB\n", m.HeapIdle/1024/1024)
		fmt.Printf("HeapInuse = %v MiB\n", m.HeapInuse/1024/1024)
		fmt.Printf("HeapReleased = %v MiB\n", m.HeapReleased/1024/1024)
		fmt.Printf("HeapObjects = %v\n", m.HeapObjects)
		fmt.Printf("StackInuse = %v\n", m.StackInuse/1024/1024)
		fmt.Printf("StackSys = %v MiB\n", m.StackSys/1024/1024)
		fmt.Printf("MSpanInuse = %v MiB\n", m.MSpanInuse/1024/1024)
		fmt.Printf("MSpanSys = %v MiB\n", m.MSpanSys/1024/1024)
		fmt.Printf("MCacheInuse = %v MiB\n", m.MCacheInuse/1024/1024)
		fmt.Printf("MCacheSys = %v MiB\n", m.MCacheSys/1024/1024)
		fmt.Printf("BuckHashSys = %v MiB\n", m.BuckHashSys/1024/1024)
		fmt.Printf("GCSys = %v MiB\n", m.GCSys/1024/1024)
		fmt.Printf("OtherSys = %v MiB\n", m.OtherSys/1024/1024)
		fmt.Println("--------------------------------")
	}
}
