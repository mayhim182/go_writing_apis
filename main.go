//main.go
package main

import (
	"fmt",
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	fmt.Println("Server is listening on: 8080...")
	http.ListenAndServe(":8000",nil)
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Println(w,"Hello, Go web")
}

