package infrastructures

import (
	"github.com/gin-gonic/gin"

	"github.com/TechDepa/c_tool/usecase"
)

type Request struct {
	*gin.Context
}

func (c Request) RenderJSON(sc usecase.Status, body interface{}) {
	c.JSON(sc.Code(), body)
}

func (c Request) RenderAbortWithError(sc usecase.Status, err error) {
	c.AbortWithError(sc.Code(), err)
}
