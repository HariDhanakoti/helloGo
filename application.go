package main

import (
	"fmt"
	"net/http"
	"time"

	"html/template"
)

//Welcome basic struct for welcome page
type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("public/index.html"))

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "index.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
