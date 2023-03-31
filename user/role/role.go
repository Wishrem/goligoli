package role

import (
	"encoding/json"

	"github.com/wishrem/goligoli/user/model"
)

type Role struct {
	roleType []string `json:"roles"`
}

func New(roleTypes ...string) *Role {
	role := new(Role)
	for _, s := range roleTypes {
		role.roleType = append(role.roleType, getRoleType(s))
	}
	return role
}

func GetRole(roles []model.Role) *Role {
	role := new(Role)
	for _, s := range roles {
		role.roleType = append(role.roleType, s.Type)
	}
	return role
}

func getRoleType(roleType string) string {
	switch roleType {
	case "user":
		return roleType
	case "admin":
		return roleType
	default:
		return "unknown"
	}
}

func (r *Role) IsAdmin() bool {
	for _, t := range r.roleType {
		if t == "admin" {
			return true
		}
	}
	return false
}

func (r *Role) Valid() bool {
	if len(r.roleType) == 0 {
		return false
	}

	for _, t := range r.roleType {
		if t == "unknown" {
			return false
		}
	}

	return true
}

func (r *Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(&r.roleType)
}

func (r *Role) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.roleType)
}
