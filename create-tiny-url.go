package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UrlObject struct {
	Url string
}

type UrlStoredObject struct {
	Url  string
	Tiny string
}

type UrlStoredObjectSet []UrlStoredObject

var filename string = "./url_mappings.json"

func generateTinyUrl(w http.ResponseWriter, r *http.Request) {
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
	urlStoredObject := ReadFromFIle(filename)

	// check if current url exists
	for _, urlObj := range urlStoredObject {
		if urlObj.Url == url {
			return urlObj.Tiny
		}
	}

	// create new mapping
	// store in file and return
	urlHashBytes := md5.Sum([]byte(url))
	urlHashString := hex.EncodeToString(urlHashBytes[:3])

	modData := urlStoredObject.toString() + `{"url":"` + url + `","tiny":"https://t.url/` + urlHashString + `"}]`
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

func getNormalUrlFromFIle(tiny string) string {
	urlStoredObject := ReadFromFIle(filename)

	for _, url := range urlStoredObject {
		if url.Tiny == tiny {
			return url.Url
		}
	}

	return ""
}
