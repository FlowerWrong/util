package util

import "io"

// IsEOF return true if this is a EOF error
func IsEOF(err error) bool {
	if err == io.EOF {
		return true
	}
	return false
}
