package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// WaitGroups

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Status)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page endpoint hit")
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("Returning Response")
	fmt.Fprintf(w, "Responses")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
