package main

import (
	"fmt"
	"time"
)

// go channel
// - 데이터를 주고받는 통로
// - 채널은 make()함수를 통해 미리 생성되어야함
// - 채널 연산자 <- 를 통해 데이터를 보내고 받음
// - 흔히 goroutine들 사이 데이터를 주고 받는데 사용
// - 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화하는데 사용
func channel1Test() {

	// 정수형 채널을 생성
	ch := make(chan int)

	go func() {
		ch <- 123
	}()

	time.Sleep(time.Second * 3)

	i := <-ch
	fmt.Println(i)
}

// go channel은 수신자와 송신자가 서로 기다리는 속성때문에, 이를 이용하여  goroutine이 끝날 때까지 기다리는 기능을 구현할 수 있다.
// 즉, 익명함수를 사용한 한 goroutine에서 어떤 작업이 실행되고 있을때, 메인 routine은 <-done에서 계속 수신 대기하고 있게 됨
// goroutine 작업이 끝난 후, done 채널에 true를 보내면, 수신자 메인 루틴은 이를 받고 프로그램 끝냄
func channel2Test() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	fmt.Println("go routine start")

	// 위의 Go routine이 끝날때까지 대기
	<-done
	fmt.Println("really done")
}

// unbuffered channel : 하나의 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있게 됨
// ------- unbuffered channel -------

// -------- buffered channel ---------
// buffered channel : 수신자가 받을 준비가 되어 있지 않더라도 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있음
// - make(chan type, buffer quantity) 함수를 통해 생성됨

// 다음은 unbuffered channel로 보냈을 때 발생할 수 있는 에러
// - 메인루틴에서 채널에 1을 보내면서 상대편 수신자를 기다리고 있는데 수신자 goroutine이 없기 때문임
func wrongBufferedChannelTest() {
	c := make(chan int)
	c <- 1           // 수신 루틴이 없으므로 데드락
	fmt.Println(<-c) // 코멘트해도 데드락 (별도의 goroutine 이 없기 때문)
}

// buffered channel로 해결가능
// - 수신자가 당장 없더라도 최대버퍼 수 까지 데이터를 보낼 수 있으므로, 에러가 발생하지 않음
func bufferedChannelTest() {
	ch := make(chan int, 1)
	// 수신자가 없더라도 보낼 수 있음
	ch <- 101
	fmt.Println(<-ch)

}

// channel parameter
//  - 채널을 함수의 파라미터로 전달할 때, 일반적으로 송수신을 모두하는 채널을 전달하지만,
//    특별히 해당 채널로 송신만 할 것인지 혹은 수신만 할 것인지 지정할 수 있음
//  - 송신 파라미터는 p chan <- int 형태,
//		수신 파라미터는 p <- chan int 형태를 사용
//  - 송신 채널 파라미터에서 수신을 한다거나, 수신 채널에 송신을 하게되면 에러가 발생함
func wrongChannelParamTest() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string) {
	ch <- "Data"
	// x := <- ch // 에러발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

// 채널 닫기
// - 데이터를 송신한 이후, close()함수를 사용하여 채널 닫을 수 있음.
// - 채널 닫게 되면, 해당 채널로는 더 이상 송신할 수 없지만, 채널이 닫힌 이후에도 계속 수신 가능
// - 채널 수신에 사용되는 <-ch은 두개의 리턴 값을 받음
//    - 첫째 : 채널 메세지, 둘째 : 수신이 제대로 되었는지 나타냄

func channelCloseTest() {
	ch := make(chan int, 2) // 버퍼 2개 보냄
	// 테스트 해보니 1개 보내면 error 떨어지지만, 2개 이상 buffer 만들었을 때는 마지막 메세지까지 이상없이 잘 보냄

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 채널 수신
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch) // 송신한 갯수 이상 받으려하니 int 여서 그런지 0을 출력함

	if _, success := <-ch; !success {
		fmt.Println("더이상 데이터 없음")
	}
}

// channel range 문
// - 송신자가 송신한 만큼 수신자가 데이터 채널 닫힐 때 까지 계속 수신하는 방법 보여줌
func channelRangeTest() {
	ch := make(chan int, 2)

	// 채널 송신
	ch <- 1
	ch <- 2

	// 채널 닫음
	close(ch)

	// 방법 1
	for {
		if i, success := <-ch; success {
			fmt.Println(i)
		} else {
			break
		}
	}

	// 방법 2
	// for i := range ch {
	// 	fmt.Println(i)
	// }
}

// channel select문
// - select문 : 복수 채널들을 기다리면서 준비된 채널을 실행하는 기능을 제공
// 	 - case channel들이 준비되지 않으면 계속 대기하게 되고, 가장 먼저 도착한 채널의 case를 실행
// 	 - 만약 복수 채널에 신호가 오면, go runtime이 랜덤하게 그 중 한 개를 선택
// 	 - 하지만 default문이 있으면, case문 채널이 준비되지 않더라도 계속 대기하지 않고 바로 default문 실행

func channelSelectTest() {
	done1 := make(chan bool)
	done2 := make(chan bool)

	go run1(done1)
	go run2(done2)

	// break label : Go에서는 해당 레이블로 이동한 후 자신이 빠져나온 루프 다음 문장을 실행하게 된다.
	// - 다른 언어의 goto문과 다름
EXIT:
	for {
		select {
		case <-done1:
			fmt.Println("run1 완료")
		case <-done2:
			fmt.Println("run2 완료")
			break EXIT
		}
	}

	fmt.Println("all process done")
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
