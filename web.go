package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func isDarwin(userAgent string) bool {
	return strings.Contains(userAgent, "mac os x") || strings.Contains(userAgent, "darwin")
}

func isWindows(userAgent string) bool {
	return strings.Contains(userAgent, "windows")
}

func guessOS(userAgent string) string {
	if isDarwin(userAgent) {
		return "darwin"
	}

	if isWindows(userAgent) {
		return "windows"
	}

	return "linux"
}

func isAmd64(userAgent string) bool {
	return strings.Contains(userAgent, "x86_64") || strings.Contains(userAgent, "amd64")
}

func guessArch(userAgent string) string {
	if isAmd64(userAgent) || isDarwin(userAgent) {
		return "amd64"
	}
	return "386"
}

func guessPlatform(userAgent string) string {
	userAgent = strings.ToLower(userAgent)
	return guessOS(userAgent) + "-" + guessArch(userAgent)
}

func initial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, guessPlatform(r.UserAgent()))
}

func main() {

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "3000")
	}
	http.HandleFunc("/", initial)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
