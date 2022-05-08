package task

import "auto_go_go/mods"

var GMgr = NewMgr()

func Init() {
	GMgr.Add(&mods.Baidu{})
	GMgr.Add(&mods.Dapan{})
	GMgr.Add(&mods.Lottery{})
}
