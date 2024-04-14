package stack

type stack []rune

func (s stack) Len() int   { return len(s) }
func (s stack) Last() rune { return s[s.Len()-1] }
func (s *stack) RemoveLast() {
	*s = (*s)[:s.Len()-1]
}

var closeToOpen = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func IsValid(s string) bool {
	stack := make(stack, 0, len(s))
	for _, chr := range s {
		if chr == '(' || chr == '{' || chr == '[' {
			stack = append(stack, chr)
		} else {
			if stack.Len() == 0 || stack.Last() != closeToOpen[chr] {
				return false
			}
			stack.RemoveLast()
		}
	}
	return stack.Len() == 0
}
