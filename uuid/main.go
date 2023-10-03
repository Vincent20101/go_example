package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New()
	fmt.Println(uuid.Domain())
	fmt.Println(uuid[0:4])
	fmt.Println(uuid.String())
	fmt.Println(uuid.ID())
	fmt.Println(fmt.Sprintf("%s-%016x", uuid.String(), 232131))
	fmt.Println(fmt.Sprintf("%s-%.16x", uuid.String(), 232131))
}
