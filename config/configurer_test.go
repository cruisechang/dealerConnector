package config

import (
	"math"
	"testing"

)
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}

func TestLoadConfigData(t *testing.T) {

	conf, err := NewConfigurer("config.json")
	if err != nil {
		t.Fatalf("TestLoadConfigData error:%v", err)
	}

	t.Logf("config version=%s", conf.Version())

}
