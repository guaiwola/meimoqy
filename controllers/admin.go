package controllers

import (
	"MAYMOE/utils"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ACCOUNT  = "admin"
	PASSWORD = "654321"
)

var ImgClassifys utils.ArrayInterface

// ============================================================================================================

func init() {
	ImgClassifys = utils.ArrayInterface{"product", "banner", "news"}
}

// ============================================================================================================

type AdminController struct {
	RequestController
}

func (c *AdminController) PageInit() {
	fmt.Print("")
}

func (c *AdminController) PageLoad() {
	page := c.GetString(":page")
	if len(page) == 0 {
		page = "index"
	}

	c.PageInit()

	switch page {
	case "index":
		c.IndexController()
	case "login":
		c.LoginController()
	case "upload-product":
		c.UploadProdController()
	case "upload-tag":
		c.UploadTagController()
	case "products":
		c.ProductsController()
	case "tags":
		c.TagsController()
	case "upload-news":
		c.UploadNewsController()
	case "news":
		c.NewsController()
	case "upload-banner":
		c.UploadBannerController()
	case "banners":
		c.BannersController()
		// case "products":
		// 	c.ProductsController()
		// case "product":
		// 	c.ProductController()
		// case "shouhou":
		// 	c.ShouhouController()
		// case "shouquan":
		// 	c.ShouquanController()
		// case "weishang":
		// 	c.WeishangController()
	}

	if page != "login" {
		a := c.GetSession("skinmesh")
		if a != "2" {
			c.Redirect("/admin/login", 302)
			return
		}
	}

	c.TplName = "admin/" + page + ".html"
}

// ============================================================================================================

func (c *AdminController) LoginController() {
}

func (this *AdminController) Login() {
	u := AdminAccount{}
	if err := this.ParseForm(&u); err != nil {
		log.Fatal(err)
	} else {
		if u.Account == ACCOUNT && u.Password == PASSWORD {
			this.SetSession("skinmesh", "2")
			this.Redirect("/admin/index", 302)
		} else {
			this.Ctx.WriteString("用户名或密码错误")
		}
	}
}

// ============================================================================================================

func (c *AdminController) ImgUpload() {
	result := new(UoloadImgResult)
	f, h, err := c.GetFile("image")
	defer f.Close()

	classify := c.GetString("classify")

	if f == nil || !ImgClassifys.InArray(classify) {
		result.Code = Failure
		result.Msg = "接口错误"
	} else {
		if err != nil {
			log.Fatal("getfile err:", err)
		} else {
			filePath := "static/upload/" + classify + "/" + utils.GetNowName(h.Filename)
			err := c.SaveToFile("image", filePath)
			if err != nil {
				log.Fatal("save image err:", err)
				result.Code = Failure
				result.Msg = "上传失败"
			} else {
				result.Code = Success
				result.Msg = "上传成功"
				result.URL = filePath
			}
		}
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

func (c *AdminController) ImgsUpload() {
	result := new(UoloadImgsResult)
	result.URLs = []string{}

	fs, err := c.GetFiles("images")
	classify := c.GetString("classify")

	if len(fs) == 0 || !ImgClassifys.InArray(classify) {
		result.Code = Failure
		result.Msg = "接口错误"
	} else {
		for _, fh := range fs {
			file, _ := fh.Open()
			filePath := "static/upload/" + classify + "/" + utils.GetNowName(fh.Filename)
			f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				log.Fatal("打开文件错误")
			}
			defer f.Close()
			_, err1 := io.Copy(f, file)
			if err1 != nil {
				log.Fatal("getfile err:", err)
			} else {
				result.URLs = append(result.URLs, "/"+filePath)
			}
		}
		if len(result.URLs) == len(fs) {
			result.Code = Success
			result.Msg = "上传成功"
		} else {
			log.Fatal("productImage err:", err)
			result.Code = Failure
			result.Msg = "上传失败"
		}
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

// ============================================================================================================

func (c *AdminController) Upload() {
	name := c.GetString(":name")
	if len(name) == 0 {
		return
	}
	var model interface{}
	var tips string

	switch name {
	case "product":
		model = new(Product)
		tips = "上传商品成功"
	case "productTag":
		model = new(ProductTag)
		tips = "上传商品分类成功"
	case "news":
		temp := new(News)
		temp.Time = time.Now().Format("2006-01-02 15:04:05")
		model = temp
		tips = "上传新闻成功"
	case "banner":
		model = new(Banner)
		tips = "上传轮播图成功"
	}

	c.RequestUploadForm(model, tips)
}

func (c *AdminController) Update() {
	name := c.GetString(":name")
	if len(name) == 0 {
		return
	}
	var model interface{}
	var tips string

	switch name {
	case "product":
		model = new(Product)
		tips = "修改商品成功"
	case "news":
		temp := new(News)
		temp.Time = time.Now().Format("2006-01-02 15:04:05")
		model = temp
		tips = "修改新闻成功"
	}

	c.RequestUpdateForm(model, tips)
}

func (c *AdminController) Delete() {
	name := c.GetString(":name")
	if len(name) == 0 {
		return
	}
	ID := c.GetString("ID")
	if len(ID) == 0 {
		return
	}

	var model interface{}
	var tips string
	ID64, _ := c.GetInt64("ID") //utils.StringToInt64(ID)

	switch name {
	case "product":
		temp := new(Product)
		temp.ID = ID64
		model = temp
		tips = "删除商品成功"
	case "productTag":
		qs := OrmGetAll("product")
		TagID := utils.StringToInt64(ID)
		if count, _ := qs.Filter("TagID", TagID).Count(); count == 0 {
			temp := new(ProductTag)
			temp.ID = ID64
			model = temp
			tips = "删除商品分类成功"
		} else {
			res := OrmResult{}
			res.ID = -1
			res.Msg = "该分类下还有商品"
			res.Code = Failure
			c.Data["json"] = &res
			c.ServeJSON()
			return
		}
	case "news":
		temp := new(News)
		temp.ID = ID64
		model = temp
		tips = "删除新闻成功"
	case "banner":
		temp := new(Banner)
		temp.ID = ID64
		model = temp
		tips = "删除轮播图成功"
	}

	c.RequestDeleteForm(model, tips)
}

// ============================================================================================================

func (c *AdminController) IndexController() {

}

// ============================================================================================================

func (c *AdminController) UploadTagController() {

}

func (c *AdminController) UploadProdController() {
	var tags []*ProductTag
	_, err := OrmGetAll("product_tag").All(&tags)
	if err != nil {
		log.Fatal(err)
	} else {
		c.Data["ProductTags"] = tags
	}

	ID := c.GetString("id")
	if len(ID) != 0 {
		product := new(Product)
		product.ID, _ = strconv.ParseInt(ID, 10, 64)
		err := OrmGetOne(product)
		if err != nil {
			return
		}
		c.Data["Product"] = &product
		c.Data["PicsDesc"] = strings.Split(product.PicsDesc, ",")
	}
}

func (c *AdminController) UploadNewsController() {
	ID, err := c.GetInt64("id")
	if err == nil {
		news := new(News)
		news.ID = ID
		err := OrmGetOne(news)
		if err != nil {
			return
		}
		c.Data["News"] = &news
	}
}

func (c *AdminController) UploadBannerController() {
	var products []*Product
	_, err := OrmGetAll("product").All(&products)
	if err != nil {
		log.Fatal(err)
	} else {
		c.Data["Products"] = products
	}
}

// ============================================================================================================
const (
	maxPage = 5
	numPage = 10 // 一页10条
)

func (c *AdminController) ProductsController() {
	qs := c.ValidPages(OrmGetAll("product"), maxPage, numPage, "product")
	if qs != nil {
		var products []*Product
		qs.All(&products)
		c.Data["Products"] = products
		c.Data["Count"] = len(products)
	}
}

// ============================================================================================================

func (c *AdminController) TagsController() {
	qs := c.ValidPages(OrmGetAll("product_tag"), maxPage, numPage, "tags")
	if qs != nil {
		var tags []*ProductTag
		qs.All(&tags)
		c.Data["ProductTags"] = tags
		c.Data["Count"] = len(tags)
	}
}

// ============================================================================================================

func (c *AdminController) NewsController() {
	qs := c.ValidPages(OrmGetAll("news"), maxPage, numPage, "news")
	if qs != nil {
		var news []*News
		qs.All(&news)
		c.Data["News"] = news
		c.Data["Count"] = len(news)
	}
}

func (c *AdminController) BannersController() {
	qs := c.ValidPages(OrmGetAll("banner"), maxPage, numPage, "banners")
	if qs != nil {
		var banners []*Banner
		qs.All(&banners)
		c.Data["Banners"] = banners
		c.Data["Count"] = len(banners)
	}
}
