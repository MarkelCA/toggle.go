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
		// NOTE: WIP, not working yet (envs missing)
		initDatabase()
	},
}

func initDatabase() {
	adminUser, err := user.NewUser("admin", "admin", "admin")
	if err != nil {
		log.Fatal(err)
	}
	err = userRepo.Upsert(*adminUser)
	if err != nil {
		log.Fatal(err)
	}

	testUser, err := user.NewUser("test", "user", "test")
	if err != nil {
		log.Fatal(err)
	}
	err = userRepo.Upsert(*testUser)
	if err != nil {
		log.Fatal(err)
	}

}
