package utils

import (
	"github.com/google/uuid"
)

func CreateUUIDToken()(string, error){
	logID, err := uuid.NewRandom()
	return logID.String(), err
}