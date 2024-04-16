package admin_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/encrypt_plugin"
	"github.com/ppoonk/AirGo/utils/response"
)

// 新建节点
func NewNode(ctx *gin.Context) {
	var node model.Node
	err := ctx.ShouldBind(&node)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminNodeSvc.NewNode(&node)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("NewNode error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("NewNode success", nil, ctx)
}

// 删除节点
func DeleteNode(ctx *gin.Context) {
	var node model.Node
	err := ctx.ShouldBind(&node)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminNodeSvc.DeleteNode(&node)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("DeleteNode error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("DeleteNode success", nil, ctx)
}

// 更新节点
func UpdateNode(ctx *gin.Context) {
	var node model.Node
	err := ctx.ShouldBind(&node)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminNodeSvc.UpdateNode(&node)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("UpdateNode error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("UpdateNode success", nil, ctx)

}

func GetNodeList(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	res, err := service.AdminNodeSvc.GetNodeList(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetNodeList error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetNodeList success", res, ctx)
}

// 获取节点列表，带流量信息
func GetNodeListWithTraffic(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	params.TableName = "node"
	res, err := service.AdminNodeSvc.GetNodeListWithTraffic(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetNodeListWithTraffic error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetNodeListWithTraffic success", res, ctx)
}

// 节点排序
func NodeSort(ctx *gin.Context) {
	var nodeArr []model.Node
	err := ctx.ShouldBind(&nodeArr)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.CommonSqlUpdateMultiLine[[]model.Node](nodeArr, "id", []string{"node_order"})
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("NodeSort error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("NodeSort success", nil, ctx)
}

// 解析
func ParseUrl(ctx *gin.Context) {
	var url model.NodeSharedReq
	err := ctx.ShouldBind(&url)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	nodeArr := service.AdminNodeSvc.ParseSubUrl(url.Url)
	response.OK("NewNodeShared success", nodeArr, ctx)
}

// 新增共享节点
func NewNodeShared(ctx *gin.Context) {
	var nodes []model.Node
	err := ctx.ShouldBind(&nodes)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	for _, v := range nodes {
		_ = service.AdminNodeSvc.NewNode(&v)

	}
	response.OK("NewNodeShared success", nil, ctx)
}

// reality x25519
func Createx25519(ctx *gin.Context) {
	str := encrypt_plugin.RandomString(43)
	pub, pri, err := encrypt_plugin.ExecuteX25519(str)
	if err != nil {
		global.Logrus.Error(err.Error())
		return
	}
	response.OK("Createx25519 success", model.AGREALITYx25519{PublicKey: pub, PrivateKey: pri}, ctx)
}

// 获取节点服务器状态
func GetNodeServerStatus(ctx *gin.Context) {
	list := service.AdminNodeSvc.GetNodesStatus()
	response.OK("GetNodeServerStatus success", list, ctx)

}
