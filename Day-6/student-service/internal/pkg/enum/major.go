package enum

type Major int64

const (
	Finance Major = iota + 1
	IT
	Medical
)

func (d Major) String() string {
	switch d {
	case Finance:
		return "Finance"
	case IT:
		return "Information Technology"
	case Medical:
		return "Medical"
	}
	return "Unknown"
}
