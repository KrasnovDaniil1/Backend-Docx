package user

import "fmt"

type User struct {
	Login, Password string
	File            []string
}

func Col() {
	fmt.Println("run")
}

// func (user *user) AddNewFile(name ,text string) {
	// 	user := file = append(user1.file, "sdfs")
	// }
