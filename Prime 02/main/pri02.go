package main

import (
	"fmt"
	"time"
)

func main() {
	var isP bool
	var primeSli []int = make([]int, 100, 200)
	//var res int
	start := time.Now().Unix()
	for num := 1; num <= 190000; num++ {
		isP = true
		for i := 2; i < num; i++ {
			if num%i == 0 { //if a number is divisible by other than 1 and itself, it isn't a prime.
				isP = false
				break
			}
		}
		if isP {
			primeSli = append(primeSli, num)
			fmt.Println(num, " is a prime ")
		}
	}
	fmt.Printf("There are %d primes in the natural numbers from 1 to 190000\n", len(primeSli))
	end := time.Now().Unix()
	fmt.Println("Finding primes with 1 process takes", end-start, "seconds")
	fmt.Println("The program will close in 10 seconds.")
	time.Sleep(time.Second * 10)
	
}
