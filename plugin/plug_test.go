package plugin_test

import (
	"errors"
	"testing"

	"github.com/vompressor/vplug/plugin"
)

func TestFunc(t *testing.T) {
	ntf := plugin.NewNTFunc("echo", func(n string) error { println(n); return errors.New(n) })

	for _, n := range ntf.InTypes {
		println(n.String())
	}
	println(ntf.Name)
	err := ntf.Call("hello")
	if err != nil {
		t.Log(err.Error())
	}
}

func TestNTFs(t *testing.T) {
	ntfs := plugin.NewNTFuncs([]plugin.NTFuncHelper{
		plugin.NewNTFunc("hello", func() error { println("hello~"); return nil }),
		plugin.NewNTFunc("bye", func() error { println("bye~"); return nil }),
		plugin.NewNTFunc("how are you", func() error { println("how are you~"); return nil }),
	})

	for k, v := range ntfs {
		println(k)
		v.Call()
	}
}
