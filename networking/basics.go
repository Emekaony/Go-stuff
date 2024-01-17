package networking

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func Run() {
	// define the timeout of the request
	var timeout time.Duration = 1000 * time.Millisecond

	// create the http client
	client := &http.Client{
		Timeout: timeout,
	}

	// define the URL for the request
	var URL string = "https://jsonplaceholder.typicode.com/posts/4"

	// create the GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		log.Panic(err)
	}

	// send the request to the server
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	body, _ := io.ReadAll(res.Body)

	// print out the response body
	fmt.Println("Response body: ", string(body))
}
