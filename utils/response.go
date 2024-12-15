package response

import (
	"encoding/json"
	"net/http"
)

type MetaErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data any               `json:"data"`
	Meta MetaErrorResponse `json:"meta"`
}

type EmptyResponse struct {
}

func ReturnResponse(
	w http.ResponseWriter,
	statusCode int,
	response any,
	err error,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)

	baseResponse := Response{
		Meta: MetaErrorResponse{},
	}

	if err != nil {
		baseResponse.Meta.Code = statusCode
		baseResponse.Meta.Message = err.Error()

		_ = encoder.Encode(baseResponse)

		return
	}

	baseResponse.Data = response

	_ = encoder.Encode(baseResponse)
}
