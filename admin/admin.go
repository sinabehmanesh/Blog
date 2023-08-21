package admin

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "authentication-status")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(w, "james lied about us!")
}

func Login(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, "authentication-status")

	adminemail := r.FormValue("adminemail")
	adminpassword := r.FormValue("adminpassword")

	if adminemail == "sina@sina.com" && adminpassword == "sinasinasina" {
		session.Values["authenticated"] = true
		err := session.Save(r, w)
		if err != nil {
			return "validation failed!"
		} else {
			return "Login Successful!"
		}
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return "Wrong input, Forbidden!"
	}

}
