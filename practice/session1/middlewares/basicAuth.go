package middlewares

import (
	"github.com/gin-gonic/gin"
)

func basicAuth() gin.HandlerFunc {
	//baisc auth take the parameters with of vails accounts and password
	return gin.BasicAuth(gin.Accounts{
		"ramakantk": " golang",
	})
}
