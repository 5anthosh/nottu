package app

import (
	"time"
)

//LoggerMW logger middleware
func LoggerMW(c *Context) {
	start := time.Now()
	path := c.Request.URL.Path
	c.Next()
	log := new(Logger)
	log.TimeStamp = time.Now()
	log.Latency = log.TimeStamp.Sub(start)
	log.Method = c.Request.Method
	log.StatusCode = c.StatusCode
	log.ClientIP = c.ClientIP()
	log.BodySize = c.Size
	log.Path = path
	log.Errors = c.Error
	go log.Print()
}
