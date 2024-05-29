package envs

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type ConnectionParams struct {
	AppPort   string
	RedisHost string
	RedisPort uint
	MongoHost string
	MongoPort uint
}

type InvalidHostError struct {
	host string
}

func (e InvalidHostError) Error() string {
	return fmt.Sprintf("Invalid host: \"%v\"", e.host)
}

type InvalidPortError struct {
	port string
}

func (e InvalidPortError) Error() string {
	return fmt.Sprintf("Invalid port: \"%v\"", e.port)
}

func validateHost(host string) bool {
	return net.ParseIP(host) != nil
}

func validatePort(port string) bool {
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false
	}
	return portInt > 0 && portInt < 65536
}

func GetConnectionParams() (*ConnectionParams, []error) {
	var errors []error
	appPort := os.Getenv("APP_PORT")
	if !validatePort(appPort) {
		errors = append(errors, InvalidPortError{port: appPort})
	}
	redisHost := os.Getenv("REDIS_HOST")
	redisIPs, err := net.LookupIP(redisHost)
	if err != nil {
		errors = append(errors, err)
	} else {
		if !validateHost(redisIPs[0].String()) {
			errors = append(errors, InvalidHostError{host: redisHost})
		}
	}
	redisPortStr := os.Getenv("REDIS_PORT")
	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		errors = append(errors, err)
	}
	mongoHost := os.Getenv("MONGO_HOST")
	mongoIPs, err := net.LookupIP(mongoHost)
	if err != nil {
		errors = append(errors, err)
	} else {
		if !validateHost(mongoIPs[0].String()) {
			errors = append(errors, InvalidHostError{host: redisHost})
		}
	}
	mongoPortStr := os.Getenv("MONGO_PORT")
	mongoPort, err := strconv.Atoi(mongoPortStr)
	if err != nil {
		errors = append(errors, err)
	}
	connParams := &ConnectionParams{AppPort: appPort, RedisHost: redisHost, MongoHost: mongoHost, RedisPort: uint(redisPort), MongoPort: uint(mongoPort)}

	if len(errors) > 0 {
		return nil, errors
	} else {
		return connParams, nil
	}

}

func PrintFatalErrors(errors []error) {
	errMsg := "Param errors have been found:\n"
	for _, err := range errors {
		errMsg += fmt.Sprintf("  - %v\n", err.Error())
	}
	log.Fatal(errMsg)
}
