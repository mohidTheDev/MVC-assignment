package main

import (
	"coclone/controllers"
	"coclone/database"
	"fmt"
	"net/http"
)

func main() {
	database.Connect()
	fmt.Println("Clash of Clans clone test")
	http.HandleFunc("/", controllers.HandleStartBattle)
	http.ListenAndServe(":8080", nil)
}
