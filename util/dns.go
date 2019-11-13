package util

import "github.com/yangyouwei/newrouter/conf"

var (
	DnsModcommand   string = "/sbin/uci set dhcp.@dnsmasq[0].port="
	UciCommit       string = "/sbin/uci commit dhcp.@dnsmasq[0].port"
	DnsmasqCommand  string = "/etc/init.d/dnsmasq restart"
	RedirectCommand string = "/etc/init.d/redirect "
)

func SwitchRedirect(s bool) {
	if s {
		StartRedict()
		return
	} else {
		StopRedirect()
		return
	}
}

func StartRedict() {
	Shellout(DnsModcommand+"0", conf.Workdir)
	Shellout(UciCommit, conf.Workdir)
	Shellout(DnsmasqCommand, conf.Workdir)
	Shellout(RedirectCommand+"restart", conf.Workdir)
}

func StopRedirect() {
	Shellout(RedirectCommand+"stop", conf.Workdir)
	Shellout(DnsModcommand+"53", conf.Workdir)
	Shellout(UciCommit, conf.Workdir)
	Shellout(DnsmasqCommand, conf.Workdir)
}
