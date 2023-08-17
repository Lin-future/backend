package model

import "time"

const LinkTable = "Links"

type Link struct {
	ID        uint      `json:"id"`         // 链接ID
	Title     string    `json:"title"`      // 链接标题
	URL       string    `json:"url"`        // 链接URL
	OwnerID   uint      `json:"owner_id"`   // 拥有者ID
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}
