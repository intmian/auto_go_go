package task

import (
	"auto_go_go/setting"
	"auto_go_go/tool"
	"time"

	"github.com/intmian/mian_go_lib/tool/xlog"
	"github.com/robfig/cron"
)

type Status int

const (
	StatusClose Status = iota
	StatusRunning
	StatusPending
)

func status2str(status Status) string {
	switch status {
	case StatusClose:
		return "关闭"
	case StatusRunning:
		return "运行"
	case StatusPending:
		return "等待"
	default:
		return "未知"
	}
}

type Task interface {
	Init()
	Do()
	GetName() string
	GetTimeStr() string
}

type Unit struct {
	c       *cron.Cron
	timeStr string
	status  Status
	name    string
	f       func()
	init    func()
}

func (u *Unit) Start() {
	if u.status != StatusClose {
		return
	}
	setting.GSettingMgr.Set(u.name+".open", true)
	setting.GSettingMgr.Save()
	u.c = cron.New()
	err := u.c.AddFunc(u.timeStr, u.do)
	if err != nil {
		tool.GLog.Log(xlog.EError, u.name, "start失败:"+err.Error())
	}
	u.c.Start()
	u.status = StatusPending
}

func (u *Unit) Stop() {
	if u.status == StatusClose {
		return
	}
	setting.GSettingMgr.Set(u.name+".open", false)
	setting.GSettingMgr.Save()
	u.c.Stop()
	u.status = StatusClose
}

func (u *Unit) Status() Status {
	return u.status
}

func (u *Unit) do() {
	defer func() {
		if err := recover(); err != nil {
			tool.GLog.Log(xlog.EError, u.name, "携程崩溃:"+err.(string))
		}
	}()
	u.status = StatusRunning
	tool.GLog.Log(xlog.ELog, u.name, "执行开始")
	ok := make(chan bool)
	begin := time.Now()
	go func() {
		u.f()
	}()
loop:
	for {
		select {
		case <-ok:
			break loop
		case <-time.After(time.Hour):
			now := time.Now()
			tool.GLog.Log(xlog.EWarning, u.name, "执行超时:"+now.Sub(begin).String())
		}
	}
	tool.GLog.Log(xlog.ELog, u.name, "执行完成")
	u.status = StatusPending

}

func (u *Unit) GetNextTime() string {
	if u.c == nil {
		return ""
	}
	return u.c.Entries()[0].Next.Format("2006-01-02 15:04:05")
}

func NewUnit(task Task) *Unit {
	return &Unit{
		timeStr: task.GetTimeStr(),
		name:    task.GetName(),
		f:       task.Do,
		init:    task.Init,
	}
}

func (u *Unit) Init() {
	u.init()
	if !setting.GSettingMgr.Exist(u.name + ".open") {
		setting.GSettingMgr.Set(u.name+".open", true)
		setting.GSettingMgr.Save()
		u.Start()
	} else {
		if setting.GSettingMgr.Get(u.name + ".open").(bool) {
			u.Start()
		} else {
			u.Stop()
		}
	}
}

func (u *Unit) check() {
	if !setting.GSettingMgr.Exist(u.name + ".open") {
		return
	}
	if !setting.GSettingMgr.Get(u.name + ".open").(bool) {
		u.Stop()
	} else {
		u.Start()
	}
}
