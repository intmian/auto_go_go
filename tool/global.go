package tool

import (
	"auto_go_go/log_cache"
	"auto_go_go/setting"

	"github.com/intmian/mian_go_lib/tool/xlog"
	"github.com/intmian/mian_go_lib/tool/xpush"
)

func Init() {
	GPush = xpush.NewMgr(&xpush.EmailToken{}, &xpush.PushDeerToken{
		Token: setting.GSettingMgr.Get("pushdeer_token").(string),
	}, "autogogo")
	GLog = xlog.SimpleNewMgr(GPush, "", "", "autogogo")
	GLog.SetLogAddr("./static/log")
	GLog.SetPrinter(log_cache.GLogCache.Add)
	GLog.Log(xlog.ELog, "INIT", "日志、推送初始化完成")
}

var GLog *xlog.Mgr
var GPush *xpush.Mgr
