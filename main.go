package main

import (
	"fmt"

	"github.com/golovers/testing-mongodb-with-gomock/pkg/users"
)

func main() {
	repo, err := users.NewMongoDB()
	if err != nil {
		panic(err)
	}
	srv := users.New()
	srv.SetRepo(repo)

	srv.AddUser(users.User{Name: "jack", Age: 22})
	srv.AddUser(users.User{Name: "sparrow", Age: 22})

	users, err := srv.GetUsers()
	if err != nil {
		panic(err)
	}
	for _, u := range users {
		fmt.Printf("Name: %s, Age: %d", u.Name, u.Age)
	}
}
