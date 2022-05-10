package http

import (
	"auto_go_go/log_cache"
	"auto_go_go/task"
	"github.com/gin-gonic/gin"
)

func getLogCache(c *gin.Context) {
	//输出json结果给调用方
	c.String(200, log_cache.GLogCache.ToString())
}

func getStatus(c *gin.Context) {
	c.String(200, task.GMgr.MakeStatusText())
}
