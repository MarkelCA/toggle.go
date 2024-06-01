package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	usersGetCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print")
	usersCmd.AddCommand(usersGetCmd)
	permissionsCmd.AddCommand(userPermissionsGetCmd)
	permissionsCmd.AddCommand(userPermissionAddCmd)
	permissionsCmd.AddCommand(userPermissionRemoveCmd)
	usersCmd.AddCommand(permissionsCmd)
	rootCmd.AddCommand(usersCmd)
}

var usersCmd = &cobra.Command{
	Use:   "user",
	Short: "User utilities",
}

var permissionsCmd = &cobra.Command{
	Use:   "permission",
	Short: "Gets all permissions",
}

var userPermissionsGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a user's permissions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You must provide a username")
			return
		}
		username := args[0]
		permissions, err := userService.GetPermissions(username)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, permission := range permissions {
			fmt.Println(permission)
		}
	},
}

var usersGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a user",
	Run: func(cmd *cobra.Command, args []string) {
		prettyPrint := cmd.Flag("pretty").Value
		if len(args) > 1 {
			fmt.Println("Too many arguments")
			return
		}

		if len(args) == 1 {
			user, err := userService.FindByUserName(args[0])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if prettyPrint.String() == "true" {
				fmt.Println(user.ToPrettyStr())
			} else {
				fmt.Println(user)
			}

		} else {
			users, err := userService.FindAll()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			for _, user := range users {
				if prettyPrint.String() == "true" {
					fmt.Println(user.ToPrettyStr())
				} else {
					fmt.Println(user)
				}
			}
		}

	},
}

var userPermissionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a permission to a user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments, expected 2 (username, permission)")
			os.Exit(1)
		}
		username := args[0]
		permission := args[1]
		err := userService.AddPermission(username, permission)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Permission added")
	},
}

var userPermissionRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a permission from a user",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments, expected 2 (username, permission)")
			os.Exit(1)
		}
		username := args[0]
		permission := args[1]
		err := userService.RemovePermission(username, permission)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Permission removed")
	},
}
