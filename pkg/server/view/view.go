package view

import "time"

type RequestCreateLog struct{
	UserID string
	Date time.Time
	WorkTime time.Time
	Concentration []float64
	LogName string
}

type ResponseCreateLog struct{
	LogID string
}
