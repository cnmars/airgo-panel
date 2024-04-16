package admin_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

// 新建折扣
func NewCoupon(ctx *gin.Context) {
	var coupon model.Coupon
	err := ctx.ShouldBind(&coupon)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminCouponSvc.NewCoupon(&coupon)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("NewCoupon error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("NewCoupon success", nil, ctx)
}

// 删除折扣
func DeleteCoupon(ctx *gin.Context) {
	var coupon model.Coupon
	err := ctx.ShouldBind(&coupon)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminCouponSvc.DeleteCoupon(&coupon)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("DeleteCoupon error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("DeleteCoupon success", nil, ctx)
}

// 更新折扣
func UpdateCoupon(ctx *gin.Context) {
	var coupon model.Coupon
	err := ctx.ShouldBind(&coupon)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminCouponSvc.UpdateCoupon(&coupon)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("UpdateCoupon error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("UpdateCoupon success", nil, ctx)
}

// 获取折扣列表
func GetCouponList(ctx *gin.Context) {
	res, err := service.AdminCouponSvc.GetCouponList()
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetCouponList error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetCouponList success", res, ctx)
}
