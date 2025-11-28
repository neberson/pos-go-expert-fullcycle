package cobracli

import (
	"fmt"
	"os"

	"stresstest/internal/stresstest"
	"stresstest/pkg/report"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stresstest",
	Short: "Ferramenta de stress test HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		if url == "" {
			fmt.Fprintln(os.Stderr, "Erro: parâmetro --url é obrigatório")
			os.Exit(1)
		}
		if requests <= 0 {
			fmt.Fprintln(os.Stderr, "Erro: --requests deve ser maior que 0")
			os.Exit(1)
		}
		if concurrency <= 0 {
			fmt.Fprintln(os.Stderr, "Erro: --concurrency deve ser maior que 0")
			os.Exit(1)
		}

		fmt.Printf("Iniciando teste de carga em %s com %d requests e %d concorrentes...\n", url, requests, concurrency)
		cfg := stresstest.Config{
			URL:         url,
			Requests:    requests,
			Concurrency: concurrency,
		}
		result := stresstest.Run(cfg)
		report.Print(cfg.Requests, result.StatusCount, result.Elapsed.String())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("url", "", "URL do serviço a ser testado")
	rootCmd.Flags().Int("requests", 1, "Número total de requests")
	rootCmd.Flags().Int("concurrency", 1, "Número de chamadas simultâneas")
}
