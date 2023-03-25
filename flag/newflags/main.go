package main

import (
	"fmt"

	"github.com/namsral/flag"
)

func main() {
	var addr = flag.String("addr", ":8080", "address to listen")
	var authAddr = flag.String("auth_addr", "localhost:8081", "address for auth service")
	var tripAddr = flag.String("trip_addr", "localhost:8082", "address for trip service")
	var profileAddr = flag.String("profile_addr", "localhost:8082", "address for profile service")
	var carAddr = flag.String("car_addr", "localhost:8084", "address for car service")

	flag.Parse()

	fmt.Println(*addr)
	fmt.Println(*authAddr)
	fmt.Println(*tripAddr)
	fmt.Println(*profileAddr)
	fmt.Println(*carAddr)
}
