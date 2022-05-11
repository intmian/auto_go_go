package http

import "github.com/gin-gonic/gin"

func InitRoot(pEngine *gin.Engine) {
	pEngine.GET("/api/log_cache", getLogCache)
	pEngine.GET("/api/title", getTitle)
	pEngine.GET("/api/task/status", getStatus)
	pEngine.GET("/api/task/start", startTask)

	pEngine.StaticFile("/", "http\\static\\index.html")
	pEngine.StaticFS("/static", gin.Dir("http\\static\\out", false))

}
