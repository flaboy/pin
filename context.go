package pin

import (
	"carthooks/galaxy/sidecar/lib/pin/usererrors"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (c *Context) RenderError(err error) error {
	message := err.Error()
	if _, ok := err.(*usererrors.Error); ok {
		return c.RenderUserError(message, message)
	}

	return c.renderError("system", message, "error.system")
}

func (c *Context) renderError(error_type, message, key string) error {
	return c.RenderResponse(&Response{
		Error: &ResponseError{
			Message: message,
			Type:    error_type,
			Key:     key,
		},
	})
}

func (c *Context) Render(data any) error {
	return c.RenderResponse(&Response{
		Data: data,
	})
}

func (c *Context) RenderUserError(message, key string) error {
	return c.renderError("user", message, key)
}

func (c *Context) RenderResponse(rsp *Response) error {
	trace_id, _ := c.Get("trace_id")
	rsp.TraceId = trace_id.(string)
	c.JSON(200, rsp)
	return nil
}
