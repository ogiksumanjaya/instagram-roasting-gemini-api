package dto

type GeminiRequest struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Roast string `json:"roast"`
}
