package entities

import "time"

type Incident struct {
	ID          int       `json:"id"`
	Severity    string    `json:"severity" binding:"required"` // low, medium, high
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"required"` // open, closed
	Timestamp   time.Time `json:"timestamp" binding:"required"`
}

// Escalate the incident
func (i *Incident) Escalate(severity string) {
	// escalate the incident
	i.Severity = severity
	i.Timestamp = time.Now()
}
