package main

import (
	"fmt"
	"html/template"
	"net/http"
)
const port = ":8080"

type Todo struct {
    Title string
    Done  bool
}



type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}


func main(){
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        
        if err:=tmpl.Execute(w,data);err!=nil{
            fmt.Println(err)
        }
	})
	fullAddr := fmt.Sprintf("http://localhost%s", port)
	fmt.Println("server in running http", fullAddr)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}