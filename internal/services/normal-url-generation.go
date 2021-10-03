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
		fmt.Println(err)
	}

	normalUrl := getNormalUrlFromFIle(urlObject.Url)

	if normalUrl == "" {
		fmt.Fprintf(w, "The particular URL does not exist")
	}

	fmt.Fprintf(w, normalUrl)
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

