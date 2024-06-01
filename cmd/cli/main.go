package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/markelca/toggles/internal/envs"
	"github.com/markelca/toggles/pkg/flags"
	"github.com/markelca/toggles/pkg/storage"
	"github.com/markelca/toggles/pkg/user"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tg",
	Short: "This is the togggles command line tool",
	Long:  `This tool offers utilities to interact with the togggles API.`,
}

var userRepo user.UserRepository
var flagService flags.FlagService
var params *envs.ConnectionParams

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var paramErr []error
	params, paramErr = envs.GetConnectionParams(envs.EnvNames{
		Mode:      envs.CliMode,
		RedisHost: "CLI_REDIS_HOST",
		RedisPort: "CLI_REDIS_PORT",
		MongoHost: "CLI_MONGO_HOST",
		MongoPort: "CLI_MONGO_PORT",
	})

	if len(paramErr) > 0 {
		envs.PrintFatalErrors(paramErr)
	}

	userRepo, err = user.NewUserMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}

	db, err := flags.NewFlagMongoRepository(params.MongoHost, params.MongoPort)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to MongoDB: %v", err))
	}

	repository := storage.NewRedisClient(params.RedisHost, params.RedisPort)
	flagService = flags.NewFlagService(repository, db)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
