package aggregate_root

import "github.com/ahnsv/vectorman/pkg/entities"

// Incident aggregate root
type IncidentRoot struct {
	incident *entities.Incident
}

// write a service of incident root to establish the business logic
func (g *IncidentRoot) Escalate(severity string) IncidentRoot {
	// escalate the incident
	g.incident.Escalate(severity)
	return *g
}
