package main

import (
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/plugins/dynamicserverapi"
)

func main() {

	proxy.Plugins = append(proxy.Plugins,
		// your plugins
		dynamicserverapi.Plugin,
	)

	gate.Execute()
}
