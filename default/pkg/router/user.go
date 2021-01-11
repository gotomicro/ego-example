// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210110_221111
package router

import (
	"default/pkg/router/api"
	"default/pkg/router/core"
	"github.com/gin-gonic/gin"
)

func InitUser(r gin.IRoutes) {
	r.GET("/api/admin/users/:uid", core.Handle(api.UserInfo))
	r.GET("/api/admin/users", core.Handle(api.UserList))
	r.POST("/api/admin/users", core.Handle(api.UserCreate))
	r.PUT("/api/admin/users/:uid", core.Handle(api.UserUpdate))
	r.DELETE("/api/admin/users/:uid", core.Handle(api.UserDelete))
}
