package app

import (
	"database/sql"
	"encoding/json"
	"net"
	"net/http"
	"strings"
)

var (
	newLine = []byte{'\n'}
)

//Context provides context for whole request/response cycle
type Context struct {
	*handlerContext
	DB         *sql.DB
	Request    *http.Request
	Response   http.ResponseWriter
	URLParams  map[string]string
	Params     map[string]string
	index      int8
	StatusCode int
	Size       int
	Error      []error
}

func (app *Application) newContext() *Context {
	return &Context{
		Params: make(map[string]string),
		DB:     app.DB,
	}
}

//Reset resets the value the context
func (c *Context) Reset() {
	c.handlerContext = nil
	c.Request = nil
	c.Response = nil
	c.Params = make(map[string]string)
	c.StatusCode = 0
	c.Size = 0
	c.Error = nil
	c.index = 0
}
func newContextPool(app *Application) func() interface{} {
	return func() interface{} {
		return app.newContext()
	}
}

//GetRequestHeader returns request header
func (c *Context) GetRequestHeader(key string) string {
	return c.Request.Header.Get(key)
}

//SetStatus set http status code
func (c *Context) HTTPStatusCode(status int) {
	c.StatusCode = status
	c.Response.WriteHeader(status)
}

func (c *Context) JSON(response interface{}) {
	c.Response.Header().Add("Content-Type", "application/json")
	jsonContentByte, err := json.Marshal(response)
	if err != nil {
		c.AppendError(err)
	}
	size, err := c.Response.Write(jsonContentByte)
	if err != nil {
		c.AppendError(err)
	}
	c.setSize(size)
	size, err = c.Response.Write(newLine)
	c.setSize(size)
	if err != nil {
		c.AppendError(err)
	}
}

func (c *Context) setSize(size int) {
	c.Size += size
}

//AppendError records error to be displayed later
func (c *Context) AppendError(err ...error) {
	if err != nil {
		c.Error = append(c.Error, err...)
	}
}

//ClientIP returns ip address of the user using request info
func (c *Context) ClientIP() string {
	clientIP := c.GetRequestHeader("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(c.GetRequestHeader("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

//Next runs the next handler
func (c *Context) Next() {
	handle := c.handlerContext.handlers[c.index]
	c.index++
	handle(c)
}

//GetURLQuery get the params in url (Eg . /?q=)
func (c *Context) GetURLQuery(query string) string {
	return c.Request.URL.Query().Get(query)
}
