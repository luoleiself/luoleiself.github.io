package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/spf13/viper"
)

var viperConfFile = "viper_test.yaml"

func main() {
	fmt.Println("使用 go test -v . [-run pattern] 运行指定匹配任务")
}
func WriteViperConf() {
	viper.Set("title", "luoleiself's blog")
	viper.Set("subtitle", "blog")
	viper.Set("description", "努力的往前飞, 再累也无所谓, 黑夜过后的光芒有多美...")

	viper.Set("bool", true)
	viper.Set("int", 127)
	viper.Set("int32", 2147483647)
	viper.Set("uint", 255)
	viper.Set("Uint16", 65535)
	viper.Set("float64", 123.456)

	var sl = []int{1, 2, 3, 4, 5}
	viper.Set("intSlice", sl)

	var sl2 = []string{"hello", "world"}
	viper.Set("stringSlice", sl2)

	var m = map[string]any{"path": "search.xml", "field": "post", "format": "html", "limit": 10000}
	viper.Set("search", m)

	var t = time.Now()
	viper.Set("time", t)

	viper.Set("duration", t.UnixNano())

	viper.SetDefault("v-default", "github.com/spf13/viper")

	// // pflag 第三方包
	// pflag.String("host", "localhost", "server listening host")
	// pflag.Int("port", 8080, "server listening port")
	// pflag.Parse()
	// viper.BindPFlag("host", pflag.Lookup("host"))
	// viper.BindPFlag("port", pflag.Lookup("port")) // 绑定指定的 pflag 参数
	// // ----
	// viper.BindPFlags(pflag.CommandLine) // 绑定 pflag 集合

	// viper.Set("host", viper.GetString("host"))
	// viper.Set("port", viper.GetInt("port"))

	// 获取配置文件所有的 key 组成的 []string, 嵌套的 key 将以点分割符的方式返回
	// [intslice stringslice search.field search.limit v-default search.format int uint subtitle time duration bool float64 int32 title search.path uint16 description]
	fmt.Println("AllKeys", viper.AllKeys())
	// 获取所有的配置信息组成的 map[string]any
	// map[boo: true int:127 intSlice:[1 2 3 4 5] search:map[field:post format:html limit:10000 path:search.xml]]
	allSettings := viper.AllSettings()
	fmt.Println("AllSettings", allSettings)

	// 监听配置文件变化
	// viper.OnConfigChange(func(in fsnotify.Event) {
	// 	fmt.Println("viper.OnConfigChange", in)
	// })

	// err := viper.SafeWriteConfig() // 仅当配置文件不存在时写入
	// err := viper.SafeWriteConfigAs("viper_test.yaml") // 仅当配置文件不存在时写入
	// err := viper.WriteConfig()
	err := viper.WriteConfigAs(viperConfFile)
	if err != nil {
		panic(fmt.Sprintln("write viper config file error", err))
	}
	fmt.Println("write success")
	fmt.Println("------------")
}

func ReadViperConf(file *os.File) {
	if file != nil {
		viper.ReadConfig(file)
	} else {
		// 关联配置文件
		// 设置配置文件名称(不包括文件扩展名)
		// viper.SetConfigName("viper_test")
		// 设置配置文件类型
		// viper.SetConfigType("yaml")
		// 添加配置文件搜索路径
		// viper.AddConfigPath(".")
		// 设置配置文件
		viper.SetConfigFile(path.Join("./", viperConfFile)) // 关联配置文件

		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Sprintln("read viper config error", err))
		}
	}
	fmt.Println("read viper config success")
	fmt.Println("viper.ConfigFileUsed()", viper.ConfigFileUsed())

	fmt.Println(viper.AllKeys())

	allSettings := viper.AllSettings()
	fmt.Println(allSettings)

	viper.SetEnvPrefix("MY") // 设置绑定环境变量前缀
	viper.AutomaticEnv()     //  自动绑定环境变量

	viper.BindEnv("HOST", "PORT") // 绑定环境变量

	fmt.Printf("viper.IsSet(\"intSlice\") %t\n", viper.IsSet("intSlice"))
	fmt.Printf("viper.InConfig(\"author\") %t\n", viper.InConfig("author"))
	fmt.Printf("viper.GetEnvPrefix() %s\n", viper.GetEnvPrefix())
	fmt.Printf("viper.GetSizeInBytes(\"stringSlice\") %d\n", viper.GetSizeInBytes("stringSlice"))
	fmt.Println("------------")
	fmt.Printf("viper.GetString(\"title\") %s\n", viper.GetString("title"))
	fmt.Printf("viper.GetString(\"subtitle\") %s\n", viper.GetString("subtitle"))
	fmt.Printf("viper.GetString(\"title\") %s\n", viper.GetString("description"))
	fmt.Println("------------")
	fmt.Printf("viper.GetBool(\"bool\") %t\n", viper.GetBool("bool"))
	fmt.Printf("viper.GetInt(\"int\") %d\n", viper.GetInt("int"))
	fmt.Printf("viper.GetInt32(\"int32\") %d\n", viper.GetInt32("int32"))
	fmt.Printf("viper.GetUint(\"uint\") %d\n", viper.GetUint("uint"))
	fmt.Printf("viper.GetUint16(\"uint16\") %d\n", viper.GetUint16("uint16"))
	fmt.Printf("viper.GetFloat64(\"float64\") %f\n", viper.GetFloat64("float64"))
	fmt.Printf("viper.GetIntSlice(\"intSlice\") %v\n", viper.GetIntSlice("intSlice"))
	fmt.Printf("viper.GetStringSlice(\"stringSlice\") %v\n", viper.GetStringSlice("stringSlice"))
	fmt.Printf("viper.GetStringMap(\"search\") %v\n", viper.GetStringMap("search"))
	fmt.Println("------------")
	fmt.Printf("viper.GetTime(\"time\") %v\n", viper.GetTime("time"))
	fmt.Printf("viper.GetDuration(\"duration\") %v\n", viper.GetDuration("duration"))
	fmt.Println("------------")
	fmt.Printf("viper.GetString(\"host\") %s \n", viper.GetString("host"))
	fmt.Printf("viper.GetString(\"port\") %d \n", viper.GetInt("port"))
	fmt.Println("------------")
	fmt.Printf("viper.Get(\"v-default\") %s \n", viper.Get("v-default"))
}

func ReadMyViperByBuffer() {
	var myViper = viper.New()
	myViper.SetConfigType("yaml")
	var yamlExmaple = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 18
eyes: brown
beard: true
`)
	myViper.ReadConfig(bytes.NewBuffer(yamlExmaple))
	fmt.Println("myViper.ReadConfig")
	fmt.Printf("myViper.Get(\"name\") %s \n", myViper.Get("name"))
	fmt.Println("myViper.RegisterAlias(\"Name\", \"name\")")
	myViper.RegisterAlias("Name", "name")
	fmt.Printf("myViper.Get(\"Name\") %s \n", myViper.Get("Name"))

	fmt.Printf("myViper.GetBool(\"Hacker\") %t \n", myViper.GetBool("Hacker"))
	fmt.Printf("myViper.GetInt(\"age\") %d\n", myViper.GetInt("age"))
	fmt.Printf("myViper.GetString(\"clothing.trousers\") %s\n", myViper.GetString("clothing.trousers"))
	fmt.Printf("myviper.GetStringSlice(\"hobbies\") %q\n", myViper.GetStringSlice("hobbies"))
	fmt.Printf("myViper.GetStringMapString(\"clothing\") %v\n", myViper.GetStringMapString("clothing"))
}
