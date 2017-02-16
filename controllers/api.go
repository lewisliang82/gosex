/**
网站入口
 */
package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"html/template"
	"fmt"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) Get() {
	//this.Ctx.Output.Body([]byte(this.URLFor(".Myext")))

	this.Data["Website"] = "gosex.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["AppName"] = beego.AppConfig.String("appname")

	Id := this.Ctx.Input.Param(":id")
	logs.Info("id %s", Id)

	this.Data["Id"] = Id
	this.Data["xsrfdata"]=template.HTML(this.XSRFFormHTML())
	fmt.Println("error1")

	logs.Debug("============")
	logs.Debug(this.URLFor("UserController.Get"))

	v := this.GetSession("asta")
	if v == nil {
		this.SetSession("asta", int(1))
		this.Data["num"] = 0
	} else {
		this.SetSession("asta", v.(int)+1)
		this.Data["num"] = v.(int)

		logs.Debug(v.(int))
	}

	this.TplName = "user.tpl"
}

func (this *ApiController) ListFood() {
	this.Data["Website"] = "gosex.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Data["AppName"] = beego.AppConfig.String("appname")


	this.Data["Id"] = "listfood"

	this.TplName = "user.tpl"
}
