package core_http_response

import (
	"fmt"
	"net/http"
)

type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

const statusCodeUnitialized = -1

func NewResponseWriter(rw http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{
		ResponseWriter: rw,
		statusCode:     statusCodeUnitialized,
	}
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.statusCode = statusCode
}

func (rw *ResponseWriter) Write(bytes []byte) (int, error) {
	bytesWrited, err := rw.ResponseWriter.Write(bytes)
	if err != nil {
		rw.statusCode = rw.GetStatusCodeOrPanic()
	}
	rw.statusCode = http.StatusOK
	return bytesWrited, err
}

func (rw *ResponseWriter) GetStatusCodeOrPanic() int {
	if rw.statusCode == statusCodeUnitialized {
		fmt.Println("\033[31mPANIC: no status code set\033[0m")
		panic("no status code set")
	}
	return rw.statusCode
}
