package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Handler handles requests for an URL
type Handler func(*Context)

type handlerContext struct {
	App      *Application
	Context  *Context
	handlers []Handler
}

//newHandlerContext creates new app handler
func newHandlerContext(app *Application, handle ...Handler) *handlerContext {
	appHandler := &handlerContext{
		App: app,
	}
	appHandler.use(LoggerMW)
	appHandler.use(handle...)
	return appHandler
}

func (appHandler *handlerContext) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := appHandler.App.contextPool.Get().(*Context)
	c.Reset()
	c.handlerContext = appHandler
	appHandler.Context = c
	c.URLParams = mux.Vars(req)
	c.Request = req
	c.Response = w
	c.Next()
	appHandler.App.contextPool.Put(c)
}

func (appHandler *handlerContext) use(handler ...Handler) {
	appHandler.handlers = append(appHandler.handlers, handler...)
}
