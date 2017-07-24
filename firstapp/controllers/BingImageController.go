package controllers

import (
	"fmt"	
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/tinynemo/nemo-go/firstapp/models"
	simplejson "github.com/bitly/go-simplejson"
	"sync"
)

type BingImgController struct {
	beego.Controller
}

var imageTotal = 0
var imageCount = 0
var ch = make(chan int, 20)
var mutex sync.Mutex

func (this *BingImgController) Post() {
	iBefore := 8
	imageTotal = iBefore
	imageCount = 0

	go func() {
		for i := 0; i < iBefore; i++ {
			baseURL := "http://cn.bing.com/HPImageArchive.aspx?format=js&idx=" + fmt.Sprintf("%d", i) + "&n=1&nc=1361089515117&FORM=HYLH1"
			resp, _ := http.Get(baseURL)

			html, err := ioutil.ReadAll(resp.Body)
			if nil == err {
				json, _ := simplejson.NewJson(html)
				images := json.Get("images").MustArray()
				//fmt.Println(images) // TODO: to logger
				for _, v := range images {
					go func() {
						var image models.BingImage
						vMap := v.(map[string]interface{})
						image.UrlPath = fmt.Sprintf("http://s.cn.bing.net/%s", vMap["url"].(string))
						temp := strings.Split(vMap["url"].(string), "/")
						image.Name = temp[len(temp)-1]
						image.BelongDate = vMap["enddate"].(string)
						image.Desc = vMap["copyright"].(string)
						fmt.Println(image)
					
						image.Save()
						image.Warehouse()

						ch <- 1
					}()
					go func() {
						mutex.Lock()
						defer mutex.Unlock()
						
						imageCount += <- ch
						// fmt.Println(imageCount)
					}()
				}
			} else {
				fmt.Println(err, string(html))
			}
		}
	}()

	this.Data["json"] = map[string]int{ "result": 0, "total": iBefore }
	this.ServeJSON();
}

func (this *BingImgController) Get() {
	this.Data["json"] = map[string]int{ "result": 0, "cnt": imageCount}
	this.ServeJSON();
}