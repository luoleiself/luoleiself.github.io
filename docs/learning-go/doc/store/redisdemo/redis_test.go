package redisdemo

import (
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

// 测试连接
func TestRedisPing(t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		pong, err := rdb.Ping(ctx).Result()
		if err != nil {
			t.Fatalf("Error rdb.Ping(ctx) %s\n", err)
		}
		t.Log("redis ping", pong)
	})
}

// 获取所有 keys
func TestRedisGetAllKeys(t *testing.T) {
	t.Run("Keys", func(t *testing.T) {
		keys, err := rdb.Keys(ctx, "*").Result()
		if err != nil {
			t.Fatalf("Error rdb.Keys(ctx, \"*\") %v\n", err)
		}
		t.Logf("rdb.Keys(ctx, \"*\") %s\n", keys)
		for _, key := range keys {
			ty, err := rdb.Type(ctx, key).Result()
			if err != nil {
				t.Fatalf("Error rdb.Type() %s\n", err)
			}
			t.Logf("%s type is %v\n", key, ty)
		}
	})
}

// 获取字符串
func TestRedisGet(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		name, err := rdb.Get(ctx, "name").Result()
		if err != nil {
			t.Fatalf("Error rdb.Get(ctx, \"name\") %s\n", err)
		}
		t.Logf("rdb.Get(ctx, \"name\") %s\n", name)
	})
}

// 保存字符串仅当没有过期时间
func TestRedisSetNX(t *testing.T) {
	t.Run("SetNX", func(t *testing.T) {
		ok, err := rdb.SetNX(ctx, "nxkey", "nxkey has expiration 6s", time.Second*6).Result()
		if err != nil {
			t.Fatalf("Error rdb.SetNX(\"nxkey\") %s\n", err)
		}
		t.Logf("rdbSetNX(\"nxkey\") %t\n", ok)
		tc := time.After(time.Second * 7)
		<-tc
		nv, err := rdb.Get(ctx, "nxkey").Result()
		t.Log(nv, err)
		if err == redis.Nil {
			t.Logf("rdb.Get(\"nxkey\") does not exist")
		} else if err != nil {
			t.Fatalf("Error rdb.Get(\"nxkey\") %s\n", err)
		} else {
			t.Logf("rdb.Get(\"nxkey\") %v\n", nv)
		}
	})
}

// Hash
func TestRedisHash(t *testing.T) {
	t.Run("Hset", func(t *testing.T) {
		len, err := rdb.HSet(ctx, "runoob", map[string]interface{}{"name": "redis", "age": 18, "addr": "beijing"}).Result()
		if err != nil {
			t.Fatalf("Error rdb.HSet() %s\n", err)
		}
		t.Logf("rdb.HSet() %d\n", len)

		l, err := rdb.HLen(ctx, "runoob").Result()
		if err != nil {
			t.Fatalf("Error rdb.HLen() %s\n", err)
		}
		t.Logf("rdb.HLen() %d\n", l)
	})
	t.Run("HKeys", func(t *testing.T) {
		keys, err := rdb.HKeys(ctx, "runoob").Result()
		if err != nil {
			t.Fatalf("Error rdb.HKeys() %s\n", err)
		}
		t.Logf("rdb.HKeys() %s\n", keys)
	})
	t.Run("HVals", func(t *testing.T) {
		vals, err := rdb.HVals(ctx, "runoob").Result()
		if err != nil {
			t.Fatalf("Error rdb.HVals() %s\n", err)
		}
		t.Logf("rdb.HVals() %s\n", vals)
	})
	t.Run("HGetAll", func(t *testing.T) {
		all, err := rdb.HGetAll(ctx, "runoob").Result()
		if err != nil {
			t.Fatalf("Error rdb.HGetAll() %s\n", err)
		}
		t.Logf("rdb.HGetAll() %v\n", all)
	})
	t.Run("HScan", func(t *testing.T) {
		keys, cursor, err := rdb.HScan(ctx, "runoob", 0, "*", 100).Result()
		if err != nil {
			t.Fatalf("Error rdb.HScan() %s\n", err)
		}
		t.Logf("keys: %s, cursor %d\n", keys, cursor)
	})
}

// List
func TestRedisList(t *testing.T) {
	t.Run("LIndex", func(t *testing.T) {
		lastV, err := rdb.LIndex(ctx, "num", -1).Result()
		if err != nil {
			t.Fatalf("Error rdb.LIndex(ctx, \"num\", -1) %s\n", err)
		}
		t.Logf("rdb.LIndex(ctx, \"num\", -1) %s %T\n", lastV, lastV)
		v2, err := strconv.ParseInt(lastV, 10, 10)
		if err != nil {
			t.Fatalf("Error strconv.ParseInt(lastV, 10, 10) %s\n", err)
		}
		t.Logf("strconv.ParseInt(lastV, 10, 10) %d %T\n", v2, v2)
	})
	t.Run("LRange", func(t *testing.T) {
		list, err := rdb.LRange(ctx, "num", 0, -1).Result()
		if err != nil {
			t.Fatalf("Error rdb.LRange(ctx, \"num\", 0, -1) %s\n", err)
		}
		t.Logf("rdb.LRange(ctx, \"num\", 0, -1) %s\n", list)
	})
	t.Run("RPop", func(t *testing.T) {
		result, err := rdb.RPop(ctx, "num").Result()
		if err != nil {
			t.Fatalf("Error rdb.RPop() %s\n", err)
		}
		t.Logf("rdb.RPop() %s\n", result)
	})
}
