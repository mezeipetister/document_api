package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Info controller
func Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello!\n")
}

// SetCookie ...
func SetCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Set Cookie
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "Name", Value: "Peti", Expires: expiration}
	http.SetCookie(w, &cookie)
	// Get Cookie
	fmt.Fprint(w, "Cookie Set\n")
}

// GetCookie ...
func GetCookie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, error := r.Cookie("Name")
	if error != nil {
		fmt.Fprint(w, "No cookie found")
	}
	// Get Cookie
	fmt.Fprint(w, cookie.Value)
}
