package main

import "APIGateway/internal/server"

func main() {
	s := server.New()
	s.Run()
}
