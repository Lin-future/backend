package model

import (
	"gopkg.in/guregu/null.v4"
)

const LinkTable = "Links"

type Link struct {
	ID        uint      `json:"ID"` //ID
	Active    bool      `json:"active"`
	Comment   string    `json:"comment"`
	EndTime   null.Time `json:"end_time"`
	Origin    string    `json:"origin"`
	Short     string    `json:"short"`
	StartTime null.Time `json:"start_time"`
	OwnerID   uint      `json:"owner_id"`
}
