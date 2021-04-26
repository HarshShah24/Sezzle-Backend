package main

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func New() Server {
	e := gin.Default()
	s := Server{
		Engine: e,
	}

	s.populate()

	return s
}

// populate: Internal method to populate handlers on the engine
func (s Server) populate() {
	s.GET("/calculations", getLastResult)
	s.POST("/calculations", doCalulcations)
}
