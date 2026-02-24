package cmd

import (
	"errors"
	"net/http"
	"time"
)

var httpClient = createHttpClient()

// Digital ocean demo example "production" ready API client (https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go#making-a-get-request)
func createHttpClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100,              // Maximum idle connections
		MaxIdleConnsPerHost: 10,               // Maximum idle connections per host
		IdleConnTimeout:     90 * time.Second, // Idle connection timeout
		DisableCompression:  false,            // Enable compression
		DisableKeepAlives:   false,            // Enable keep-alives
	}

	return &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Custom redirect handling
			if len(via) >= 2 {
				return errors.New("stopped after 2 redirects!")
			}
			return nil
		},
	}
}
