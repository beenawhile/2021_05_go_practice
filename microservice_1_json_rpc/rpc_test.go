package main

import (
	"microservice_2_http_rpc/server"
	"microservice_2_json_rpc/client"
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
