package loader

import (
	plg "plugin"

	"github.com/vompressor/vplug/vplugin"
)

func Load(plugPath string, sym string) (*vplugin.VPlugin, error) {
	p, err := plg.Open(plugPath)
	if err != nil {
		return nil, err
	}

	s, err := p.Lookup(sym)

	if err != nil {
		return nil, err
	}

	d := s.(**vplugin.VPlugin)
	// if !ok {
	// 	return nil, errors.New("type casting error")
	// }

	return *d, nil
}

// TODO::
// Load plugins
// func LoadDir(plugDir string) {
// 	entry, err := os.ReadDir(plugDir)
// }
