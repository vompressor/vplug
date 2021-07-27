package loader

import (
	"errors"
	plg "plugin"

	"github.com/vompressor/vplug/plugin"
)

func Load(plugPath string, sym string) (*plugin.NTFuncs, *plugin.PluginInfo, error) {
	p, err := plg.Open(plugPath)
	if err != nil {
		return nil, nil, err
	}
	var d *plugin.NTFuncs
	var i *plugin.PluginInfo
	var ok bool

	s, err := p.Lookup("PluginInfo")
	if err != nil {
		return nil, nil, err
	}

	i, ok = s.(*plugin.PluginInfo)
	if !ok {
		return nil, nil, errors.New("plugin info casting error")
	}

	s, err = p.Lookup(sym)

	if err != nil {
		return nil, i, err
	}

	d, ok = s.(*plugin.NTFuncs)
	if !ok {
		return nil, i, errors.New("type casting error")
	}

	return d, i, nil
}

// TODO::
// Load plugins
// func LoadDir(plugDir string) {
// 	entry, err := os.ReadDir(plugDir)
// }
