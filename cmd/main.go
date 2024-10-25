package main

import "github.com/n0o01lh/ml-url-shortener/internals/server"

func main() {

	server := server.NewServer()

	server.Initialize()
}
