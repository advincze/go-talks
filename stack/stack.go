package collection //package declaration

type Stack struct { // ~ java class
	data []interface{} // ~ Arraylist<Object>()
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x) //builtin
}

func (s *Stack) Pop() interface{} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res
}

func (s *Stack) Size() int {
	return len(s.data) //builtin
}
