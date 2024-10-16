package main

import (
	"easy-life-back-go/internal/server"
)

const port string = "8080"

func main() {
	server.Start(port)
}
