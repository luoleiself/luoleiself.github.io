package redisdemo

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.ClusterClient

// var masterRdb *redis.Client
// var clusterRdb *redis.ClusterClient

func init() {
	rdb = redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"172.31.218.169:26379", "172.31.218.169:26380", "172.31.218.169:26381"},
		// 读写分离配置
		RouteByLatency: true, // 从最近的节点读取
		RouteRandomly:  true, // 随机选择从节点读取
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Println("sentinel, new connection established...")
			return nil
		},
	})

	// rdb = redis.NewClient(&redis.Options{
	// 	Addr: "172.31.218.169:6379",
	// 	// Username: "",
	// 	Password:   "",
	// 	DB:         0,
	// 	MaxRetries: 3, // 默认 3, 重试次数
	// 	// PoolSize: 10, // 默认 10, 每个可用的 CPU 最大连接数
	// 	MinIdleConns: 10, // 最小空闲连接数
	// 	OnConnect: func(ctx context.Context, cn *redis.Conn) error {
	// 		log.Println("new connection established...")
	// 		return nil
	// 	},
	// })

	// 读写分离配置
	// 创建两个分别连接到 redis 主节点和从节点的客户端
}

// func init() {
// 	// 读写分离配置, 使用 cluster 模式
// 	clusterRdb = redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs:          []string{"node1:6379", "node2:6379", "node3:6379"},
// 		ReadOnly:       true, // 从节点读取
// 		RouteByLatency: true, // 从低延迟节点读取
// 	})
// 	log.Println(clusterRdb)
// }
