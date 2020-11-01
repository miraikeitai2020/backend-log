package view

import (
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/server/model"
	"github.com/miraikeitai2020/ap2-merihariko-backend/pkg/utils"
	"strconv"
	"strings"
	"time"
)

type RequestCreateLog struct {
	UserID        string
	Date          time.Time
	WorkTime      int
	Concentration []float64
	LogName       string
}

type ResponseCreateLog struct {
	LogID string
}

type Logs struct{
	LogID string
	LogName string
}

type ResponseGetLogs struct {
	Logs []Logs
}

type ResponseGetLog struct {
	Date time.Time
	WorkTime int
	Concentration []float64
	LogName string
}

func (req *RequestCreateLog)ParseModel(log *model.Log){
	// requestで受け取ったデータをそのままブチ込める奴はここでブチ込む
	log.UserID = req.UserID
	log.LogName = req.LogName
	log.WorkTime = req.WorkTime

	// requestで受け取ったdatetime型をstringに変換してLogにブチ込む
	log.Date = req.Date.String()

	// concentrationを文字列化し，結合する
	var list []string
	for i := range req.Concentration {
		list = append(list, strconv.FormatFloat(req.Concentration[i], 'f', -1, 64))
	}
	log.Concentration = strings.Join(list, ",")
}

func NewResponseCreateLog(log *model.Log) *ResponseCreateLog{
	response := &ResponseCreateLog{
		LogID: log.LogID,
	}
	return response
}

func NewResponseGetLog(log *model.Log) *ResponseGetLog{
	response := &ResponseGetLog{
		Date: utils.StringToTime(log.Date),
		WorkTime: log.WorkTime,
		Concentration: utils.StringToFloatList(log.Concentration),
		LogName: log.LogName,
	}
	return response
}
/*
func NewResponseGetLogs(log *model.Log) *ResponseGetLogs{
	response := &ResponseGetLogs{
	}
	return response
}
*/