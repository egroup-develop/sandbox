package hello

import (
	"html/template"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	input, err := template.ParseFiles("input.html")
	if err == nil {
		contents := make(map[string]string, 1)
		input.Execute(w, contents)
	}
}

func sign(w http.ResponseWriter, r *http.Request) {
	err := signTemplate.Execute(w, r.FormValue("content"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var signTemplate = template.Must(template.New("sign").Parse(signTemplateHTML))

const signTemplateHTML = `
<html>
  <body>
      <p>You wrote:</p>
	      <pre>{{.}}</pre>
		    </body>
			</html>
			`
