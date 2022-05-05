package task

import (
	"auto_go_go/tool"

	"github.com/intmian/mian_go_lib/tool/xlog"
	"github.com/robfig/cron"
)

type Status int

const (
	StatusClose Status = iota
	StatusRunning
	StatusPending
)

type Unit struct {
	c       *cron.Cron
	timeStr string
	status  Status
	name    string
	f       func()
}

func (u *Unit) Start() {
	u.c = cron.New()
	err := u.c.AddFunc(u.timeStr, u.do)
	if err != nil {
		tool.GLog.Log(xlog.EError, u.name, "start失败:"+err.Error())
	}
	u.c.Start()
}

func (u *Unit) Stop() {
	u.c.Stop()
}

func (u *Unit) Status() Status {
	return u.status
}

func (u *Unit) do() {
	u.status = StatusRunning
	tool.GLog.Log(xlog.ELog, u.name, "执行开始")
	u.f()
	tool.GLog.Log(xlog.ELog, u.name, "执行完成")
	u.status = StatusPending
}

func (u *Unit) GetNextTime() string {
	if u.c == nil {
		return ""
	}
	return u.c.Entries()[0].Next.Format("2006-01-02 15:04:05")
}
