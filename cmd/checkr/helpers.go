package main

import (
	"fmt"
	"net/http"
	"runtime"
)

// PackageResponse Type
type PackageResponse struct {
	packageManager string
	statusCode     int
}

// Get status code from package manager
func getPackageStatus(name, url string, ch chan<- PackageResponse) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- PackageResponse{name, 500}
		return
	}
	ch <- PackageResponse{name, resp.StatusCode}
	resp.Body.Close()
}

// Get full url of the package manager based on name
func getFullURL(url, manager, name string) string {
	fullURL := url + name
	switch manager {
	case "pip":
		fullURL = fullURL + "/json"
	case "gem":
		fullURL = fullURL + ".json"
	}
	return fullURL
}

// Get success and error symbols
func getSymbols() (successSym, errorSym string) {
	if runtime.GOOS == "windows" {
		successSym = "\u001b[32m" + "√" + "\u001b[39m"
		errorSym = "\u001b[31m" + "×" + "\u001b[39m"
	} else {
		successSym = "\u001b[32m" + "✔" + "\u001b[39m"
		errorSym = "\u001b[31m" + "✖" + "\u001b[39m"
	}
	return
}

// Get the message to print on console
func getMessage(resp PackageResponse, name string) string {
	var msg string
	ss, es := getSymbols()
	switch resp.statusCode {
	case 404:
		msg = fmt.Sprintf("%s %s is available on '%s' \n", ss, name, resp.packageManager)
	case 200:
		msg = fmt.Sprintf("%s %s is unavailable on '%s' \n", es, name, resp.packageManager)
	default:
		msg = fmt.Sprintf("%s oops! something bad happened in '%s' \n", es, resp.packageManager)
	}
	return msg
}
