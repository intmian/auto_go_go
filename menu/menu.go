package menu

import (
	"auto_go_go/log_cache"
	"auto_go_go/setting"
	"auto_go_go/task"
	"auto_go_go/tool"
	"github.com/intmian/mian_go_lib/tool/menu"
	"github.com/intmian/mian_go_lib/tool/xlog"
)

const BindText = `
{
    "nodes" : [
        {
            "id":0,
            "name" : "主面板",
            "child" : [1,2,3]
        },
		{
			"id":1,
			"name" : "状态",
			"child" : []
		},
		{
			"id":2,
			"name" : "日志",
			"child" : []
		},
		{
			"id":3,
			"name" : "设置",
			"child" : []
		}
    ],
    "root" : 0
}
`

func Do() {
	m, done := Init()
	if done {
		return
	}
	m.Do()
}

func Init() (menu.Menu, bool) {
	tool.Init()
	tool.GLog.Log(xlog.ELog, "sys", "初始化开始")
	task.Init()
	tool.GLog.Log(xlog.ELog, "sys", "task初始化完成")
	m := menu.Menu{}
	suc := m.Init(menu.BindInfo{
		LogicBindText: BindText,
		FuncBindList: []menu.FuncBind{
			{
				ID: 1,
				FUNC: menu.MakeUntilPressForShowFunc(func() {
					println(task.GMgr.MakeStatusText())
				}, 1),
			},
			{
				ID: 2,
				FUNC: menu.MakeUntilPressForShowFunc(func() {
					println(log_cache.GLogCache.ToString())
				}, 1),
			},
			{
				ID: 3,
				FUNC: menu.MakeUniListInputFunc(setting.GSettingMgr.Data(), func() {
					task.GMgr.Check()
					setting.GSettingMgr.Save()
					tool.GLog.Log(xlog.ELog, "menu", "设置更新")
				}),
			},
		},
	})
	if !suc {
		tool.GLog.Log(xlog.ELog, "sys", "初始化失败")
		return menu.Menu{}, true
	}
	tool.GLog.Log(xlog.ELog, "sys", "初始化完成")
	return m, false
}
