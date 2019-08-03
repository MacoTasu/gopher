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
	assert.Equal(t, 0.5, float64(percentage)/100)
}

func setMeminfo(s string) {
	meminfoFunc = func() (string, error) {
		return s, nil
	}
}
