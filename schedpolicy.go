package schedpolicy

import "fmt"

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
