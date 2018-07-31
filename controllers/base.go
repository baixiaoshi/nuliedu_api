package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

//返回json
func (b *BaseController) Response(success bool, message string, data interface{}) {
	b.Data["json"] = map[string]interface{}{
		"success": success,
		"message": message,
		"data":    data,
	}
	b.ServeJSON()
}
