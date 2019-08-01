package controllers

import (
	"log"
	"strconv"
	"strings"
)

func init() {

}

// ============================================================================================================

type MainController struct {
	RequestController
}

func (c *MainController) PageInit() {
	var tags []*ProductTag
	_, err := OrmGetAll("product_tag").All(&tags)
	if err != nil {
		log.Fatal(err)
	} else {
		c.Data["ProductTags"] = tags
	}
}

func (c *MainController) PageLoad() {
	page := c.GetString(":page")
	if len(page) == 0 {
		page = "index"
	}

	c.PageInit()

	switch page {
	case "about":
		c.AboutController()
	case "contact":
		c.ContactController()
	case "fangwei":
		c.FangweiController()
	case "feedback":
		c.FeedbackController()
	case "index":
		c.IndexController()
	case "news":
		c.NewsController()
	case "news_detail":
		c.NewsDetailController()
	case "partners":
		c.PartnersController()
	case "products":
		c.ProductsController()
	case "product":
		c.ProductController()
	case "shouhou":
		c.ShouhouController()
	case "shouquan":
		c.ShouquanController()
	case "weishang":
		c.WeishangController()

	}

	c.TplName = page + ".html"
}

func (c *MainController) OtherPagesLoad() {
	URI := c.Ctx.Request.RequestURI
	if strings.Index(URI, "/") == 0 {
		URI = strings.TrimLeft(URI, "/")
	}
	c.TplName = URI
}

// ============================================================================================================

func (c *MainController) AboutController() {

}

// ============================================================================================================

func (c *MainController) ContactController() {

}

// ============================================================================================================

func (c *MainController) FangweiController() {

}

// ============================================================================================================

func (c *MainController) FeedbackController() {

}

// ============================================================================================================

func (c *MainController) IndexController() {
	var banners []*Banner
	OrmGetAll("banner").All(&banners)
	c.Data["Banners"] = banners

	var tags4 []*ProductTag
	tags4 = c.Data["ProductTags"].([]*ProductTag)
	count := len(tags4)
	if count > 4 {
		count = 4
	}
	tags4 = tags4[:count]

	var productsHot []map[string]string
	if len(tags4) != 0 {
		for index := 0; index < len(tags4); index++ {
			e := tags4[index]
			productsHot = append(productsHot, map[string]string{"Title": e.Name, "Pic": e.Pics, "URL": "/p_products?id=" + strconv.FormatInt(e.ID, 10), "Index": strconv.FormatInt(int64(index)+1, 10)})
		}
		c.Data["ProductsHot"] = productsHot
	}
}

// ============================================================================================================

func (c *MainController) NewsController() { // 新闻每页2篇
	qs := c.ValidPages(OrmGetAll("news"), 3, 2, "p_news")
	if qs != nil {
		var news []*News
		qs.All(&news)
		c.Data["News"] = news
	}
}

func (c *MainController) NewsDetailController() {
	ID, err := c.GetInt64("id")
	if err != nil {
		return
	}
	news := new(News)
	news.ID = ID
	err = OrmGetOne(news)
	if err == nil {
		c.Data["News"] = news
	}
}

// ============================================================================================================

func (c *MainController) PartnersController() {

}

// ============================================================================================================

func (c *MainController) ProductsController() {
	ID := c.GetString("id")
	if len(ID) == 0 {
		tags := c.Data["ProductTags"].([]*ProductTag)
		if len(tags) == 0 {
			return
		} else {
			tag := tags[0]
			ID = strconv.FormatInt(tag.ID, 10)
		}
	}

	c.Data["TagID"], _ = strconv.ParseInt(ID, 10, 64)

	var products []*Product
	qs := OrmGetAll("product")
	qs.Filter("TagID", ID).All(&products)

	c.Data["Products"] = products
}

func (c *MainController) ProductController() {
	ID := c.GetString("id")
	if len(ID) == 0 {
		return
	}
	product := new(Product)
	product.ID, _ = strconv.ParseInt(ID, 10, 64)
	OrmGetOne(product)

	c.Data["Product"] = &product
	c.Data["PicsDesc"] = strings.Split(product.PicsDesc, ",")
}

// ============================================================================================================

func (c *MainController) ShouhouController() {

}

// ============================================================================================================

func (c *MainController) ShouquanController() {

}

// ============================================================================================================

func (c *MainController) WeishangController() {

}

// ============================================================================================================

func (c *MainController) OtherPagesController() {

}
