package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type Engine struct {
	*gin.Engine
}

func TraceIDParser(c *gin.Context) {
	trace_id := c.Request.Header.Get("X-Trace-Id")
	if trace_id == "" {
		trace_id = uuid.NewV4().String()
	}
	c.Set("trace_id", trace_id)
}

type HandlerFunc func(c *Context) error

func HandleFunc(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{c}
		if err := handler(ctx); err != nil {
			ctx.RenderError(err)
		}
	}
}
