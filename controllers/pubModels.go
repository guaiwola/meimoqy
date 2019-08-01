// 前端models
//

package controllers

import (
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

type ModelOrmBase struct {
	ID int64 `form:"ID"`
}

type ProductTag struct {
	ModelOrmBase
	Name string `form:"product-tag"`
	Pics string `form:"pics"`
}

type Product struct {
	ModelOrmBase
	Name     string  `form:"product-name"`
	Pics     string  `form:"product-pics"`
	Price    float32 `form:"product-price"`
	Desc     string  `orm:"type(text)" form:"product-desc"`
	PicsDesc string  `orm:"type(text)" form:"product-pics-desc"`
	TagID    int64   `form:"product-tag"`
	TagName  string  `form:"product-tag-name"`
}

type ProductGroup struct {
	Product
	Products []*Product `orm:"rel(m2m)"`
}

// type News struct {
// 	ModelOrmBase
// 	Name     string `form:"name"`
// 	Time     string `form:"-"`
// 	Desc     string `form:"desc"`
// 	Pics     string `form:"pics"`
// 	PicsDesc string `orm:"type(text)" form:"pics-desc"`
// }

type News struct {
	ModelOrmBase
	Name string `form:"name"`
	Time string `form:"-"`
	Desc string `orm:"type(text)" form:"desc"`
}

type Banner struct {
	ModelOrmBase
	Pics string `form:"pics"`
	URL  string `form:"URL"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	var regSQL string
	if beego.BConfig.RunMode == "dev" {
		regSQL = "root:123456@/GOORM?charset=utf8mb4"
	} else {
		regSQL = "goorm:b3f2l^$63Tl2OI8O@tcp(rm-uf648aptvfl0ocgqg.mysql.rds.aliyuncs.com)/goorm?charset=utf8mb4"
	}
	orm.RegisterDataBase("default", "mysql", regSQL)

	orm.RegisterModel(new(ModelOrmBase), new(Product), new(ProductGroup), new(News), new(ProductTag), new(Banner))
	orm.RunSyncdb("default", false, true)
}

// ========================================================================================================================

// func OrmGetAll(result []interface{}, table string, num, page int64) error {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable(table)

// 	var offset int64
// 	offset = 0
// 	var err error
// 	var index int64 = 0
// 	for ; index < page; index++ {
// 		var temp []interface{}
// 		offset, err = qs.Limit(num, offset).All(&temp)
// 		result = append(result, temp...)
// 		if err != nil || offset <= num*index+num {
// 			break
// 		}
// 	}

// 	return err
// }

func OrmGetAll(table string) orm.QuerySeter {
	o := orm.NewOrm()
	return o.QueryTable(table)
}

func OrmGetOne(this interface{}) error {
	o := orm.NewOrm()
	err := o.Read(this, "ID")
	return err
}

func OrmAdd(this interface{}) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(this)
	if err != nil {
		return -1, err
	}

	return id, err
}

func OrmDelete(this interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Delete(this)
}

func OrmUpdate(this interface{}) (int64, error) {
	o := orm.NewOrm()
	return o.Update(this)
}

// ========================================================================================================================
