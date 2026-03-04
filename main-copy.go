package main

// import (
// 	"fmt"
// 	"time"
// )

// func sleep() {
// 	fmt.Println("Start:", time.Now())
// 	// Sleep for 2 seconds
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("End:", time.Now())
// 	panic("something broke")
// }

// func main() {
// 	// Create a new channel of type string

// 	messages := make(chan string)
// 	errChan := make(chan string)

// 	safeSleep := func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				errChan <- "something broke"
// 			}
// 		}()
// 		sleep()
// 		messages <- "ping"
// 	}
// 	// Receive the message from the channel and print it
// 	go safeSleep()
// 	select {
// 	case err := <-errChan:
// 		if err != "" {
// 			fmt.Println(err)
// 		}
// 	case msg := <-messages:
// 		fmt.Println(msg)
// 	}

// 	fmt.Println("Waiting for message...")

// }
