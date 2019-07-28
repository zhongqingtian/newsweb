package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)
//声明全局变量
var   db orm.Ormer
type User struct{
	//首写大写 全局 才被别的文件调用
	Id int
	Name string
	Password string
}

func init()  {
	orm.Debug = true
	orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/newsweb?charset=utf8",30)

    orm.RegisterModel(new(User))
	orm.RunSyncdb("default",true,false) //生成数据表 verbose = true 表示检查
	// 是否有这个表，然后创建
	db = orm.NewOrm()
}
//插入用户账号
func (this *User)InsertUser() (n int64,err error) {
	n,err = db.Insert(this)
	return
}