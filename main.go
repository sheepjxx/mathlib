package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func randomAddQuestion() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var a int
	var b int
	var input int
	var result int
	var promote string
	a = r.Intn(10)
	b = r.Intn(10)
	result = a + b
	promote = strconv.Itoa(a) + "+" + strconv.Itoa(b) + "="
	fmt.Print(promote)
	fmt.Scanln(&input)
	if input == result {
		fmt.Println("correct")
	} else {
		fmt.Printf("Wrong, correct answer is %d\n", result)
	}
}

func main() {

	for {
		randomAddQuestion()
		time.Sleep(1 * time.Second)
	}
}
