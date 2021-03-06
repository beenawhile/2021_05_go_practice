package main

import (
	"fmt"
	"microservice_1_http_rpc/client"
	"microservice_1_http_rpc/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
