package utils

import (
	"create-tiny-url/internal/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadFromFIle(filename string) models.StoredUrlMappings {
	hash, readErr := ioutil.ReadFile(filename)
	var storedUrlMappings models.StoredUrlMappings

	if readErr != nil {
		fmt.Println("[Func: ReadFromFIle]", readErr)
	}

	err := json.Unmarshal(hash, &storedUrlMappings)

	if err != nil {
		fmt.Println("[Func: ReadFromFIle]", err)
	}

	return storedUrlMappings
}

