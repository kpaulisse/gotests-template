package examples

import "fmt"

func exampleFuncWithError(s string) error {
	return fmt.Errorf("Error: %s", s)
}

func exampleFuncWithResult(s string) string {
	return fmt.Sprintf("Hello: %s", s)
}

func exampleFuncWithResultAndError(s string) (int, error) {
	return 2, fmt.Errorf("Hello: %s", s)
}

type stru struct {
	foo string
}

func (s *stru) exampleFuncWithNoParams() {
	s.foo = "Hello"
}

func (s *stru) exampleFuncWithParam(x string) {
	s.foo = x
}

func (s stru) exampleFuncWithReturn() string {
	return s.foo
}
