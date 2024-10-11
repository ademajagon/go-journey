package main

import (
	"fmt"
	"strconv"
)

func ConvertToString[T int64 | uint64](value T) string {
	switch any(value).(type) {
	case int64:
		return strconv.FormatInt(int64(value), 10)
	case uint64:
		return strconv.FormatUint(uint64(value), 10)
	}
	return ""
}

func main() {
	var signedValue int64 = -12345
	var unsignedValue uint64 = 12345

	fmt.Println("Signed int64 to string", ConvertToString(signedValue))
	fmt.Println("Unsigned uint64 to string", ConvertToString(unsignedValue))
}
