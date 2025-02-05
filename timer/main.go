// Golang program to illustrate the usage of
// AfterFunc() function

// Including main package
package main

// Importing fmt and time
import (
	"fmt"
	"time"
)

// Main function
func main() {

	TimerFunc()
}

func TimerFunc() {

	DurationOfTime := time.Duration(3) * time.Second

	f := func() {

		fmt.Println("Function called by " +
			"AfterFunc() after 3 seconds")
	}

	Timer1 := time.AfterFunc(DurationOfTime, f)

	//defer Timer1.Stop()

	time.Sleep(3 * time.Second)
	Timer1.Stop()
	fmt.Println("stop finish")
	time.Sleep(5 * time.Second)
}

func TimerFunc2() {

	DurationOfTime := time.Duration(3) * time.Second

	f := func() {

		fmt.Println("Function called by " +
			"AfterFunc() after 3 seconds")
	}

	Timer1 := time.AfterFunc(DurationOfTime, f)

	time.Sleep(1 * time.Second)
	f2 := func() {

		fmt.Println("22222222Function called by " +
			"AfterFunc() after 3 seconds")
	}
	Timer1 = time.AfterFunc(DurationOfTime, f2)
	Timer1.Stop()
	fmt.Println("stop finish")
	time.Sleep(5 * time.Second)
}

// Golang program to illustrate the usage of
// AfterFunc() function

func main01() {

	// Defining duration parameter of
	// AfterFunc() method
	DurationOfTime := time.Duration(3) * time.Second

	// Defining function parameter of
	// AfterFunc() method
	f := func() {

		// Printed when its called by the
		// AfterFunc() method in the time
		// stated above
		fmt.Println("Function called by " +
			"AfterFunc() after 3 seconds")
	}

	// Calling AfterFunc() method with its
	// parameter
	Timer1 := time.AfterFunc(DurationOfTime, f)

	// Calling stop method
	// w.r.to Timer1
	defer Timer1.Stop()

	// Calling sleep method
	time.Sleep(10 * time.Second)
}

// Golang program to illustrate the usage of
// AfterFunc() function

func main02() {

	// Creating channel using
	// make keyword
	mychan := make(chan int)

	// Calling AfterFunc() method
	// with its parameters
	time.AfterFunc(6*time.Second, func() {

		// Printed after stated duration
		// by AfterFunc() method is over
		fmt.Println("6 seconds over....")

		// loop stops at this point
		mychan <- 30
	})

	// Calling for loop
	for {

		// Select statement
		select {

		// Case statement
		case n := <-mychan:

			// Printed after the loop stops
			fmt.Println(n, "is arriving")
			fmt.Println("Done!")
			return

		// Returned by default
		default:

			// Printed until the loop stops
			fmt.Println("time to wait")

			// Sleeps for 3 seconds
			time.Sleep(3 * time.Second)
		}
	}
}
