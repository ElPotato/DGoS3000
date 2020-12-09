package main

import (
	"net/http"
	"fmt"
	"flag"
	"io/ioutil"
	"strings"
	"errors"
)

func main() {
	fDebug := flag.String("mode", "debug", "toggle debug mode")
	fWorkers := flag.Int("workers", 4, "amount of workers per URL")
	fListPath := flag.String("list", "list.txt", "path to list with urls")
	flag.Parse()

	// read urls from list file
	urls, err := readList(*fListPath)
	if err != nil {
		panic(err)
	}

	// thread with job server
	go jobServer(urls, *fWorkers, *fDebug)

	select{} // prevent program exit
}

// serve job for workers / spawn workers
func jobServer(urls []string, workers int, mode string) {
	for _, url := range urls {
		for i := 0; i < workers; i++ {
			go requestURL(url, mode)
		}
	}
}

// request from server the provided file
func requestURL(url string, mode string) {
	for {
		response, err := http.Get(url)
		if err != nil {
			if mode != "silent" {
				fmt.Printf("error: %s", err.Error())
			}
		}

		if mode == "debug" {
			fmt.Printf("%v", response)
		}
	}
	
}

// load, return slice with urls
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