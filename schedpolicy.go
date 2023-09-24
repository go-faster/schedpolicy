package schedpolicy

//go:generate go run -modfile=_tools/go.mod github.com/dmarkham/enumer -type Policy -transform lower -text -output policy_gen.go

// Policy is a scheduling policy.
type Policy int

// Possible values for sched_policy.
//
// https://github.com/torvalds/linux/blob/d90b0276af8f25a0b8ae081a30d1b2a61263393b/include/uapi/linux/sched.h#L111-L120
const (
	Normal   Policy = 0
	FIFO     Policy = 1
	RR       Policy = 2
	Batch    Policy = 3
	Idle     Policy = 5
	Deadline Policy = 6
)
