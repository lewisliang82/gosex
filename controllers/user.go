/**
网站入口
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "gosex.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["AppName"] = beego.AppConfig.String("appname")
	c.TplName = "user.tpl"
}
