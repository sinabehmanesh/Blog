package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/gorilla/mux"
)

//ORM functions
// func insertmessage(message string)  {
// 	dsn := DBUSER":pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// }

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

	http.ListenAndServe(":3000", r)
}
