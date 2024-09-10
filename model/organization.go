package model

import "time"

type OrganizationType int

const (
	IC OrganizationType = iota
	LLC
	JSC
)

type organization struct {
	Id               int64  `json:"id" db:"id"`
	Name             string `json:"name" db:"name"`
	Description      string `json:"description" db:"description"`
	OrganizationType OrganizationType
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type OrganizationResponsible struct {
	Id             int64 `json:"id" db:"id"`
	OrganizationId int64 `json:"organization_id" db:"organization_id"`
	UserId         int64 `json:"user_id" db:"user_id"`
}
