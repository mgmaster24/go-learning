package gol_conversions

import (
	"math"
	"strconv"
)

func StringsToFloat64s(strs []string) ([]float64, error) {
	floats := make([]float64, len(strs))
	for stringIdx, stringVal := range strs {
		val, e := strconv.ParseFloat(stringVal, 64)
		if e != nil {
			return nil, e
		}

		floats[stringIdx] = val
	}

	return floats, nil
}

func TruncateFloat(val float64) float64 {
	return math.Round(val*100) / 100
}
