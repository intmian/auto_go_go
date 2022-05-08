package task

import (
	"auto_go_go/tool/time"
	"fmt"
)

type Mgr struct {
	Units map[string]*Unit
}

// Add 新增单元
func (mgr *Mgr) Add(task Task) {
	t := NewUnit(task)
	t.Init()
	mgr.Units[task.GetName()] = t
}

// Del 删除单元
func (mgr *Mgr) Del(name string) {
	if v, ok := mgr.Units[name]; ok {
		v.Stop()
		delete(mgr.Units, name)
	}
}

func (mgr *Mgr) AllStart() {
	for _, unit := range mgr.Units {
		unit.Start()
	}
}

func (mgr *Mgr) AllStop() {
	for _, unit := range mgr.Units {
		unit.Stop()
	}
}

func (mgr *Mgr) Start(name string) {
	if unit, ok := mgr.Units[name]; ok {
		unit.Start()
	}
}

func (mgr *Mgr) StopUnit(name string) {
	if unit, ok := mgr.Units[name]; ok {
		unit.Stop()
	}
}

func (mgr *Mgr) StartUnit(name string) {
	if unit, ok := mgr.Units[name]; ok {
		unit.do()
	}
}

func (mgr *Mgr) MakeStatusText() string {
	var text string
	text = time.GTime.GetTimeStr()
	title := fmt.Sprintf("%-10s%-10s%-20s", "任务", "状态", "下次调用")
	text += title + "\n"
	for _, unit := range mgr.Units {
		str := "%10s%10s%20s\n"
		text += fmt.Sprintf(str, unit.name, status2str(unit.status), unit.GetNextTime())
	}
	return text
}

func NewMgr() *Mgr {
	return &Mgr{
		Units: make(map[string]*Unit),
	}
}

func (mgr *Mgr) Check() {
	for _, unit := range mgr.Units {
		unit.check()
	}
}
