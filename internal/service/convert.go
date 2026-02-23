package service

import (
	"fmt"
	"go-unit-converter/internal/model"
	"go-unit-converter/internal/util"
)

type convertService struct {
}

func NewConvertService() *convertService {
	return &convertService{}
}

func (s *convertService) ConvertLength(req model.ConversionRequest) (model.ConversionResponse, error) {
	valueInMeters, err := util.ConvertToMeters(req.FromValue, req.FromUnit)
	if err != nil {
		return model.ConversionResponse{}, fmt.Errorf("ConvertToMeters: %w", err)
	}

	result, err := util.ConvertFromMeters(valueInMeters, req.ToUnit)
	if err != nil {
		return model.ConversionResponse{}, fmt.Errorf("ConvertFromMeters: %w", err)
	}

	resp := model.NewConversionResponse(result)
	return resp, nil
}

func (s *convertService) ConvertWeight(req model.ConversionRequest) (model.ConversionResponse, error) {
	valueInKg, err := util.ConvertToKilograms(req.FromValue, req.FromUnit)
	if err != nil {
		return model.ConversionResponse{}, fmt.Errorf("ConvertToKilograms: %w", err)
	}

	result, err := util.ConvertFromKilograms(valueInKg, req.ToUnit)
	if err != nil {
		return model.ConversionResponse{}, fmt.Errorf("ConvertFromKilograms: %w", err)
	}

	resp := model.NewConversionResponse(result)
	return resp, nil
}

func (s *convertService) ConvertTemperature(req model.ConversionRequest) (model.ConversionResponse, error) {
	valueInCelsius, err := util.ConvertToCelsius(req.FromValue, req.FromUnit)
	if err != nil {
		return model.ConversionResponse{}, fmt.Errorf("ConvertToCelsius: %w", err)
	}

	result, err := util.ConvertFromCelsius(valueInCelsius, req.ToUnit)
	if err != nil {
		return model.ConversionResponse{}, fmt.Errorf("ConvertFromCelsius: %w", err)
	}

	resp := model.NewConversionResponse(result)
	return resp, nil
}
