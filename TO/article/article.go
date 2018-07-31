package article

type ListReq struct {
	Title    string
	KeyWord  string
	Page     int
	PageSize int
}

type Request struct {
	SeriesId int
	Title    string
	Desc     string
	KeyWord  string
	Content  string
}
