package models

import (
	"github.com/astaxie/beego/orm"
)

func InitModels() {
	// 注册表
	orm.RegisterModelWithPrefix("t_", new(Users), new(BingImage), new(Module))
	
	// 同步表
	orm.RunSyncdb("default", false, true)

	// 初始化表	
	InitModule()
}