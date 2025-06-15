package pin

import (
	"fmt"

	"github.com/flaboy/pin/usererrors"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

var errorHandler ErrorHandler

type ErrorHandler func(c *gin.Context, err error) error

func SetErrorHandler(h ErrorHandler) {
	errorHandler = h
}

func RenderError(c *gin.Context, err error) error {
	pinCtx := Context{c}
	return pinCtx.RenderError(err)
}

func Render(c *gin.Context, data any) error {
	pinCtx := Context{c}
	return pinCtx.Render(data)
}

func (c *Context) RenderError(err error) error {
	fmt.Println(333333)
	if errorHandler != nil {
		if err := errorHandler(c.Context, err); err != nil {
			return c.RenderError(err)
		}
	}
	message := err.Error()
	if userErr, ok := err.(*usererrors.Error); ok {
		return c.RenderUserError(userErr.Message(), userErr.Code())
	}
	return c.renderError("system", message, "error.system")
}

func (c *Context) renderError(error_type, message, key string) error {
	code := 200
	codeV, ok := c.Get("pin.error_code." + error_type)
	if ok {
		code = codeV.(int)
	}

	return c.RenderResponse(&Response{
		Error: &ResponseError{
			Message: message,
			Type:    error_type,
			Key:     key,
		},
	}, code)
}

func (c *Context) Render(data any) error {
	return c.RenderResponse(&Response{
		Data: data,
	}, 200)
}

func (c *Context) RenderUserError(message, key string) error {
	return c.renderError("user", message, key)
}

func (c *Context) RenderResponse(rsp *Response, code int) error {
	trace_id, _ := c.Get("trace_id")
	switch trace_id := trace_id.(type) {
	case string:
		rsp.TraceId = trace_id
	}
	c.JSON(code, rsp)
	return nil
}
