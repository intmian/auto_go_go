package http

import "github.com/gin-gonic/gin"

func InitRoot(pEngine *gin.Engine) {
	pEngine.GET("/api/log_cache", getLogCache)
	pEngine.GET("/api/status", getStatus)

	//pEngine.LoadHTMLFiles("http\\static\\*")
	pEngine.GET("/", getIndex)
}
