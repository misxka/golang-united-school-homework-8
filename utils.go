package main

import (
	"errors"
	"flag"
)

func parseArgs() Arguments {
	fileName := flag.String(FileName, "./users.json", "Name of the file containing users.")
	operation := flag.String(Operation, "add", "Operation to be performed: 'add', 'list', 'findById' or 'remove'.")
	item := flag.String(Item, `{"id": "1", "email": "email@test.com", "age": 23}`, "User to be added in JSON format.")
	id := flag.String(Id, "1", "Id of user to be found.")

	flag.Parse()

	arguments := make(map[string]string)

	arguments[FileName] = *fileName
	arguments[Operation] = *operation
	arguments[Item] = *item
	arguments[Id] = *id

	return arguments
}

func isOperationValid(operation string) bool {
	matched := 0
	possibleValues := [4]string{"add", "list", "findById", "remove"}
	for _, v := range possibleValues {
		if v == operation {
			matched++
			break
		}
	}
	return matched > 0
}

func argumentMissingError(argName string) error {
	return errors.New("-" + argName + " flag has to be specified")
}
