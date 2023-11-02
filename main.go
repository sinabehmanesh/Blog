package main

import (
	adminapi "blog/admin"
	contactapi "blog/database"

	"fmt"
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

	contactemail := r.FormValue("email")
	contactsubject := r.FormValue("subject")
	contactmessage := r.FormValue("message")

	tmpl.Execute(w, struct{ Success bool }{true})

	//call contact database function to submit message to database
	err := contactapi.Insert_message(contactemail, contactsubject, contactmessage)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func adminhandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/admin.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{true})

	adminapi.Login(w, r)
}

func thiswebsitehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/thiswebsite.html"))
	tmpl.Execute(w, nil)
}

func aboutmehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./html/aboutme.html"))
	tmpl.Execute(w, nil)
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
	r.HandleFunc("/aboutme", aboutmehandler)
	http.Handle("/", r)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Can not load .env file!! Err: %s", err)
	}

	port := os.Getenv("API_PORT")

	fmt.Println("starting server on port:", port)
	http.ListenAndServe(":"+port, r)
}
