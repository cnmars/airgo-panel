package admin_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

func DeleteTicket(ctx *gin.Context) {
	var ticket model.Ticket
	err := ctx.ShouldBind(&ticket)
	if err != nil {
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.CommonSqlDelete[model.Ticket, model.Ticket](ticket)
	if err != nil {
		response.Fail("DeleteTicket error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("DeleteTicket success", nil, ctx)

}

func UpdateTicket(ctx *gin.Context) {
	var ticket model.Ticket
	err := ctx.ShouldBind(&ticket)
	if err != nil {
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.CommonSqlSave[model.Ticket](ticket)
	if err != nil {
		response.Fail("UpdateTicket error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("UpdateTicket success", nil, ctx)

}

func GetTicketList(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	data, total, err := service.CommonSqlFindWithFieldParams(&params)
	if err != nil {
		response.Fail("GetTicketList error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetTicketList success", model.CommonDataResp{
		Total: total,
		Data:  data,
	}, ctx)
}

func FirstTicket(ctx *gin.Context) {
	var params model.Ticket
	err := ctx.ShouldBind(&params)
	if err != nil {
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	ticket, err := service.AdminTicketSvc.FirstTicket(&model.Ticket{ID: params.ID})

	if err != nil {
		response.Fail("FirstTicket error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("FirstTicket success", ticket, ctx)
}
func SendTicketMessage(ctx *gin.Context) {
	var msg model.TicketMessage
	err := ctx.ShouldBind(&msg)
	if err != nil {
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	msg.IsAdmin = true
	err = service.AdminTicketSvc.NewTicketMessage(&msg)
	if err != nil {
		response.Fail(err.Error(), nil, ctx)
		return
	}
	response.OK("SendTicketMessage success", nil, ctx)
}
