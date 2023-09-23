//go:build !linux

package schedpolicy

import (
	"errors"
)

// Set is not implemented.
func Set(pid int, policy Policy, priority int) error {
	return errors.New("not implemented")
}

// Get is not implemented.
func Get(pid int) (Policy, error) {
	return 0, errors.New("not implemented")
}
