package main

import (
	"fmt"
	"log"

	"github.com/markelca/toggles/pkg/user"
	"github.com/spf13/cobra"
)

func init() {
	dbGetCmd.AddCommand(usersCmd)
	dbGetCmd.AddCommand(permissionsCmd)

	databaseCmd.AddCommand(dbGetCmd)
	databaseCmd.AddCommand(initCmd)

}

var databaseCmd = &cobra.Command{
	Use:   "db",
	Short: "Database utilities",
}

var dbGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets info from the database",
}
var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initializes the database",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		adminPermissions := []string{"get_flags", "create_flag", "update_flag", "delete_flag", "get_users", "create_users", "update_users", "delete_users", "get_me"}
		adminUser, err := user.NewUser("admin", "admin", adminPermissions)
		if err != nil {
			log.Fatal(err)
		}
		err = userRepo.Upsert(*adminUser)
		if err != nil {
			log.Fatal(err)
		}

		userPermissions := []string{"get_flags", "get_flag", "get_me"}
		testUser, err := user.NewUser("test", "test", userPermissions)
		if err != nil {
			log.Fatal(err)
		}
		err = userRepo.Upsert(*testUser)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "Gets all users",
	Run: func(cmd *cobra.Command, args []string) {
		users, err := userRepo.FindAll()
		if err != nil {
			panic(err)
		}

		for _, user := range users {
			fmt.Println(user.ToPrettyStr())
		}
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
		username := args[0]
		permissions, err := userRepo.GetPermissions(username)
		if err != nil {
			panic(err)
		}

		for _, permission := range permissions {
			fmt.Println(permission)
		}
	},
}
