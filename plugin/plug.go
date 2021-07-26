package plugin

import (
	"fmt"
	"reflect"
)

// Plugin info part

type PluginInfo struct {
	Name    string
	Version string
}

// Function Part
type NonTypeFunc func(...interface{}) error
type NTFuncHelper struct {
	Name    string
	InTypes []reflect.Type
	ff      reflect.Value
}

func NewNTFunc(name string, f interface{}) NTFuncHelper {
	x := reflect.TypeOf(f)
	nin := x.NumIn()

	h := NTFuncHelper{}

	h.InTypes = make([]reflect.Type, nin)
	for i := range h.InTypes {
		h.InTypes[i] = x.In(i)
	}

	h.ff = reflect.ValueOf(f)
	h.Name = name
	return h
}

func (n *NTFuncHelper) Call(a ...interface{}) error {
	refVal := make([]reflect.Value, len(a))
	for i, v := range a {
		refVal[i] = reflect.ValueOf(v)
	}
	e := n.ff.Call(refVal)[0]
	if e.IsNil() {
		return nil
	}
	err, ok := e.Interface().(error)
	if !ok {
		panic("return error")
	}
	return err
}

func (n *NTFuncHelper) TypeCheck(a ...interface{}) error {
	inTypeLen := len(n.InTypes)
	inputLen := len(a)

	if inTypeLen != inputLen {
		return fmt.Errorf("input args: %d, need args: %d", inputLen, inTypeLen)
	}

	for i, v := range a {
		inputType := reflect.TypeOf(v)
		if inputType != n.InTypes[i] {
			return fmt.Errorf("num: %d, input type: %s, need type: %s", i, inputType.String(), n.InTypes[i].String())
		}
	}
	return nil
}

type GetNTFuncs interface {
	Get() *NTFuncs
}

type NTFuncs map[string]NTFuncHelper

func NewNTFuncs(f []NTFuncHelper) NTFuncs {
	ret := make(NTFuncs)

	for _, v := range f {
		ret[v.Name] = v
	}

	return ret
}

func (n *NTFuncs) Get() *NTFuncs {
	return n
}

// Value part
// TODO::
// Values
// type NonTypeVal interface{}

// type NTValHelper struct {
// 	val  reflect.Value
// 	Type reflect.Kind
// }

// func NewNTVal(v interface{}) NTValHelper {
// 	return NTValHelper{
// 		val:  reflect.ValueOf(v),
// 		Type: reflect.TypeOf(v).Kind(),
// 	}
// }

// func (nt NTValHelper) Get() interface{} {
// 	if nt.val.IsNil() {
// 		return nil
// 	}
// 	return nt.val.Interface()
// }

// type NTVals []NTValHelper
