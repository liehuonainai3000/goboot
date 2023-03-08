package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liehuonainai3000/goboot/frame/web"
	"github.com/liehuonainai3000/goboot/internal/entity"
	"github.com/liehuonainai3000/goboot/internal/service"
)

type officeController struct{}

var Office = &officeController{}

func (o *officeController) RegitsRouter(r *gin.RouterGroup) {

	r.POST("/office/add", Office.Add)
	r.POST("/office/update", Office.Update)
	r.POST("/office/get", Office.Get)
	r.POST("/office/delete", Office.Delete)
	r.POST("/office/list", Office.List)
	r.POST("/office/page", Office.Page)
}

func (o *officeController) Add(c *gin.Context) {

	office := &entity.Office{}
	err := binder.BindBodyExcept(c, office, "Id")
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	err = service.OfficeService.Insert(office)
	if err != nil {
		logger.Errorf("Add office error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespSucc(c, "新增用户信息成功")
}

func (o *officeController) List(c *gin.Context) {
	office := &entity.Office{}
	err := binder.BindBody(c, office)
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, err := service.OfficeService.QueryList(office)
	if err != nil {
		logger.Errorf("get office error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespData(c, resp)
}

func (o *officeController) Page(c *gin.Context) {
	office := &entity.Office{}
	err := binder.BindBody(c, office)
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, total, err := service.OfficeService.FindPage(office)
	if err != nil {
		logger.Errorf("get office error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}
	web.RespData(c, web.PageResult[entity.Office]{
		List:     resp,
		PageInfo: office.GetPageInfo().SetRowCount(total),
	})

}

func (o *officeController) Get(c *gin.Context) {

	office := &entity.Office{}
	err := binder.BindBody(c, office, "Id")
	if err != nil {
		logger.Errorf("parameter bind error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	resp, err := service.OfficeService.GetByPk(office)
	if err != nil {
		logger.Errorf("get office error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}

	web.RespData(c, resp)
}

func (o *officeController) Update(c *gin.Context) {

	office := &entity.Office{}
	err := binder.BindBodyExcept(c, office)
	if err != nil {
		logger.Errorf("office parameter error:%v", err)
		web.RespErr(c, 501, err.Error())
		return
	}

	err = service.OfficeService.UpdateByPk(office)
	if err != nil {
		logger.Errorf("update office error:%v", err)
		web.RespErr(c, 500, err.Error())
		return
	}

	web.RespSucc(c, "修改机构信息成功")

}

func (o *officeController) Delete(c *gin.Context) {

	office := &entity.Office{}
	err := binder.BindBody(c, office, "Id")
	if err != nil {
		logger.Errorf("office parameter error:%v", err)
		web.RespErr(c, 1, err.Error())
		return
	}

	err = service.OfficeService.DeleteByPk(office)
	if err != nil {
		logger.Errorf("delete office error:%v", err)
		web.RespErr(c, 3, err.Error())
		return
	}

	web.RespSucc(c, "删除用户信息成功")

}
