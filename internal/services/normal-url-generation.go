package services

import (
	"create-tiny-url/internal/models"
	"create-tiny-url/internal/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateNormalUrl(w http.ResponseWriter, r *http.Request) {
	var urlObject models.RequestJson
	reqBody, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(reqBody, &urlObject)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	normalUrl := getNormalUrlFromFIle(urlObject.Url)

	if normalUrl == "" {
		http.Error(w, "The particular URL does not exist", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, normalUrl)

	defer r.Body.Close()
}

func getNormalUrlFromFIle(tinyUrl string) string {
	storedUrlMappings := utils.ReadFromFIle(filename)

	splitUrl := strings.Split(tinyUrl, "/")
	tinyHash := splitUrl[len(splitUrl) - 1]

	for url, tiny := range storedUrlMappings {
		if tinyHash == tiny {
			return url
		}
	}

	return ""
}

