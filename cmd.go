package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	Numbers []int
}

func (s *stack) Push(data int) {
	s.Numbers = append(s.Numbers, data)
}

func (s *stack) Pop() int {
	if s.Isempty() {
		return 0
	}
	last := len(s.Numbers) - 1
	s.Numbers = s.Numbers[:len(s.Numbers)-1]
	return last

}
func (s *stack) lastI() int {
	last := s.Numbers[len(s.Numbers)-1]

	return last
}
func (s *stack) Isempty() bool {
	if len(s.Numbers) != 0 {
		return false
	}
	return true
}

func main() {

	var Num = stack{}
	Num.Push(1)
	Num.Push(2)
	Num.Push(3)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		//fmt.Println("you have wrote dj", input)

		switch input {
		case "-last":
			print(Num.lastI())

		case "-d":
			Num.Pop()
			print(Num.lastI())

		case "-hist":
			fmt.Println(Num.Numbers)

		default:
			fmt.Println(input)

		}
	}

	//fmt.Println(Num)
}
