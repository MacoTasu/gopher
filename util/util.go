package util

import (
	"../cmd"
	"strconv"
	"strings"
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

	memories := strings.Split(string(result), " ")

	totalMem, err := strconv.Atoi(memories[7])
	if err != nil {
		return 0, err
	}

	freeMem, err := strconv.Atoi(memories[21])
	if err != nil {
		return 0, err
	}

	return float64((freeMem / totalMem)), nil
}
