package main

import (
	"net/http"
	"opc-site/app/controller"
	"os"
)

const publicDir = "app/public/"

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, publicDir+"index.html")
}

func main() {
	appIp := os.Getenv("APP_IP")
	appPort := os.Getenv("APP_PORT")

	sqlHandler := controller.NewSqlHandler()
	defer sqlHandler.Db.Close()
	userController := controller.NewUserController(*sqlHandler)

	http.HandleFunc("/login", userController.HandleLogin)
	http.HandleFunc("/logout", userController.HandleLogout)
	http.HandleFunc("/registration", userController.HandleRegistration)

	apiController := controller.NewApiController(*sqlHandler)
	http.HandleFunc("/api/", apiController.HandleApi)

	http.HandleFunc("/index", handleIndex)

	err := http.ListenAndServe(appIp+":"+appPort, nil)
	if err != nil {
		return
	}
}
