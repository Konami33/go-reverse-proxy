// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// )

// // ProxyServer represents a reverse proxy server.
// type ProxyServer struct {
// 	targets map[string]*url.URL
// 	proxy   *httputil.ReverseProxy
// }

// // NewProxyServer creates a new ProxyServer instance.
// func NewProxyServer() (*ProxyServer, error) {
// 	targets := map[string]*url.URL{
// 		"/origin1/app": parseURL("http://127.0.0.1:8081"),
// 		"/origin2/app":  parseURL("http://127.0.0.1:8082"),
// 	}

// 	//fmt.Printf("request url path: %v\n", req.URL.Path)

// 	reverseProxy := &httputil.ReverseProxy{
// 		Director: func(req *http.Request) {
// 			fmt.Printf("request url path: %v\n", req.URL.Path)

// 			//extracting the <namespace>
// 			target := targets[req.URL.Path]

// 			//checking if valid otherwise setting up the target URL info
// 			if target != nil {
// 				req.URL.Scheme = target.Scheme
// 				req.URL.Host = target.Host
// 				req.URL.Path = target.Path
// 			}
// 		},
// 	}

// 	return &ProxyServer{
// 		targets: targets,
// 		proxy:   reverseProxy,
// 	}, nil
// }

// // ServeHTTP handles incoming HTTP requests and forwards them to the target server.
// //Proxyserver struct method
// func (p *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("Proxying request for %s to: %s\n", r.URL.Path, p.targets[r.URL.Path])
// 	p.proxy.ServeHTTP(w, r)
// }

// func parseURL(rawURL string) *url.URL {
// 	u, err := url.Parse(rawURL)
// 	if err != nil {
// 		log.Fatalf("Failed to parse URL: %v", err)
// 	}
// 	return u
// }

// func main() {
// 	proxyServer, err := NewProxyServer()
// 	if err != nil {
// 		log.Fatalf("Failed to create proxy server: %v", err)
// 	}
// 	fmt.Println("Proxy server running on :8080")
// 	log.Fatal(http.ListenAndServe(":8080", proxyServer))  //listening at port 8080
// }

package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Define the targets for different URL paths
	targets := map[string]*url.URL{
		"/origin1/app": parseURL("http://127.0.0.1:8081"),
		"/origin2/app": parseURL("http://127.0.0.1:8082"),
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
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		// Forward the request to the target
		proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = target.Path
		}}
		proxy.ServeHTTP(w, r)
	})
}
