package setting

import (
	"github.com/intmian/mian_go_lib/tool/misc"
	"sync"
)

type Mgr struct {
	j      *misc.TJsonTool
	data   map[string]string // 为了方便通用的存储机制暂时将
	rwLock *sync.RWMutex
}

func NewMgr() *Mgr {
	m := &Mgr{
		j:      nil,
		data:   make(map[string]string),
		rwLock: new(sync.RWMutex),
	}
	m.j = misc.NewTJsonTool("setting.json", m.data)
	return m
}

func (m *Mgr) ReadToSetting() map[string]string {
	m.rwLock.Lock()
	return m.data
}

func (m *Mgr) SettingReady() {
	m.j.Save()
	m.rwLock.Unlock()
}

func (m *Mgr) Load() {
	m.rwLock.Lock()
	m.j.Load("setting.json")
	m.rwLock.Unlock()
}

func (m *Mgr) Get(key string) string {
	m.rwLock.RLock()
	defer m.rwLock.RUnlock()
	if v, ok := m.data[key]; ok {
		return v
	} else {
		return ""
	}
}

var GSettingMgr = NewMgr()
