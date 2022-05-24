package main

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) peek() string {
	if s.isEmpty() {
		return ""
	} else {
		return (*s)[len(*s)-1]
	}
}

func (s *Stack) pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		ele := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return ele, true
	}
}

func (s *Stack) push(ele string) {
	*s = append(*s, ele)
}

func main() {

}

func polishNotation() {

}
