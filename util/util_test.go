package util

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestFreeMemoryPercentage(t *testing.T) {
	setMeminfo("MemTotal 300 kb MemFree 100 kb MemAvaialble 150 kb")

	percentage, err := FreeMemoryPercentage()
	if err != nil {
		t.Log("error:", err)
	}
	assert.Equal(t, 50, percentage)
}

func setMeminfo(s string) {
	meminfoFunc = func() (string, error) {
		return s, nil
	}
}
