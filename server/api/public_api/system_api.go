package public_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service"
	"github.com/ppoonk/AirGo/utils/response"
)

// 主题配置
func GetThemeConfig(ctx *gin.Context) {
	if cache, ok := global.LocalCache.Get(constant.CACHE_THEME); ok {
		response.OK("GetThemeConfig success", cache, ctx)
		return
	}
	theme, _, err := service.CommonSqlFirst[model.Theme, model.Theme, model.Theme](model.Theme{ID: 1})
	if err != nil {
		response.Fail("GetThemeConfig error: "+err.Error(), nil, ctx)
		return
	}
	global.LocalCache.SetNoExpire(constant.CACHE_THEME, theme)
	response.OK("GetThemeConfig success", theme, ctx)

}

// 获取公共系统设置
func GetPublicSetting(ctx *gin.Context) {
	var ps = model.PublicSystem{
		EnableRegister:          global.Server.Website.EnableRegister,
		AcceptableEmailSuffixes: global.Server.Website.AcceptableEmailSuffixes,
		EnableEmailCode:         global.Server.Website.EnableEmailCode,
		EnableLoginEmailCode:    global.Server.Website.EnableLoginEmailCode,
		BackendUrl:              global.Server.Subscribe.BackendUrl,
		CommissionRate:          global.Server.Finance.CommissionRate,
		WithdrawThreshold:       global.Server.Finance.WithdrawThreshold,
		EnableLottery:           global.Server.Finance.EnableLottery,
		Jackpot:                 global.Server.Finance.Jackpot,
	}
	response.OK("GetPublicSetting success", ps, ctx)
}
