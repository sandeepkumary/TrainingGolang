package main

/*
This program explains the channels and go-routinees with the help of real
life example of a relay race
There are 5 runners on the track on various position waiting for previous runner to handover the baton to start running
a runner can not start running without receiving the baton from previous runner
first runner get the baton to mark the start of the race
*/
import (
	"fmt"
	"sync"
	"time"
)

//declare a wait group to wait untill the complete race is over
var wg sync.WaitGroup

func main() {
	//declare the unbuffered channel to synchronise the relay run
	baton := make(chan int)

	//Add one wait group to mark one race
	wg.Add(1)

	// start the run
	go Runner(baton)
	//pass the baton to the first runner to mark the start of the race
	baton <- 1

	wg.Wait()
}

//Function to simulate the run

func Runner(baton chan int) {
	var newrunner int
	//runner is waitng for the baton
	runner := <-baton
	fmt.Printf("Runner %d is running with baton \n", runner)

	//check if all runners are done running, if not then start a go-routine to mark a runner waiting
	if runner != 5 {
		newrunner = runner + 1
		fmt.Printf("runner %d is waiting for the baton \n", newrunner)
		go Runner(baton)
	}

	//simulating the time delay for each runner to complete the run
	time.Sleep(3 * time.Second)

	// If all runners have finished running
	if runner == 5 {
		fmt.Printf("Runner %d finished , race is over!!!\n", runner)
		//complete the wait group and return to main function
		wg.Done()
		return
	}

	fmt.Printf("runner %d passed the baton to runner %d \n", runner, newrunner)
	//Pass the baton to the next runner
	baton <- newrunner
}
