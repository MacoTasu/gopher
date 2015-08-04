package util

import (
	"../cmd"
	"strconv"
	"strings"
)

// Linux only
func FreeMemoryPercentage() (int, error) {
	c := cmd.Cmd{
		Name: "free",
		Args: []string{"|", "grep", "-E", "Mem"},
	}

	result, err := c.Exec()
	if err != nil {
		return 0, err
	}

	memories := strings.Split(string(result), " ")

	totalMem, err := strconv.Atoi(memories[1])
	if err != nil {
		return 0, err
	}

	freeMem, err := strconv.Atoi(memories[3])
	if err != nil {
		return 0, err
	}

	bufferMem, err := strconv.Atoi(memories[5])
	if err != nil {
		return 0, err
	}

	cachedMem, err := strconv.Atoi(memories[6])
	if err != nil {
		return 0, err
	}

	return totalMem/freeMem + bufferMem + cachedMem, nil
}
