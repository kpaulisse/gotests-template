package examples

import (
	"errors"
	"fmt"
)

func ExampleFuncWithStructInput(s stru) string {
	return s.ExampleFuncWithReturn()
}

func ExampleFuncWithError(s string) error {
	return fmt.Errorf("Error: %s", s)
}

func ExampleFuncWithResult(s string) string {
	return fmt.Sprintf("Hello: %s", s)
}

func ExampleFuncWithResultAndError(s string) (int, error) {
	return 2, fmt.Errorf("Hello: %s", s)
}

func ExampleFuncWithMultipleResults(s string) (string, int) {
	return fmt.Sprintf("1: %s", s), len(s)
}

func ExampleFuncWithMultipleResultsAndError(s, e string) (string, int, error) {
	if e == "" {
		return fmt.Sprintf("1: %s", s), len(s), nil
	} else {
		return "", len(s), errors.New(e)
	}
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
