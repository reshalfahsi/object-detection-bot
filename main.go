package main

import (
	"log"
	"net/http"     
)

func main() {
	port := getenv("PORT", "8080")

	http.HandleFunc("/login", Login)
	// http.HandleFunc("/predict", Predict)
	http.HandleFunc("/refresh", Refresh)
	// http.HandleFunc("/logout", Logout)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
