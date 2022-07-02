package main

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

func Perform(args Arguments, writer io.Writer) error {
	operation := args[Operation]
	if operation == "" {
		return argumentMissingError(Operation)
	}
	if !isOperationValid(operation) {
		return errors.New("Operation " + operation + " not allowed!")
	}

	fileName := args[FileName]
	if fileName == "" {
		return argumentMissingError(FileName)
	}

	f, openError := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if openError != nil {
		panic(openError)
	}

	data, readError := ioutil.ReadAll(f)
	if readError != nil {
		panic(readError)
	}
	defer f.Close()

	switch operation {
	case "list":
		_, err := writer.Write(data)
		return err
	case "add":
		item := args[Item]
		if item == "" {
			return argumentMissingError(Item)
		}

		var users []User
		var user User

		if len(data) > 0 {
			err := json.Unmarshal(data, &users)
			if err != nil {
				panic(err)
			}
		}

		json.Unmarshal([]byte(item), &user)

		for _, v := range users {
			if v.Id == user.Id {
				writer.Write([]byte("Item with id " + v.Id + " already exists"))
			}
		}

		users = append(users, user)
		usersString, _ := json.Marshal(users)
		ioutil.WriteFile(fileName, usersString, 0644)
	case "findById":
		id := args[Id]
		if id == "" {
			return argumentMissingError(Id)
		}

		var users []User
		if len(data) > 0 {
			err := json.Unmarshal(data, &users)
			if err != nil {
				panic(err)
			}
		}

		for _, v := range users {
			if v.Id == id {
				foundUser, _ := json.Marshal(v)
				writer.Write(foundUser)
			}
		}
	case "remove":
		id := args[Id]
		if id == "" {
			return argumentMissingError(Id)
		}

		var users []User
		if len(data) > 0 {
			err := json.Unmarshal(data, &users)
			if err != nil {
				panic(err)
			}
		}

		position := -1
		for i, v := range users {
			if v.Id == id {
				position = i
				break
			}
		}

		if position == -1 {
			writer.Write([]byte("Item with id " + id + " not found"))
			return nil
		}

		users = append(users[:position], users[position+1:]...)
		usersString, _ := json.Marshal(users)
		ioutil.WriteFile(fileName, usersString, 0644)
	}

	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
