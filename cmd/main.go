package main

import (
	"fmt"
	"github.com/treddy73/go-fullstack/internal/server"
	"os"
)

func main() {
	srv, err := server.New(server.NewConfig())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	srv.Start() // blocking
}
