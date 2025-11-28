package report

import (
	"fmt"
)

func Print(totalRequests int, result map[int]int, elapsed string, errorTypes map[string]int) {
	fmt.Println("\n--- Relatório ---")
	fmt.Printf("Tempo total: %v\n", elapsed)
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Printf("HTTP 200: %d\n", result[200])
	fmt.Println("Distribuição de status:")
	for code, count := range result {
		if code == 200 {
			continue
		}
		if code == 0 {
			fmt.Printf("Erros de request: %d\n", count)
			if len(errorTypes) > 0 {
				fmt.Println("Tipos de erro:")
				for errStr, errCount := range errorTypes {
					fmt.Printf("- %s: %d\n", errStr, errCount)
				}
			}
		} else {
			fmt.Printf("HTTP %d: %d\n", code, count)
		}
	}
}
