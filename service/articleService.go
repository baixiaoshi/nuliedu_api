package service

import (
	"nuliedu_api/TO/article"
	"nuliedu_api/models"
	"github.com/pkg/errors"
	"time"
	"fmt"
)

type ArticleService struct {
	BaseService
}

type ArticleList []*models.Article

//获取文章列表
func (a *ArticleService) GetArticleList(req *article.ListReq) ArticleList {

	list := models.GetArticlePage(req.Page, req.PageSize, req.Title, req.KeyWord)

	return list

}

//添加文章逻辑
//1.看是否有文章标题一样的文章
//2.直接往db里面插入好了
func (a *ArticleService) AddArticle(req *article.Request) error {

	fmt.Println("title=", req.Title)
	count, err := models.CountArticleByTitle(req.Title)
	fmt.Println("count=", count)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the article exists")
	}

	m := models.Article{}
	m.SeriesId = req.SeriesId
	m.Title = req.Title
	m.Desc = req.Desc
	m.Keyword = req.KeyWord
	m.Content = req.Content
	t := time.Now().Unix()
	m.CreateTime = int(t)
	m.UpdateTime = int(t)

	_, err = models.InsertArticle(&m)
	if err != nil {
		return err
	}

	return nil
}
