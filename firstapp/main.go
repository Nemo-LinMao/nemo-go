package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/tinynemo/nemo-go/firstapp/routers"
)

type t_users struct {
	Id         int
	Name       string
	Addr       string
	Birth      string
	Createtime string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/nemo_go?charset=UTF8")
	orm.RegisterModel(new(t_users))
}
func main() {
	o := orm.NewOrm()
	o.Using("default")

	u := t_users{Id: 1}
	err := o.Read(&u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
	beego.Run()
}
