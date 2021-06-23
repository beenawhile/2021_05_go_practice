package main

import (
	"fmt"
	"strings"
)

func stringsTest() {

	m1 := "my name"
	m2 := "name"
	fmt.Println(strings.Contains(m1, m2))

	fmt.Println(strings.ReplaceAll(m1, "m", "NO"))

	fmt.Println(strings.Split(m1, " "), m1+m2)
}
