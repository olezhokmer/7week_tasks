package stack

type Stack []int

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() int {
	l := len(*s)
	if l != 0 {
		n := (*s)[l-1]
		*s = (*s)[:l-1]
		return n
	}
	return 0
}

func (s *Stack) Push(n int) {
	*s = append(*s, n)
}
