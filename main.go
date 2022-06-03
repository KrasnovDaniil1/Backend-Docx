package main

import (
	// "os/user"
	// "http"
	"app/user"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("We're live !!!"))
}

func main() {
	var allUsers = []user.User{} // все пользователи

	user.NewUser(&allUsers, "Daniil", "password")
	// user.NewUser(&allUsers, "Daniil", "password")
	// user.NewUser(&allUsers, "Danil", "password")

	fmt.Println(allUsers)

	http.HandleFunc("/", HomePageHandler)
	fmt.Printf("Starting application on port %v\n", portNumber)
	http.ListenAndServe(portNumber, nil)

}
