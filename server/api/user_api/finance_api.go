package user_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/api"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

func GetBalanceStatementList(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	uIDInt, ok := api.GetUserIDFromGinContext(ctx)
	if !ok {
		response.Fail("user id error", nil, ctx)
		return
	}
	data, err := service.FinanceSvc.GetBalanceStatementList(&params, uIDInt)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("GetBalanceStatementList error:", err.Error(), ctx)
		return
	}
	response.OK("GetBalanceStatementList success", data, ctx)
}

func GetCommissionStatementList(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	uIDInt, ok := api.GetUserIDFromGinContext(ctx)
	if !ok {
		response.Fail("user id error", nil, ctx)
		return
	}
	data, err := service.FinanceSvc.GetCommissionStatementList(&params, uIDInt)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("GetCommissionStatementList error:", err.Error(), ctx)
		return
	}
	response.OK("GetCommissionStatementList success", data, ctx)
}

func WithdrawToBalance(ctx *gin.Context) {
	uIDInt, ok := api.GetUserIDFromGinContext(ctx)
	if !ok {
		response.Fail("user id error", nil, ctx)
		return
	}
	err := service.FinanceSvc.WithdrawToBalance(uIDInt)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("WithdrawToBalance error:", err.Error(), ctx)
		return
	}
	response.OK("WithdrawToBalance success", nil, ctx)
}

func GetCommissionSummary(ctx *gin.Context) {
	uIDInt, ok := api.GetUserIDFromGinContext(ctx)
	if !ok {
		response.Fail("user id error", nil, ctx)
		return
	}
	data, err := service.FinanceSvc.GetCommissionSummary(uIDInt)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("GetCommissionSummary error:", err.Error(), ctx)
		return
	}
	response.OK("GetCommissionSummary success", data, ctx)

}

func GetInvitationUserList(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	uIDInt, ok := api.GetUserIDFromGinContext(ctx)
	if !ok {
		response.Fail("user id error", nil, ctx)
		return
	}
	data, err := service.FinanceSvc.GetInvitationUserList(&params, uIDInt)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("GetInvitationUserList error:", err.Error(), ctx)
		return
	}
	response.OK("GetInvitationUserList success", data, ctx)
}
