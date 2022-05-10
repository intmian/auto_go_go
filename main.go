package main

import (
	"auto_go_go/http"
	"auto_go_go/menu"
	"auto_go_go/setting"
	"auto_go_go/task"
	"auto_go_go/tool"
	"github.com/gin-gonic/gin"
	"github.com/intmian/mian_go_lib/tool/xlog"
	"strconv"
)

func main() {
	tool.Init()
	tool.GLog.Log(xlog.ELog, "SYS", "初始化开始")
	task.Init()
	tool.GLog.Log(xlog.ELog, "SYS", "task初始化完成")
	mode := setting.GSettingMgr.Get("system.web_mode")
	if mode == nil || mode.(bool) {
		r := gin.Default()
		http.InitRoot(r)
		tool.GLog.Log(xlog.ELog, "SYS", "web初始化完成")
		tool.GLog.Log(xlog.ELog, "SYS", "初始化完成")
		var err error
		if setting.GSettingMgr.Get("web.port") == nil {
			err = r.Run(":8080")
		} else {
			port := ":" + strconv.Itoa(setting.GSettingMgr.Get("web.port").(int))
			err = r.Run(port)
		}
		if err != nil {
			tool.GLog.Log(xlog.ELog, "SYS", "web启动失败")
		}
	} else {
		menu.Init()
		tool.GLog.Log(xlog.ELog, "SYS", "初始化完成")
		menu.Do()
	}

}
