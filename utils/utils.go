package utils

import (
	"strconv"
	"time"
)

type ArrayInterface []string

func (array *ArrayInterface) InArray(target interface{}) bool {
	for _, element := range *array {
		if target == element {
			return true
		}
	}

	return false
}

func GetNowName(name string) string {
	return strconv.FormatInt(time.Now().Unix(), 10) + "_" + name
}

func StringToInt64(a string) int64 {
	i, _ := strconv.ParseInt(a, 10, 64)
	return i
}
