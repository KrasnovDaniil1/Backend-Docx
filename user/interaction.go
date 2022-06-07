package user

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*возвращает данные пользователя*/

func GetUser(allUsers []User, login, password string) (string, string) {
	var message, err string
	_, permission := PermissionUser(allUsers, login, password)
	if permission {
		userCreateFolder(login) // навсякий случай создаём папку если ранее не создали
		message = "Авторизация прошла успешна"
		return message, err
	}
	err = "Неверный логин или пароль"
	return message, err
}

/*создаёт нового пользователя*/

func NewUser(allUsers *[]User, login, password string) (string, string) {
	_, permission := PermissionUser(*allUsers, login, password)
	var message, err string
	if !permission {
		user := User{Login: login, Password: password, File: []string{}}
		*allUsers = append(*allUsers, user)
		userCreateFolder(login)
		message = "Новый пользователь создан"
		return message, err
	}
	err = "Такой пользователь уже есть"
	return message, err
}

func UserCreateNewFile(allUsers []User, login string, password string, filename string, text string) (string, string) {
	var message, err string
	key, permission := PermissionUser(allUsers, login, password)
	if permission {
		message, err = allUsers[key].AddNewFile(filename, text)
	} else {
		err = "Что то не так, не найден пользователь"
	}
	return message, err
}

/*создаёт папку*/

func userCreateFolder(login string) {
	folderCreate := false
	files, _ := ioutil.ReadDir("userFolder")
	for _, file := range files {
		if file.Name() == login {
			folderCreate = true
		}
	}
	if !folderCreate {
		os.Mkdir(("userFolder/" + login), 0755)
		fmt.Printf("Папка %v создана", login)
	}
}

/*проверяет есть ли такой пользователь*/

func PermissionUser(allUsers []User, login string, password string) (int, bool) {
	for i, v := range allUsers {
		if v.Login == login && v.Password == password {
			return i, true
		}
	}
	return -1, false
}
