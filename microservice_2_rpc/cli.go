package main

import (
	"fmt"
	"microservice_2_rpc/client"
	"microservice_2_rpc/server"
)

// RPC
// - 인터페이스를 준수할 필요가 없다?
// - 구조체 필드 및 관련 메서드를 가지고 있다면, 이것을 RPC 서버에 등록할 수 있음

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
