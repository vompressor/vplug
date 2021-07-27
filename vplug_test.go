package main_test

import (
	"fmt"
	"testing"

	"github.com/vompressor/vplug/loader"
	"github.com/vompressor/vplug/vplugin"
)

func TestLoad(t *testing.T) {
	var p *vplugin.VPlugin

	p, err := loader.Load("vplug.so", "VPlugin")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%#v\n", p.FuncMap)
}
