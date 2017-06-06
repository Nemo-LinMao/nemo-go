package routers

import (
	"github.com/astaxie/beego"
	"github.com/tinynemo/nemo-go/firstapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/helloworld", &controllers.HelloController{})
}
