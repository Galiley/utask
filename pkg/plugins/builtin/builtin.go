package builtin

import (
	"github.com/ovh/utask/engine/step"
	"github.com/ovh/utask/pkg/plugins"
	pluginapiovh "github.com/ovh/utask/pkg/plugins/builtin/apiovh"
	pluginbatch "github.com/ovh/utask/pkg/plugins/builtin/batch"
	plugincallback "github.com/ovh/utask/pkg/plugins/builtin/callback"
	pluginecho "github.com/ovh/utask/pkg/plugins/builtin/echo"
	pluginemail "github.com/ovh/utask/pkg/plugins/builtin/email"
	pluginhttp "github.com/ovh/utask/pkg/plugins/builtin/http"
	pluginnotify "github.com/ovh/utask/pkg/plugins/builtin/notify"
	pluginping "github.com/ovh/utask/pkg/plugins/builtin/ping"
	pluginscript "github.com/ovh/utask/pkg/plugins/builtin/script"
	pluginssh "github.com/ovh/utask/pkg/plugins/builtin/ssh"
	pluginsubtask "github.com/ovh/utask/pkg/plugins/builtin/subtask"
	plugintag "github.com/ovh/utask/pkg/plugins/builtin/tag"
	"github.com/ovh/utask/pkg/plugins/taskplugin"
)

// RegisterInit takes all builtin init plugins and registers them
func RegisterInit(service *plugins.Service) error {
	for pluginName, pluginSymbol := range map[string]plugins.InitializerPlugin{
		"callback": plugincallback.Init,
	} {
		if err := plugins.RegisterInit(pluginName, pluginSymbol, service); err != nil {
			return err
		}
	}
	return nil
}

// Register takes all builtin plugins and registers them as step executors
func Register() error {
	for _, p := range []taskplugin.PluginExecutor{
		pluginssh.Plugin,
		pluginhttp.Plugin,
		pluginapiovh.Plugin,
		pluginsubtask.Plugin,
		pluginnotify.Plugin,
		pluginecho.Plugin,
		pluginemail.Plugin,
		pluginping.Plugin,
		pluginscript.Plugin,
		plugintag.Plugin,
		plugincallback.Plugin,
		pluginbatch.Plugin,
	} {
		if err := step.RegisterRunner(p.PluginName(), p); err != nil {
			return err
		}
	}
	return nil
}
