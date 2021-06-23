package client

import (
	"fmt"
	"log"
	"microservice_2_rpc/contract"
	"net/rpc"
)

const port = 1234

// HTTP 프로토콜을 사용하지 않고, 메시지 형식도 JSON이 아니므로 클라이언트에 요청할 때 간단하게 curl을 사용할 수 없음
func CreateClient() *rpc.Client {
	// client 자체를 생성
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	return reply
}
