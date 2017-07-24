package models

import (
	"database/sql"
	"fmt"

	"io/ioutil"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BingImage struct {
	Id		   int
	UrlPath    string
	Name       string
	BelongDate string
	Desc       string
}

func (image *BingImage) Save() bool {
	path := fmt.Sprintf("/mnt/f/BingDesktop/%s", image.Name)
	resp, err := http.Get(image.UrlPath)
	pic, err := ioutil.ReadAll(resp.Body)

	fout, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if nil != err {
		fmt.Println("open file error", err)
		return false
	}
	defer fout.Close()

	_, err = fout.Write(pic)
	if nil != err {
		fmt.Println("save pic 2 disk error", err)
		return false
	}

	return true
}

func (image *BingImage) Warehouse() bool {
	db, err := sql.Open("mysql", "root:root@/nemo_go?charset=UTF8")
	if nil != err {
		fmt.Println("open database error", err)
		return false
	}
	stmt, err := db.Prepare(` insert into t_bing_image(url_path, image_name, belong_date, description) 
					         values(?,?,?,?) on duplicate 	key update image_name=?`)
	if nil != err {
		fmt.Println("prepare error", err)
		return false
	}
	stmt.Exec(image.UrlPath, image.Name, image.BelongDate, image.Desc, image.Name)

	return true
}

type ImageTest struct {
	Id  	int
	ImageName  	string
}

func (image *BingImage) TestOrm() (int, []*ImageTest) {
	o := orm.NewOrm()
	//var r RawSeter
	result := make([]*ImageTest, 10)
	r := o.Raw("select id, image_name from t_bing_image limit 10")
	num, err := r.QueryRows(&result)
	if nil != err {
		fmt.Println("query error", err)
	}
	return int(num), result
}