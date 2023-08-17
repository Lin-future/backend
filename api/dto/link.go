// route/link.go​
package dto

import "gopkg.in/guregu/null.v4"

// >>>>>>>>>> LinkCreate >>>>>>>>>>​
type LinkCreateReq struct {
	Comment   string    `json:"comment"`    // 备注信息​
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC​
	Origin    string    `json:"origin"`     // 原始链接​
	Short     string    `json:"short"`      // 短链ID，全局唯一​
	StartTime null.Time `json:"start_time"` // 起始时间，UTC​
}

type LinkCreateResp struct {
	Active    bool      `json:"active"`     // 服务状态​
	Comment   string    `json:"comment"`    // 备注信息​
	EndTime   null.Time `json:"end_time"`   // 到期时间，UTC​
	Origin    string    `json:"origin"`     // 原始链接​
	Short     string    `json:"short"`      // 短链ID，全局唯一​
	StartTime null.Time `json:"start_time"` // 起始时间，UTC​
}
