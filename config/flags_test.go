package config

import (
	"github.com/galenliu/chip/platform"
	"testing"
)

func TestFlags(t *testing.T) {
	t.Log(platform.GetFatConFile())
	t.Log(platform.GetSysConFile())
	t.Log(platform.GetLocalConFile())
}
