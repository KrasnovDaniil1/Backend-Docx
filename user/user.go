package user

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

var AllUsers = []User{} // все пользователи
type User struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	File     []string `json:"filename"`
}

/*создаст новый файл*/

func (user *User) AddNewFile(name string) {
	user.File = append(user.File, name)
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
