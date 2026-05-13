package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	cache        = make(map[string]*User) // 模拟缓存
	mu           sync.RWMutex             // 保护缓存
	requestGroup singleflight.Group       // singleflight 实例
)

type User struct {
	Id    int
	Name  string
	Email string
}

// 模拟从数据库获取数据
func GetUserFromDB(userName string) *User {
	fmt.Printf("\033[1;32mQuerying DB\033[0m for key: %s\n", userName)

	time.Sleep(time.Second * 1)

	id, _ := strconv.Atoi(userName[len(userName)-3:])
	fakerUser := &User{
		Id:    int(id),
		Name:  userName,
		Email: userName + "@singleflight.com",
	}
	return fakerUser
}

// 获取数据, 先从缓存读取, 如果缓存未命中, 则从数据库读取
func GetUser(key string) *User {
	// 先从缓存中读取
	mu.RLock()
	val, ok := cache[key]
	mu.RUnlock()
	if ok {
		return val
	}
	fmt.Printf("User \033[1;34m%s\033[0m not in cache\n", key)

	result, _, _ := requestGroup.Do(key, func() (interface{}, error) {
		val := GetUserFromDB(key) // 模拟从数据获取数据

		mu.Lock()
		cache[key] = val
		mu.Unlock()
		return val, nil
	})
	return result.(*User)
}

func main() {
	/*
		第一轮并发请求中, 缓存中为空, fmt.Printf("User \033[1;34m%s\033[0m not in cache\n", key) 的日志打印了 3 次
		fmt.Printf("\033[1;32mQuerying DB\033[0m for key: %s\n", userName) 打印了 2 次, key 为 user_123 的 2 次查询请求被 singleflight 合并了
		第二轮并发请求中, 缓存中已存在数据, 只会打印 fmt.Printf("2. Get user for key: %s -> %+v\n", k, GetUser(k)) 日志

		User user_456 not in cache
		Querying DB for key: user_456
		User user_123 not in cache
		Querying DB for key: user_123
		User user_123 not in cache
		1. Get user for key: user_456 -> &{Id:456 Name:user_456 Email:user_456@singleflight.com}
		1. Get user for key: user_123 -> &{Id:123 Name:user_123 Email:user_123@singleflight.com}
		1. Get user for key: user_123 -> &{Id:123 Name:user_123 Email:user_123@singleflight.com}
		=========================================
		2. Get user for key: user_456 -> &{Id:456 Name:user_456 Email:user_456@singleflight.com}
		2. Get user for key: user_123 -> &{Id:123 Name:user_123 Email:user_123@singleflight.com}
		2. Get user for key: user_123 -> &{Id:123 Name:user_123 Email:user_123@singleflight.com}
	*/

	var wg sync.WaitGroup
	keys := []string{"user_123", "user_123", "user_456"}

	// 第一轮并发查询, 缓存中没有数据, 使用 singleflight 减少 DB 查询
	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			fmt.Printf("1. Get user for key: %s -> %+v\n", k, GetUser(k))
		}(key)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("=========================================")

	for _, key := range keys {
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			fmt.Printf("2. Get user for key: %s -> %+v\n", k, GetUser(k))
		}(key)
	}

	wg.Wait()
}
