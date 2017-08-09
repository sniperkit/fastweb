package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
	"log"
	"time"
)

var (
	url     string        = "http://httpbin.org/"
	timeOut time.Duration = 5 * time.Second
)

func main() {
	fastClientGet(url)
	fastClientPost("http://httpbin.org/post")
}
func fastClientGet(url string) error {

	client := &fasthttp.Client{}
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(request)
	defer fasthttp.ReleaseResponse(response)

	request.SetConnectionClose()
	request.SetRequestURI(url)

	if err := client.DoTimeout(request, response, timeOut); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(response.Header.Header()))
		fmt.Println(string(response.Body()))
	}
	return nil
}

func fastClientPost(url string) error {
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.Add("User-Agent", "Test-Agent")

	println(req.Header.String())
	// GET http://127.0.0.1:61765 HTTP/1.1
	// User-Agent: fasthttp
	// User-Agent: Test-Agent

	req.Header.SetMethod("POST")
	req.SetBodyString("p=q")

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	if err := client.DoTimeout(req, resp, timeOut); err != nil {
		println("Error:", err.Error())
	} else {
		bodyBytes := resp.Body()
		println(string(bodyBytes))
	}
	return nil
}

func hostClient() {
	// Perpare a client, which fetches webpages via HTTP proxy listening
	// on the localhost:8080.
	c := &fasthttp.HostClient{
		Addr: "localhost:8080",
	}

	// Fetch google page via local proxy.
	statusCode, body, err := c.Get(nil, "http://google.com/foo/bar")
	if err != nil {
		log.Fatalf("Error when loading google page through local proxy: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	useResponseBody(body)

	// Fetch foobar page via local proxy. Reuse body buffer.
	statusCode, body, err = c.Get(body, "http://foobar.com/google/com")
	if err != nil {
		log.Fatalf("Error when loading foobar page through local proxy: %s", err)
	}
	if statusCode != fasthttp.StatusOK {
		log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
	}
	useResponseBody(body)
}

func useResponseBody(body []byte) {
	// Do something with body :)
	println(string(body))
}

// design and code by tsingson
