## Event source를 이용한 chatting server

- Event Source : server-sent events에 대한 웹 컨텐츠 인터페이스
  - Event source는 text/event-stream 포맷으로 이벤트를 보내는 HTTP 서버에 지속적인 연결을 함
  - 무슨말인지 이해하기 힘들다

### 설명
- 과거의 웹은 request를 보내면 이에 대한 문서를 만들어서 response 반환하는 것이 역할이 끝이었음
- 현대는 동적 웹에 대한 니즈가 강해져 html5에 websocket, event source 라는 기술이 추가됨

- websocket : 소켓처럼 하겠다?
  - network의 tcp socket처럼 양쪽에 socket을 연결해서 데이터를 주고 받을 수 있도록 함
  - 단건 통신만 하고 연결을 끊어버리는 것이 아닌, 지속적으로 연결해놓고 데이터를 주고 받을 수 있게 함

- event source
  - (단순하게) server만 지속적으로 데이터를 send
