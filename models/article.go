package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Article struct {
	Id         int    `orm:"column(id)"`
	SeriesId   int    `orm:"column(series_id)"`
	Title      string `orm:"column(title)"`
	Desc       string `orm:"column(desc)"`
	Keyword    string `orm:"column(keyword)"`
	Content    string `orm:"column(content)"`
	CreateTime int    `orm:"column(create_time)"`
	UpdateTime int    `orm:"column(update_time)"`
}

//获取表名
func (a *Article) TableName() string {
	return "nl_article"
}

//获取用户分页
func GetArticlePage(page, pageSize int, title, keyword string) []*Article {


	var m []*Article
	qs := orm.NewOrm().QueryTable("article")
	if title != "" {
		qs = qs.Filter("title__icontains", title)
	}
	if keyword != "" {
		qs = qs.Filter("keyword__icontains", keyword)
	}
	offset := (page - 1) * pageSize
	qs = qs.Offset(offset)
	qs = qs.Limit(pageSize).OrderBy("-update_time")
	qs.All(&m)

	return m
}

//插入文章
func InsertArticle(m *Article) (int, error) {
	lastId, err := orm.NewOrm().Insert(m)

	return int(lastId), err
}

//通过标题统计
func CountArticleByTitle(title string) (int64, error) {

	fmt.Println("mytitle=", title)
	q := orm.NewOrm().QueryTable("nl_article")
	q = q.Filter("title", title)
	count, err := q.Count()
	return count, err
}
