package main

import "net/http"

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc()
}
