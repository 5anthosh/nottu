package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
	"yrnote/config"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

// Application represents main application
type Application struct {
	router      *mux.Router
	contextPool *sync.Pool
	DB          *sql.DB
}

//Handle registers http handler
func (app *Application) Handle(path string, handlerFunc ...Handler) *mux.Route {
	appHandler := newHandlerContext(app, handlerFunc...)
	return app.router.Handle(path, appHandler)
}

//HandleStatic registers a new handler to handle static content such as img, css, html, js.
func (app *Application) HandleStatic(path string, dir string) {
	app.router.PathPrefix(path).Handler(http.FileServer(http.Dir(dir)))
}

//New creates new application
func New() *Application {
	app := &Application{}
	// app.Config = config.GetConfig(config.ConfigPath)
	// app.DB = db.Connection(app.Config, db.DevMode)
	app.contextPool = &sync.Pool{
		New: newContextPool(app),
	}
	app.router = NewRouter()
	return app
}

//Run starts the application
func (app *Application) Run() {
	port := config.DefaultPort
	serverAdd := ":" + port
	fmt.Println("🚀  Starting server....")

	//	go open(localAddress + "/#/pltm/container")
	protocal := "http"
	var err error
	localAddress := protocal + "://localhost" + serverAdd
	fmt.Println("🌠 Ready on " + localAddress)
	err = http.ListenAndServe(serverAdd, handlers.RecoveryHandler()(app.router))
	if err != nil {
		fmt.Println("Stopping the server" + err.Error())
	}
}
