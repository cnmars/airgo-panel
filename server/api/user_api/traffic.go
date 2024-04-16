package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

func GetSubTrafficList(ctx *gin.Context) {
	var params model.CustomerService
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	list, err := service.TrafficSvc.GetSubTrafficList(&params)
	if err != nil {
		response.Fail(err.Error(), nil, ctx)
		return
	}
	response.OK("Success", list, ctx)
}
