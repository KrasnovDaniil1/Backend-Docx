package main

import (
	// "os/user"
	// "http"
	"app/user"
    "app/handler"
	"fmt"

	"net/http"
)

// const portNumber = ":8000"

func main() {
	var allUsers = []user.User{} // все пользователи

	user.NewUser(&allUsers, "Daniil", "password")
	fmt.Println(allUsers)

	http.Handle("/user/", http.HandlerFunc(handler.UserGetPost))
	http.Handle("/user/filename", http.HandlerFunc(handler.UserFilenameGetPostPutDelete))

	http.ListenAndServe(":8081", nil)

}


