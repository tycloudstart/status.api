package model

import (
	"fmt"
)

// request struct of querying for one judge status
type JudgeStatusQuery struct {
	Uuid     string `json:"uuid"`
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
}

func NewJudgeStatusQuery(uuid, endpoint, counter string) *JudgeStatusQuery {
	return &JudgeStatusQuery{uuid, endpoint, counter}
}

func (this *JudgeStatusQuery) ToString() string {
	return fmt.Sprintf("{JudgeStatusQuery Uuid=%s Endpoint=%s Counter=%s}",
		this.Uuid, this.Endpoint, this.Counter)
}

// response struct of querying for one judge status
type JudgeStatus struct {
	Uuid     string `json:"uuid"`
	Endpoint string `json:"endpoint"`
	Counter  string `json:"counter"`
	Status   int    `json:"status"` // 0:OK, 1:ERROR
}

func NewJudgeStatus(uuid, endpoint, counter string, status int) *JudgeStatus {
	return &JudgeStatus{uuid, endpoint, counter, status}
}

func (this *JudgeStatus) ToString() string {
	return fmt.Sprintf("{JudgeStatus Uuid=%s Endpoint=%s Counter=%s Status=%d}",
		this.Uuid, this.Endpoint, this.Counter, this.Status)
}
