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
	if errorHandler != nil {
		if err := errorHandler(c.Context, err); err != nil {
			return c.RenderError(err)
		}
	}
	message := err.Error()
	if userErr, ok := err.(*usererrors.Error); ok {
		return c.RenderResponse(&Response{
			Error: &ResponseError{
				Message: userErr.Message(),
				Type:    "user",
				Key:     userErr.Code(),
			},
		}, userErr.HttpStatus())
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

func (c *Context) Render(data any, metaPairs ...any) error {
	rsp := &Response{Data: data}

	// Parse meta key-value pairs
	if len(metaPairs) > 0 {
		if len(metaPairs)%2 != 0 {
			return fmt.Errorf("meta pairs must be even number")
		}
		meta := make(map[string]interface{})
		for i := 0; i < len(metaPairs); i += 2 {
			key, ok := metaPairs[i].(string)
			if !ok {
				return fmt.Errorf("meta key must be string")
			}
			meta[key] = metaPairs[i+1]
		}
		rsp.Meta = meta
	}

	return c.RenderResponse(rsp, 200)
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
