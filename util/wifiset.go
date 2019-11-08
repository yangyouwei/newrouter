package util

var(
	WifiSetCommandSSID string = "wireless.default_radio0.ssid="
	WifiSetCommandENCRY string = "wireless.default_radio0.encryption="
	WifiSetCommandPassword string = "wireless.default_radio0.key="
	WifiCommitCommand string = "/sbin/uci commit wireless"
	WifiRestartCommand string = "/sbin/wifi "
)

func SETWIFI(n string,p string)  {
	Shellout(WifiSetCommandSSID+n,"/")
	Shellout(WifiSetCommandENCRY+"psk2","/")
	Shellout(WifiSetCommandPassword+p,"/")
	Shellout(WifiCommitCommand,"/")
	Shellout(WifiRestartCommand+"down","/")
	Shellout(WifiRestartCommand,"/")
}