package main

import "ml-url-shortener/internals/server"

func main() {

	server := server.NewServer()

	server.Initialize()
}
