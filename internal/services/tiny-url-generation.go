package services

import (
	"create-tiny-url/internal/models"
	"create-tiny-url/internal/utils"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// note: paths are resolved relative to root directory
var filename string = "./internal/url_mappings.json"

func GetTinyUrl(w http.ResponseWriter, r *http.Request) {
	var urlObject models.RequestJson
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &urlObject)

	if err != nil {
		fmt.Println(err)
	}

	tinyUrl := createHashForUrl(urlObject.Url)

	fmt.Fprint(w, "https://t.url/"+tinyUrl)
}

func createHashForUrl(url string) string {
	storedUrlMappings := utils.ReadFromFIle(filename)

	// check if current url exists
	if _, ok:= storedUrlMappings[url]; ok {
		return storedUrlMappings[url]
	}

	// create new mapping
	// store in file and return
	urlHashBytes := md5.Sum([]byte(url))
	urlHashString := hex.EncodeToString(urlHashBytes[:3])

	modData := storedUrlMappings.String() + `"` + url + `":"` + urlHashString + `"}`
	ioutil.WriteFile(filename, []byte(modData), 066)

	return urlHashString
}