// API返回值
//

package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type ResultCode int

const ( // 返回状态
	Success ResultCode = iota
	Failure
	RemoteFailure
	SessionExpired
)

type Result struct {
	Code ResultCode
	Msg  string
}

type UoloadImgResult struct {
	Result
	URL string
}

type UoloadImgsResult struct {
	Result
	URLs []string
}

type OrmResult struct {
	Result
	ID int64
}

type AdminAccount struct {
	Account  string `form:"account"`
	Password string `form:"password"`
	Remember bool   `form:"remember"`
}

type RequestController struct { // 请求类基础
	beego.Controller
}

func (c *RequestController) RequestUploadForm(model interface{}, msgSuccess string) { // 通用上传表单方法
	res := OrmResult{}
	if err := c.ParseForm(model); err != nil {
		log.Fatal(err)
		res.Code = Failure
		res.Msg = err.Error()
		res.ID = -1
	} else {
		ID, err := OrmAdd(model)
		if err != nil {
			res.Code = Failure
			res.Msg = err.Error()
			res.ID = -1
		} else {
			res.Code = Success
			res.Msg = msgSuccess
			res.ID = ID
		}
	}

	c.Data["json"] = &res
	c.ServeJSON()
}

func (c *RequestController) RequestUpdateForm(model interface{}, msgSuccess string) {
	res := OrmResult{}
	if err := c.ParseForm(model); err != nil {
		log.Fatal(err)
		res.Code = Failure
		res.Msg = err.Error()
		res.ID = -1
	} else {
		ID, err := OrmUpdate(model)
		if err != nil {
			res.Code = Failure
			res.Msg = err.Error()
			res.ID = -1
		} else {
			res.Code = Success
			res.Msg = msgSuccess
			res.ID = ID
		}
	}

	c.Data["json"] = &res
	c.ServeJSON()
}

func (c *RequestController) RequestDeleteForm(model interface{}, msgSuccess string) {
	res := OrmResult{}
	ID, err := OrmDelete(model)
	if err != nil {
		res.Code = Failure
		res.Msg = err.Error()
		res.ID = -1
	} else {
		res.Code = Success
		res.Msg = msgSuccess
		res.ID = ID
	}

	c.Data["json"] = &res
	c.ServeJSON()
}

// ========================================================================================================================

type Page struct {
	Page int64
	Curr bool
	URL  string
}

// 模板里面按range .Pages的Page来显示页面，上一页和下一页分别是.PreURL和.NextURL
func (c *RequestController) ValidPages(qs orm.QuerySeter, maxPage, num int64, pageURL string) orm.QuerySeter {
	if maxPage < 3 {
		log.Fatalln("最小页码要有3")
	}

	curPage, err := c.GetInt64("page")
	if err != nil {
		curPage = 0
	}
	count, _ := qs.Count()

	if curPage*num+1 > count {
		fmt.Println(pageURL + "超过内容数量上限")
		return nil
	}

	var pages []Page

	for index := curPage - maxPage/2; index < curPage && index >= 0; index++ {
		URL := pageURL + "?page=" + strconv.FormatInt(index, 10)
		if index == curPage-1 {
			c.Data["PreURL"] = URL
		}
		pages = append(pages, Page{Page: index, Curr: false, URL: URL})
	}

	pages = append(pages, Page{Page: curPage, Curr: true, URL: pageURL + "?page=" + strconv.FormatInt(curPage, 10)})

	for index := curPage + 1; index <= curPage+maxPage/2; index++ {
		if index*num+1 > count {
			break
		}
		URL := pageURL + "?page=" + strconv.FormatInt(index, 10)
		if index == curPage+1 {
			c.Data["NextURL"] = URL
		}
		pages = append(pages, Page{Page: index, Curr: false, URL: URL})
	}

	c.Data["Pages"] = pages

	return qs.Limit(num, curPage*num)
}
