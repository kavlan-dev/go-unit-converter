package service

import (
	"go-unit-converter/internal/model"
	"go-unit-converter/internal/util"
)

func (s *Service) ConvertLength(req *model.ConversionRequest) (*model.ConversionResponse, error) {
	valueInMeters, err := util.ConvertToMeters(req.FromValue, req.FromUnit)
	if err != nil {
		return nil, err
	}

	result, err := util.ConvertFromMeters(valueInMeters, req.ToUnit)
	if err != nil {
		return nil, err
	}

	return &model.ConversionResponse{Result: result}, nil
}

func (s *Service) ConvertWeight(req *model.ConversionRequest) (*model.ConversionResponse, error) {
	valueInKg, err := util.ConvertToKilograms(req.FromValue, req.FromUnit)
	if err != nil {
		return nil, err
	}

	result, err := util.ConvertFromKilograms(valueInKg, req.ToUnit)
	if err != nil {
		return nil, err
	}

	return &model.ConversionResponse{Result: result}, nil
}

func (s *Service) ConvertTemperature(req *model.ConversionRequest) (*model.ConversionResponse, error) {
	valueInCelsius, err := util.ConvertToCelsius(req.FromValue, req.FromUnit)
	if err != nil {
		return nil, err
	}

	result, err := util.ConvertFromCelsius(valueInCelsius, req.ToUnit)
	if err != nil {
		return nil, err
	}

	return &model.ConversionResponse{Result: result}, nil
}
