package main
import (
	"fmt"
	"net/http"
	openapi "github.com/VadimSmD/KR/internal/openapi"
	
)

func main() {
	service := $UserRepo{
		users:map[int64]openapi.User{},
	}
	srv, err := openapi.NewServer(service)
	if err != nil {
		log.Fatal(err)
	}
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal(err)
	}
}
