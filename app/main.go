package main

import (
	"fmt"
	"net/http"
	"opc-site/app/controller"
	"strings"
)

func handleApi(w http.ResponseWriter, r *http.Request) {
	/**
	TODO вынести API Handler в свое пространство
	*/
	/*sessionController := &controller.SessionController{}
	if !sessionController.CheckSession(&w, r) {
		http.Error(w, "401", http.StatusUnauthorized)
	}*/

	path := r.URL.Path
	trimmedPath := path[len("/api/"):]
	trimmedPath = strings.TrimSuffix(trimmedPath, "/")

	switch trimmedPath {
	case "test":
		// some method...
	}

	fmt.Println("Запрос к:", trimmedPath)
}

func main() {
	sqlHandler := controller.NewSqlHandler()
	defer sqlHandler.Db.Close()
	userController := controller.NewUserController(*sqlHandler)

	http.HandleFunc("/login", userController.HandleLogin)
	http.HandleFunc("/logout", userController.HandleLogout)
	http.HandleFunc("/api/", handleApi)

	err := http.ListenAndServe(":25565", nil)
	if err != nil {
		return
	}
}
