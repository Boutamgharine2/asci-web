package ascii

import (
	"fmt"
	"html/template"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("tmplt/tmplt.html"))
	r.ParseForm()

	banner := r.Form.Get("select")
	input := r.Form.Get("user_input")
	fmt.Println(input)
	v := Printing(input, banner)

	tmpl.Execute(w, v)
}
