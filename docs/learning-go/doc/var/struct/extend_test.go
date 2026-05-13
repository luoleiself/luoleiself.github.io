package structdoc

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"
)

type Speaker interface {
	Speak() string
}

type Dog struct{}
type Mouse struct{}

func (d Dog) Speak() string {
	return "汪汪汪"
}
func (c Mouse) Speak() string {
	return "吱吱吱"
}

func TestExtendImplInter(t *testing.T) {
	t.Run("接口类型变量可以赋值给任意实现该接口的对象", func(t *testing.T) {
		var speaker Speaker

		speaker = Dog{}
		t.Log("speaker.Speak()", speaker.Speak())

		speaker = Mouse{}
		t.Log("speaker.Speak()", speaker.Speak())
		t.Log("-------------")
	})
}

type Eater interface {
	Eat() string
}

type Animal struct {
	Name string
}

func (a Animal) Eat() string {
	var builder = strings.Builder{}
	builder.WriteString(a.Name)
	builder.Write([]byte(" is eating"))
	builder.WriteByte(byte('.'))
	builder.WriteRune('\n')
	return builder.String()
}

type Duck struct {
	Animal
	Breed string
}

func (d Duck) Ho() {
	fmt.Printf("%s the %s is Ho.\n", d.Name, d.Breed)
}

type Pig struct {
	a    Animal
	Bark string
}

func (p Pig) Ho() {
	fmt.Printf("%s the %s is Ho.\n", p.a.Name, p.Bark)
}

func TestExtendEmbedAnonymousStruct(t *testing.T) {
	t.Run("继承-结构体嵌套匿名结构体字段, 可以赋值给匿名结构体字段实现的接口类型变量", func(t *testing.T) {
		var eater Eater = Duck{Animal{"Buddy"}, "Golden Retriever"}

		t.Log("调用聚合的方法", eater.Eat())

		t.Log("接口类型变量赋值给继承对象使用类型断言调用自己的方法")
		if duck, ok := eater.(Duck); ok {
			duck.Ho()
		}
		t.Log("-------------")
	})
	t.Run("继承-结构体嵌套具名结构体字段, 不能赋值给具名结构体字段实现的接口类型变量, 编译报错", func(t *testing.T) {
		// cannot use Pig{…} (value of type Pig) as Eater value in variable declaration: Pig does not implement Eater (missing method Eat)compiler
		// var ani Eater = Pig{a: Animal{Name: "Pig buddy"}, Bark: "pig buddy"}
		t.Log("cannot use Pig{…} (value of type Pig) as Eater value in variable declaration: Pig does not implement Eater (missing method Eat)compiler")
	})
}

// 嵌入匿名接口字段实现策略模式
type Notifier interface {
	Notify(message string) error
}
type EmailNotifier struct{}

func (en EmailNotifier) Notify(message string) error {
	fmt.Printf("Sending email: %s\n", message)
	return nil
}

type SMSNotifier struct{}

func (sn SMSNotifier) Notify(message string) error {
	fmt.Printf("Seding SMS: %s\n", message)
	return nil
}

type WeChatNotifier struct{}

func (wn WeChatNotifier) Notify(message string) error {
	fmt.Printf("Sending WeChat: %s\n", message)
	return nil
}

// 包含 Notifier 接口的结构体
type NotificationService struct {
	Notifier // 匿名接口字段
	sendAt   time.Time
}

// 使用 NotificationService 发送消息
func (ns *NotificationService) SendNotification(message string) error {
	if ns.Notifier == nil {
		return errors.New("not notifier set")
	}
	ns.sendAt = time.Now()
	fmt.Printf("At: %s ", ns.sendAt.Format(time.RFC3339Nano))
	return ns.Notifier.Notify(message)
}

func TestExtendEmbedInter(t *testing.T) {
	t.Run("结构体内嵌匿名接口字段, 可以接收任意实现了该匿名接口的对象", func(t *testing.T) {
		emailNotifier := EmailNotifier{}
		smsNotifier := SMSNotifier{}
		wechatNotifier := WeChatNotifier{}

		// 创建 NotificationService 实例并设置不同的 Notifier
		service := NotificationService{Notifier: emailNotifier}
		service.SendNotification("Hello via email")

		service.Notifier = smsNotifier
		service.SendNotification("Hello via SMS")

		service.Notifier = wechatNotifier
		service.SendNotification("Hello via wechat")
	})
}
