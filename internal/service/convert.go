package service

import (
	"go-unit-converter/internal/model"
	"go-unit-converter/internal/util"
)

type convertService struct {
}

func NewConvertService() *convertService {
	return &convertService{}
}

func (s *convertService) ConvertLength(req *model.ConversionRequest) (*model.ConversionResponse, error) {
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

func (s *convertService) ConvertWeight(req *model.ConversionRequest) (*model.ConversionResponse, error) {
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

func (s *convertService) ConvertTemperature(req *model.ConversionRequest) (*model.ConversionResponse, error) {
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
