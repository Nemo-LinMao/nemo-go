package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Data["title"] = "nemo"
	this.TplName = "home.html"

}
