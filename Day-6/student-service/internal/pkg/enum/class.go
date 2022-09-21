package enum

type Class int64

const (
	A Class = iota + 1
	B
)

func (r Class) String() string {
	switch r {
	case A:
		return "A"
	case B:
		return "B"
	}
	return "Unknown"
}
