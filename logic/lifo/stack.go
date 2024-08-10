package lifo

type Stack struct {
	Numbers []int
}

func (s *Stack) Push(data int) {
	s.Numbers = append(s.Numbers, data)
}

func (s *Stack) Pop() int {
	if s.Isempty() {
		return 0
	}
	last := len(s.Numbers) - 1
	s.Numbers = s.Numbers[:len(s.Numbers)-1]
	return last

}
func (s *Stack) lastI() int {
	last := s.Numbers[len(s.Numbers)-1]

	return last
}
func (s *Stack) Isempty() bool {
	if len(s.Numbers) != 0 {
		return false
	}
	return true
}
