package util

import (
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/models"
	"strings"
)

type LineConf struct {
	Ipaddr string
	TCPPort string
	UDPPort string
}
var Port LineConf

func init()  {
	Port.getpoart()
}

var IptablesFull string = `/usr/sbin/ip rule add fwmark 0x01/0x01 table 100
/usr/sbin/ip route add local 0.0.0.0/0 dev lo table 100
/usr/sbin/iptables -t mangle -N SS-UDP
/usr/sbin/iptables -t mangle -A SS-UDP -d $ssserver -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 192.168/16 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 255.255.255.255/32 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 224/4 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 240/4 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 0/8 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 127/8 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 10/8 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 169.254/16 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 172.16/12 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -p udp -j TPROXY --tproxy-mark 0x01/0x01 --on-port ${UDPPort}
/usr/sbin/iptables -t mangle -A PREROUTING -p udp -j SS-UDP
/usr/sbin/iptables -t nat -N SS-TCP
/usr/sbin/iptables -t nat -A SS-TCP -d $ssserver -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 192.168/16 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 224/4 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 240/4 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 255.255.255.255/32 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 0/8 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 127/8 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 10/8 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 169.254/16 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 172.16/12 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -p tcp -j REDIRECT --to-ports ${TCPPort}
/usr/sbin/iptables -t nat -A OUTPUT -p tcp -j SS-TCP
/usr/sbin/iptables -t nat -A PREROUTING -p tcp -j SS-TCP`

var IpatablesMulti string = `ip rule add fwmark 0x01/0x01 table 100
/usr/sbin/ip route add local 0.0.0.0/0 dev lo table 100
/usr/sbin/ptables -t mangle -N SS-UDP
/usr/sbin/ptables -t mangle -A SS-UDP -d $ssserver -j RETURN
/usr/sbin/ptables -t mangle -A SS-UDP -p udp -j TPROXY --tproxy-mark 0x01/0x01 --on-port ${UDPPort}
/usr/sbin/ptables -t mangle -A PREROUTING -p udp -m set --match-set foreign dst -j SS-UDP
/usr/sbin/ptables -t nat -N SS-TCP
/usr/sbin/ptables -t nat -A SS-TCP -d $ssserver -j RETURN
/usr/sbin/ptables -t nat -A SS-TCP -p tcp -j REDIRECT --to-ports ${TCPPort}
/usr/sbin/ptables -t nat -A PREROUTING -p tcp -m set --match-set foreign dst -j SS-TCP`

var IptablesDomestic string = `ip rule add fwmark 0x01/0x01 table 100
/usr/sbin/ip route add local 0.0.0.0/0 dev lo table 100
/usr/sbin/iptables -t mangle -N SS-UDP
/usr/sbin/iptables -t mangle -A SS-UDP -d $ssserver -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -m set --match-set domestic dst -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 192.168/16 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 255.255.255.255/32 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 224/4 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 240/4 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 0/8 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 127/8 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 10/8 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 169.254/16 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -d 172.16/12 -j RETURN
/usr/sbin/iptables -t mangle -A SS-UDP -p udp -j TPROXY --tproxy-mark 0x01/0x01 --on-port ${UDPPort}
/usr/sbin/iptables -t mangle -A PREROUTING -p udp -j SS-UDP
/usr/sbin/iptables -t nat -N SS-TCP
/usr/sbin/iptables -t nat -A SS-TCP -d $ssserver -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -m set --match-set domestic dst -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 192.168/16 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 224/4 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 240/4 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 255.255.255.255/32 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 0/8 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 127/8 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 10/8 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 169.254/16 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -d 172.16/12 -j RETURN
/usr/sbin/iptables -t nat -A SS-TCP -p tcp -j REDIRECT --to-ports ${TCPPort}
/usr/sbin/iptables -t nat -A OUTPUT -p tcp -j SS-TCP
/usr/sbin/iptables -t nat -A PREROUTING -p tcp -j SS-TCP`

func (p *LineConf)getpoart()  {
	i := models.Line{}
	i.GetUseLine()
	ip := strings.Split(i.Ipaddr,":")
	p.Ipaddr = ip[1]
	p.TCPPort = "8001"
	p.UDPPort = "8001"
}

func ChSpeedMod(m string,conf LineConf)  {
	switch  {
	case m == "fullspeed":
		FullSpeed(conf)
	case m == "multispeed":
		MultiSpeed(conf)
	case m == "domsticspeed":
		DomesticSpeed(conf)
	default:
		Stopspeed()
	}
}

func FullSpeed(l LineConf)  {
	//关闭dnsmasq
	Shellout("/etc/init.d/dnsmasq stop",conf.Workdir)
	//启动 restart redirect
	Shellout("/etc/init.d/redirect restart",conf.Workdir)
	//清理防火墙
	Stopspeed()
	//加载防火墙规则
	iptables := strings.Split(IptablesFull,"\n")
	for _,i := range iptables {
		if strings.Contains(i,"$ssserver") {
			c := strings.Replace(i,"$ssserver",Port.Ipaddr,-1)
			Shellout(c,conf.Workdir)
		}else if strings.Contains(i,"${UDPPort}") {
			c := strings.Replace(i,"${UDPPort}",Port.UDPPort,-1)
			Shellout(c,conf.Workdir)
		} else if strings.Contains(i,"${TCPPort}") {
			c := strings.Replace(i,"${TCPPort}",Port.TCPPort,-1)
			Shellout(c,conf.Workdir)
		}else {
			Shellout(i,conf.Workdir)
		}
	}
}

func MultiSpeed(l LineConf)  {
	//关闭dnsmasq
	Shellout("/etc/init.d/dnsmasq stop",conf.Workdir)
	//启动 restart redirect
	Shellout("/etc/init.d/redirect restart",conf.Workdir)
	//清理防火墙
	Stopspeed()
	//加载防火墙规则
	iptables := strings.Split(IptablesFull,"\n")
	for _,i := range iptables {
		if strings.Contains(i,"$ssserver") {
			c := strings.Replace(i,"$ssserver",Port.Ipaddr,-1)
			Shellout(c,conf.Workdir)
		}else if strings.Contains(i,"${UDPPort}") {
			c := strings.Replace(i,"${UDPPort}",Port.UDPPort,-1)
			Shellout(c,conf.Workdir)
		} else if strings.Contains(i,"${TCPPort}") {
			c := strings.Replace(i,"${TCPPort}",Port.TCPPort,-1)
			Shellout(c,conf.Workdir)
		}else {
			Shellout(i,conf.Workdir)
		}
	}
}

func DomesticSpeed(l LineConf)  {
	//关闭dnsmasq
	Shellout("/etc/init.d/dnsmasq stop",conf.Workdir)
	//启动 restart redirect
	Shellout("/etc/init.d/redirect restart",conf.Workdir)
	//清理防火墙
	Stopspeed()
	//加载防火墙规则
	iptables := strings.Split(IptablesFull,"\n")
	for _,i := range iptables {
		if strings.Contains(i,"$ssserver") {
			c := strings.Replace(i,"$ssserver",Port.Ipaddr,-1)
			Shellout(c,conf.Workdir)
		}else if strings.Contains(i,"${UDPPort}") {
			c := strings.Replace(i,"${UDPPort}",Port.UDPPort,-1)
			Shellout(c,conf.Workdir)
		} else if strings.Contains(i,"${TCPPort}") {
			c := strings.Replace(i,"${TCPPort}",Port.TCPPort,-1)
			Shellout(c,conf.Workdir)
		}else {
			Shellout(i,conf.Workdir)
		}
	}
}

func Stopspeed()  {
	//关闭 restart redirect
	Shellout("/etc/init.d/redirect stop",conf.Workdir)
	//启动dnsmasq
	Shellout("/etc/init.d/dnsmasq start",conf.Workdir)
	//重启防火墙
	Shellout("/etc/init.d/firewall restart",conf.Workdir)
	//删除默认路由
	Shellout("ip rule  del fwmark 0x1/0x1",conf.Workdir)
	Shellout("ip route del local 0.0.0.0/0 dev lo table 100",conf.Workdir)
}