package main

import (
	"fmt"

	"github.com/vroomy/plugins"
)

type vpm struct {
	p   *plugins.Plugins
	cfg plugins.Config
}

func (v *vpm) getPluginsMatchingAny(pluginNames ...string) (plugins []string) {
	// Unfiltered, return all plugins
	if len(pluginNames) == 0 {
		return v.cfg.Plugins
	}

	// Filter only plugins contained in pluginNames
	for _, pluginKey := range v.cfg.Plugins {
		// Match name to plugin key suffix (`as <name>`, or last path component `/name`)
		if keyHasSuffixInAny(pluginKey, pluginNames...) {
			plugins = append(plugins, pluginKey)
		}
	}

	return
}

func (v *vpm) addPlugins(pluginNames ...string) (err error) {
	for _, pluginKey := range v.getPluginsMatchingAny(pluginNames...) {
		if err = v.addPlugin(pluginKey); err != nil {
			return
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

func (v *vpm) listPlugins(pluginNames ...string) {
	for _, p := range v.getPluginsMatchingAny(pluginNames...) {
		out.Notification(p)
	}
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

func (v *vpm) testPlugins(pluginNames ...string) (err error) {
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

	if err = v.p.TestAsync(q); err != nil {
		err = fmt.Errorf("error testing plugins: %v", err)
		return
	}

	return
}
