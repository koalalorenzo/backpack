package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func downloadIntoTempFile(url string) (newPath string, err error) {
	file, err := ioutil.TempFile(os.TempDir(), "backpack-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	response, err := http.Get(url)
	if err != nil {
	}
	defer response.Body.Close()
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return file.Name(), nil
}

func whatIsInput(v string) string {
	if isValidUrl(v) {
		return "url"
	}

	if _, err := os.Stat(v); err == nil {
		return "file"
	}

	return "unknown"
}

// getAUsablePathOfFile will try to understand what is v, and if it is a URL
// it will try to obtain it in one way or another. if it is a file, it will
// leave the path as is. This will help to support multiple path formats in CLI
func getAUsablePathOfFile(v string) string {
	switch whatIsInput(v) {
	case "url":
		np, err := downloadIntoTempFile(v)
		if err != nil {
			log.Fatalf("Error downloading the url %s: %s", v, err)
		}
		return np
	case "file":
		return v
	default:
		return v
	}
}
