package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/qm64/backpack/pkg"
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

func getPackFromCLIInput(cmd *cobra.Command, args []string) pkg.Pack {
	b := pkg.Pack{}

	readFromDir, err := cmd.Flags().GetBool("unpacked")
	if err != nil {
		log.Fatalf("Error parsing CLI flags: %s", err)
	}

	if !readFromDir {
		// get a file from URL or Path
		p := getAUsablePathOfFile(args[0])

		b, err = pkg.GetPackFromFile(p)
		if err != nil {
			log.Fatalf("Error parsing the pack: %s", err)
		}
	} else {
		// If we have to read from directory instead args[0] is a path
		d, err := pkg.GetPackFromDirectory(args[0])
		if err != nil {
			log.Fatalf("Error parsing the unpacked pack: %s", err)
		}
		b = *d
	}

	return b
}

func getValuesFromCLIInput(cmd *cobra.Command) pkg.ValuesType {
	vfPath, err := cmd.Flags().GetString("values")
	if err != nil {
		log.Fatalf("Error parsing CLI flags: %s", err)
	}

	values := pkg.ValuesType{}
	if vfPath != "" {
		values, err = pkg.ValuesFromFile(vfPath)
		if err != nil {
			log.Fatalf("Error reading the value file: %s", err)
		}
	}
	return values
}

func sanitizeUUIDPrefix(s string) string {
	return strings.Split(s, "-")[0]
}
