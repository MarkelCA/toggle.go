package main

import (
	"fmt"

	"github.com/markelca/toggles/internal/envs"
	"github.com/markelca/toggles/pkg/user"
	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use: "database",
	// Short:   "",
	Aliases: []string{"db"},
	Args:    cobra.ExactArgs(1),
}
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
	params, paramErr := envs.GetConnectionParams()
	if len(paramErr) > 0 {
		envs.PrintFatalErrors(paramErr)
	}

	userRepo, err := user.NewUserMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}
	userRepo.Upsert(user.User{
		UserName: "admin",
		Role:     "admin",
		Password: "admin",
	})
	userRepo.Upsert(user.User{
		UserName: "test",
		Role:     "user",
		Password: "test",
	})

}
