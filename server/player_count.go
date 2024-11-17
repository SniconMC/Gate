package server

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/redisdb"
	"go.minekube.com/gate/server/aggregator"
)

// Mutex to avoid race conditions
var playerCountMutex sync.Mutex

// InitPlayerCountUpdater initializes periodic updates for per-proxy player counts.
func InitPlayerCountUpdater(p *proxy.Proxy) error {
	logger := log.Default()

	// Get proxy name from the environment variable
	proxyName := os.Getenv("PROXY_NAME")
	if proxyName == "" {
		logger.Println("Environment variable PROXY_NAME is not set.")
		return nil
	}
	proxyPlayerCountKey := proxyName + ":playercount"

	// Start a goroutine to update player count every 5 seconds
	go func() {
		for {
			time.Sleep(5 * time.Second)
			updateProxyPlayerCount(p, proxyPlayerCountKey, logger)
		}
	}()
	aggregator.StartNetworkAggregator("proxy", 10*time.Second)
	logger.Println("Per-proxy player count updater initialized.")
	return nil
}

// updateProxyPlayerCount updates the current proxy's player count in Redis.
func updateProxyPlayerCount(p *proxy.Proxy, key string, logger *log.Logger) {
	playerCountMutex.Lock()
	defer playerCountMutex.Unlock()

	ctx := context.Background()

	// Get the current player count
	count := p.PlayerCount()

	// Save the updated count to Redis
	err := redisdb.RedisClient.Set(ctx, key, count, 0).Err()
	if err != nil {
		logger.Printf("Failed to update player count for key '%s' in Redis: %v", key, err)
		return
	}

	logger.Printf("Updated player count for key '%s': %d", key, count)
}
