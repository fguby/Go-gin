package util

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	Data       interface{}
}

//封装分页方法
func GetPage(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, Data: list}
}
