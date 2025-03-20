package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	//"k8s.io/apimachinery/pkg/util/rand"
)

var newRand *rand.Rand

func init() {
	newRand = rand.New(rand.NewSource(time.Now().UnixNano()))

}
func main() {

	fmt.Println(ShakeRandIntn(1, 10))
	fmt.Println(ShakeRandIntn(1, 10))
	fmt.Println(ShakeRandIntn(1, 10))
	fmt.Println(ShakeRandIntn(1, 10))
	fmt.Println(ShakeRandIntn(1, 10))
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

func ShakeRandIntn(min, max int) int {
	return min + newRand.Intn(max-min)
}
