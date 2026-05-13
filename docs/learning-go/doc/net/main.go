package main

import (
	"fmt"
	"net"
	"strings"
)

func init() {
	fmt.Println("main.go init()...")
}

var tab = strings.Repeat("", 2)

func main() {
	fmt.Println("net")
	fmt.Println(tab, "func Dial(network, address string) (Conn, error) // 连接 address 并返回一个 Conn 接口")
	fmt.Println(tab, "func Interfaces() ([]Interface, error) // 返回该系统的网络接口列表")
	fmt.Println(tab, "func InterfaceAddrs() ([]Addr, error) // 返回该系统的网络接口的地址列表")
	fmt.Println(tab, "func InterfaceByIndex(index int) (*Interface, error) // 返回指定索引的网络接口")
	fmt.Println(tab, "func InterfaceByName(name string) (*Interface, error) // 返回指定名字的网络接口")
	fmt.Println(tab, "func SplitHostPort(hostport string) (host, port string, err error) // 分割网络地址")
	fmt.Println(tab, "func JoinHostPort(host, port string) string // 合并网络地址")
	fmt.Println(tab, "func IPv4(a, b, c, d byte) IP // 返回包含一个IPv4地址 a.b.c.d 的IP地址(16字节格式)")
	fmt.Println(tab, "func ParseIP(s string) IP // 解析 ip 地址")
	fmt.Println(tab, "func ParseCIDR(s string) (IP, *IPNet, error) // 解析 CIDR (无类型域间路由) 的IP地址和掩码字符串")
	fmt.Println(tab, "func CIDRMask(ones, bits int) IPMask // 返回一个 IPMask 类型值, ip 地址的掩码")
	fmt.Println(tab, "func IPv4Mask(a, b, c, d byte) IPMask // 返回一个 4 字节格式的 IPv4 掩码 a.b.c.d")
	fmt.Println(tab, "func ParseMAC(s string) (hw HardwareAddr, err error) // 解析一个硬件地址")
	fmt.Println(tab, "func ResolveIPAddr(network, address string) (*IPAddr, error) // 将 addr 作为一个格式为\"host\"或\"ipv6-host%zone\"的 IP 地址来解析, network 必须是\"ip\"、\"ip4\"或\"ip6\"")
	fmt.Println(tab, "func ResolveTCPAddr(network, address string) (*TCPAddr, error) // 将 addr 作为 TCP 地址解析并返回, network 必须是\"tcp\"、\"tcp4\"或\"tcp6\"")
	fmt.Println(tab, "func ResolveUDPAddr(network, address string) (*UDPAddr, error) // 将 addr 作为 TCP 地址解析并返回, network 必须是\"udp\"、\"udp4\"或\"udp6\"")
	fmt.Println(tab, "func ResolveUnixAddr(network, address string) (*UnixAddr, error) // 将addr作为Unix域socket地址解析, network 必须是\"unix\"、\"unixgram\"或\"unixpacket\"")
	fmt.Println(tab, "")
	fmt.Println("----------------------------------------")

	fmt.Println("net.IPv4len = ", net.IPv4len)
	fmt.Println("net.IPv6len = ", net.IPv6len)
	fmt.Println("---------------------")

	fmt.Println("1 << 0 =>", 1<<0)
	fmt.Println("1 << 1 =>", 1<<1)
	fmt.Println("1 << 2 =>", 1<<2)
	fmt.Println("1 << 3 =>", 1<<3)
	fmt.Println("1 << 4 =>", 1<<4)
	fmt.Println("---------------------")

	fmt.Println("net.FlagUp", net.FlagUp)
	fmt.Println("net.FlagBroadcast", net.FlagBroadcast)
	fmt.Println("net.FlagLoopback", net.FlagLoopback)
	fmt.Println("net.FlagPointToPoint", net.FlagPointToPoint)
	fmt.Println("net.FlagMulticast", net.FlagMulticast)
	fmt.Println("----------------------------------------")

	fmt.Print("net.LookupIP(\"localhost\") ")
	fmt.Println(net.LookupIP("localhost"))
	fmt.Print("net.LookupCNAME(\"www.baidu.com\") ")
	fmt.Println(net.LookupCNAME("www.baidu.com"))
	fmt.Print("net.LookupAddr(\"fe80::e5fa:1b7c:6f1e:fc93\") ")
	fmt.Println(net.LookupAddr("fe80::e5fa:1b7c:6f1e:fc93"))
	fmt.Print("net.LookupHost(\"KLVC-WXX9\") ")
	fmt.Println(net.LookupHost("KLVC-WXX9"))
	fmt.Print("net.LookupPort(\"localhost\", \"8080\") ")
	fmt.Println(net.LookupPort("localhost", "8080"))
	fmt.Println("---------------------")

	ipAddr, _ := net.ResolveIPAddr("ip", "192.168.1.20")
	fmt.Println("ipAddr, _ := net.ResolveIPAddr(\"ip\", \"192.168.1.20\")", ipAddr)
	fmt.Println("ipAddr.Network()", ipAddr.Network())
	fmt.Println("ipAddr.String()", ipAddr.String())
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	fmt.Println("tcpAddr, _ := net.ResolveTCPAddr(\"tcp\", \"localhost:8080\")", tcpAddr)
	fmt.Println("tcpAddr.Network()", ipAddr.Network())
	fmt.Println("tcpAddr.String()", ipAddr.String())
	fmt.Println("----------------------------------------")

	host, port, _ := net.SplitHostPort("127.0.0.1:8080")
	fmt.Println("host, port, _ := net.SplitHostPort(\"127.0.0.1:8080\")", host, port)
	fmt.Println("net.JoinHostPort(\"127.0.0.1", "8080\")", net.JoinHostPort("127.0.0.1", "8080"))
	fmt.Println("----------------------------------------")

	fmt.Println("------------maskFunc()------------")
	fmt.Println("---------------------")
	maskFunc()

	fmt.Println("------------parseFunc()------------")
	fmt.Println("---------------------")
	parseFunc()

	fmt.Println("------------interfaceFunc()------------")
	fmt.Println("---------------------")
	interfaceFunc()
}
