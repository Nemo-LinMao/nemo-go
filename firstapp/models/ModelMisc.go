package models

import (
	"github.com/astaxie/beego/orm"
)

// Init 初始化主页的所有模块
func InitModules() {
	o := orm.NewOrm()
	o.Using("default")
	new(Module).Trunc()
	
	modules := []Module {   {Url:"slideshow", Name:"Memory Gallery", Img:"static/img/modules/slide.png"},
							{Url:"image", Name:"Images", Img:"static/img/modules/image.png"},
							{Url:"color", Name:"Color", Img:"static/img/modules/color.png"},
							{Url:"video", Name:"Video", Img:"static/img/modules/video.png"},
							{Url:"video", Name:"Video1", Img:"static/gallery1/AddoElephants_ZH-CN13744097225_1920x1080.jpg"}, }
	for _, val := range modules {
		o.Insert(&val)
	}
}

func InitModels() {
	// 注册表
	orm.RegisterModelWithPrefix("t_", new(Users), new(BingImage), new(Module))
	
	// 同步表
	orm.RunSyncdb("default", false, true)

	// 初始化表	
	InitModules()
}