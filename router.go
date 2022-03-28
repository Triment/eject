package eject

import (
	"fmt"
	"strings"
)

type Router struct {
	Tree          *Trie
	CurrentPath   string
	CurrentHandle func(*Context)
	CurrentMethod string
	Handler       map[string]func(*Context)
}

func GetPath(path string) ([]string, int) {
	paths := strings.Split(path, "/")[1:]
	return paths, len(paths)
}

func CreateRouter() *Router {
	return &Router{Tree: &Trie{Part: ".", Children: map[string]*Trie{}}, CurrentMethod: "", CurrentPath: "", CurrentHandle: nil, Handler: map[string]func(*Context){}}
}

func (r *Router) RegistRouter(method string, path string, handler func(*Context)) {
	r.Tree.Insert(path)
	r.Handler[method+"-"+path] = handler
}

func (r *Router) Before(interceptor func(*Context) bool) *Router{
	wrap := func(wrap func(*Context) bool, dest func(*Context)) func(*Context) {
		return func(ctx *Context) {
			if wrap(ctx){
				dest(ctx)
			}
		}
	}
	r.RegistRouter(r.CurrentMethod, r.CurrentPath, wrap(interceptor, r.CurrentHandle))
	r.CurrentHandle = nil
	r.CurrentMethod = ""
	r.CurrentPath = ""
	return r
}

func (r *Router) clearPreviousHandle() {
	if r.CurrentPath != "" && r.CurrentHandle != nil && r.CurrentMethod != "" {
		r.RegistRouter(r.CurrentMethod, r.CurrentPath, r.CurrentHandle)
	}
}

func (r *Router) HEAD(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "HEAD"
	r.CurrentHandle = handler
	return r
}
func (r *Router) CONNECT(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "CONNECT"
	r.CurrentHandle = handler
	return r
}
func (r *Router) OPTIONS(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "OPTIONS"
	r.CurrentHandle = handler
	return r
}
func (r *Router) GET(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "GET"
	r.CurrentHandle = handler
	return r
}

func (r *Router) POST(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "POST"
	r.CurrentHandle = handler
	return r
}

func (r *Router) PUT(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "PUT"
	r.CurrentHandle = handler
	return r
}

func (r *Router) DELETE(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "DELETE"
	r.CurrentHandle = handler
	return r
}

func (r *Router) PATCH(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "PATCH"
	r.CurrentHandle = handler
	return r
}

func (r *Router) TRACE(path string, handler func(*Context)) *Router{
	r.clearPreviousHandle()
	r.CurrentPath = path
	r.CurrentMethod = "TRACE"
	r.CurrentHandle = handler
	return r
}

func (r *Router) Accept() func(*Context) {
	r.clearPreviousHandle()
	return func(c *Context) {
		paths := strings.Split(c.Req.URL.Path, "/")[1:]
		node := r.Tree.Search(paths, len(paths), 0, c.Params)
		if node != nil {
			handler := r.Handler[c.Req.Method+"-"+node.FullPath]
			if handler != nil {
				handler(c)
			} else {
				c.ERROR(fmt.Sprintf("method error %s", c.Req.Method), 404)
			}
		} else {
			c.ERROR(fmt.Sprintf("Not found %s", c.Req.URL.Path), 404)
		}
	}
}
