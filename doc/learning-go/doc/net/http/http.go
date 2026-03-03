package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

/*
launch
	配置文件启动任务指向的当前文件(必须包含 main 函数), 测试文件中运行 debug test 进入 debugger 模式, 在 DEBUG CONSOLE 面板查看调试信息

attach
	使用命令 'go build' 编译当前包(必须包含 main 函数), 在当前命令行中执行编译后的文件, 配置文件启动任务选择当前执行文件的 processID,
	测试文件中运行 debug test 进入 debugger 模式, 在当前命令行中查看调试信息
*/

type ResStruct struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type CountHandler struct {
	mu sync.Mutex
	n  int
}

func (h *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	llog.Println("----start----")
	llog.Printf("ServeMux. Implement the ServeHTTP method of the http.Handler interface. count is %d, request method is %v, request url is %v", h.n, r.Method, r.URL)

	// GET 请求参数
	if r.Method == http.MethodGet {
		params, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			llog.Fatal(err.Error())
		}
		llog.Println(params)
	}

	w.Header().Add("sm1", "1")
	var s = strings.Builder{}
	s.WriteString(strconv.Itoa(http.StatusOK))
	s.WriteString(" cout is ")
	s.WriteString(strconv.Itoa(h.n))
	w.Write([]byte(s.String()))
	llog.Println("----end----")
}

var (
	llog   = log.Default()
	server *http.Server
	sm     *http.ServeMux
)

func main() {
	llog.SetPrefix("http-")
	sm = http.NewServeMux()
	sm.Handle("/sm1", new(CountHandler))

	sm.HandleFunc("/sm5", func(w http.ResponseWriter, r *http.Request) {
		llog.Println("ServeMux. Use http.HandleFunc regist the handler function. request method is", r.Method, "request url is", r.URL)
		w.Header().Add("sm5", "5")
		w.Write([]byte(strconv.Itoa(http.StatusOK)))
	})

	// http.HandlerFunc 适配器函数, 将一个普通函数作为 HTTP handler
	// http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	sm.Handle("/sm2", middlewareChain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// POST 请求参数
		if r.Method == http.MethodPost {
			llog.Printf("req.FormValue(name) : %v\n", r.FormValue("name"))
			llog.Printf("req.PostFormValue(name) : %v\n", r.PostFormValue("name"))
			llog.Printf("req.FormValue(age) : %v\n", r.FormValue("age"))
			llog.Printf("req.PostFormValue(age) : %v\n", r.PostFormValue("age"))
			llog.Printf("req.FormValue(addr) : %v\n", r.FormValue("addr"))
			llog.Printf("req.PostFormValue(addr) : %v\n", r.PostFormValue("addr"))
		}
		w.Header().Add("sm2", "2")
		w.Write([]byte(strconv.Itoa(http.StatusOK)))
	})))

	sm.Handle("/sm3", middlewareChain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 将扩展名和 mimetype 建立偶联; 扩展名应以点号开始, 例如 .html
		// mime.AddExtensionType(ext, typ string) error
		// 将媒体类型和参数 param 连接为一个 mime 媒体类型, 类型和参数都采用小写字母
		// mime.FormatMediaType(t string, param map[string]string) string
		// 函数返回与扩展名偶联的 mime 类型, 扩展名应以 点号开始,
		// mime.TypeByExtension(ext string) string
		// 解析一个媒体类型以及可能的参数, 媒体类型值一般应为 Content-Type 和 Content-Disposition 头域的值, 成功的调用会返回一个小写字母、去空格的媒体类型和一个非空的 map
		// mime.ParseMediaType(v string) (mediaType string, params map[string]string, err error)
		mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			llog.Fatal(err)
		}
		// POST 表单请求
		if r.Method == http.MethodPost && strings.HasPrefix(mediaType, "multipart/") {
			// llog.Println("r.MultipartForm.File", r.MultipartForm.File["files"])

			mr := multipart.NewReader(r.Body, params["boundary"])

			for {
				p, err := mr.NextPart()
				if err == io.EOF {
					break
				}
				if err != nil {
					llog.Fatal(err)
				}
				slurp, err := io.ReadAll(p)
				if err != nil {
					llog.Fatal(err)
				}
				llog.Printf("Part %q: %s\n", p.FormName(), slurp)
			}
		}
		w.Header().Add("sm3", "3")
		w.Write([]byte(strconv.Itoa(http.StatusOK)))
	})))

	sm.Handle("/sm4", middlewareChain(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mediaType, _, err := mime.ParseMediaType(r.Header.Get("Accept"))
		if err != nil {
			llog.Fatal(err)
		}
		if strings.ToLower(mediaType) == "application/json" {
			w.Header().Add("Content-Type", "application/json")
			res := ResStruct{Code: 10000, Msg: "success", Data: []map[string]string{{"name": "张三", "addr": "beijing"}, {"name": "lisi", "addr": "shanghai"}}}
			resData, err := json.Marshal(res)
			if err != nil {
				llog.Fatalln("Marshal error", err)
			}
			llog.Printf("response application/json data is %s\n", resData)
			w.Write(resData)
		} else {
			w.Write([]byte(strconv.Itoa(http.StatusOK)))
		}
	})))

	server = &http.Server{
		Addr:         "localhost:3000",
		Handler:      sm,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// 反向代理
	// proxy := httputil.ReverseProxy{
	// 	Director: func(r *http.Request) {
	// 		r.URL.Scheme = "http"
	// 		r.URL.Host = backendHost
	// 		r.Header.Set("X-Proxy", "GoReverseProxy")
	// 	},
	// }
	// proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
	// 	log.Printf("Proxy error: %s\n", err.Error())
	// 	w.WriteHeader(http.StatusBadGateway)
	// }

	// graceful shutdown
	serverError := make(chan error, 1)
	go func() {
		log.Printf("服务运行中: %s", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			serverError <- err
		}
	}()

	stop := make(chan os.Signal, 1)

	fmt.Println("os.Interrupt == syscall.SIGINT", os.Interrupt == syscall.SIGINT)
	fmt.Println("os.Kill == syscall.SIGKILL", os.Kill == syscall.SIGKILL)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverError:
		log.Printf("服务错误: %s\n", err.Error())
	case sig := <-stop:
		log.Printf("收到关机信号: %s\n", sig)
	}

	log.Println("服务正在关闭...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Shutdown gracefully shuts down the server without interrupting any active connections.
	// 首先关闭所有打开的侦听器, 然后关闭所有空闲连接, 然后无限期等待连接返回空闲状态, 然后关机
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("服务关闭错误: %s", err.Error())
		return
	}
	log.Println("服务已正常退出...")
}
