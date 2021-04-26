package main

import "log"

func main() {
	server := New()

	err := server.Run(":8000")

	if err != nil {
		log.Fatal(err)
	}
}
