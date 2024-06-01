package main

import (
	"fmt"
	"log"

	"github.com/markelca/toggles/pkg/user"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initializes the database",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing database:")
		initDatabase()
	},
}

func initDatabase() {
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
}
