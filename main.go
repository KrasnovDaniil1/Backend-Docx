package main

import (
	// "os/user"
	// "http"
	"app/user"
	"fmt"
)

func main() {

	allUsers := []user.User{}
	
	user1 := user.User{Login: "Daniil", Password: "qwer", File: []string{}}

	// user1.File = append(user1.file, "sdfs")

	allUsers = append(allUsers, user1)
	allUsers = append(allUsers, user1)


	fmt.Println(allUsers)

}
