package utils

import "strconv"

func ToString(obj interface{}) string {
	if m, ok := obj.(string); ok {
		return m
	}
	if m, ok := obj.(float64); ok {
		return strconv.FormatFloat(m, 'f', -1, 64)
	}
	if m, ok := obj.(int); ok {
		return strconv.Itoa(m)
	}
	if m, ok := obj.(int64); ok {
		return strconv.FormatInt(m, 10)
	}
	if m, ok := obj.(bool); ok {
		return strconv.FormatBool(m)
	}

	return ""
}
