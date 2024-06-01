package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets info from the database",
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Gets all users",
	Run: func(cmd *cobra.Command, args []string) {
		getUsers()
	},
}

var permissionsCmd = &cobra.Command{
	Use:   "permissions",
	Short: "Gets all permissions",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("You must provide a username")
			return
		}
		getPermissions(args[0])
	},
}

func getPermissions(username string) {
	permissions, err := userRepo.GetPermissions(username)
	if err != nil {
		panic(err)
	}

	for _, permission := range permissions {
		fmt.Println(permission)
	}
}

func getUsers() {
	users, err := userRepo.FindAll()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.ToPrettyStr())
	}
}
