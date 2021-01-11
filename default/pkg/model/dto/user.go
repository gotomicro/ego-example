// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210110_221111
package dto

type UserCreate struct {
	Uid int `json:"id" binding:""` // id

	UserName string `json:"userName" binding:""` // 昵称

}

type UserUpdate struct {
	Uid int `json:"id" binding:""` // id

	UserName string `json:"userName" binding:""` // 昵称

}
