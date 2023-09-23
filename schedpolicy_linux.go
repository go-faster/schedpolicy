//go:build linux

package schedpolicy

import (
	"syscall"
	"unsafe"
)

// Set sets the scheduling policy and parameters for the process specified by pid.
//
// Calls sched_setscheduler.
// https://man7.org/linux/man-pages/man2/sched_setscheduler.2.html
func Set(pid int, policy Policy, priority int) error {
	p := struct{ Priority int32 }{
		Priority: int32(priority),
	}
	if _, _, errno := syscall.Syscall(
		syscall.SYS_SCHED_SETSCHEDULER,
		uintptr(pid),
		uintptr(policy),
		uintptr(unsafe.Pointer(&p)),
	); errno != 0 {
		return errno
	}
	return nil
}

// Get gets the scheduling policy and parameters for the process specified by pid.
//
// Calls sched_getscheduler.
// https://man7.org/linux/man-pages/man2/sched_setscheduler.2.html
func Get(pid int) (Policy, error) {
	currentPolicy, _, errno := syscall.Syscall(
		syscall.SYS_SCHED_GETSCHEDULER,
		uintptr(pid),
		0, 0)
	if errno != 0 {
		return 0, errno
	}
	return Policy(currentPolicy), nil
}
