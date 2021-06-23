package main

import (
	"microservice_1_json_rpc/client"
	"microservice_1_json_rpc/server"
	"testing"
)

func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest()
	}
}

func init() {
	// start the server
	go server.StartServer()
}
