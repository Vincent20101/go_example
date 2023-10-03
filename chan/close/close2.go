package main

import "fmt"

func main() {
	var stopCh = make(chan bool)
	close(stopCh)
	select {
	case _, ok := <-stopCh:
		fmt.Println("p.stopCh has been closed:", ok)
	default:
		fmt.Println("p.stopCh is still open, need to close it")
		//close(stopCh)
	}

	//if _, ok := <-stopCh; ok {
	//	fmt.Println("p.stopCh is still open, need to close it")
	//	close(stopCh)
	//} else {
	//	fmt.Println("false:", ok)
	//}
}
