package helpers

import (
	"log"
	"strconv"
)

// ValidateInputArguments : To validate the number of arguments to the request and the type of data passed.
func ValidateInputArguments(arguments []string) (string, int32) {
	if len(arguments) > 3 || len(arguments) < 3 || len(arguments) != 3 {
		log.Fatal("Invalid numbers of parameters passed. Only 2 parameters are allowed (Chart URL and numbers of movies required).")
	}

	if arguments[1] == "" {
		log.Fatal("Invalid Chart URL passed. Invalid Argument : ", arguments[1])
	}

	items_count, err := strconv.Atoi(arguments[2])
	if err != nil {
		log.Fatal("Invalid Items Count passed.", err)
	}
	if items_count < 0 || items_count == 0 {
		log.Fatal("Invalid Items Count passed. Value of Items count must be greater than Zero.")
	}

	return arguments[1], int32(items_count)
}
