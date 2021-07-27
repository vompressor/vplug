package main

import "github.com/vompressor/vplug/vplugin"

func main() {}

var VPlugin = vplugin.NewVPlugin().
	AddVPVal("pi", 3.14).
	AddVPFunc("hello", func() error { println("hello"); return nil })
