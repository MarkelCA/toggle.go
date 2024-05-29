package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tg",
	Short: "This is the togggles command line tool",
	Long:  `This tool offers utilities to interact with the togggles API.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the first cobra example")
	},
}

func main() {
	databaseCmd.AddCommand(initCmd)
	rootCmd.AddCommand(databaseCmd)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
