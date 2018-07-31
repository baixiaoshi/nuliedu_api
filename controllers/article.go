package controllers

import (
	"nuliedu_api/TO/article"
	"github.com/astaxie/beego/validation"
	"nuliedu_api/service"
)

// 文章控制器
type ArticleController struct {
	BaseController
}

func (a *ArticleController) Index() {
	a.Response(false, "hello 小白", nil)
}

func (a *ArticleController) GetList() {

	title := a.GetString("title", "")
	keyword := a.GetString("keyword", "")
	page, err := a.GetInt("page", 1)

	if err != nil {
		a.Response(false, err.Error(), nil)
		return
	}
	pageSize, err := a.GetInt("pageSize", 20)
	if err != nil {
		a.Response(false, err.Error(), nil)
		return
	}

	req := article.ListReq{}
	req.Title = title
	req.KeyWord = keyword
	req.Page = page
	req.PageSize = pageSize

}

func (a *ArticleController) Add() {
	seriesID, err := a.GetInt("series_id", 0)
	if err != nil {
		a.Response(false, err.Error(), nil)
		return
	}
	title := a.GetString("title", "")
	desc := a.GetString("desc", "")
	keyword := a.GetString("keyword", "")
	content := a.GetString("content", "")

	req := article.Request{}
	req.SeriesId = seriesID
	req.Title = title
	req.Desc = desc
	req.KeyWord = keyword
	req.Content = content


	//参数校验
	valid := validation.Validation{}
	valid.Min(req.SeriesId, 0, "series_id")
	valid.Required(req.Title, "title")
	valid.Range(len(req.Title), 1, 256, "title_len")
	valid.Required(req.Desc, "desc")
	valid.Required(req.KeyWord, "keyword")
	valid.Required(req.Content, "content")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			a.Response(false, err.Key+":"+err.Message, nil)
			return
		}
	}

	//添加文章
	err = service.Article.AddArticle(&req)
	if err != nil {
		a.Response(false, "添加文章失败:"+err.Error(), nil)
		return
	}

	a.Response(true, "ok", nil)
}
