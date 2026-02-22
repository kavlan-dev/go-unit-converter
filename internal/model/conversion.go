package model

type Conversion struct {
	FromValue float64 `json:"from_value"`
	FromUnit  string  `json:"from_unit"`
	ToUnit    string  `json:"to_unit"`
	Result    float64 `json:"result"`
}

type ConversionRequest struct {
	FromValue float64 `json:"from_value"`
	FromUnit  string  `json:"from_unit"`
	ToUnit    string  `json:"to_unit"`
}

type ConversionResponse struct {
	Result float64 `json:"result"`
}

func NewConversionResponse(result float64) *ConversionResponse {
	return &ConversionResponse{Result: result}
}
