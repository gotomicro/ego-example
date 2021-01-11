// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210110_221111
package transport

type ReqPage struct {
	Current  int    `json:"current" form:"current"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Sort     string `json:"sort" form:"sort"`
}
