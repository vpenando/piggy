package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	totalErrs   int32
	hostBaseURL string

	host    = "http://localhost"
	port    = 8081
	count   = 100
	verbose = false
)

func init() {
	flag.StringVar(&host, "host", host, "Host name")
	flag.IntVar(&port, "port", port, "Host port")
	flag.IntVar(&count, "count", count, "Requests count")
	flag.BoolVar(&verbose, "verbose", verbose, "Verbose mode")
	hostBaseURL = fmt.Sprintf("%s:%d", host, port)
	flag.Parse()
	if verbose {
		log.Println("Host set to", host)
		log.Println("Port set to", port)
		log.Println("Count set to", count)
		log.Println("Verbose set to", verbose)
	}
}

func get(route string, status int) {
	url := fmt.Sprintf("%s%s", hostBaseURL, route)
	client := http.Client{Timeout: 10 * time.Second}
	if verbose {
		log.Println("GET", route, "with expected status code", status)
	}
	var errs int32
	start := time.Now()
	for i := 0; i < count; i++ {
		result, err := client.Get(url)
		if result != nil && result.StatusCode != status {
			if verbose {
				log.Printf("Error while sending request (status code %d, expected %d)\n", result.StatusCode, status)
			}
			errs++
		} else if err != nil {
			if verbose {
				log.Println("Error:", err)
			}
			errs++
		}
	}
	end := time.Now()
	elapsed := end.Sub(start)
	if verbose {
		log.Printf("Finished in %s: %d/%d errors\n\n", elapsed.Truncate(time.Millisecond), errs, count)
	}
	totalErrs += errs
}

type testCase struct {
	test   func(string, int)
	route  string
	status int
}

var testCases = []testCase{
	// Pages
	{get, "/", http.StatusOK},
	{get, "/edit?year=2020&month=12", http.StatusOK},
	{get, "/edit?year=2020&month=0", http.StatusUnprocessableEntity},
	{get, "/edit?year=2020&month=13", http.StatusUnprocessableEntity},
	{get, "/settings", http.StatusOK},

	// Resources
	{get, "/static/css/piggy.css", http.StatusOK},
	{get, "/static/images/favicon.ico", http.StatusOK},
	{get, "/static/scripts/piggy.js", http.StatusOK},

	// Misc
	{get, "/months", http.StatusOK},
	{get, "/reports?year=2020&month=12", http.StatusOK},
	{get, "/reports?year=2020&month=0", http.StatusUnprocessableEntity},
	{get, "/reports?year=2020&month=13", http.StatusUnprocessableEntity},

	// Operations
	{get, "/operations?year=2020&month=12", http.StatusOK},
	{get, "/operations?year=2020&month=0", http.StatusUnprocessableEntity},
	{get, "/operations?year=2020&month=13", http.StatusUnprocessableEntity},

	// Categories
	{get, "/categories", http.StatusOK},
}

func main() {
	start := time.Now()
	for _, tc := range testCases {
		tc.test(tc.route, tc.status)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("Finished in %s: %d errors", elapsed.Truncate(time.Millisecond), totalErrs)
}
