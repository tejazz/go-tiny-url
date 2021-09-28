package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (u UrlStoredObjectSet) toString() string {
	raw := "["

	for _, url := range u {
		raw = raw + `{"url":"` + url.Url + `","tiny":"` + url.Tiny + `"},`
	}

	return raw
}

func ReadFromFIle(filename string) UrlStoredObjectSet {
	hash, _ := ioutil.ReadFile(filename)
	var urlStoredObject UrlStoredObjectSet

	err := json.Unmarshal(hash, &urlStoredObject)

	if err != nil {
		fmt.Println(err)
	}

	return urlStoredObject
}
