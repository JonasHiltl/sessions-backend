package datastruct

type Provider int64

const (
	Google Provider = iota
	Apple
)

func (p Provider) String() string {
	switch p {
	case Google:
		return "google"
	case Apple:
		return "apple"
	}
	return "unknown"
}

func (p Provider) EnumIndex() int {
	return int(p)
}
