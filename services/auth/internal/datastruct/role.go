package datastruct

type Role int64

const (
	UserRole Role = iota
	AdminRole
)

func (r Role) String() string {
	switch r {
	case UserRole:
		return "user"
	case AdminRole:
		return "admin"
	}
	return "unknown"
}

func (r Role) EnumIndex() int {
	return int(r)
}

func (r Role) IsNil() bool {
	if r.String() == "unknown" {
		return true
	} else {
		return false
	}
}
