package gohelpers

import (
	"fmt"
	"os"
	"strconv"
)

func setFromStringEnv(variable *string, env string) {
	result, exists := os.LookupEnv(env)
	if exists {
		*variable = result
	}
}

func setFromBoolEnv(variable *bool, env string) {
	result, exists := os.LookupEnv(env)
	if exists {
		if result == "true" {
			*variable = true
			return
		}
		if result == "false" {
			*variable = false
			return
		}
	}
}

func setFromIntEnv(variable *int, env string) {
	result, exists := os.LookupEnv(env)
	if exists {
		result, err := strconv.Atoi(result)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert %v to integer for variable %v", result, env))
		}

		*variable = result
	}
}
