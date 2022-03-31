package datastruct

type Role int64

const (
	UserRole Role = iota
	AdminRole
)

func (p Role) String() string {
	switch p {
	case UserRole:
		return "user"
	case AdminRole:
		return "admin"
	}
	return "unknown"
}

func (p Role) EnumIndex() int {
	return int(p)
}
