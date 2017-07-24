package controllers

import (
	"github.com/astaxie/beego"
	"github.com/tinynemo/nemo-go/firstapp/models"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var modules1 []*models.Module
	o := orm.NewOrm()
	o.QueryTable("t_module").All(&modules1)
	this.Data["modules"] = modules1
	this.TplName = "index.html"
}

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	var image models.BingImage
	_, res := image.TestOrm()
	this.Data["images"] = res
	this.Data["title"] = "nemo"
	this.Data["json"] = res
	this.ServeJSON();
	//this.TplName = "hello.tpl"
}

type ImgcurController struct {
	beego.Controller
}

func (this *ImgcurController) Get() {
	this.Data["title"] = "Image Carousel"
	this.TplName = "test.html"
}

type SlideshowController struct {
	beego.Controller
}

func (this *SlideshowController) Get() {
	this.Data["title"] = "Gallery"
	this.TplName = "index-gallery.html"
}

type IndexImageController struct {
	beego.Controller
}

func (this *IndexImageController) Get() {
	this.Data["title"] = "Index Image"
	this.TplName = "index-image.html"
}

type IndexColorController struct {
	beego.Controller
}

func (this *IndexColorController) Get() {
	this.Data["title"] = "Index Color"
	this.TplName = "index-color.html"
}

type IndexVideoController struct {
	beego.Controller
}

func (this *IndexVideoController) Get() {
	this.Data["title"] = "Index Video"
	this.TplName = "index-video.html"
}