package vplugin

import (
	"fmt"
	"reflect"
)

type VFuncMap map[string]VPluginFunc
type VValMap map[string]VPluginVal

type VPluginInfo struct {
	Name        string
	Version     string
	Description string
}

type VPlugin struct {
	Info    VPluginInfo
	FuncMap VFuncMap
	ValMap  VValMap
}

func NewVPlugin(name string, version string, description string) *VPlugin {
	vp := &VPlugin{}
	vp.FuncMap = make(VFuncMap)
	vp.ValMap = make(VValMap)

	vp.Info.Name = name
	vp.Info.Description = description
	vp.Info.Version = version

	return vp
}

type Vf interface{}

type VPluginFunc struct {
	Name    string
	InTypes []reflect.Type
	ff      reflect.Value
}

func (vp *VPlugin) AddVPFunc(name string, f Vf) *VPlugin {
	x := reflect.TypeOf(f)
	if x.Kind() != reflect.Func {
		panic("name" + name + " is not function")
	}

	nin := x.NumIn()

	h := VPluginFunc{}

	h.InTypes = make([]reflect.Type, nin)
	for i := range h.InTypes {
		h.InTypes[i] = x.In(i)
	}

	h.ff = reflect.ValueOf(f)
	h.Name = name

	vp.FuncMap[name] = h

	return vp
}

func (n VPluginFunc) Call(a ...interface{}) error {
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
		panic("return type is not error")
	}
	return err
}

func (n VPluginFunc) TypeCheck(a ...interface{}) error {
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

type VPluginVal struct {
	Name string
	Type reflect.Type
	Val  interface{}
}

func (vp *VPlugin) AddVPVal(name string, v interface{}) *VPlugin {
	z := VPluginVal{}

	z.Type = reflect.TypeOf(v)
	z.Name = name
	z.Val = v

	vp.ValMap[name] = z

	return vp
}
