package util

import "strings"

// @example `go run cmd/main.go -ids "1,2"`

// SliceFlag ...
type SliceFlag []string

// NewSliceFlag ...
func NewSliceFlag(vals []string, p *[]string) *SliceFlag {
	*p = vals
	return (*SliceFlag)(p)
}

// Set ...
func (s *SliceFlag) Set(val string) error {
	*s = SliceFlag(strings.Split(val, ","))
	return nil
}

// Get ...
func (s *SliceFlag) Get() interface{} {
	return []string(*s)
}

// String ...
func (s *SliceFlag) String() string {
	return strings.Join([]string(*s), ",")
}
