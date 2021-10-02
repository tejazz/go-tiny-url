package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type UrlObject struct {
	Url string
}

type UrlStoredObject struct {
	Url  string
	Tiny string
}

var filename string = "./url_mappings.json"

func getTinyUrl(w http.ResponseWriter, r *http.Request) {
	var urlObject UrlObject
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &urlObject)

	if err != nil {
		fmt.Println(err)
	}

	tinyUrl := createHashForUrl(urlObject.Url)

	fmt.Fprint(w, "https://t.url/"+tinyUrl)
}

func createHashForUrl(url string) string {
	storedUrlMappings := ReadFromFIle(filename)

	// check if current url exists
	if _, ok:= storedUrlMappings[url]; ok {
		return storedUrlMappings[url]
	}

	// create new mapping
	// store in file and return
	urlHashBytes := md5.Sum([]byte(url))
	urlHashString := hex.EncodeToString(urlHashBytes[:3])

	modData := storedUrlMappings.String() + `"` + url + `":"` + urlHashString + `"}`
	ioutil.WriteFile("./url_mappings.json", []byte(modData), 066)

	return urlHashString
}

func generateNormalUrl(w http.ResponseWriter, r *http.Request) {
	var urlObject UrlObject
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &urlObject)

	if err != nil {
		fmt.Println(err)
	}

	normalUrl := getNormalUrlFromFIle(urlObject.Url)

	if normalUrl == "" {
		fmt.Fprintf(w, "The particular URL does not exist")
	}

	fmt.Fprintf(w, normalUrl)
}

func getNormalUrlFromFIle(tinyUrl string) string {
	storedUrlMappings := ReadFromFIle(filename)

	tinyHash := strings.Split(tinyUrl, "/")[len(strings.Split(tinyUrl, "/")) - 1]

	for url, tiny := range storedUrlMappings {
		if tinyHash == tiny {
			return url
		}
	}

	return ""
}
