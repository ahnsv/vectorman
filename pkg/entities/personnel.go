package entities

type OnCallPersonnel struct {
	ID           int    `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NotifyMethod string `json:"notify_method" binding:"required"`
}
