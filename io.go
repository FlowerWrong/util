package util

import (
	"io"
	"net"
)

// IsEOF return true if this is a EOF error
func IsEOF(err error) bool {
	if err == io.EOF {
		return true
	}
	return false
}

// IsTimeout return true if this is a Timeout error
func IsTimeout(err error) bool {
	if err, ok := err.(net.Error); ok && err.Timeout() {
		return true
	}
	return false
}
