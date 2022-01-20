package eject

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Res    http.ResponseWriter
	Req    *http.Request
	Params map[string]string
}

func CreateContext() *Context {
	return new(Context)
}

func (c *Context) JSON(v interface{}) {
	c.Res.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(c.Res)
	if err := encoder.Encode(v); err != nil {
		c.ERROR(err.Error(), 500)
	}
}
func (c *Context) ERROR(str string, code int) {
	http.Error(c.Res, fmt.Sprintf("Error: %v", str), code)
}
