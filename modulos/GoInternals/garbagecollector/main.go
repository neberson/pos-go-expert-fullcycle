package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//Habilitar o tracing do GC
	// debug.SetGCPercent(-1) // Desabilita o GC
	// Ajustar o percentual do GC (por exemplo, para 300%)
	// debug.SetGCPercent(300)

	// função para alocar memória
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}

	// Alocando memória para observar o comportamento do GC
	for i := 0; i < 10; i++ {
		allocateMemory(10 * 1024 * 1024) // Aloca 10MB
		time.Sleep(time.Second)
	}

	//Exibindo o uso de memória
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc = %v MiB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys = %v MiB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC = %v\n", m.NumGC)
}
