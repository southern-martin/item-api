package http_util

import (
	"encoding/json"
	"github.com/southern-martin/util-go/rest_error"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err rest_error.RestErr) {
	RespondJson(w, err.Status(), err)
}
