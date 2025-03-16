package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  string
	IsSuper  bool
}

func (Manager) TableName() string {
	return "manager"
}
