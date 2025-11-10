package main

import (
	"fmt"
	"net/http"
)
const port = ":8080"

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	fullAddr := fmt.Sprintf("http://localhost%s", port)
	fmt.Println("server in running http", fullAddr)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}