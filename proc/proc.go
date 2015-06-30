package proc

import (
	nproc "github.com/niean/gotools/proc"
	"log"
)

// 统计指标的整体数据
var (
	// judge status query cnt
	JudgeStatusQuery   = nproc.NewSCounterQps("JudgeStatusQuery")
	JudgeStatusesQuery = nproc.NewSCounterQps("JudgeStatusesQuery")
)

func Start() {
	log.Println("proc.Start, ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	ret = append(ret, JudgeStatusQuery.Get())
	ret = append(ret, JudgeStatusesQuery.Get())

	return ret
}
