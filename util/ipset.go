package util

var (
	IPSETforeignCMD string = "ipset create foreign hash:net maxelem 1000000"
	IPSETdomesicCMD string = "ipset create domestic hash:net maxelem 1000000 "
	IpsetAdd string
	IpsetClear string
)


func IpsetAddIP(s string,ipdir string)  {

}

func IpsetClearIP()  {

}
