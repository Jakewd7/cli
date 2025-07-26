package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gojake",
	Short: "Custom CLI scaffolding tool for Go projects",
	Long:  `Gojake is a modular CLI tool that helps scaffold scalable Golang application structures.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
