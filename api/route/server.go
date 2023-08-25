package route

import (
	"go-svc-tpl/internal/controller"
)

type ServerCtlWrapper struct { //Wrapper类隔离接口具体逻辑
	ctl controller.IServerController
}
