package queue_utils

type node interface{}

type Stack struct {
	Items []node
}

func (s *Stack) Push(item node) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() (node, bool) {
	if len(s.Items) == 0 {
		return 0, false
	}
	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return item, true
}
