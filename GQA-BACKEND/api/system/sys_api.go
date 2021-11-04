package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiApi struct {
}

func (a *ApiApi) GetApiList(c *gin.Context) {
	var pageInfo global.RequestPage
	if err := c.ShouldBindJSON(&pageInfo); err != nil{
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败！"+err.Error(), c)
		return
	}
	if err, apiList, total := service.GroupServiceApp.ServiceSystem.GetApiList(pageInfo); err != nil {
		global.GqaLog.Error("获取API列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取API列表失败，" + err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:     apiList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiApi) EditApi(c *gin.Context) {}

func (a *ApiApi) AddApi(c *gin.Context) {}

func (a *ApiApi) DeleteApi(c *gin.Context) {}

func (a *ApiApi) QueryApiById(c *gin.Context) {}

