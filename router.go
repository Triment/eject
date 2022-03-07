package eject

import (
	"fmt"
	"strings"
)

type Router struct {
	Tree    *Trie
	Handler map[string]func(*Context)
}

func GetPath(path string) ([]string, int) {
	paths := strings.Split(path, "/")[1:]
	return paths, len(paths)
}

func CreateRouter() *Router {
	return &Router{Tree: &Trie{Part: ".", Children: map[string]*Trie{}}, Handler: map[string]func(*Context){}}
}

func (r *Router) RegistRouter(method string, path string, handler func(*Context)) {
	r.Tree.Insert(path)
	r.Handler[method+"-"+path] = handler
}

func (r *Router) HEAD(path string, handler func(*Context)) {
	r.RegistRouter("HEAD", path, handler)
}
func (r *Router) CONNECT(path string, handler func(*Context)) {
	r.RegistRouter("CONNECT", path, handler)
}
func (r *Router) OPTIONS(path string, handler func(*Context)) {
	r.RegistRouter("OPTIONS", path, handler)
}
func (r *Router) GET(path string, handler func(*Context)) {
	r.RegistRouter("GET", path, handler)
}

func (r *Router) POST(path string, handler func(*Context)) {
	r.RegistRouter("POST", path, handler)
}

func (r *Router) PUT(path string, handler func(*Context)) {
	r.RegistRouter("PUT", path, handler)
}

func (r *Router) DELETE(path string, handler func(*Context)) {
	r.RegistRouter("DELETE", path, handler)
}

func (r *Router) PATCH(path string, handler func(*Context)) {
	r.RegistRouter("PATCH", path, handler)
}

func (r *Router) TRACE(path string, handler func(*Context)) {
	r.RegistRouter("TRACE", path, handler)
}

func (r *Router) Accept() func(*Context) {
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
