package tool

import (
	"auto_go_go/log"
	"auto_go_go/setting"
	"github.com/intmian/mian_go_lib/tool/push"
	"github.com/intmian/mian_go_lib/tool/xlog"
)

func Init() {
	GPush = push.NewMgr(&push.EmailToken{}, &push.PushDeerToken{
		Token: setting.GSettingMgr.Get("pushdeer_token"),
	}, "autogogo")
	GLog = xlog.SimpleNewMgr(GPush, "", "", "autogogo")
	GLog.SetPrinter(log.GLog.Add)
	GLog.Log(xlog.ELog, "INIT", "日志、推送初始化完成")
}

var GLog *xlog.Mgr
var GPush *push.Mgr
