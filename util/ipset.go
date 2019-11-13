package util

var (
	IPSETforeignCMD string = "/usr/sbin/ipset create foreign hash:net maxelem 1000000"
	IPSETdomesicCMD string = "/usr/sbin/ipset create domestic hash:net maxelem 1000000 "
	IpsetAdd string = "/usr/sbin/ipset restore -f "
	IpsetClear string = "/usr/sbin/ipset destroy "

)

func InitIpset(m string,c string)  {
	//判断加速模式
	//获取国家代码
	//生成ipset 临时文件
	//创建并加载ip
}

func IpetClear()  {
	//清除本国ip
	//清理国外ip
}