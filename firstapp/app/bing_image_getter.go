package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/tinynemo/nemo-go/firstapp/models"
)

func getImage(iBefore int, strSavePath string) {
	for i := 0; i < iBefore; i++ {
		baseURL := "http://cn.bing.com/HPImageArchive.aspx?format=js&idx=" + fmt.Sprintf("%d", i) + "&n=2&nc=1361089515117&FORM=HYLH1"
		resp, _ := http.Get(baseURL)

		html, err := ioutil.ReadAll(resp.Body)
		if nil == err {
			json, _ := simplejson.NewJson(html)
			images := json.Get("images").MustArray()
			fmt.Println(images) // TODO: to logger
			for _, v := range images {
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
			}
		} else {
			fmt.Println(err, string(html))
		}
	}

}

func main() {
	getImage(1, "")
}
