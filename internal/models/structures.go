package models

type StoredUrlMappings map[string]string

type RequestJson struct {
	Url string
}

func (u StoredUrlMappings) String() string {
	raw := "{"

	for url, tiny := range u {
		raw = raw + `"` + url + `":"` + tiny + `",`
	}

	return raw
}
