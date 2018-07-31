package models

import "github.com/astaxie/beego/orm"

type Imgs struct {
	Id         int    `orm:"column(id)"`
	ImgPath    string `orm:"column(imgpath)"`
	QiNiuPath     string `orm:"column(qiniu_path)"`
	ImgWidth   int    `orm:"column(img_width)"`
	ImgHeight  int    `orm:"column(img_height)"`
	ImgSize    int    `orm:"column(img_size)"`
	IsDelete   int    `orm:"column(is_delete)"`
	CreateTime int    `orm:"column(create_time)"`
	UpdateTime int    `orm:"column(update_time)"`
}


func InsertImgs(m *Imgs) (int, error) {

	lastId, err := orm.NewOrm().Insert(m)
	return int(lastId), err
}
