package server

import (
	"fmt"
	"log"
	"microservice_2_rpc/contract"
	"net"
	"net/rpc"
)

const port = 1234

type HelloWorldHandler struct{}

// Register : 지정된 인터페이스의 일부 메서드를 기본 서버에 공개하고 클라이언트가 해당 메서드를 호출할 수 있게 함
// 이름은 구체화된 타입(concrete type)의 이름을 사용
// 구체화된 이름을 사용하고 싶지 않다면, 매개 변수로 제공된 이름을 대신 사용하는 RegisterName 함수를 사용해 다른 이름으로 등록할 수 있음
func StartServer() {
	helloworld := &HelloWorldHandler{}
	// handler를 rpc 서버에 등록
	rpc.Register(helloworld)

	// net/http 와 달리 RPC를 사용할 때는 수작업을 거쳐줘야함
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	// ListenAndServe와 다르게 RPC 서버는 각 연결을 개별적으로 처리하며 처음 연결을 처리하고 나서 Accept 를 호출해 후속 연결을 처리하거나 애플리케이션을 종료
	//  => for loop을 돌아줘야함
	for {
		// http.Listener interface의 Accept method : 연결 수신
		// Accpet : 블로킹 메서드 => 현재 서비스에 연결하려는 클라이언트가 없는 경우 Accpet 메서드는 연결이 이루어 질 때까지 대기
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
		// ServeConn : 지정된 연결에서 DefaultServer 메서드를 실행하고 클라이언트가 완료될때까지 대기
		// go 구문을 사용하여 연결을 종료할 때 까지 대기하지 않고 즉시 대기 중인 다음 연결을 처리할 수 있음
		// 통신 프로토콜 측면에서 ServeConn은 Gob wire 타입을 사용

		// Gob : Go 프로세스 사이의 통신을 용이하게 하기 위해 특별히 고안됐으며, Protocol Buffer와 같은 것 보다 사용하기 쉬우면서 좀 더 효율적일 수 있는 것을 만들려는 아이디어를 바탕으로 설계
		// - gob을 사용할 때 원본 데이터와 대상 데이터 값과 타입이 정확하게 일치할 필요가 없다

		// HTTP 프로토콜을 사용하지 않고, 메시지 형식도 JSON이 아니므로 클라이언트에 요청할 때 간단하게 curl을 사용할 수 없음
	}
}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
