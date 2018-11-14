package util

import "fmt"

// ArrayFlags ...
type ArrayFlags []string

func (i *ArrayFlags) String() string {
	return fmt.Sprint(*i)
}

// Set ...
func (i *ArrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}
