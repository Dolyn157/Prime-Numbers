package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {

	for i := 1; i <= 190000; i++ {
		intChan <- i
	}
	close(intChan)
}

func priNum(intChan chan int, priChan chan int, exitChan chan bool) {
	for {
		num, ok := <-intChan
		//Check if there remains number in intChan channel.
		if !ok {
			break
		}
		var isP bool = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isP = false
				break
			}
		}
		if isP {
			priChan <- num
		}
	}
	fmt.Print("a priNum coroutine ended because it failed to get numbers from intChan\n")
	exitChan <- true
}

func main() {
	start := time.Now().Unix()

	theintChan := make(chan int, 20000)
	theprimeChan := make(chan int, 30000)
	theexitChan := make(chan bool, 4)

	// Start a coroutine to put natural numbers into the theintChan.
	go putNum(theintChan)

	/* Start 4 coroutines to check if the number from theintChan is a prime, and put them into the result channel
	(theprimeChan)
	*/
	for i := 0; i < 4; i++ {
		go priNum(theintChan, theprimeChan, theexitChan)
	}

	//Wait for the four coroutines to finish their checks.
	go func() {
		for i := 0; i < 4; i++ {
			<-theexitChan
		}
		close(theprimeChan)
	}()

	//Traverse the result channel

	for {
		res, ok := <-theprimeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
	fmt.Println("The main process ended.")
	end := time.Now().Unix()
	fmt.Println("Finding primes with 4 coroutines takes", end-start, "seconds")
	fmt.Println("The program will close in 12 seconds.")
	time.Sleep(time.Second * 12)
	
}

/*
func priNum(intChan chan int, priChan chan int, exitChan chan bool) {
	var isP bool = true
	for num := range intChan {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isP = false
				break
			}
		}
		if isP {
			priChan <- num
		}
	}
}

*/
