package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gregor-pifko/learn-go-with-tests/concurrency"
)

func main() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	var wc concurrency.WebsiteChecker = func(url string) bool {
		resp, err := http.Get(url)
		if err != nil {
			return false
		}

		defer resp.Body.Close()
		_, err = io.ReadAll(resp.Body)
		fmt.Println("URL Request", resp.Request.URL, resp.Status)
		return resp.StatusCode == 200
	}

	check := concurrency.CheckWebsites(wc, websites)

	fmt.Println(check)
}
