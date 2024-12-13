package routes

import (
	"net/http"
	"user_management/controllers"
)

func SetupUserRoutes() {
	http.HandleFunc("/users/create", controllers.CreateUser)
	http.HandleFunc("/users/get", controllers.GetUser)
	http.HandleFunc("/users/update", controllers.UpdateUser)
	http.HandleFunc("/users/delete", controllers.DeleteUser)
	http.HandleFunc("/users/list", controllers.ListUsers)
}
