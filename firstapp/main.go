package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/tinynemo/nemo-go/firstapp/routers"

	"github.com/tinynemo/nemo-go/firstapp/models"
)

func Init() {
	// regist database
	orm.RegisterDataBase("default", "mysql", "root:root@/nemo_go?charset=UTF8")
	// init models
	models.InitModels()

	// regist static file path
	beego.SetStaticPath("/assets","assets")
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.TemplateLeft = "<<<"
	beego.BConfig.WebConfig.TemplateRight = ">>>"
	
	beego.Run()
}
