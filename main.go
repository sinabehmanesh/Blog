package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, "")

}

func sysloghandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/syslog.html"))
	tmpl.Execute(w, "")
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/contact.html"))
	tmpl.Execute(w, "")
}
func main() {

	//fileserver
	fs := http.FileServer(http.Dir("html/"))
	http.Handle("/html/", http.StripPrefix("/html/", fs))

	//Routes
	r := mux.NewRouter()
	r.HandleFunc("/", homehandler)

	// r.HandleFunc("/admin", adminhandler)
	r.HandleFunc("/syslog", sysloghandler)
	r.HandleFunc("/contact", contacthandler)
	// r.HandleFunc("/thiswebsite", thiswebsitehandler)
	http.Handle("/", r)

	http.ListenAndServe(":3000", r)
}

//dev-mux branch for mux development
