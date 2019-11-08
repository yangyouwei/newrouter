package util

var GetNetInfo string = "/bin/ubus call network.interface.wan status"

type NetINFO struct {
	WanInterfaceName string
	WanInterfaceIp string
	//on/off
	WanInterfaceSstat string
	//full domestic multiconty
	SpeedSstat string
	Contry string
}

