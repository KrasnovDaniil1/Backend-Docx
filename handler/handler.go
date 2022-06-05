package handler

import (
	"app/user"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func UserGetPost(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var userBody user.User

	bodyBytes, _ := io.ReadAll(r.Body)

	err := json.Unmarshal([]byte(bodyBytes), &userBody)

	fmt.Println(userBody, err)

	if r.Method == http.MethodGet {
		fmt.Println("UserGetPost")
	} else if r.Method == http.MethodPost {
		fmt.Println("UserGetPost", r.Method, http.MethodPost)
	} else {
		fmt.Println("UserGetPost", r.Method, http.MethodDelete)
	}

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
