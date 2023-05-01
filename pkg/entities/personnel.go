package entities

type OnCallPersonnel struct {
	ID           int    `json:"id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	Email        string `json:"email" binding:"required"`
	NotifyMethod string `json:"notify_method" binding:"required"`
}

var onCallPersonnel []OnCallPersonnel

func (p *OnCallPersonnel) Create() OnCallPersonnel {
	onCallPersonnel = append(onCallPersonnel, *p)
	return *p
}

func GetOncallPersonnelByID(id int) OnCallPersonnel {
	for i := range onCallPersonnel {
		if onCallPersonnel[i].ID == id {
			return onCallPersonnel[i]
		}
	}
	return OnCallPersonnel{}
}
