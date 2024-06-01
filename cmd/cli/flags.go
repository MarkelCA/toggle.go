package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/markelca/toggles/pkg/flags"
	"github.com/spf13/cobra"
)

func init() {
	flagsGetCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print")
	flagsCmd.AddCommand(flagsGetCmd)
	flagsCmd.AddCommand(flagsCreateCmd)
	flagsCmd.AddCommand(flagsDeleteCmd)
	flagsCmd.AddCommand(flagsUpdateCmd)
	rootCmd.AddCommand(flagsCmd)
}

var flagsCmd = &cobra.Command{
	Use:   "flag",
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
			fmt.Println("Invalid number of arguments, expected 0 or 1")
			os.Exit(1)
		}
	},
}

var flagsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a flag",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments, expected 2 (name,value)")
			os.Exit(1)
		}

		name := args[0]
		value := args[1]
		valueBool, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatal("Invalid value, expected a boolean")
		}

		flag := flags.Flag{Name: name, Value: valueBool}
		err = flagService.Create(flag)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Flag created: %s\n", flag)
	},
}

var flagsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a flag",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Invalid number of arguments, expected 1 (name)")
			os.Exit(1)
		}

		name := args[0]
		err := flagService.Delete(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Flag deleted: %s\n", name)
	},
}

var flagsUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a flag",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments, expected 2 (name,value)")
			os.Exit(1)
		}

		name := args[0]
		value := args[1]
		valueBool, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatal("Invalid value, expected a boolean")
		}

		err = flagService.Update(name, valueBool)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Flag updated: %s\n", flags.Flag{Name: name, Value: valueBool})
	},
}
