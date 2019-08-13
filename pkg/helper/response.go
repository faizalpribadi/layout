package helper

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//APIResponse helper
func APIResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	value, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
	}
	rw.Write(value)
}
