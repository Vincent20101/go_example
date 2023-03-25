package main

import (
	"fmt"
	"math"
	"time"

	"k8s.io/apimachinery/pkg/util/rand"
)

func main() {
	//rand.Seed(time.Now().Unix())
	//for i := 0; i < 10; i++ {
	//	fmt.Println(rand.Intn(50))
	//}

	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Uint64())
	//for i := 0; i < 10; i++ {
	//	fmt.Println(r2.Int())
	//}

	rand.Seed(time.Now().UnixNano())
	randPort := uint16(rand.Intn(math.MaxUint16-math.MaxUint16/2) + math.MaxUint16/2)
	fmt.Println(randPort)
	fmt.Println(math.MaxUint16)
	//fmt.Println(rand.Int(rand.Reader, big.NewInt(time.Now().UnixNano())))

	var s *int
	fmt.Println(s)
	s = new(int)
	fmt.Println(s)
}
