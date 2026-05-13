package redisdemo

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

var luaScript = `
	local key = KEYS[1]

	if redis.call("EXISTS", key) == 1 then
		local keyType = redis.call("TYPE", key)
		if keyType["ok"] == "string" then
			return redis.call("GET", key)
		elseif keyType["ok"] == "hash" then
			return redis.call("HGETALL", key)
		elseif keyType["ok"] == "list" then
			local start = ARGV[1]
			local stop = ARGV[2]
			return redis.call("LRANGE", key, start, stop)
		elseif keyType["ok"] == 'zset' then
			local member = ARGV[1]
			local radius = ARGV[2]
			local unit = ARGV[3]
			local ordered = ARGV[4]
			local withCoord = ARGV[5]
			local withDist = ARGV[6]
			local withHash = ARGV[7]	
			return redis.call("GEOSEARCH", key, "FROMMEMBER", member, "BYRADIUS", radius, unit, ordered, withCoord, withDist, withHash)
		else 
			return "Unsupported key type"
		end
		return "Updated"
	else
		return "Key not exists"
	end
`

func TestRedisLua(t *testing.T) {
	t.Run("luaScript-GET", func(t *testing.T) {
		keys := []string{"name"}
		// 方式1
		v1, err := rdb.Eval(ctx, luaScript, keys).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v1", v1)

		// 方式2
		script := redis.NewScript(luaScript)
		v2, err := script.Run(ctx, rdb, keys).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v2", v2)
		// 方式2.1
		v3 := script.Eval(ctx, rdb, keys, 0, -1).Val()
		t.Log("v3", v3)

		v4, err := rdb.Eval(ctx, luaScript, []string{"newname"}).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v4", v4)
	})
	t.Run("luaScript-hash", func(t *testing.T) {
		keys := []string{"runoob"}
		// 方式1
		v1, err := rdb.Eval(ctx, luaScript, keys).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v1", v1)

		// 方式2
		script := redis.NewScript(luaScript)
		v2, err := script.Run(ctx, rdb, keys).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v2", v2)
		// 方式2.1
		v3 := script.Eval(ctx, rdb, keys, 0, -1).Val()
		t.Log("v3", v3)
	})
	t.Run("luaScipt-list", func(t *testing.T) {
		keys := []string{"num"}
		// 方式1
		v1, err := rdb.Eval(ctx, luaScript, keys, 0, -1).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v1", v1)

		// 方式2
		script := redis.NewScript(luaScript)
		v2, err := script.Run(ctx, rdb, keys, 0, -1).Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Log("v2", v2)
		// 方式2.1
		v3 := script.Eval(ctx, rdb, keys, 0, -1).Val()
		t.Log("v3", v3)
	})
	t.Run("luaScript-geo", func(t *testing.T) {
		keys := []string{"citys"}
		v1, err := rdb.Eval(ctx, luaScript, keys, "nanchang", 1200, "KM", "DESC", "WITHCOORD", "WITHDIST", "WITHHASH").Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Logf("v1 %+v\n", v1)

		v2, err := rdb.Eval(ctx, luaScript, keys, "zhengzhou", 5000, "mi", "ASC", "WITHCOORD", "WITHDIST", "WITHHASH").Result()
		if err != nil {
			t.Fatalf("Error rdb.Eval() %s\n", err)
		}
		t.Logf("v2 %+v\n", v2)
	})
}
