package eject_test

import (
	"fmt"
	"testing"

	"github.com/Triment/eject"
)

func TestCreateRouter(t *testing.T) {
	router := eject.CreateRouter()
	router.GET("/hello", func(c *eject.Context) {
		fmt.Fprintf(c.Res, "hello")
	})
	paths, length := eject.GetPath("/hello")
	node := router.Tree.Search(paths, length, 0, map[string]string{})
	if node == nil {
		t.Errorf("节点未找到")
	}
}
