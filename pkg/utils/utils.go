package utils

import (
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
	"time"
)
// 2017-02-01 03:15:45 +0000 UTC
var layout = "2006-01-02 15:04:05 +0000 UTC"

func CreateUUIDToken()(string, error){
	logID, err := uuid.NewRandom()
	return logID.String(), err
}

func StringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func StringToFloatList(str string) []float64{
	list := strings.Split(str,",")
	var f []float64
	for i := range list {
		n , err := strconv.ParseFloat(list[i], 64)
		if err != nil {
			log.Println(err)
		}
		f = append(f, n)
	}
	return f
}