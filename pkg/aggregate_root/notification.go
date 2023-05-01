package aggregate_root

import "github.com/ahnsv/vectorman/pkg/entities"

// Notification aggregate root
type NotificationRoot struct {
	notification *entities.Notification
}

func (n *NotificationRoot) Send() NotificationRoot {
	n.notification.Send()
	return *n
}
