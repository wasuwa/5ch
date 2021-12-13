package model

import "time"

type base struct {
	id        uint
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
