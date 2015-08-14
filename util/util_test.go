package util

import (
	"github.com/github/hub/Godeps/_workspace/src/github.com/bmizerany/assert"

	"testing"
)

func TestFreeMemoryPercentage(t *testing.T) {
	setMeminfo("MemTotal 300 kb MemFree 100 kb MemAvaialble 150 kb")

	percentage, _ := FreeMemoryPercentage()
	assert.Equal(t, 0.5, percentage)
}

func setMeminfo(s string) {
	meminfoFunc = func() (string, error) {
		return s, nil
	}
}
