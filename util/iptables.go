package util

type LineConf struct {
	Ipaddr string
	TCPPort string
	UDPPort string
}

var IptablesFull string = `ip rule add fwmark 0x01/0x01 table 100
ip route add local 0.0.0.0/0 dev lo table 100
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
ip route add local 0.0.0.0/0 dev lo table 100
/usr/sbin/ptables -t mangle -N SS-UDP
/usr/sbin/ptables -t mangle -A SS-UDP -d $ssserver -j RETURN
/usr/sbin/ptables -t mangle -A SS-UDP -p udp -j TPROXY --tproxy-mark 0x01/0x01 --on-port ${UDPPort}
/usr/sbin/ptables -t mangle -A PREROUTING -p udp -m set --match-set foreign dst -j SS-UDP
/usr/sbin/ptables -t nat -N SS-TCP
/usr/sbin/ptables -t nat -A SS-TCP -d $ssserver -j RETURN
/usr/sbin/ptables -t nat -A SS-TCP -p tcp -j REDIRECT --to-ports ${TCPPort}
/usr/sbin/ptables -t nat -A PREROUTING -p tcp -m set --match-set foreign dst -j SS-TCP`

var IptablesDomestic string = `ip rule add fwmark 0x01/0x01 table 100
ip route add local 0.0.0.0/0 dev lo table 100
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
	
}

func MultiSpeed(l LineConf)  {
	
}

func DomesticSpeed(l LineConf)  {

}

func Stopspeed()  {

}