package main

import (
	"log"
	"net/http"
	"os"
	"unsafe"
)

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func echoResponseHeader(writer http.ResponseWriter, request *http.Request) {
	reqHeader := request.Header
	for headerName, headerValue := range reqHeader {
		for _, val := range headerValue {
			writer.Header().Add(headerName, val)
		}
	}
	versionEnv := os.Getenv("VERSION")
	writer.Header().Add("VERSION", versionEnv)

	logStr := "server: " + request.Host + " & remote ip,port: " + request.RemoteAddr

	writer.WriteHeader(http.StatusAccepted)
	writer.Write(str2bytes(logStr))
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/response_header", echoResponseHeader)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
