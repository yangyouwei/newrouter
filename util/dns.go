package util

var (
	DnsModcommand string = "uci set dhcp.@dnsmasq[0].port="
	UciCommit string = "uci commit dhcp.@dnsmasq[0].port="
	DnsmasqCommand string = "/etc/init.d//etc/init.d/dnsmasq restart"
	RedirectCommand string = "/etc/init.d/redirect "
)

func SwitchRedirect(s bool)  {
	if s {
		StartRedict()
	}else {
		StopRedirect()
	}
}

func StartRedict()  {
	Shellout(DnsModcommand+"0","/")
	Shellout(UciCommit,"/")
	Shellout(DnsmasqCommand,"/")
	Shellout(RedirectCommand+"restart","/")
}

func StopRedirect()  {
	Shellout(RedirectCommand+"stop","/")
	Shellout(DnsModcommand+"53","/")
	Shellout(UciCommit,"/")
	Shellout(DnsmasqCommand,"/")
}