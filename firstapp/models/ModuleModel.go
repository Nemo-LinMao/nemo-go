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

// InitModule 初始化主页的所有模块
func InitModule() {
	o := orm.NewOrm()
	o.Using("default")

	modules := []Module {   {Url:"slideshow", Name:"Memory Gallery", Img:"static/img/modules/slide.png"},
							{Url:"image", Name:"Images", Img:"static/img/modules/image.png"},
							{Url:"color", Name:"Color", Img:"static/img/modules/color.png"},
							{Url:"video", Name:"Video", Img:"static/img/modules/video.png"} }
	for _, val := range modules {
		o.Insert(&val)
	}
}