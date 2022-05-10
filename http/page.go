package http

import (
	"github.com/gin-gonic/gin"
)

func getIndex(c *gin.Context) {
	//c.HTML(200, `index.html`, gin.H{})
	c.File(`http\static\index.html`)
}
