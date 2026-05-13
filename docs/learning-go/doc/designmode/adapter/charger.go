package adapter

import (
	"fmt"
	"strings"
)

type PhoneCharger interface {
	Output5V()
}

type HuaWeiCharger struct{}

func NewHuaWeiCharger() *HuaWeiCharger {
	return &HuaWeiCharger{}
}
func (h *HuaWeiCharger) Output5V() {
	fmt.Println("华为手机充电器输出电压 5V...")
}

type XiaoMiCharger struct{}

func NewXiaoMiCharger() *XiaoMiCharger {
	return &XiaoMiCharger{}
}
func (x *XiaoMiCharger) Output5V() {
	fmt.Println("小米手机充电器输出电压 5V...")
}

type MacBookCharger struct{}

func NewMacBookCharger() *MacBookCharger {
	return &MacBookCharger{}
}
func (m *MacBookCharger) Output28V() {
	fmt.Println("苹果笔记本充电器输出电压 28V...")
}

// 笔记本适配器类
type MacBookChargerAdapter struct {
	core *MacBookCharger
}

func NewMacBookChargerAdapter(m *MacBookCharger) *MacBookChargerAdapter {
	return &MacBookChargerAdapter{core: m}
}
func (m *MacBookChargerAdapter) Output5V() {
	m.core.Output28V()
	fmt.Println("适配器将输出电压调整为 5V...")
}

// 手机类型声明
type Phone interface {
	Charge(phoneCharger PhoneCharger)
}
type HuaWeiPhone struct{}

func NewHuaWeiPhone() Phone {
	return &HuaWeiPhone{}
}
func (h *HuaWeiPhone) Charge(phoneCharger PhoneCharger) {
	fmt.Println("华为手机准备开始充电...")
	phoneCharger.Output5V()
}

type XiaoMiPhone struct{}

func NewXiaoMiPhone() Phone {
	return &XiaoMiPhone{}
}
func (x *XiaoMiPhone) Charge(phoneCharger PhoneCharger) {
	fmt.Println("小米手机准备开始充电...")
	phoneCharger.Output5V()
}

var tab = strings.Repeat("", 2)

func Doc() {
	fmt.Println("适配器模式: 能够实现两个不兼容或弱兼容接口之间的适配桥接作用, 通常包含以下几个角色")
	fmt.Println("  目标 target: 是一类含有指定功能的接口")
	fmt.Println("  使用方 client: 需要使用 target 的用户")
	fmt.Println("  被适配的类 adaptee: 和目标类 target 功能类似, 但不完全吻合")
	fmt.Println("  适配器类 adapter: 能够将 adaptee 适配转换城 target 的功能类")
	fmt.Println("-----------------------------")
}
