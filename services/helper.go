package services

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func CheckDomain(url string) bool {
	if url == os.Getenv("DOMAIN") {
		return false
	}

	tempURL := strings.Replace(url, "http://", "", 1)
	tempURL = strings.Replace(tempURL, "https://", "", 1)
	tempURL = strings.Replace(tempURL, "www.", "", 1)
	tempURL = strings.Split(tempURL, "/")[0]

	return tempURL != os.Getenv("DOMAIN")
}
