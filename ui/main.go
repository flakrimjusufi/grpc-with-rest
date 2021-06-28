package main


import (
	"html/template"
	"net/http"
)



type Todo struct {
	Title string
	Done  bool
}



type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}



func main() {
	tmpl := template.Must(template.ParseFiles("todo.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: true},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: false},
			},
		}
		err := tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}