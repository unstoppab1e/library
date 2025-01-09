package main

import "github.com/unstoppab1e/library/src/cmd/api"

func main() {
	server := api.NewAPIServer(":8080", nil)
	if server.Run():

}
