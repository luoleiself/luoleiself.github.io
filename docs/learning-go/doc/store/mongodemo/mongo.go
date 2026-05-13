package mongodemo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

var mg *mongo.Client
var mClient *mongo.Client
var ctx = context.Background()

func init() {
	fmt.Println("ObjectId() 一个 12 字节的对象标识符, 4 个字节时间戳表示创建时间的秒数, 每个进程生成一次 5 个字节随机值, 3 个字节递增计数器")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	var uri = "mongodb://root:123456@172.31.218.169:27017/?maxPoolSize=10"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error mongo connecting: %s\n", err)
	}
	mg = client

	// 读写分离配置
	// 创建两个分别连接到 mysql 主库和从库的客户端
}

func init() {
	// 读写分离配置
	// 连接副本集所有节点, 客户端自动路由
	mClient, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://master-host:27017,slave1-host:27017,slave2-host:27017").SetReplicaSet("myReplicaSet").SetReadPreference(readpref.Secondary()).SetWriteConcern(writeconcern.Majority()))
	if err != nil {
		log.Fatalf("Error mongo connecting: %s\n", err)
	}
	log.Print(mClient)
}
