package user

import "fmt"

type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	File     []string `json:"filename"`
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
		user := User{Login: login, Password: password, File: []string{}}
		*allUsers = append(*allUsers, user)
		fmt.Println("Новый пользователь создан")
	}
}
