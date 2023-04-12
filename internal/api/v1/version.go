package v1

import "github.com/gin-gonic/gin"

func (V1) GetVersionPrefix() string {
	return "v1"
}

func (V1) AddRoutes(router gin.RouterGroup) {

}
