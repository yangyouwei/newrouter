package util

import "fmt"

var (
	DnsModcommand string = "/sbin/uci set dhcp.@dnsmasq[0].port="
	UciCommit string = "/sbin/uci commit dhcp.@dnsmasq[0].port"
	DnsmasqCommand string = "/etc/init.d/dnsmasq restart"
	RedirectCommand string = "/etc/init.d/redirect "
)

func SwitchRedirect(s bool)  {
	if s {
		StartRedict()
		return
	}else {
		StopRedirect()
		return
	}
}

func StartRedict()  {
	a,b,c := Shellout(DnsModcommand+"0","/")
	fmt.Println(DnsModcommand+"0")
	fmt.Println(a,b,c)
	a,b,c = Shellout(UciCommit,"/")
	fmt.Println(UciCommit)
	fmt.Println(a,b,c)
	a,b,c = Shellout(DnsmasqCommand,"/")
	fmt.Println(DnsmasqCommand)
	fmt.Println(a,b,c)
	a,b,c = Shellout(RedirectCommand+"restart","/")
	fmt.Println(RedirectCommand+"restart")
	fmt.Println(a,b,c)
}

func StopRedirect()  {
	a,b,c := Shellout(RedirectCommand+"stop","/")
	fmt.Println(RedirectCommand+"stop")
	fmt.Println(a,b,c)
	a,b,c = Shellout(DnsModcommand+"53","/")
	fmt.Println(DnsModcommand+"53")
	fmt.Println(a,b,c)
	a,b,c = Shellout(UciCommit,"/")
	fmt.Println(UciCommit)
	fmt.Println(a,b,c)
	a,b,c = Shellout(DnsmasqCommand,"/")
	fmt.Println(DnsmasqCommand)
	fmt.Println(a,b,c)
}