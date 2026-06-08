package main

import (
	"coclone/controllers" // Import your custom package
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Clash of Clans clone test")
	http.HandleFunc("/", controllers.HandleStartBattle)
	http.ListenAndServe(":8080", nil)
}
