package schedpolicy

import (
	"fmt"
	"syscall"
	"unsafe"
)

// Policy is a scheduling policy.
type Policy int

func (p Policy) String() string {
	switch p {
	case Normal:
		return "normal"
	case FIFO:
		return "fifo"
	case RR:
		return "rr"
	case Batch:
		return "batch"
	case Idle:
		return "idle"
	case Deadline:
		return "deadline"
	default:
		return fmt.Sprintf("unknown(%d)", int(p))
	}
}

// Possible values for sched_policy.
//
// https://github.com/torvalds/linux/blob/d90b0276af8f25a0b8ae081a30d1b2a61263393b/include/uapi/linux/sched.h#L111-L120
const (
	Normal   = 0
	FIFO     = 1
	RR       = 2
	Batch    = 3
	Idle     = 5
	Deadline = 6
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
