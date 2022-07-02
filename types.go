package main

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int64  `json:"age"`
}

type Arguments map[string]string
