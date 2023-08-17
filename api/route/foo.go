package route

import (
	"go-svc-tpl/api/dto"
	"go-svc-tpl/internal/controller"

	"github.com/gin-gonic/gin"
)

func setupFooController(r *gin.RouterGroup) {
	cw := FooCtlWrapper{
		ctl: controller.NewFooController(),
	}
	p := r.Group("/foo")
	p.GET("/get", cw.GetFoo)
}

type FooCtlWrapper struct {
	ctl controller.IFooController
}

// >>>>>>>>>>>>>>>>>> Controller >>>>>>>>>>>>>>>>>>

func (w FooCtlWrapper) GetFoo(c *gin.Context) {
	var req dto.GetFooReq
	if err := dto.BindReq(c, &req); err != nil {
		dto.ResponseFail(c, err)
		return
	}
	resp, err := w.ctl.GetFoo(c, &req)
	if err != nil {
		dto.ResponseFail(c, err)
		return
	}
	dto.ResponseSuccess(c, resp)
}
