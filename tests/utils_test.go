package redmine_test

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func loadEnvsFromDotenv(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Log("Dotenv file not found")
	} else {
		t.Log("Dotenv file found; environment variables loaded")
	}
}

func getRedmineHostApikFromEnvs() (stat bool, host string, apik string) {
	host = os.Getenv("REDMINE_HOST")
	apik = os.Getenv("REDMINE_API_KEY")
	stat = len(host) > 0 && len(apik) > 0
	return
}

func doesRedmineHostExist(redmineAddress string) bool {
	// Perform a GET request to the host address
	resp, err := http.Get(redmineAddress)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	// Convert the response body to a string
	html := string(body)

	// Look for the Redmine meta tag in the HTML
	return strings.Contains(html, `<meta name="description" content="Redmine" />`)
}
