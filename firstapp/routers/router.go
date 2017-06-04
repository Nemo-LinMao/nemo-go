package routers

import (
	"github.com/tinynemo/nemo-go/firstapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
