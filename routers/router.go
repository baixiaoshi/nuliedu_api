
// all route config
package routers

import (
	"nuliedu_api/controllers"
	"github.com/astaxie/beego"
)

func init() {

	//user
	beego.Router("/login", &controllers.UserController{}, "*:Login")

	// 默认登录
	beego.Router("/", &controllers.ArticleController{}, "*:Index")
	beego.Router("/article/add", &controllers.ArticleController{}, "*:Add")

	//文件上传

	beego.Router("/upload/img", &controllers.UploadController{}, "*:Img")

	//beego.Router("/home", &controllers.HomeController{}, "*:Index")
	//beego.Router("/home/start", &controllers.HomeController{}, "*:Start")
	//beego.Router("/home/ajaxallinfo", &controllers.HomeController{}, "*:AjaxAllInfo") //ajax刷新面板信息
}
