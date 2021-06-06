// select : 여러 개의 channel에서 결과를 받을 수 있을 경우, 로직을 나눠주는 용도
//  - 여러개의 channel을 동시에 기다릴 수 있게 해줌
//  - switch case 문과 비슷함

package main

import (
	"fmt"
	"strconv"
	"time"
)

type Plane struct {
	val string
}

func MakeCertainTire(inCarChan chan Car, inPlaneChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {

		select {
		case car := <-inCarChan:
			car.val += "Tire_C, "
			outCarChan <- car
		case plane := <-inPlaneChan:
			plane.val += "Tire_P"
			outPlaneChan <- plane
		}

	}
}

func MakeCertainEngine(inCarChan chan Car, inPlaneChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-inCarChan:
			car.val += "Engine_C, "
			outCarChan <- car

		case plane := <-inPlaneChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}

	}
}

func StartCarWork(inCarChan chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		inCarChan <- Car{val: "Car " + strconv.Itoa(i)}
		i++
	}
}

func StartPlaneWork(inCarChan chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		inCarChan <- Plane{val: "Plane " + strconv.Itoa(i)}
		i++
	}
}

func channelSelectTest() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)

	go MakeCertainTire(carChan1, planeChan1, carChan2, planeChan2)
	go MakeCertainEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
	}

}
