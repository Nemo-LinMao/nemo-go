package routers

import (
	"github.com/astaxie/beego"
	"github.com/tinynemo/nemo-go/firstapp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/helloworld", &controllers.HelloController{})
	beego.Router("/gallery", &controllers.ImgcurController{})
	beego.Router("/crawlpic", &controllers.BingImgController{})
	beego.Router("/slideshow", &controllers.SlideshowController{})
	beego.Router("/image", &controllers.IndexImageController{})
	beego.Router("/color", &controllers.IndexColorController{})
	beego.Router("/video", &controllers.IndexVideoController{})
}
