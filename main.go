package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Sentence struct {
	Text string
}

func main() {

	sentence := Sentence{"Imersão Full Cycle"}

	templates := template.Must(template.ParseFiles("templates/home.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := templates.ExecuteTemplate(w, "home.html", sentence); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Serving on PORT 8000")
	fmt.Println(http.ListenAndServe(":8000", nil))

}
