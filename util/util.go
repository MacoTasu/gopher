package util

import (
	"../cmd"
	"regexp"
	"strconv"
)

// Linux only
func FreeMemoryPercentage() (float64, error) {
	c := cmd.Cmd{
		Name: "cat",
		Args: []string{"/proc/meminfo"},
	}

	result, err := c.Exec()
	if err != nil {
		return 0, err
	}

	pattern := `^([0-9]+)`
	memories := regexp.MustCompile(pattern).FindStringSubmatch(result)

	totalMem, err := strconv.Atoi(memories[0])
	if err != nil {
		return 0, err
	}

	availableMem, err := strconv.Atoi(memories[2])
	if err != nil {
		return 0, err
	}

	return float64((availableMem / totalMem)), nil
}
