package view

import (
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
	"time"
)

type RequestCreateLog struct {
	UserID        string
	Date          time.Time
	WorkTime      time.Time
	Concentration []float64
	LogName       string
}

type ResponseCreateLog struct {
	LogID string
}

func NewResponseCreateLog(log *model.Log) *ResponseCreateLog{
	response := &ResponseCreateLog{
		LogID: log.LogID,
	}
	return response
}
