package aggregator

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"go.minekube.com/gate/redisdb"
)

const networkPlayerCountKey = "network:playercount"

// StartNetworkAggregator starts a periodic aggregation of player counts.
func StartNetworkAggregator(proxyPrefix string, interval time.Duration) {
	logger := log.Default()

	go func() {
		for {
			time.Sleep(interval)
			aggregateNetworkPlayerCount(proxyPrefix, logger)
		}
	}()
}

// aggregateNetworkPlayerCount sums up player counts from all proxies and updates the network-wide count in Redis.
func aggregateNetworkPlayerCount(proxyPrefix string, logger *log.Logger) {
	ctx := context.Background()

	// Get all keys for proxies
	keys, err := redisdb.RedisClient.Keys(ctx, proxyPrefix+"*").Result()
	if err != nil {
		logger.Printf("Failed to fetch proxy keys from Redis: %v", err)
		return
	}

	// Sum up player counts
	totalCount := 0
	for _, key := range keys {
		countStr, err := redisdb.RedisClient.Get(ctx, key).Result()
		if err != nil && !strings.Contains(err.Error(), "redis: nil") {
			logger.Printf("Failed to fetch player count for key '%s': %v", key, err)
			continue
		}
		logger.Printf(countStr)
		// Convert count to an integer
		count := 0
		if countStr != "" {
			count, err = strconv.Atoi(countStr)
			if err != nil {
				logger.Printf("Invalid player count for key '%s': %v", key, err)
				continue
			}
		}

		totalCount += count
	}

	// Update the network-wide player count in Redis
	err = redisdb.RedisClient.Set(ctx, networkPlayerCountKey, totalCount, 0).Err()
	if err != nil {
		logger.Printf("Failed to update network player count in Redis: %v", err)
		return
	}

	logger.Printf("Updated network player count in Redis: %d", totalCount)
}
