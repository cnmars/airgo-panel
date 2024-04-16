package admin_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/api"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

// 更新主题
func UpdateThemeConfig(ctx *gin.Context) {
	var theme model.Theme
	err := ctx.ShouldBind(&theme)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.CommonSqlSave[model.Theme](theme)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("UpdateThemeConfig error:"+err.Error(), nil, ctx)
		return
	}
	//重新加载主题设置
	global.LocalCache.SetNoExpire(constant.CACHE_THEME, theme)
	response.OK("UpdateThemeConfig success", nil, ctx)
}

// 获取系统设置
func GetSetting(ctx *gin.Context) {
	res, _, err := service.CommonSqlFind[model.Server, string, model.Server]("id = 1")
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetSetting error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetSetting success", res, ctx)
}

// 更新系统设置
func UpdateSetting(ctx *gin.Context) {
	var setting model.Server
	err := ctx.ShouldBind(&setting)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = service.AdminServerSvc.UpdateSetting(&setting)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("UpdateSetting error:"+err.Error(), nil, ctx)
		return
	}
	global.Server = setting
	//重新加载email
	global.GoroutinePool.Submit(func() {
		global.Logrus.Info("重新加载 email")
		service.EmailSvc.Reload()
	})
	//重新加载tg bot
	if global.Server.Notice.EnableTGBot {
		if global.Server.Notice.BotToken != "" {
			global.GoroutinePool.Submit(func() {
				global.Logrus.Info("重新加载tg bot")
				//关闭
				service.TgBotSvc.TGBotCloseListen()
				//重启
				service.TgBotSvc.TGBotStart()
			})
		}
	} else {
		global.GoroutinePool.Submit(func() {
			global.Logrus.Info("停止 tg bot")
			//关闭
			service.TgBotSvc.TGBotCloseListen()
		})
	}
	//重新加载通知消息时的管理员id
	global.GoroutinePool.Submit(func() {
		global.Logrus.Info("重新加载通知消息时的管理员id")
		service.AdminServerSvc.AdminAccountHandler()
	})
	response.OK("UpdateSetting success", nil, ctx)
}

func GetCurrentVersion(ctx *gin.Context) {
	response.OK("GetCurrentVersion success", gin.H{"version": constant.V}, ctx)
}

func GetLatestVersion(ctx *gin.Context) {
	v, err := service.AdminServerSvc.GetLatestVersion()
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetLatestVersion error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetLatestVersion success", gin.H{"version": v}, ctx)
}

func UpdateLatestVersion(ctx *gin.Context) {
	api.SSE(ctx)
	err := service.AdminServerSvc.DownloadLatestVersion(ctx)
	if err != nil {
		response.ResponseSSE("message error", err.Error(), ctx)
		return
	}
	response.ResponseSSE("success", "success", ctx)
}
