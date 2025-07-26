package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goapp",
	Short: "CLI scaffolding tool seperti breeze untuk Golang",
	Long:  `GoApp adalah CLI scaffolding tool untuk membantu generate struktur modular di proyek Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
