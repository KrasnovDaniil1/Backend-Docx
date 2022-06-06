package user

var AllUsers = []User{} // все пользователи
type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	File     []string `json:"filename"`
}

// func (user *User) AddNewFile(file0 *file) {
// 	AllUsers = append(AllUsers, "sdfs")
// }

/*проверяет на идентичность логина*/

func checkUser(allUsers []User, login string) (bool, string) {
	var err string
	for _, v := range allUsers {
		if v.Login == login {
			err = "Такой логин уже есть"
			return false, err
		}
	}
	return true, err
}

/*создаёт нового пользователя*/

func NewUser(allUsers *[]User, login, password string) (string, string) {
	status, err := checkUser(*allUsers, login)
	var message string
	if status {
		user := User{Login: login, Password: password, File: []string{}}
		*allUsers = append(*allUsers, user)
		message = "Новый пользователь создан"
	}
	return message, err
}

/*получение пользователя*/

func GetUser(allUsers *[]User, login, password string) (string, string) {
	var message, err string
	for _, v := range *allUsers {
		if v.Login == login {
			if v.Password == password {
				message = "Авторизация прошла успешна"
				return message, err
			}
			err = "Неверный пароль"
			return message, err
		}
	}
	err = "Неверный логин или пароль"
	return message, err

}
