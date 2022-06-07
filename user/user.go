package user

import (
	"fmt"
	"os"
)

var AllUsers = []User{} // все пользователи
type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	File     []string `json:"filename"`
}

/*создаст новый файл*/

func (user *User) AddNewFile(filename string, text string) (string, string) {
	var message, err string
	for _, v := range user.File {
		if v == filename {
			err = "Файл с таким именем уже есть"
			return message, err
		}
	}
	user.File = append(user.File, filename)

	// file, e := os.Create(user.Login + "/" + filename + ".txt")
	file, e := os.Create("userFolder/" + user.Login + "/" + filename + ".txt")

	if e != nil {
		fmt.Println("Unable to create file:", err)
		err = "Файл не был создан причина в железе"
	} else {
		file.WriteString(text)
		message = "Файл создан"
	}
	defer file.Close()

	return message, err
}

// func CreateFile(allUsers []User, login string, password string) {
// 	if FindUser(allUsers, login, password) >= 0 {

// 	}
// 	file, err := os.Create("Daniil/hello.txt")
// 	if err != nil {
// 		fmt.Println("Unable to create file:", err)
// 		os.Exit(1)
// 	}
// 	file.WriteString("Hello Gold!")
// 	defer file.Close()
// }
