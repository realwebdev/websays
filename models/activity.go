package models

import (
	"time"
)

type Activity struct {
	ID            uint
	ProjectID     uint          `json:"projectid"`
	UserID        uint          `json:"userid"`
	Activity_name string        `json:"activity_name"`
	Start_time    time.Time     `json:"start_time"`
	End_time      time.Time     `json:"end_time"`
	Total_time    time.Duration `json:"total_time"`
}
