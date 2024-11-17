package main

import (
	controller "github.com/VadimSmD/KR/internal"
)

func main() {
	userController := &controller.UserController{}
	http.HandleFunc("/users", userController)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
