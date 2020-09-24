package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "276502494@qq.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Post()  {
	name := c.Ctx.Request.FormValue("name")
	age := c.Ctx.Request.FormValue("age")
	sex :=c.Ctx.Request.FormValue("sex")
	fmt.Println(name,age,sex)
	
}

func (c * MainController) post() {
	for i := 0;i<10; i++ {
		fmt.Printf("",i)
	}
}