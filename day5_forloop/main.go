package main

import (
	"fmt"
	"time"
)

func main() {
	// basic for loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//condition based for loop (while loop)
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// for {
	// 	fmt.Println("runnning")
	// }

	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("inside if ", i)
			continue //skip this iteration
		}
		if i == 5 {
			fmt.Println("inside if ", i)
			break //exit loop
		}
		fmt.Println(i)
	}
	// lopping over the collections
	numbers := []int{1, 2, 3, 4, 5, 6}
	for i, v := range numbers {
		fmt.Println(i, v)
	}
	// if only values
	for _, v := range numbers {
		fmt.Println(v)
	}
	// if only index,keys
	for k, _ := range numbers {
		fmt.Println(k)
	}

	// loopin over the map
	m := map[string]int{
		"anil": 28,
		"Hn":   27,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	s := "anil"
	// range looping over the string -->>UTF-8 ,not in bytes
	for i, v := range s {
		fmt.Println(i, v)
	}
	// range internals
	for i := range numbers {
		numbers[i] = numbers[i] * 2
	}
	fmt.Println("right way to modify the slice", numbers)

	// nested for loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println(i, j)
			// break it is only internal break
		}
	}
	// labelled break and contine
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outer
			}
			fmt.Println(i, j)
		}
	}
	// for loop with channels
	// ch := make(chan int)
	// go func() {
	// 	for i := 1; i <= 3; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// }()
	// Key Rule
	// range on channel automatically stops when channel is closed
	// for v := range ch {
	// 	fmt.Println(v)
	// }
	// Event Loop with for + select + timeout
	ch := make(chan string)
	go producer(ch)
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				fmt.Println("channel close ,existing")
				return
			}
			fmt.Println("recived", msg)
		case <-time.After(2 * time.Second):
			fmt.Println("timeout: no message received")
			return
		}
	}
}
func producer(ch chan string) {
	message := []string{"order_created", "order_paid", "order_shipped"}
	for _, msg := range message {
		time.Sleep(1 * time.Second)
		ch <- msg
	}
	close(ch)
}
