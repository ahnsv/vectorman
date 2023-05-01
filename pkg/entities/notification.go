package entities

import (
	"fmt"
	"time"
)

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
	// get personnel phone number and email
	println("get personnel phone number and email")
	personnel := GetOncallPersonnelByID(n.PersonnelID)
	if personnel == (OnCallPersonnel{}) {
		println("personnel not found")
		return
	}

	// send the notification
	fmt.Printf("send the notification %s to %s:%s:%s", n.Severity, personnel.Name, personnel.NotifyMethod, personnel.Phone)
}
