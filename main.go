package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.redit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls {
		hitURL(url)
	}
}

var errRequestFailed = errors.New("Error, Request Failed")

func hitURL(url string) error {
	fmt.Println("Checking :", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
