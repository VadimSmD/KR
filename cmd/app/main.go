package main

import (
	"fmt"
	"net/http"

	controller "github.com/VadimSmD/KR/internal/user_controller.go"
)

func main() {
	userController := &UserController{}
	http.HandleFunc("/users", userController)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
