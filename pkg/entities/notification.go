package entities

import "time"

// Notification entity
type Notification struct {
	ID           int       `json:"id" binding:"required"`
	IncidentID   int       `json:"incident_id" binding:"required"`
	PersonnelID  int       `json:"personnel_id" binding:"required"`
	Severity     string    `json:"severity" binding:"required"`
	NotifyMethod string    `json:"notify_method" binding:"required"`
	Timestamp    time.Time `json:"timestamp" binding:"required"`
}

// Send the notification
func (n *Notification) Send() {
	// send the notification
	println("send the notification")
}
