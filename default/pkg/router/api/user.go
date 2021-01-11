// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210110_221111
package api

import (
	"default/pkg/invoker"
	"default/pkg/model/dto"
	"default/pkg/model/mysql"
	"default/pkg/model/transport"
	"default/pkg/router/core"
	"github.com/spf13/cast"
)

// UserInfo 查询单条记录
func UserInfo(c *core.Context) {
	id := cast.ToInt(c.Param("uid"))
	if id == 0 {
		c.JSONE(1, "请求错误", nil)
		return
	}
	info, _ := mysql.UserInfo(invoker.Db, id)
	c.JSONOK(info)
}

// UserList 查询多条带分页记录
func UserList(c *core.Context) {
	req := &transport.ReqPage{}
	if err := c.Bind(req); err != nil {
		c.JSONE(1, "参数错误", err)
		return
	}

	query := mysql.Conds{}
	if v := c.Query("uid"); v != "" {
		query["uid"] = v
	}
	if v := c.Query("user_name"); v != "" {
		query["user_name"] = v
	}

	total, list := mysql.UserListPage(query, req)
	c.JSONPage(list, core.Pagination{
		Current: req.Current, PageSize: req.PageSize, Total: total,
	})
}

// UserCreate 创建记录
func UserCreate(c *core.Context) {
	req := &dto.UserCreate{}
	if err := c.Bind(req); err != nil {
		c.JSONE(1, "参数错误: "+err.Error(), err)
		return
	}

	create := &mysql.User{

		Uid: req.Uid,

		UserName: req.UserName,
	}

	err := mysql.UserCreate(invoker.Db, create)
	if err != nil {
		c.JSONE(1, "创建失败", err)
		return
	}
	c.JSONOK(req)
}

// UserUpdate 更新指定记录
func UserUpdate(c *core.Context) {
	req := &dto.UserUpdate{}
	if err := c.Bind(req); err != nil {
		c.JSONE(1, "参数错误: "+err.Error(), err)
		return
	}

	id := cast.ToInt(c.Param("uid"))
	if id == 0 {
		c.JSONE(1, "请求错误", nil)
		return
	}

	err := mysql.UserUpdate(invoker.Db, id, mysql.Ups{
		"uid":       req.Uid,
		"user_name": req.UserName,
	})
	if err != nil {
		c.JSONE(1, "更新失败", err)
		return
	}
	c.JSONOK()
}

// UserDelete 删除指定记录
func UserDelete(c *core.Context) {
	id := cast.ToInt(c.Param("uid"))
	if id == 0 {
		c.JSONE(1, "请求uid错误", nil)
		return
	}

	err := mysql.UserDelete(invoker.Db, id)
	if err != nil {
		c.JSONE(1, "删除失败", err)
		return
	}
	c.JSONOK()
}
