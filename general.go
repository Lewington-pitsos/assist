package help

import (
	"strconv"
	"time"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Wait(milliseconds int) {
	if milliseconds > 0 {
		time.Sleep(time.Millisecond * time.Duration(milliseconds))
	}
}

// func CurrentDay() string {
// 	return time.Now().Local().Format("01-02-2006")
// }

func ToFloat32(str string) float32 {
	f, err := strconv.ParseFloat(str, 32)
	Check(err)
	return float32(f)
}

func Flattened(values [][]int) []int {
	final := make([]int, 10000)
	for _, list := range values {
		for _, value := range list {
			final = append(final, value)
		}
	}

	return final
}
