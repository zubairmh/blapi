package main

import (
	"github.com/zubairmh/blapi/internal/server"
)

func main() {
	server.CreateHTTPServer(5000).ListenAndServe()
}
