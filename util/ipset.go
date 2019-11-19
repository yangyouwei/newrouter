package util

import (
	"fmt"
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/models"
	"io/ioutil"
	"os"
	"strings"
)

type SpeedMod struct {
	Full bool
	OnlyForeigen bool
	DesignatedCountry bool
	StopSpeed bool
}

var SpeedModInit SpeedMod

var (
	IPSETforeignCMD string = "/usr/sbin/ipset create foreign hash:net maxelem 1000000"
	IPSETdomesicCMD string = "/usr/sbin/ipset create domestic hash:net maxelem 1000000 "
	IpsetAdd string = "/usr/sbin/ipset restore -f "
	Tmpdir string = "/tmp/"
)

func InitIpset(s *models.Sysstr) {
	fmt.Println("creating ipset .")
	//获取国家ip列表文件名称
	cs := GetContryNames(conf.ContryIPlist)
	if cs == nil {
		//不存在的话。改为加速模式为不加速。
		fmt.Println("没有找到ip列表，请下载本国家ip.否则加速模式，只能使用全局加速模式。")
		return
	}

	//判断国家ip文件是否存在
	rs := IfFileIsExist(conf.ContryIPlist + s.Contry)
	if !rs {
		fmt.Println("请下载本国国家ip,否则加速模式，不能使用仅加速国外模式")
		return
	}

	//判断临时文件是否存在。存在的话删除
	rs1 := IfFileIsExist(conf.IpsetFfile  + s.Contry)
	if rs1 {
		Delfile(conf.IpsetFfile +  s.Contry)
	}

	//创建ipset文件
	ips, err := ReadAlLFromFile(conf.ContryIPlist + s.Contry)
	if err != nil {
		fmt.Println(err)
	}
	IP := strings.Split(ips, "\n")
	var tmpstr string
	for _, i := range IP {
		//创建临时文件
		tmpstr = "add domestic " + i + "\n"
		//ipset restore
		Shellout(IpsetAdd + Tmpdir + s.Contry, conf.Workdir)
	}

	linefile, err := os.OpenFile(Tmpdir + s.Contry, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	linefile.Write([]byte(tmpstr))

	//创建国外ipset文件
	rs2 := IfFileIsExist(conf.IpsetFfile  + "foreign")
	if rs2 {
		Delfile(conf.IpsetFfile +  "foreign")
	}

	str := strings.Split(s.Foreigencontry, ";")

	if str == nil {
		fmt.Println("not foreigen config.")
		return
	}else {
		for _,cns := range str {
			r := IfFileIsExist(conf.ContryIPlist  + cns )
			if r {

			}else {
				fmt.Println("contry file not found.please download contry file.")
			}
		}
	}
}

func GetContryNames(d string) (n []string ) {
	files, _ := ioutil.ReadDir(conf.ContryIPlist)
	for _, f := range files {
		n = append(n,f.Name())
	}
	return
}

func RstoreIpset(s *models.Sysstr)  {
	//判断国家ip文件是否存在
	rs := IfFileIsExist(conf.ContryIPlist + s.Contry)
	if !rs {
		fmt.Println("请下载本国家ip,请重新初始化ipset")
		SpeedModInit.OnlyForeigen = false
		return
	}else {
		SpeedModInit.OnlyForeigen = true
		Shellout(IPSETdomesicCMD, conf.Workdir)
		Shellout(IpsetAdd + conf.IpsetFfile + s.Contry, conf.Workdir)
	}

	rs2 := IfFileIsExist(conf.IpsetFfile  + "foreign")
	if rs2 {
		SpeedModInit.DesignatedCountry = true
		Shellout(IPSETforeignCMD, conf.Workdir)
		Shellout(IpsetAdd + conf.IpsetFfile + "foreign",conf.Workdir)
	}else {
		fmt.Println("请初始化ipset")
		SpeedModInit.DesignatedCountry = false
		return
	}
	SpeedModInit.Full = true
	SpeedModInit.StopSpeed = true
}