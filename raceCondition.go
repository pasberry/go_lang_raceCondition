// The following code has a race condition that may cause the value of the printCounter function statement  to vary
//between 0 or 2 based on how the underlining system interleaves the assembly code. If the two underlining go routines
//incrementCounter occur before printCounter the value will be 2. If the system performs the printCounter function before
//the value is incremented the resulting output will be 0.
// Else if the program interleaves in a manner expected by the code execution manner, the output of the program will be 1.
package main

import (
	"fmt"
	"sync"
)

var counter = 0

func incrementCounter(group *sync.WaitGroup) {

	counter = counter + 1
	group.Done()
}

func printCounter(group *sync.WaitGroup) {

	fmt.Println(counter)
	group.Done()
}

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(3)

	go incrementCounter(&waitGroup)
	go printCounter(&waitGroup)
	go incrementCounter(&waitGroup)

	waitGroup.Wait()
}

