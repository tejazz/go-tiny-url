package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type StoredUrlMappings map[string]string

func (u StoredUrlMappings) String() string {
	raw := "{"

	for url, tiny := range u {
		raw = raw + `"` + url + `":"` + tiny + `",`
	}

	return raw
}

func ReadFromFIle(filename string) StoredUrlMappings {
	hash, readErr := ioutil.ReadFile(filename)
	var storedUrlMappings StoredUrlMappings

	if readErr != nil {
		fmt.Println("[Func: ReadFromFIle]", readErr)
	}

	err := json.Unmarshal(hash, &storedUrlMappings)

	if err != nil {
		fmt.Println("[Func: ReadFromFIle]", err)
	}

	return storedUrlMappings
}
