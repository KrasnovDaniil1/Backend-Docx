package handler

import (
	"app/user"
	"encoding/json"

	// "errors"
	"fmt"
	"io"
	"net/http"
)

type UserResp struct {
	Login    string   `json:"login"`
	Password string   `json:"password"`
	File     []string `json:"filename"`
	Status   string   `json:"status"`
	Error    string   `json:"error"`
}

func UserGetPost(w http.ResponseWriter, r *http.Request) {
	var userBody UserResp

	bodyBytes, _ := io.ReadAll(r.Body)

	json.Unmarshal([]byte(bodyBytes), &userBody)

	if r.Method == http.MethodGet {
		userBody.Status, userBody.Error = user.GetUser(&user.AllUsers, userBody.Login, userBody.Password)
	} else if r.Method == http.MethodPost {
		userBody.Status, userBody.Error = user.NewUser(&user.AllUsers, userBody.Login, userBody.Password)
	} else {
		fmt.Println("Немогу обработать запрос")
	}
	userWrite, _ := json.Marshal(userBody)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userWrite)

}

func UserFilenameGetPostPutDelete(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodGet {
		fmt.Println("UserFilenameGetPostPutDelete", r.Method, http.MethodGet)
	} else if r.Method == http.MethodPost {
		fmt.Println("UserFilenameGetPostPutDelete", r.Method, http.MethodPost)
	} else if r.Method == http.MethodPut {
		fmt.Println("UserFilenameGetPostPutDelete", r.Method, http.MethodPut)
	} else if r.Method == http.MethodDelete {
		fmt.Println("UserFilenameGetPostPutDelete", r.Method, http.MethodDelete)
	} else {
		fmt.Println("UserFilenameGetPostPutDelete", r.Method, http.MethodDelete)
	}
}
