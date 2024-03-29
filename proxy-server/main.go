package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define the targets for different URL paths
	targets := map[string]*url.URL{
		"/origin1/app": parseURL("http://host.docker.internal:8081"),
		"/origin2/app": parseURL("http://host.docker.internal:8082"),
	}

	// Create a reverse proxy
	proxy := NewReverseProxy(targets)

	// Start the HTTP server
	http.Handle("/", proxy)
	fmt.Println("Proxy server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func parseURL(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		panic(err)
	}
	return u
}

// NewReverseProxy creates a reverse proxy handler that forwards requests to different targets based on the URL path
func NewReverseProxy(targets map[string]*url.URL) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		target, ok := targets[r.URL.Path]
		if !ok {
			// http.Error(w, "Proxy-server is running", http.StatusNotFound)
			// return
			log.Printf("Target not found for path: %s", r.URL.Path)
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		log.Printf("Forwarding request to target: %s", target.String())

		// Forward the request to the target
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path
		}}
		proxy.ServeHTTP(w, r)
	})
}
