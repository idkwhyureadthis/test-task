package formatter

import (
	"encoding/json"
	"net/http"
)

func JsonifyError(w http.ResponseWriter, err error) {
	type ErrorResp struct {
		Error string `json:"error"`
	}
	resp := ErrorResp{Error: err.Error()}
	w.WriteHeader(500)
	dat, _ := json.Marshal(resp)
	w.Write(dat)
}

func JsonifyMessage(w http.ResponseWriter, message string) {
	type ErrorResp struct {
		Message string `json:"data"`
	}
	resp := ErrorResp{Message: message}
	w.WriteHeader(200)
	dat, _ := json.Marshal(resp)
	w.Write(dat)
}
