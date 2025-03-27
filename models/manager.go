package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	//关联role表单的Id
	Role Role `gorm:"foreignkey:RoleId;references:Id"`
}

func (Manager) TableName() string {
	return "manager"
}
