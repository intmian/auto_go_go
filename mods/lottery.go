package mods

import (
	"auto_go_go/tool"
	"github.com/intmian/mian_go_lib/tool/spider"
	"github.com/intmian/mian_go_lib/tool/xlog"
)

type Lottery struct {
}

func (l *Lottery) Init() {
}

func (l *Lottery) Do() {
	lotteries := spider.GetLottery()
	if lotteries == nil {
		tool.GLog.Log(xlog.EWarning, l.GetName(), "接口失效")
		return
	}
	s := spider.ParseLotteriesToMarkDown(lotteries)
	tool.GPush.PushPushDeer("彩票", s, true)
}

func (l *Lottery) GetName() string {
	return "LOTTERY"
}

func (l *Lottery) GetTimeStr() string {
	return "0 0 21 * * ?"
}
