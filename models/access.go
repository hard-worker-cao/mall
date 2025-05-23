package models

type Access struct {
	Id          int
	ModuleName  string //模块名称
	ActionName  string //操作名称
	Type        int    //节点类型: 1、表示模块2、表示菜单3、操作
	Url         string //路由跳转地址
	ModuleId    int    //此module_id和当前模型的id关联  module_id= 0 表示模块
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id"`
	Checked     bool     `gorm:"-"` // 忽略本字段
}

func (Access) TableName() string {
	return "access"
}
