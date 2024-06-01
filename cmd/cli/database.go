package main

import (
	"log"

	"github.com/markelca/toggles/pkg/user"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(dbGetCmd)
	databaseCmd.AddCommand(initCmd)
	rootCmd.AddCommand(databaseCmd)
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
		err = userService.Upsert(*adminUser)
		if err != nil {
			log.Fatal(err)
		}

		userPermissions := []string{"get_flags", "get_flag", "get_me"}
		testUser, err := user.NewUser("test", "test", userPermissions)
		if err != nil {
			log.Fatal(err)
		}
		err = userService.Upsert(*testUser)
		if err != nil {
			log.Fatal(err)
		}
	},
}
