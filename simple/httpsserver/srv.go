package main

import (
	"io"
	"net/http"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Path == "/favicon.ico" {
		Printfln("Request for icon detected - returning 404")
		writer.WriteHeader(http.StatusNotFound)
		return

	}
	// Printfln("Method: %v", request.Method)
	// Printfln("URL: %v", request.URL)
	// Printfln("HTTP Version: %v", request.Proto)
	// Printfln("Host: %v", request.Host)

	// for name, val := range request.Header {
	// 	Printfln("Header: %v, Value: %v", name, val)
	// }

	// Printfln("---")
	Printfln("Request for %v", request.URL.Path)

	io.WriteString(writer, sh.message)
}
func runsrv() {

	err := http.ListenAndServe(":5000", StringHandler{message: "Hello, World"})
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}
