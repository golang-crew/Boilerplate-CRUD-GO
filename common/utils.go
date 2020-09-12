package common

import "strconv"

func ConvertStringToUint(data string) (uintData uint, err error) {
	u64, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return
	}

	return uint(u64), nil
}

func ConvertStringToInt(data string) (uintData int, err error) {
	u64, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		return
	}

	return int(u64), nil
}
