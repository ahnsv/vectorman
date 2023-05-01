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

var onCallSchedules []OnCallSchedule

// Get current on-call schedule
func GetCurrentOncallSchedule() OnCallSchedule {
	// based on time.Now() and rotation, find the current on-call schedule
	for i := range onCallSchedules {
		if onCallSchedules[i].Start.Before(time.Now()) && onCallSchedules[i].End.After(time.Now()) {
			return onCallSchedules[i]
		}
	}
	return OnCallSchedule{}
}
