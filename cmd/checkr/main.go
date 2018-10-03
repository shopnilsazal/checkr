package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/shopnilsazal/checkr/spinner"
)

var urls = map[string]string{
	"npm":    "https://registry.npmjs.org/",
	"pip":    "https://pypi.org/pypi/",
	"crates": "https://crates.io/api/v1/crates/",
	"gem":    "https://rubygems.org/api/v1/gems/",
	"hex":    "https://hex.pm/api/packages/",
}

func main() {

	name := flag.String("n", "react", "Provide a package name you want to search.")
	managers := flag.String("m", "npm", "Provide package managers where you want to search.")
	flag.Parse()

	spin := spinner.New("%s Loading...")

	ch := make(chan PackageResponse)
	_, errorSym := getSymbols()

	allManagers := strings.Split(*managers, ",")
	if len(allManagers) == 1 {
		manager := allManagers[0]
		if manager == "all" {
			spin.Start()
			defer spin.Stop()

			for key, url := range urls {
				fullURL := getFullURL(url, key, *name)
				go getPackageStatus(key, fullURL, ch)
			}

			for range urls {
				res := <-ch
				spin.Stop()
				fmt.Printf(getMessage(res, *name))
				spin.Start()
			}

		} else if u, ok := urls[manager]; ok {
			spin.Start()
			fullURL := getFullURL(u, manager, *name)
			resp, err := http.Get(fullURL)
			if err != nil {
				fmt.Printf("%s oops! something bad happened in '%s' \n", errorSym, manager)
			}
			spin.Stop()
			defer resp.Body.Close()
			fmt.Printf(getMessage(PackageResponse{manager, resp.StatusCode}, *name))

		} else {
			fmt.Printf("%s '%s' is not a valid package manager.\n", errorSym, manager)
		}

	} else {
		var validManagers []string
		spin.Start()
		defer spin.Stop()

		for _, m := range allManagers {
			if u, ok := urls[m]; ok {
				fullURL := getFullURL(u, m, *name)
				go getPackageStatus(m, fullURL, ch)
				validManagers = append(validManagers, m)
			} else {
				fmt.Printf("%s '%s' is not a valid package manager.\n", errorSym, m)
			}
		}

		for range validManagers {
			res := <-ch
			spin.Stop()
			fmt.Printf(getMessage(res, *name))
			spin.Start()
		}
	}
}
