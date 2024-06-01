package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	flagsGetCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print")
	flagsCmd.AddCommand(flagsGetCmd)
	rootCmd.AddCommand(flagsCmd)
}

var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "Flag utilities",
}

var flagsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets flags",
	Long:  "It gets all flags if no flag name is specified. If a flag name is specified, it gets the flag with that name",
	Run: func(cmd *cobra.Command, args []string) {
		prettyPrint := cmd.Flag("pretty").Value
		if len(args) == 1 {
			flag, err := flagService.Get(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(flag)
		} else if len(args) == 0 {
			flags, err := flagService.List()
			if err != nil {
				fmt.Println(err)
				return
			}
			var jsonBody []byte
			if prettyPrint.String() == "true" {
				jsonBody, err = json.MarshalIndent(flags, "", "  ")
			} else {
				jsonBody, err = json.Marshal(flags)

			}
			if err != nil {
				fmt.Println(flags)
				return
			}
			fmt.Println(string(jsonBody))
		} else {
			log.Fatal("Invalid number of arguments, expected 0 or 1")
		}
	},
}
