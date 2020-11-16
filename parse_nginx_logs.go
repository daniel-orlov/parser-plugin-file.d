package parser_plugin_file_d

import (
	"fmt"
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
	//p.config = config.(*Config)
}

func (p *Plugin) Stop() {

}

func (p *Plugin) Do(event *pipeline.Event) pipeline.ActionResult {
	logger.Info("THIS FUCKING WORKS!")
	fmt.Println("THIS PROBABLY WORKS...")
	fmt.Println(event.Root.EncodeToString() + "ANOTHER TIME")
	return pipeline.ActionPass//currently just passes every action it encounters
}
