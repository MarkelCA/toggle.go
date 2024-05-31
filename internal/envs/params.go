package envs

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type AppMode int

const (
	ApiMode AppMode = iota
	CliMode
	WSMode
)

type EnvNames struct {
	Mode      AppMode
	AppPort   string
	RedisHost string
	RedisPort string
	MongoHost string
	MongoPort string
}

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

func GetConnectionParams(en EnvNames) (*ConnectionParams, []error) {
	var errors []error
	appPort := os.Getenv(en.AppPort)
	if en.Mode != CliMode && !validatePort(appPort) {
		errors = append(errors, InvalidPortError{port: appPort})
	}
	redisHost := os.Getenv(en.RedisHost)
	redisIPs, err := net.LookupIP(redisHost)
	if err != nil {
		errors = append(errors, err)
	} else {
		if !validateHost(redisIPs[0].String()) {
			errors = append(errors, InvalidHostError{host: redisHost})
		}
	}
	redisPortStr := os.Getenv(en.RedisPort)
	redisPort, err := strconv.Atoi(redisPortStr)
	if err != nil {
		errors = append(errors, err)
	}
	mongoHost := os.Getenv(en.MongoHost)
	mongoIPs, err := net.LookupIP(mongoHost)
	if err != nil {
		errors = append(errors, err)
	} else {
		if !validateHost(mongoIPs[0].String()) {
			errors = append(errors, InvalidHostError{host: redisHost})
		}
	}
	mongoPortStr := os.Getenv(en.MongoPort)
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
