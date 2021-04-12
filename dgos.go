package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fDebug := flag.String("mode", "debug", "toggle debug mode")
	fWorkers := flag.Int("workers", 4, "amount of workers per URL")
	fListPath := flag.String("list", "", "path to list with urls")
	flag.Parse()

	// read urls from list file
	urls, err := readList(*fListPath)
	if err != nil {
		panic(err)
	}

	// thread with job server
	go jobServer(urls, *fWorkers, *fDebug)

	select {} // prevent program exit
}

// jobServer serve job for workers / spawn workers.
func jobServer(urls []string, workers int, mode string) {
	for _, url := range urls {
		for i := 0; i < workers; i++ {
			go requestURL(url, mode)
		}
	}
}

// requestURL requests from server the provided url.
func requestURL(url string, mode string) {
	for {
		response, err := http.Get(url)
		if err != nil {
			if mode != "silent" {
				fmt.Printf("error: %s", err.Error())
			}
		}

		if response != nil {
			defer response.Body.Close()
		}

		if mode == "debug" {
			fmt.Printf("%v", response)
		}
	}

}

// readList load, return slice with urls.
func readList(listPath string) ([]string, error) {
	if listPath == "" {
		return nil, errors.New("the list is empty")
	}

	rawData, err := ioutil.ReadFile(listPath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(rawData), "\n"), nil
}
