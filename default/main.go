// @EgoctlOverwrite NO
// @EgoctlGenerateTime 20210110_221111
package main

import (
	"default/pkg/invoker"
	"default/pkg/router"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egovernor"
)

func main() {
	if err := ego.New().
		Invoker(invoker.Init).
		Serve(
			egovernor.Load("server.governor").Build(),
			router.ServeHTTP(),
		).
		Run(); err != nil {
		elog.Panic(err.Error())
	}
}
