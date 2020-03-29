package main

import (
	"fmt"
	"strings"

	"github.com/vroomy/plugins"
)

type vpm struct {
	p   *plugins.Plugins
	cfg plugins.Config
}

func (v *vpm) addPlugins(pluginNames ...string) (err error) {
	for _, pluginKey := range v.cfg.Plugins {
		if len(pluginNames) == 0 {
			if err = v.addPlugin(pluginKey); err != nil {
				return
			}
		} else {
			// Filter
			for _, name := range pluginNames {
				if strings.HasSuffix(pluginKey, name) {
					if err = v.addPlugin(pluginKey); err != nil {
						return
					}

					break
				}
			}
		}
	}

	return
}

func (v *vpm) addPlugin(pluginKey string) (err error) {
	if _, err = v.p.New(pluginKey, true); err != nil {
		err = fmt.Errorf("error creating new plugin for key \"%s\": %v", pluginKey, err)
		return
	}

	return
}

func (v *vpm) updatePlugins(pluginNames ...string) (err error) {
	if v.p, err = plugins.New("plugins"); err != nil {
		err = fmt.Errorf("error initializing plugins manager: %v", err)
		return
	}

	if len(v.cfg.Plugins) == 0 {
		return
	}

	if err = v.addPlugins(pluginNames...); err != nil {
		return
	}

	if err = v.p.Retrieve(); err != nil {
		err = fmt.Errorf("error retrieving plugins: %v", err)
		return
	}

	if err = v.p.BuildAsync(q); err != nil {
		err = fmt.Errorf("error building plugins: %v", err)
		return
	}

	return
}

func (v *vpm) buildPlugins(pluginNames ...string) (err error) {
	if v.p, err = plugins.New("plugins"); err != nil {
		err = fmt.Errorf("error initializing plugins manager: %v", err)
		return
	}

	if len(v.cfg.Plugins) == 0 {
		return
	}

	if err = v.addPlugins(pluginNames...); err != nil {
		return
	}

	if err = v.p.BuildAsync(q); err != nil {
		err = fmt.Errorf("error building plugins: %v", err)
		return
	}

	return
}
