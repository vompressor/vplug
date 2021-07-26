package loader

import (
	"errors"
	plg "plugin"

	"github.com/vompressor/vplug/plugin"
)

func Load(plugPath string, sym string) (*plugin.NTFuncs, error) {
	p, err := plg.Open(plugPath)
	if err != nil {
		return nil, err
	}

	s, err := p.Lookup(sym)

	if err != nil {
		return nil, err
	}
	var d *plugin.NTFuncs
	var ok bool
	d, ok = s.(*plugin.NTFuncs)
	if !ok {
		return nil, errors.New("load symbolic error")
	}

	return d, nil
}

// TODO::
// Load plugins
// func LoadDir(plugDir string) {
// 	entry, err := os.ReadDir(plugDir)
// }
