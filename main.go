package main

import (
	contactapi "blog/database"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

/////////////////
// handlers defined on routes
/////////////////

func homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)

}

func sysloghandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/syslog.html"))
	tmpl.Execute(w, nil)
}

func contacthandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/contact.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	contactinfo := contact{
		email:   r.FormValue("email"),
		subject: r.FormValue("subject"),
		message: r.FormValue("message"),
	}
	tmpl.Execute(w, struct{ Success bool }{true})

	_ = contactinfo.message

	contactapi.Insert_message()
}

func adminhandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/admin.html"))
	tmpl.Execute(w, nil)
}

func thiswebsitehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/thiswebsite.html"))
	tmpl.Execute(w, nil)
}

// ///////////////
// Types
// ///////////////
type contact struct {
	email   string
	subject string
	message string
}

func main() {

	//fileserver
	fs := http.FileServer(http.Dir("html/"))
	http.Handle("/html/", http.StripPrefix("/html/", fs))

	//Routes
	r := mux.NewRouter()
	r.HandleFunc("/", homehandler)

	r.HandleFunc("/admin", adminhandler)
	r.HandleFunc("/syslog", sysloghandler)
	r.HandleFunc("/contact", contacthandler)
	r.HandleFunc("/thiswebsite", thiswebsitehandler)
	http.Handle("/", r)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can not load .env file!! Err: %s", err)
	}

	port := os.Getenv("API_PORT")
	http.ListenAndServe(":"+port, r)
}
