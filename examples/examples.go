package examples

import (
	"fmt"
)

func ExampleFuncWithError(s string) error {
	return fmt.Errorf("Error: %s", s)
}

func ExampleFuncWithResult(s string) string {
	return fmt.Sprintf("Hello: %s", s)
}

func ExampleFuncWithResultAndError(s string) (int, error) {
	return 2, fmt.Errorf("Hello: %s", s)
}

type stru struct {
	foo string
}

func (s *stru) ExampleFuncWithNoParams() {
	s.foo = "Hello"
}

func (s *stru) ExampleFuncWithParam(x string) {
	s.foo = x
}

func (s stru) ExampleFuncWithReturn() string {
	return s.foo
}

func (s stru) ExampleFuncWithParamAndReturnAndError(x string) (string, error) {
	if x == "" {
		return s.foo, nil
	} else {
		return "", fmt.Errorf("Error: %s", x)
	}
}
