package util

import (
	"../cmd"
	"regexp"
	"strconv"
)

var meminfoFunc = meminfo

// Linux only
func FreeMemoryPercentage() (int, error) {
	result, err := meminfoFunc()
	if err != nil {
		return 0, err
	}

	pattern := `([0-9]+)`
	memories := regexp.MustCompile(pattern).FindAllStringSubmatch(result, -1)
	totalMem, err := strconv.ParseFloat(memories[0][0], 64)
	if err != nil {
		return 0, err
	}

	availableMem, err := strconv.ParseFloat(memories[2][0], 64)
	if err != nil {
		return 0, err
	}

	return (availableMem / totalMem) * 100, nil
}

func meminfo() (string, error) {
	c := cmd.Cmd{
		Name: "cat",
		Args: []string{"/proc/meminfo"},
	}

	return c.Exec()
}
