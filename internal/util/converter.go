package util

import "errors"

func ConvertToMeters(value float64, unit string) (float64, error) {
	switch unit {
	case "millimeter", "mm":
		return value / 1000, nil
	case "centimeter", "cm":
		return value / 100, nil
	case "meter", "m":
		return value, nil
	case "kilometer", "km":
		return value * 1000, nil
	case "inch", "in":
		return value * 0.0254, nil
	case "foot", "ft":
		return value * 0.3048, nil
	case "yard", "yd":
		return value * 0.9144, nil
	case "mile", "mi":
		return value * 1609.34, nil
	default:
		return 0, errors.New("unsupported length unit: " + unit)
	}
}

func ConvertFromMeters(value float64, unit string) (float64, error) {
	switch unit {
	case "millimeter", "mm":
		return value * 1000, nil
	case "centimeter", "cm":
		return value * 100, nil
	case "meter", "m":
		return value, nil
	case "kilometer", "km":
		return value / 1000, nil
	case "inch", "in":
		return value / 0.0254, nil
	case "foot", "ft":
		return value / 0.3048, nil
	case "yard", "yd":
		return value / 0.9144, nil
	case "mile", "mi":
		return value / 1609.34, nil
	default:
		return 0, errors.New("unsupported length unit: " + unit)
	}
}

func ConvertToKilograms(value float64, unit string) (float64, error) {
	switch unit {
	case "milligram", "mg":
		return value / 1_000_000, nil
	case "gram", "g":
		return value / 1000, nil
	case "kilogram", "kg":
		return value, nil
	case "ounce", "oz":
		return value * 0.0283495, nil
	case "pound", "lb":
		return value * 0.453592, nil
	default:
		return 0, errors.New("unsupported weight unit: " + unit)
	}
}

func ConvertFromKilograms(value float64, unit string) (float64, error) {
	switch unit {
	case "milligram", "mg":
		return value * 1_000_000, nil
	case "gram", "g":
		return value * 1000, nil
	case "kilogram", "kg":
		return value, nil
	case "ounce", "oz":
		return value / 0.0283495, nil
	case "pound", "lb":
		return value / 0.453592, nil
	default:
		return 0, errors.New("unsupported weight unit: " + unit)
	}
}

func ConvertToCelsius(value float64, unit string) (float64, error) {
	switch unit {
	case "celsius", "C":
		return value, nil
	case "fahrenheit", "F":
		return (value - 32) * 5 / 9, nil
	case "kelvin", "K":
		return value - 273.15, nil
	default:
		return 0, errors.New("unsupported temperature unit: " + unit)
	}
}

func ConvertFromCelsius(value float64, unit string) (float64, error) {
	switch unit {
	case "celsius", "C":
		return value, nil
	case "fahrenheit", "F":
		return value*9/5 + 32, nil
	case "kelvin", "K":
		return value + 273.15, nil
	default:
		return 0, errors.New("unsupported temperature unit: " + unit)
	}
}
