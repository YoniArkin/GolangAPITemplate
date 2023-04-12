package api

import "github.com/gin-gonic/gin"

type Version interface {
	GetVersionPrefix() string
	AddRoutes(router gin.RouterGroup)
}
