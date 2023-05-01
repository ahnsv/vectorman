package aggregate_root

import "github.com/ahnsv/vectorman/pkg/entities"

// On-call schedule aggregate root
type OnCallScheduleRoot struct {
	schedule  *entities.OnCallSchedule
	personnel []*entities.OnCallPersonnel
}
