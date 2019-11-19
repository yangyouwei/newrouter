package util

import (
	"fmt"
)



func SpeedCtl(m string)  {
	speedmode := m
	switch {
	case speedmode == "full":
		fmt.Println("mode is full")
		//启动redirect
		if SpeedModInit.Full {
			SwitchRedirect(true)
			//加载防火墙
			fmt.Println("loading iptables")
			ChSpeedMod("fullspeed")
		}else {
			return
		}

	case speedmode == "foreigen":
		fmt.Println("mode is foreigen")
		//启动redirect
		if SpeedModInit.OnlyForeigen {
			SwitchRedirect(true)
			//加载防火墙
			fmt.Println("loading iptables")
			ChSpeedMod("domsticspeed")
		}else {
			return
		}

	case speedmode == "multicontry":
		fmt.Println("mode is multicontry")
		//启动redirect
		if SpeedModInit.DesignatedCountry {
			SwitchRedirect(true)
			//加载防火墙
			fmt.Println("loading iptables")
			ChSpeedMod("multispeed")
		}else {
			return
		}

	case speedmode == "stopspeed":
		fmt.Println("mode is stopspeed")
		//清空防火墙
		Stopspeed()
		//启动关闭加速
		SwitchRedirect(false)
	default:
		fmt.Println("mode is stopspeed")
		//清空防火墙
		Stopspeed()
		//启动关闭加速
		SwitchRedirect(false)
	}
}



