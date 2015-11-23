package main

import (
	"net/http"
	"strings"
)

// DoServer starts HTTP server and handles redirecting image requests
func DoServer(path string, hostAndPort string, searcher func(string, string) string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		subdomain := strings.Split(r.Host, ".")[0]
		redirect := searcher(r.RemoteAddr, subdomain)
		http.Redirect(w, r, redirect, 307)
	})
	http.ListenAndServe(hostAndPort, nil)
}
