// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210110_221111
package mysql

import (
	"default/pkg/invoker"
	"default/pkg/model/transport"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type User struct {
	Uid int `gorm:"AUTO_INCREMENT"` // id

	UserName string `gorm:"not null"` // 昵称

}

type Users []*User

// TableName 设置表明
func (t User) TableName() string {
	return "user"
}

// UserCreate 创建一条记录
func UserCreate(db *gorm.DB, data *User) (err error) {
	if err = db.Create(data).Error; err != nil {
		invoker.Logger.Error("create user error", zap.Error(err))
		return
	}
	return
}

// UserUpdate 根据主键更新一条记录
func UserUpdate(db *gorm.DB, paramId int, ups Ups) (err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramId}

	if err = db.Model(User{}).Where(sql, binds...).Updates(ups).Error; err != nil {
		invoker.Logger.Error("user update error", zap.Error(err))
		return
	}
	return
}

// UserUpdateX Update的扩展方法，根据Cond更新一条或多条记录
func UserUpdateX(db *gorm.DB, conds Conds, ups Ups) (err error) {
	sql, binds := BuildQuery(conds)
	if err = db.Model(User{}).Where(sql, binds...).Updates(ups).Error; err != nil {
		invoker.Logger.Error("user update error", zap.Error(err))
		return
	}
	return
}

// UserDelete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func UserDelete(db *gorm.DB, paramId int) (err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramId}

	if err = db.Model(User{}).Where(sql, binds...).Delete(&User{}).Error; err != nil {
		invoker.Logger.Error("user delete error", zap.Error(err))
		return
	}

	return
}

// UserDeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func UserDeleteX(db *gorm.DB, conds Conds) (err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Model(User{}).Where(sql, binds...).Delete(&User{}).Error; err != nil {
		invoker.Logger.Error("user delete error", zap.Error(err))
		return
	}

	return
}

// UserInfo 根据PRI查询单条记录
func UserInfo(db *gorm.DB, paramId int) (resp User, err error) {
	var sql = "`uid`= ?"
	var binds = []interface{}{paramId}

	if err = db.Model(User{}).Where(sql, binds...).First(&resp).Error; err != nil {
		invoker.Logger.Error("user info error", zap.Error(err))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func UserInfoX(db *gorm.DB, conds Conds) (resp User, err error) {
	sql, binds := BuildQuery(conds)

	if err = db.Model(User{}).Where(sql, binds...).First(&resp).Error; err != nil {
		invoker.Logger.Error("user info error", zap.Error(err))
		return
	}
	return
}

// UserList 查询list，extra[0]为sorts
func UserList(conds Conds, extra ...string) (resp []*User, err error) {

	sql, binds := BuildQuery(conds)
	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = invoker.Db.Model(User{}).Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		invoker.Logger.Error("user info error", zap.Error(err))
		return
	}
	return
}

// UserListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func UserListMap(conds Conds) (resp map[int]*User, err error) {

	sql, binds := BuildQuery(conds)
	mysqlSlice := make([]*User, 0)
	resp = make(map[int]*User, 0)
	if err = invoker.Db.Model(User{}).Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		invoker.Logger.Error("user info error", zap.Error(err))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Uid] = value
	}
	return
}

// UserListPage 根据分页条件查询list
func UserListPage(conds Conds, reqList *transport.ReqPage) (total int, respList Users) {
	respList = make(Users, 0)

	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := BuildQuery(conds)

	db := invoker.Db.Model(User{}).Where(sql, binds...)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
