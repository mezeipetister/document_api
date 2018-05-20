package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Info controller
func Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello!\n")
}
