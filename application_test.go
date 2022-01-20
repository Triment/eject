package eject_test

import (
	"fmt"
	"testing"

	"github.com/Triment/eject"
)

func TestCreateApp(t *testing.T) {
	app := eject.CreateApp()
	router := eject.CreateRouter()
	router.GET("/hello", func(c *eject.Context) {
		fmt.Fprintf(c.Res, "hello")
	})
	app.Inject(router.Accept())
	//app.Listen(":8080")
}
