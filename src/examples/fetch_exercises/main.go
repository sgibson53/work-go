package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)

		fmt.Print("Status: " + resp.Status)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatal(err)
		}

	}
}
