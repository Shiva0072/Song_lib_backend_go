package entities

import "time"

type Entity interface {
	GetId() uint
	GetUpdatedAt() time.Time
	GetCreatedAt() time.Time
	GetStatus() string
}
