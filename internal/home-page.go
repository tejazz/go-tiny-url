package internal

import (
	"fmt"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TinyURL Generator: Welcome")
}
