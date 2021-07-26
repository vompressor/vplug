package main

import "github.com/vompressor/vplug/plugin"

func main() {}

var V = plugin.NewNTFuncs([]plugin.NTFuncHelper{
	plugin.NewNTFunc("hello", func() error { println("hello~"); return nil }),
	plugin.NewNTFunc("bye", func() error { println("bye~"); return nil }),
	plugin.NewNTFunc("how are you", func() error { println("how are you~"); return nil }),
})
