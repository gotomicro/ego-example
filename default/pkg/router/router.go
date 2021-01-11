// @EgoctlOverwrite YES
// @EgoctlGenerateTime 20210110_221111
package router

import (
	"default/pkg/invoker"
	"github.com/gotomicro/ego/server/egin"
)

func ServeHTTP() *egin.Component {
	r := invoker.Gin

	InitUser(r)

	return r
}
