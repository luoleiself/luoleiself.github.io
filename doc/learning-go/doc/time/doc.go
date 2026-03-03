package main

import (
	"fmt"
	"time"
)

func readme() {
	fmt.Println("func (t Time) In(loc *Location) Time // 返回时间在指定时区上的值")
	fmt.Println("--------------------------------------")

	fmt.Println("time.Layout 格式使用英文时间习惯:")
	fmt.Println(tab, "1 -> 月")
	fmt.Println(tab, "2 -> 日")
	fmt.Println(tab, "3 -> 时")
	fmt.Println(tab, "4 -> 分")
	fmt.Println(tab, "5 -> 秒")
	fmt.Println(tab, "6 -> 年")
	fmt.Println(tab, "7 -> 时区")
	fmt.Println("--------------------------------------")

	fmt.Print("time.Now() 当前时间: ")
	fmt.Println(time.Now())                     // 2023-06-03 16:50:22.488995921 +0800
	fmt.Println("time.Layout", time.Layout)     // 01/02 03:04:05PM '06 -0700
	fmt.Println("time.RFC822", time.RFC822)     // 02 Jan 06 15:04 MST
	fmt.Println("time.RFC822Z", time.RFC822Z)   // 02 Jan 06 15:04 -0700
	fmt.Println("time.RFC1123", time.RFC1123)   // Mon, 02 Jan 2006 15:04:05 MST
	fmt.Println("time.RFC1123Z", time.RFC1123Z) // Mon, 02 Jan 2006 15:04:05 -0700
	fmt.Println("time.RFC3339", time.RFC3339)   // 2006-01-02T15:04:05Z07:00
	fmt.Print("time.RFC3339Nano ")
	fmt.Println(time.RFC3339Nano)                        // 2006-01-02T15:04:05.999999999Z07:00
	fmt.Println("小时 time.Hour", time.Hour)               // 1h0m0s
	fmt.Println("分钟 time.Minute", time.Minute)           // 1m0s
	fmt.Println("秒 time.Second", time.Second)            // 1s
	fmt.Println("毫秒 time.Millisecond", time.Millisecond) // 1ms
	fmt.Println("微秒 time.Microsecond", time.Microsecond) // 1µs
	fmt.Println("纳秒 time.Nanosecond", time.Nanosecond)   // 1ns
	fmt.Println("--------------------------------------")

	fmt.Println("func LoadLocation(name string) (*Location, error) // 返回具有指定名称的 Location")
	locUSALa, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println("time.LoadLocation(\"America/Los_Angeles\")", locUSALa) // America/Los_Angeles
	fmt.Println("--------------------")

	fmt.Println("func FixedZone(name string, offset int) *Location // 返回一个始终使用给定区域名称偏移量的 Location, 自定义时区名称")
	locUTC8 := time.FixedZone("UTC+8", +(8 * 60 * 60))
	fmt.Println("time.FixedZone(\"UTC+8\", +(8 * 60 * 60))", locUTC8) // UTC+8
	fmt.Println("--------------------")

	t, _ := time.Parse(time.RFC3339, "1990-03-05T16:07:08+08:00")
	fmt.Println("func Parse(layout, value string) (Time, error) // 根据指定的格式格式化日期时间字符串为日期时间")
	fmt.Print("time.Parse(time.RFC3339, \"1990-03-05T16:07:08+08:00\") ")
	fmt.Println(t)                                                        // 1990-03-05 16:07:08 +0800 CST
	fmt.Println("t.Format(time.RFC3339)", t.Format(time.RFC3339))         // 1990-03-05T16:07:08+08:00
	fmt.Println("t.Format(time.RFC3339Nano)", t.Format(time.RFC3339Nano)) // 1990-03-05T16:07:08+08:00
	fmt.Println("t.Format(time.RFC1123)", t.Format(time.RFC1123))         // Mon, 05 Mar 1990 16:07:08 CST
	fmt.Println("t.Format(time.RFC1123Z)", t.Format(time.RFC1123Z))       // Mon, 05 Mar 1990 16:07:08 +0800
	fmt.Println("t.Format(time.RFC822)", t.Format(time.RFC822))           // 05 Mar 90 16:07 CST
	fmt.Println("t.Format(time.RFC822Z)", t.Format(time.RFC822Z))         // 05 Mar 90 16:07 +0800
	fmt.Println("t.Format(time.RFC850)", t.Format(time.RFC850))           // Monday, 05-Mar-90 16:07:08 CST
	fmt.Print("t.Zone() ")
	fmt.Println(t.Zone())                     // CST 28800
	fmt.Println("t.Location()", t.Location()) // Local
	fmt.Println("t.Local()", t.Local())       // 1990-03-05 16:07:08 +0800 CST
	fmt.Println("t.UTC()", t.UTC())           // 1990-03-05 08:07:08 +0000 UTC
	fmt.Println("t.Year()", t.Year())         // 1990
	fmt.Println("t.YearDay()", t.YearDay())   // 64
	// 本地时间 1990-03-05 16:07:08 +0800 CST 先转换为 UTC 时间 1990-03-05T08:07:08Z
	// 然后再转换为归属时区的时间 1990-03-05 00:07:08 -0800 PST
	fmt.Println("t.In(locUSALa)", t.In(locUSALa)) // 1990-03-05 00:07:08 -0800 PST
	fmt.Println("t.In(locUTC8)", t.In(locUTC8))   // 1990-03-05 16:07:08 +0800 UTC+8
	fmt.Println("--------------------------------------")

	t2, _ := time.Parse("2006/01/02 15:04:05", "2022/11/08 10:54:55")
	fmt.Print("time.Parse(\"2006/01/02 15:04:05\", \"2022/11/08 10:54:55\") ")
	fmt.Println(t2)                       // 2022-11-08 10:54:55 +0000 UTC
	fmt.Println("t2.IsDST()", t2.IsDST()) // false
	fmt.Print("t2.Zone() ")
	fmt.Println(t2.Zone())                      // UTC 0
	fmt.Println("t2.Location()", t2.Location()) // UTC
	fmt.Println("t2.Local()", t2.Local())       // 2022-11-08 18:54:55 +0800 CST
	fmt.Println("t2.UTC()", t2.UTC())           // 2022-11-08 10:54:55 +0000 UTC
	fmt.Print("t2.Date() ")
	fmt.Println(t2.Date())                    // 2022 November 8
	fmt.Println("t2.Year()", t2.Year())       // 2022
	fmt.Println("t2.YearDay()", t2.YearDay()) // 312
	fmt.Println("t2.Month()", t2.Month())     // November
	fmt.Println("t2.Day()", t2.Day())         // 8
	fmt.Print("t2.Clock() ")
	fmt.Println(t2.Clock())                                     // 10 54 55
	fmt.Println("t2.Hour()", t2.Hour())                         // 10
	fmt.Println("t2.Minute()", t2.Minute())                     // 54
	fmt.Println("t2.Second()", t2.Second())                     // 55
	fmt.Println("t2.Add(time.Hour * 10)", t2.Add(time.Hour*10)) // 2022-11-08 20:54:55 +0000 UTC
	fmt.Println("t2.AddDate(1, 2, 3)", t2.AddDate(1, 2, 3))     // 2024-01-11 10:54:55 +0000 UTC
	fmt.Println("--------------------")
	fmt.Println("func (t Time) After(u time) bool 返回时间点 t 是否在时间点 u 的后面")
	fmt.Printf("t2.After(time.Now()) %t\n", t2.After(time.Now())) // false
	fmt.Println("func (t Time) Before(u time) bool 返回时间点 t 是否在时间点 u 的前面")
	fmt.Printf("t2.Before(time.Now()) %t\n", t2.Before(time.Now())) // true
	fmt.Println("--------------------")
	fmt.Print("t2.Format(\"2006_01_02 15_04_05\") ")
	fmt.Println(t2.Format("2006_01_02 15_04_05")) // 2022_11_08 10_54_55
	// UTC 时间 2022-11-08 10:54:55 +0000 UTC
	// 转换为归属时区的时间 2022-11-08 02:54:55 -0800 PST
	fmt.Println("t2.In(locUSALa)", t2.In(locUSALa)) // 2022-11-08 02:54:55 -0800 PST
	fmt.Println("t2.In(locUTC8)", t2.In(locUTC8))   // 2022-11-08 18:54:55 +0800 UTC+8
	fmt.Println("--------------------------------------")

	fmt.Println("func (t Time) Format(layout string) string // 格式化日期时间的值为文本表示形式")
	fmt.Print("time.Now().Format(\"2006/01/02 15:04:05\") ")
	fmt.Println(time.Now().Format("2006/01/02 15:04:05")) // 2023/06/03 17:17:37
	fmt.Println("--------------------------------------")

	fmt.Println("func Unix(sec int64, nsec int64) Time // 返回以 sec 为基准的 nsec(纳秒) 偏移后的时间")
	fmt.Println(tab, "如果 sec(秒) 为 0, 则以 1970/01/01 00:00:00 UTC 时间为基准, 计算 nsec(纳秒) 偏移后的时间")
	fmt.Println(tab, "如果 sec(秒) 不为 0, 则以 sec(秒) 时间为基准, 计算 nsec(纳秒) 偏移后的时间")
	fmt.Println("--------------------")

	fmt.Print("time.Unix(1663158154, 0) ")
	fmt.Println(time.Unix(1663158154, 0)) // 2022-09-14 20:22:34 +0800 CST
	fmt.Print("time.Unix(0, 1663158154e+9) ")
	fmt.Println(time.Unix(0, 1663158154e+9)) // 2022-09-14 20:22:34 +0800 CST
	fmt.Print("time.Unix(0, 1663158154e9) ")
	fmt.Println(time.Unix(0, 1663158154e9)) // 2022-09-14 20:22:34 +0800 CST
	fmt.Println("--------------------")
	fmt.Print("time.Unix(1663158154, 60e9) ")   // 60 second * 1000 Millisecond * 1000 Microsecond * 1000 Nanosecond
	fmt.Println(time.Unix(1663158154, 60e9))    // 2022-09-14 20:23:34 +0800 CST
	fmt.Print("time.Unix(1663158154, -60e+9) ") // 60 second * 1000 Millisecond * 1000 Microsecond * 1000 Nanosecond
	fmt.Println(time.Unix(1663158154, -60e+9))  // 2022-09-14 20:21:34 +0800 CST
	fmt.Println("--------------------")
	fmt.Print("time.Unix(0, -1663158154e+9) ")
	fmt.Println(time.Unix(0, -1663158154e+9)) // 1917-04-19 19:37:26 +0800 CST
	fmt.Println("--------------------------------------")

	t1 := time.UnixMilli(1663158154301)
	fmt.Println("time.UnixMilli(1663158154301)", t1) // 2022-09-14 20:22:34.301 +0800 CST
	fmt.Print("t1.Zone() ")
	fmt.Println(t1.Zone())                                          // CST 28800
	fmt.Println("t1.Location()", t1.Location())                     // Local
	fmt.Println("t1.Local()", t1.Local())                           // 2022-09-14 20:22:34.301 +0800 CST
	fmt.Println("t1.UTC()", t1.UTC())                               // 2022-09-14 12:22:34.301 +0000 UTC
	fmt.Println("t1.Year()", t1.Year())                             // 2022
	fmt.Println("t1.Hour()", t1.Hour())                             // 20
	fmt.Println("t1.Minute()", t1.Minute())                         // 22
	fmt.Println("t1.Second()", t1.Second())                         // 34
	fmt.Println("t1.Unix()", t1.Unix())                             // 1663158154
	fmt.Println("t1.UnixMilli()", t1.UnixMilli())                   // 1663158154301
	fmt.Println("t1.UnixMicro()", t1.UnixMicro())                   // 1663158154301000
	fmt.Println("t1.UnixNano()", t1.UnixNano())                     // 1663158154301000000
	fmt.Println("t1.Format(time.RFC3339)", t1.Format(time.RFC3339)) // 2022-09-14T20:22:34+08:00
	// 本地时间 2022-09-14 20:22:34.301 +0800 CST 先转换为 UTC 时间 2022-09-14 12:22:34.301 +0000 UTC
	// 然后再转为归属时区的时间 2022-09-14 05:22:34.301 -0700 PDT
	fmt.Println("t1.In(locUSALa)", t1.In(locUSALa)) // 2022-09-14 05:22:34.301 -0700 PDT
	fmt.Println("t1.In(locUTC8)", t1.In(locUTC8))   // 2022-09-14 20:22:34.301 +0800 UTC+8
	fmt.Println("--------------------------------------")

	fmt.Println("type Duration int64 // 以纳秒为单位表示两个时间点之间经过的时间, 最大可表示约 290 年")
	fmt.Println(tab, "func ParseDuration(s string) (Duration, error) ")
	fmt.Println(tab, tab, "解析一个时间段字符串, 可以包含正负号, 十进制数, 小数部分和单位, 单位支持 \"ns\", \"us\", \"ms\", \"s\", \"m\", \"h\"")
	fmt.Print("func (t Time) Sub(u Time) Duration 计算两个时间点之间的差异, 返回一个时间段 t-u")
	fmt.Println("go-staticheck, should use time.Until instead of time.Sub(time.Now())")
	fmt.Println(t.Sub(time.Now())) // -291425h30m47.153950498s
	fmt.Println(time.Until(t))
	td := time.Since(t)
	fmt.Print(tab, "time.Since(t) 返回从 t 到现在经过的时间, 等价于time.Now().Sub(t), 通常用于计算过去的时间")
	fmt.Println(td) // 291425h30m47.154004975s
	fmt.Print(tab, "time.Until(t) 返回从现在到 t 的经过的时间. 等价于t.Sub(time.Now()), 通常用于计算未来的时间")
	fmt.Println(time.Until(t))                          // -291425h30m47.154236568s
	fmt.Println("td.Hours()", td.Hours())               // 291425.5130983347
	fmt.Println("td.Minutes()", td.Minutes())           // 1.7485530785900082e+07
	fmt.Println("td.Milliseconds()", td.Milliseconds()) // 1049131847154
	fmt.Println("td.Microseconds()", td.Microseconds()) // 1049131847154004
	fmt.Println("td.Nanoseconds()", td.Nanoseconds())   // 1049131847154004975
	fmt.Println("td.String()", td.String())             // 291425h30m47.154004975s
	fmt.Println("td.Truncate(d Duration) 返回舍入到 d 的结果 ")
	fmt.Println(td.Truncate(time.Minute)) // 291425h30m0s
	fmt.Println("--------------------------------------")

	td2, _ := time.ParseDuration("20h150m300s45ms")
	fmt.Println("time.ParseDuration(\"20h150m300s45ms\") ", td2)    // 22h35m0.045s
	fmt.Println("td2.Hours()", td2.Hours())                         // 22.583345833333333
	fmt.Println("td2.Minutes()", td2.Minutes())                     // 1355.00075
	fmt.Println("td2.Seconds()", td2.Seconds())                     // 81300.045
	fmt.Println("td2.Milliseconds()", td2.Milliseconds())           // 81300045
	fmt.Println("td2.Microseconds()", td2.Microseconds())           // 81300045000
	fmt.Println("td2.Nanoseconds()", td2.Nanoseconds())             // 81300045000000
	fmt.Println("td2.Round(time.Minute)", td2.Round(time.Minute))   // 22h35m0s
	fmt.Println("td2.Truncate(time.Hour)", td2.Truncate(time.Hour)) // 22h0m0s
	fmt.Println("td2.String()", td2.String())                       // 22h35m0.045s
	fmt.Println("--------------------------------------")
}
