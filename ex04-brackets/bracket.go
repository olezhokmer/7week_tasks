package brackets

type Stack []string

func New() *Stack {
	return &Stack{}
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Peek() string {
	l := len(*s)
	if l != 0 {
		return (*s)[l-1]
	}
	return ""
}

func (s *Stack) Pop() string {
	l := len(*s)
	if l != 0 {
		a := (*s)[l-1]
		*s = (*s)[:l-1]
		return a
	}
	return ""
}

func Bracket(s string) (bool, error) {
	st := New()
	f := true
	for _, ch := range s {
		if ch == '{' || ch == '(' || ch == '[' {
			st.Push(string(ch))
			f = false
		} else if (ch == '}' || ch == ')' || ch == ']') && !f {
			l := st.Peek()
			if (l == "{" && ch == '}') || (l == "(" && ch == ')') || (l == "[" && ch == ']') {
				st.Pop()
				f = true
				continue
			}
			return false, nil
		}
	}
	if f {
		return true, nil
	} else {
		return false, nil
	}
}
