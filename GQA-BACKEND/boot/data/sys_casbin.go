package data

import (
	"fmt"
	"gin-quasar-admin/global"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysCasbin = new(sysCasbin)

type sysCasbin struct{}

var sysCasbinData = []gormadapter.CasbinRule{
	// 超级管理员组
	{Ptype: "g", V0: "admin", V1: "super-admin"},
	// user组
	{Ptype: "p", V0: "super-admin", V1: "/user/user-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-id", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-menu", V2: "GET"},
	{Ptype: "p", V0: "super-admin", V1: "/user/user-role", V2: "GET"},
	// role组
	{Ptype: "p", V0: "super-admin", V1: "/role/role-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-id", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-menu", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-menu-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-api", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-api-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-user", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-user-remove", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/role/role-user-add", V2: "POST"},
	// menu组
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/menu/menu-id", V2: "POST"},
	// dept组
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/dept/dept-id", V2: "POST"},
	// dict组
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-id", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/dict/dict-parent-id", V2: "POST"},

	// api组
	{Ptype: "p", V0: "super-admin", V1: "/api/api-list", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-edit", V2: "PUT"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-add", V2: "POST"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-delete", V2: "DELETE"},
	{Ptype: "p", V0: "super-admin", V1: "/api/api-id", V2: "POST"},
}

func (s *sysCasbin) Init() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&[]gormadapter.CasbinRule{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> casbin_rule 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("casbin_rule 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysCasbinData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> casbin_rule 表初始数据成功！")
		global.GqaLog.Error("casbin_rule 表初始数据成功！")
		return nil
	})
}