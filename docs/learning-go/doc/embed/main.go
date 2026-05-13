package main

import (
	"embed"
	"fmt"
	"net/http"
)

// 嵌入目录
//
//go:embed messages/*
var messages embed.FS

//go:embed static/*
var staticDir embed.FS

// 嵌入文件
//
//go:embed templates/home.html
var homePageTemplate string

//go:embed templates/about.html
var aboutPageTemplate string

type Templates struct {
	Home  string
	About string
}

func main() {
	fmt.Println("//go:embed 指令告诉编译器在构建时将文件和文件夹包含到编译后的二进制文件中, 这意味着应用程序可以直接从内存访问这些资源, 而无需再运行时从磁盘上读取。")
	fmt.Println("--------------------")
	fmt.Println("go:embed file")
	t := Templates{
		Home:  homePageTemplate,
		About: aboutPageTemplate,
	}
	fmt.Println("Home template: ", t.Home)
	fmt.Println("About template: ", t.About)

	fmt.Println("--------------------")
	fmt.Println("go:embed dir")
	files, _ := messages.ReadDir("messages")
	for _, file := range files {
		data, _ := messages.ReadFile("messages/" + file.Name())
		fmt.Printf("File: %s\nContent: %s\n\n", file.Name(), string(data))
	}
	fmt.Println("--------------------")

	fmt.Println("go:embed Web")
	http.Handle("/", http.FileServer(http.FS(staticDir)))
	http.ListenAndServe(":8080", nil)
}
