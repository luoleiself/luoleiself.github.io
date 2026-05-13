package main

import (
	"net/http"
	"time"
)

// loggingMiddleware 中间件函数
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		llog.Println("----start----")
		llog.Printf("loggingMiddleware. request method is %v, request url is %v\n", r.Method, r.URL)
		next.ServeHTTP(w, r) // Call the next handler in the chain
		duration := time.Since(start)
		llog.Printf("[%s] %s %s %v\n", r.Method, r.URL.Path, r.URL.RawQuery, duration)
		llog.Println("----end----")
	})
}

// recoveryMiddleware 中间件函数
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				llog.Printf("recoveryMiddleware. panic: %v\n", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware 中间件函数
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Language, Cache-Control")
		// w.Header().Add("Access-Control-Expose-Headers", "")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,DELETE,OPTIONS")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Max-Age", "86400")
		w.Header().Add("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		w.Header().Add("Content-Security-Policy", "default-src 'self'")
		next.ServeHTTP(w, r)
	})
}

// compose combine multiple middleware functions into one
func compose(middlewares ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	// 采用逆序应用的方式, 每个中间件都会包装其后的处理器,
	// 如果按照正序应用的方式, 那么最终的结果将是第一个中间包裹了所有后续的中间件和处理器
	// 在每次遍历中, 当前的 h 都会被传递给当前索引位置对应的中间件函数, 并且中间件函数返回新的 http.Handler 将更新 h 的值
	return func(h http.Handler) http.Handler {
		for i := range middlewares {
			h = middlewares[len(middlewares)-1-i](h)
		}
		return h
	}
}

var middlewareChain = compose(corsMiddleware, loggingMiddleware, recoveryMiddleware)
