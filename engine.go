package pin

import (
	"github.com/gin-gonic/gin"
	"github.com/twinj/uuid"
)

type Engine struct {
	*gin.Engine
}

func New(g *gin.Engine) *Engine {
	r := &Engine{g}
	r.Use(func(c *gin.Context) {
		trace_id := c.Request.Header.Get("X-Trace-Id")
		if trace_id == "" {
			trace_id = uuid.NewV4().String()
		}
		c.Set("trace_id", trace_id)
	})
	return r
}

type HandlerFunc func(c *Context) error

func (p *Engine) handle(c *gin.Context, handler HandlerFunc) {
	ctx := &Context{c}
	if err := handler(ctx); err != nil {
		ctx.RenderError(err)
	}
}

func (p *Engine) Get(path string, handler HandlerFunc) {
	p.Engine.GET(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Post(path string, handler HandlerFunc) {
	p.Engine.POST(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Put(path string, handler HandlerFunc) {
	p.Engine.PUT(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Delete(path string, handler HandlerFunc) {
	p.Engine.DELETE(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Patch(path string, handler HandlerFunc) {
	p.Engine.PATCH(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Options(path string, handler HandlerFunc) {
	p.Engine.OPTIONS(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Head(path string, handler HandlerFunc) {
	p.Engine.HEAD(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Any(path string, handler HandlerFunc) {
	p.Engine.Any(path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}

func (p *Engine) Match(methods []string, path string, handler HandlerFunc) {
	p.Engine.Match(methods, path, func(c *gin.Context) {
		p.handle(c, handler)
	})
}
