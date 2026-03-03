package adapter

import (
	"fmt"
	"testing"
)

func TestDoc(t *testing.T) {
	Doc()
}

func TestAdapter(t *testing.T) {
	b1 := t.Run("HuaWeiPhone Charge", func(t *testing.T) {
		// 创建一个华为手机实例
		huaWeiPhone := NewHuaWeiPhone()

		// 使用华为手机充电器进行充电
		huaWeiCharger := NewHuaWeiCharger()
		huaWeiPhone.Charge(huaWeiCharger)

		fmt.Println("----------")

		// 使用适配器转换后的 macBook 充电器进行充电
		macBookCharger := NewMacBookCharger()
		macBookChargerAdapter := NewMacBookChargerAdapter(macBookCharger)
		huaWeiPhone.Charge(macBookChargerAdapter)
	})

	if !b1 {
		t.Errorf("error... %v\n", b1)
	} else {
		t.Logf("success... %v\n", b1)
	}

	b2 := t.Run("XiaoMiPhone Charge", func(t *testing.T) {
		// 创建一个小米手机实例
		xiaoMiPhone := NewXiaoMiPhone()

		// 使用小米手机充电器进行充电
		xiaoMiCharger := NewXiaoMiCharger()
		xiaoMiPhone.Charge(xiaoMiCharger)
		fmt.Println("----------")

		// 使用适配器转换的 macBook 充电器进行充电
		macBookCharger := NewMacBookCharger()
		macBookChargerAdapter := NewMacBookChargerAdapter(macBookCharger)
		xiaoMiPhone.Charge(macBookChargerAdapter)
	})
	if !b2 {
		t.Errorf("error... %v\n", b2)
	} else {
		t.Logf("success...%v\n", b2)
	}
}
