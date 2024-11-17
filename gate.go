package main

import (
	"context"

	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/plugins/dynamicserverapi"
	"go.minekube.com/gate/redisdb"
	"go.minekube.com/gate/server"
)

func main() {
	// Initialize Redis client
	redisdb.RedisClient = redisdb.InitRedis()

	// Add plugins to Gate
	proxy.Plugins = append(proxy.Plugins,
		dynamicserverapi.Plugin,
		proxy.Plugin{
			Name: "PlayerCountUpdater",
			Init: func(ctx context.Context, p *proxy.Proxy) error {
				// Initialize the player count updater
				return server.InitPlayerCountUpdater(p)
			},
		},
	)

	// Run Gate
	gate.Execute()
}
