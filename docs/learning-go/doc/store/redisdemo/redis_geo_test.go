package redisdemo

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

// geo
func TestRedisGeo(t *testing.T) {
	t.Run("GeoAdd", func(t *testing.T) {
		len, err := rdb.GeoAdd(ctx, "citys",
			[]*redis.GeoLocation{
				{Name: "beijing", Longitude: 116.405285, Latitude: 39.904989},
				{Name: "shanghai", Longitude: 121.472644, Latitude: 31.231706},
				{Name: "guangzhou", Longitude: 113.280637, Latitude: 23.125178},
				{Name: "shenzhen", Longitude: 114.05571, Latitude: 22.52245},
				{Name: "chongqing", Longitude: 106.504962, Latitude: 29.533155},
				{Name: "sansha", Longitude: 112.34882, Latitude: 16.831039},
				{Name: "zhengzhou", Longitude: 113.665412, Latitude: 34.757975},
				{Name: "lanzhou", Longitude: 103.823557, Latitude: 36.058039},
				{Name: "xi'an", Longitude: 108.948024, Latitude: 34.263161},
				{Name: "nanchang", Longitude: 115.892151, Latitude: 28.676493},
				{Name: "wuhan", Longitude: 114.311586, Latitude: 30.598467},
				{Name: "shijiazhuang", Longitude: 114.521529, Latitude: 38.048312},
				{Name: "ji'nan", Longitude: 117.126399, Latitude: 36.656554},
				{Name: "changsha", Longitude: 112.945470, Latitude: 28.234890},
				{Name: "nanjing", Longitude: 118.802422, Latitude: 32.064652},
				{Name: "hangzhou", Longitude: 120.215510, Latitude: 30.253082},
				{Name: "hefei", Longitude: 117.233441, Latitude: 31.826578},
				{Name: "huhehaote", Longitude: 111.755512, Latitude: 40.848422},
				{Name: "taiyuan", Longitude: 112.556396, Latitude: 37.876990},
			}...,
		).Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoAdd() %s\n", err)
		}
		t.Logf("GeoAdd success %d\n", len)

		len, err = rdb.ZCard(ctx, "citys").Result()
		if err != nil {
			t.Fatalf("Error rdb.ZCard() %s\n", err)
		}
		t.Logf("rdb.ZCard() %d\n", len)
	})
	t.Run("ZRange", func(t *testing.T) {
		keys, err := rdb.ZRange(ctx, "citys", 0, -1).Result()
		if err != nil {
			t.Fatalf("Error rdb.ZRange() %s\n", err)
		}
		t.Logf("rdb.ZRange() %s\n", keys)

		len, err := rdb.GeoAdd(ctx, "citys", &redis.GeoLocation{
			Name:      "shenzhen",
			Longitude: 113.88311,
			Latitude:  22.55371,
		}).Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoAdd() %s\n", err)
		}
		t.Logf("GeoAdd success %d\n", len)

		len, err = rdb.ZCard(ctx, "citys").Result()
		if err != nil {
			t.Fatalf("Error rdb.ZCard() %s\n", err)
		}
		t.Logf("rdb.ZCard() %d\n", len)

		keys, err = rdb.ZRange(ctx, "citys", 0, -1).Result()
		if err != nil {
			t.Fatalf("Error rdb.ZRange() %s\n", err)
		}
		t.Logf("rdb.ZRange() %s\n", keys)
	})
	t.Run("GeoPos", func(t *testing.T) {
		pos, err := rdb.GeoPos(ctx, "citys", []string{"beijing", "chongqing"}...).Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoPos() %s\n", err)
		}
		for _, v := range pos {
			t.Log(v)
		}
	})
	t.Run("GeoHash", func(t *testing.T) {
		h, err := rdb.GeoHash(ctx, "citys", []string{"beijing", "chongqing", "hello"}...).Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoHash() %s\n", err)
		}
		t.Logf("rdb.GeoHash() %s\n", h)
	})
	t.Run("GeoDist", func(t *testing.T) {
		dist, err := rdb.GeoDist(ctx, "citys", "beijing", "shanghai", "KM").Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoDist() %s\n", err)
		}
		t.Logf("rdb.GeoDist() %fKM\n", dist)
	})
	t.Run("GeoSearch", func(t *testing.T) {
		searchQuery := redis.GeoSearchQuery{
			Member:     "nanchang",
			Radius:     1200,
			RadiusUnit: "KM",
			Sort:       "DESC",
		}
		s, err := rdb.GeoSearch(ctx, "citys", &searchQuery).Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoSearch() %s\n", err)
		}
		t.Logf("rdb.GeoSearch() %s\n", s)
	})
	t.Run("GeoSearchLocation", func(t *testing.T) {
		searchQuery := redis.GeoSearchLocationQuery{
			GeoSearchQuery: redis.GeoSearchQuery{
				Member:     "nanchang",
				Radius:     1200,
				RadiusUnit: "KM",
				Sort:       "DESC",
			},
			WithCoord: true,
			WithDist:  true,
			WithHash:  true,
		}
		loc, err := rdb.GeoSearchLocation(ctx, "citys", &searchQuery).Result()
		if err != nil {
			t.Fatalf("Error rdb.GeoSearchLocation() %s\n", err)
		}
		for _, v := range loc {
			t.Logf("%+v\n", v)
		}
	})
}
