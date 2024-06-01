package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var flagsCmd = &cobra.Command{
	Use:   "flags",
	Short: "Flag utilities",
}

var flagsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets flags",
	Long:  "It gets all flags if no flag name is specified. If a flag name is specified, it gets the flag with that name",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			flags, err := flagService.List()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(flags)
		} else {
			flag, err := flagService.Get(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(flag)
		}
	},
}
