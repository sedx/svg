package path

// Command is SVG path command
type Command string

const (
	// M path command
	M Command = "M"
	// MR is m path command
	MR Command = "m"
	// L path command
	L = "L"
	// LR is l path command
	LR = "l"
	// Q path command
	Q = "Q"
	// T path command
	T = "T"
	// C path comand
	C = "C"
	// CR is c path comand
	CR = "c"
	// S path command
	S = "S"
	// SR is s path command
	SR = "s"
	// A path command
	A = "A"
	// AR is a path command
	AR = "a"
	// Z is Z path command
	Z = "Z"
)
