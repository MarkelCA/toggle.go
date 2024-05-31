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

func getUsers() {
	users, err := userRepo.FindAll()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Println(user.ToPrettyStr())
	}
}
