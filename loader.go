package main

import (
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/s0rbus/bookshop-micro-rpg/api"
)

var p *plugin.Plugin

func LoadPlugins(dir string, name string) (api.Expansion, error) {
	// Glob - Gets the plugin to be loaded
	plugins, err := filepath.Glob(fmt.Sprintf("%s/%s.so", dir, name))
	if err != nil {
		return nil, err
	}
	// Open - Loads the plugin
	fmt.Printf("Loading expansion plugin %s\n", plugins[0])
	p, err = plugin.Open(plugins[0])
	if err != nil {
		return nil, err
	}
	GetExpansion, err := p.Lookup("GetExpansion")
	if err != nil {
		return nil, err
	}
	exp, err := GetExpansion.(func() (api.Expansion, error))()
	if err != nil {
		return nil, err
	}
	return exp, nil
}
