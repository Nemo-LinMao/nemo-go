package models

import (
	"github.com/astaxie/beego/orm"
)

// Module  模块信息定义
type Module struct {
	Id			int
	Name    	string `orm:"unique";size(128)`
	Url      	string `orm:"size(256)"`
	Img      	string `orm:"size(256)"`
}

func (this *Module) TableName() string {
	return "module"
}

func (this *Module) Trunc() {
	o := orm.NewOrm()
	o.Using("defualt")

	o.Raw("trunc table t_" + this.TableName()).Exec()
}