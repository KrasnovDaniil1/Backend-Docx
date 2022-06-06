package main

import (
	// "os/user"
	// "http"
	// "app/user"
    "app/handler"
	// "fmt"

	"net/http"
	// "os"
)

// const portNumber = ":8000"

func main() {

	// text := "Hello Gold!"
    // file, _ := os.Create("hello.txt")
     
    // if err != nil{
    //     fmt.Println("Unable to create file:", err) 
    //     os.Exit(1) 
    // }
    // defer file.Close() 
    // file.WriteString(text)
	http.Handle("/user/", http.HandlerFunc(handler.UserGetPost))
	http.Handle("/user/filename", http.HandlerFunc(handler.UserFilenameGetPostPutDelete))



	http.ListenAndServe(":8081", nil)

}


