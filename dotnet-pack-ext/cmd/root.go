package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "dotnet pack-ext",
	Short:   "Dotnet Pack Extensions Command Line",
	Long:    `Dotnet Pack Extensions Command Line`,
	Version: "0.0.1",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
