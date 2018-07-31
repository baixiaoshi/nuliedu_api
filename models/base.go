package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/url"
)

// 连接数据库
func init() {

	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s", dbuser, dbpassword, dbhost, dbport, dbname, url.QueryEscape(timezone))
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(Article))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}
