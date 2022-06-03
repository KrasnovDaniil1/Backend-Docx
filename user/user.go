package user

import "fmt"

type User struct {
	ID              int
	Login, Password string
	File            []string
}

// type file struct {
// 	name,text string
// }

// func (user *User) AddNewFile(file0 *file) {
// 	AllUsers = append(AllUsers, "sdfs")
// }

/*проверяет на идентичность логина*/

func checkUser(allUsers []User, login string) bool {
	for _, v := range allUsers {
		if v.Login == login {
			fmt.Println("Такой логин уже есть")
			return false
		}
	}
	fmt.Println("Уникальный логин")
	return true
}

/*создаёт нового пользователя*/

func NewUser(allUsers *[]User, login, password string) {
	if checkUser(*allUsers, login) {
		user := User{ID: len(*allUsers), Login: login, Password: password, File: []string{}}
		*allUsers = append(*allUsers, user)
		fmt.Println("Новый пользователь создан")
	}
}
