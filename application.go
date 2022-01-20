package eject

import "net/http"

type Application struct {
	Middle []func(*Context)
}

func CreateApp() *Application {
	return new(Application)
}

func (app *Application) Inject(middle func(*Context)) {
	app.Middle = append(app.Middle, middle)
}

func (app *Application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := CreateContext()
	context.Req = r
	context.Res = w
	for _, mid := range app.Middle {
		mid(context)
	}
}

func (app *Application) Listen(addr string) error {
	return http.ListenAndServe(addr, app)
}
