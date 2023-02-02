package routers

import (
	"fmt"
	"github.com/gorilla/mux"
	"student_crudapp/middleware"
)

func Router() *mux.Router {
	fmt.Println("Server at 8090")
	router := mux.NewRouter()
	fmt.Println(router)
	router.HandleFunc("/api/student/{id}", middleware.GetStudent).Methods("GET", "OPTIONS") //.Schemes("http") //.Methods("GET", "OPTIONS")
	router.HandleFunc("/api/student", middleware.GetAllStudents).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/student", middleware.CreateStudent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/student/{id}", middleware.UpdateStudent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/student/{id}", middleware.DeleteStudent).Methods("DELETE", "OPTIONS")

	return router
}
