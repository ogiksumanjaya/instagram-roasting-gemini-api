package handler

import (
	response "instagram-roasting/utils"
	"net/http"
)

func Healthness(w http.ResponseWriter, r *http.Request) {
	response.ReturnResponse(w, http.StatusOK, response.EmptyResponse{}, nil)
}
