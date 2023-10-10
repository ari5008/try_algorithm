package try

type Stack struct {
    stack []interface{}
}

func (s *Stack) Push(data interface{}) {
    s.stack = append(s.stack, data)
}

func (s *Stack) Pop() interface{} {
    if len(s.stack) > 0 {
        data := s.stack[len(s.stack)-1]
        s.stack = s.stack[:len(s.stack)-1]
        return data
    }
    return nil
}

func (s *Stack) GetStack() []interface{} {
    return s.stack
}