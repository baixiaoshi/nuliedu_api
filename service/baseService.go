package service

var (
	Article ArticleService
)

//在此处初始化所有的service
func init() {

	Article = ArticleService{}

}

type BaseService struct {
}
