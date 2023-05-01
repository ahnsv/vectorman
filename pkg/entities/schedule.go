package entities

import "time"

// On-call schedule entity
type OnCallSchedule struct {
	ID       int       `json:"id"`
	Start    time.Time `json:"start" binding:"required"`
	End      time.Time `json:"end" binding:"required"`
	Rotation []int     `json:"rotation" binding:"required"`
	TimeZone string    `json:"timezone" binding:"required"`
}
