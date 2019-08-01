package routers

import (
	"MAYMOE/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{}, "get:PageLoad")
	beego.Router("/p_:page:string", &controllers.MainController{}, "get:PageLoad")

	beego.Router("/admin/?:page", &controllers.AdminController{}, "get:PageLoad")

	beego.Router("/api/ImgUpload", &controllers.AdminController{}, "post:ImgUpload")   // 上传图片
	beego.Router("/api/ImgsUpload", &controllers.AdminController{}, "post:ImgsUpload") // 上传多个图片

	beego.Router("/api/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/api/upload/?:name", &controllers.AdminController{}, "post:Upload") // 上传
	beego.Router("/api/update/?:name", &controllers.AdminController{}, "post:Update") //
	beego.Router("/api/delete/?:name", &controllers.AdminController{}, "post:Delete") //

	beego.Router("/baidu_verify_NRX92oGNK8.html", &controllers.MainController{}, "get:OtherPagesLoad") // 百度
}
