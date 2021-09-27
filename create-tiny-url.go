package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func generateTinyUrl(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}
