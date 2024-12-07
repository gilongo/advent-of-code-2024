package utils

import (
	"bytes"
	"os"
	"runtime"
)

func GetBreakLineToken() string {
	var breakLineToken string
	if runtime.GOOS == "windows" {
		breakLineToken = "\r\n"
	} else {
		breakLineToken = "\n"
	}
	return breakLineToken
}

func ReadInputFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return data
}

func SplitData(data []byte, token string) [][]byte {
	if len(data) < 0 {
		panic("data is empty")
	}

	return bytes.Split(data, []byte(token))
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
