package main

import (
	"fmt"
	"go-make-me-a-pdf/internal/server"
)

func main() {
	server := server.NewServer()

	fmt.Println("Sever running at port", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
