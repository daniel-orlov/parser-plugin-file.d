package parser_plugin_file_d

import (
	"github.com/ozonru/file.d/fd"
	"github.com/ozonru/file.d/logger"
	"github.com/ozonru/file.d/pipeline"
)

type Plugin struct {
	config *Config
}

type Config struct {
}

func init() {
	fd.DefaultPluginRegistry.RegisterAction(&pipeline.PluginStaticInfo{
		Type:    "parse_nginx_logs",
		Factory: factory,
	})
}

func factory() (pipeline.AnyPlugin, pipeline.AnyConfig) {
	return &Plugin{}, &Config{}
}

func (p *Plugin) Start(_ pipeline.AnyConfig, _ *pipeline.ActionPluginParams) {

}

func (p *Plugin) Stop() {

}

func (p *Plugin) Do(_ *pipeline.Event) pipeline.ActionResult {
	logger.Infof("THIS FUCKING WORKS!")
	return pipeline.ActionDiscard //currently just discards every action it encounters
}
