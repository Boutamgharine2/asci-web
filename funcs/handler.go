package ascii

import (
	"html/template"
	"net/http"
)

func Asciiart(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmplt/tmplt.html"))
	r.ParseForm()
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusMethodNotAllowed)
		return
	}
	banner := r.Form.Get("select")
	input := r.Form.Get("user_input")
	if banner == "" || input == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	v := Printing(input, banner)

	tmpl.Execute(w, v)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmplt/tmplt.html"))
	r.ParseForm()
	if r.URL.Path != "/" {
		http.Error(w, "page not found", 404)
		return
	}
	// 	if r.Form.Get("submit") {
	// Asciiart(w,r)
	// }

	tmpl.Execute(w, nil)
}
