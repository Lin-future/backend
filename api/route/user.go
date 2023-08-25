package route

import (
	"go-svc-tpl/internal/controller"
)

type UserCtlWrapper struct { //Wrapper类隔离接口具体逻辑
	ctl controller.IUserController
}
