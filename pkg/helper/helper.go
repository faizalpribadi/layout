package helper

import (
	uuid "github.com/satori/go.uuid"
)

//GenerateUID random uid
func GenerateUID() string {
	uid := uuid.NewV4()
	return uid.String()
}
