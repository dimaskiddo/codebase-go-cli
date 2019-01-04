package service

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CLI Version Variable Structure
var workdirCLI = &cobra.Command{
	Use:   "workdir",
	Short: "Show the Go CLI working directory",
	Long:  "Show the Go CLI working directory information",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error to Get Current Working Directory!")
			os.Exit(1)
		}

		fmt.Println("Current Working Directory:", pwd)
	},
}
