package main

import "fmt"

type school struct {
	countNum CountNum
}

func (s *school) SetCount(c *int) {
	s.countNum.count = c
}

func (s *school) GetCount() *int {
	return s.countNum.count
}

func (s *school) FinalCount() *int {
	return s.countNum.count
}

type CountNum struct {
	count *int
}

func main() {
	s := school{}
	cc := s.GetCount()
	fmt.Println(cc)
	s.SetCount(&[]int{3}[0])

	cc = s.GetCount()
	*cc--
	fmt.Println(*s.FinalCount())
}
