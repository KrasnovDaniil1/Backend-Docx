package main

import (
	"fmt"
	// "os/user"
	// "http"
	"backend-Docx/internal/user"
	// "bainternal/user/user.go"
)

func main() {

	allUsers := []user.User{}
	// func (user *user) AddNewFile(name ,text string) {
	// 	user := file = append(user1.file, "sdfs")
	// }

	user1 := user.User{login: "Daniil", password: "qwer", file: []string{}}
	user1.file = append(user1.file, "sdfs")

	allUsers = append(allUsers, user1)

	// fmt.Println(user1)
	fmt.Println(allUsers)

}
