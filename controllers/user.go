package controllers

import (
	"github.com/astaxie/beego"
	"news/models"
	_ "github.com/go-sql-driver/mysql"
)

type UserController struct {
	beego.Controller
}
//显示注册页面
func (this *UserController)ShowReg()  {
    this.TplName = "register.html"
}

func (this *UserController)HandleRge() {
	name := this.GetString("userName")
	password := this.GetString("password")
	if name ==""||password ==""{
		beego.Info("用户密码错误！")
		this.TplName = "register.html"
		return
	}
//创建 user 对象
	user := models.User{}
	user.Name = name
	user.Password = password

	_,err :=(&user).InsertUser() //调用在model 改造后的 insert方法
	this.TplName = "content.html"
	if err != nil{
		beego.Info(err)
	}

}