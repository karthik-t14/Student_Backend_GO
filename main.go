package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"student_crudapp/routers"
)

func main() {
	r := routers.Router()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	fmt.Println("Server at 8090")
	http.Handle("/", r)
	handler := cors.Handler(r)
	err := http.ListenAndServe(":8090", handler)
	if err != nil {
		log.Fatal("There's an error with the server", err)
	}
	fmt.Println("Server at 8090")
}
