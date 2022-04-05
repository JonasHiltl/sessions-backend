package comtypes

import "encoding/json"

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

func (t *Role) FromString(role string) Role {
	return map[string]Role{
		"user":  UserRole,
		"admin": AdminRole,
	}[role]
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

func (t Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Role) UnMarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = t.FromString(s)
	return nil
}
