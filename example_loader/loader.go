package main

import (
	"github.com/vompressor/vplug/loader"
	"github.com/vompressor/vplug/plugin"
)

func main() {
	var pf plugin.NTFuncs

	pf, err := loader.Load("../example_plug/plug.so", "V")
	if err != nil {
		panic(err.Error())
	}

	for _, n := range pf {
		println(n.Name)
	}

}
