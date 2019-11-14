package util

import (
	"fmt"
	"github.com/yangyouwei/newrouter/conf"
	"github.com/yangyouwei/newrouter/models"
	"io/ioutil"
	"os"
	"strings"
)

var (
	IPSETforeignCMD string = "/usr/sbin/ipset create foreign hash:net maxelem 1000000"
	IPSETdomesicCMD string = "/usr/sbin/ipset create domestic hash:net maxelem 1000000 "
	IpsetAdd string = "/usr/sbin/ipset restore -f "
	Tmpdir string = "/tmp"
)

func InitIpset(s *models.Sysstr) {
	fmt.Println("creating ipset .")
	//获取国家ip列表文件名称
	cs := GetContryNames(conf.ContryIPlist)
	if cs == nil {
		//不存在的话。改为加速模式为不加速。
		fmt.Println("下载国家ip")
		s.SpeedMod = "stopspeed"
		return
	}

	//判断国家ip文件是否存在
	rs := IfFileIsExist(conf.ContryIPlist + "/" + s.Contry)
	if !rs {
		fmt.Println("下载国家ip")
		s.SpeedMod = "stopspeed"
		return
	}

	//判断临时文件是否存在。存在的话删除
	rs1 := IfFileIsExist(Tmpdir + "/" + s.Contry)
	if rs1 {
		Delfile(Tmpdir + "/" + s.Contry)
	}

	ips, err := ReadAlLFromFile(conf.ContryIPlist + "/" + s.Contry)
	if err != nil {
		fmt.Println(err)
	}
	IP := strings.Split(ips, "\n")
	for _, i := range IP {
		//创建临时文件
		linefile, err := os.OpenFile(Tmpdir+"/"+s.Contry, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println(err)
		}
		i = "add domestic " + i
		linefile.Write([]byte(i))
		//ipset restore
		Shellout(IpsetAdd+"/"+s.Contry, conf.Workdir)
	}
	//create ipset
	Shellout(IPSETdomesicCMD, conf.Workdir)
	//restore ipset
	Shellout(IPSETdomesicCMD+Tmpdir+"/"+s.Contry, conf.Workdir)
	//删除
	Delfile(Tmpdir + "/" + s.Contry)


	//创建国外ip列表
	rs2 := IfFileIsExist(Tmpdir + "/" + "foreign")
	if rs2 {
		Delfile(Tmpdir + "/" + "foreigen")
	}

	Shellout(IPSETforeignCMD, conf.Workdir)
	for _, f := range cs {
		for _, i := range s.Foreigencontry {
			if i == f {
				//读取文件内容
				c,err := ReadAlLFromFile(i)
				if err != nil {
					fmt.Println("create foreigen failed.")
					return
				}
				IPs := strings.Split(c, "\n")
				for _, onecontry := range IPs {
					linefile, err := os.OpenFile(Tmpdir + "/" + "foreigen", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
					if err != nil {
						fmt.Println(err)
					}
					onecontry := "add domestic " + onecontry
					linefile.Write([]byte(onecontry))
				}
			}else {
				continue
			}
		}
	}
	//ipset restore
	Shellout(IpsetAdd+"foreign",conf.Workdir)
	//删除临时文件
	Delfile(Tmpdir+"/" + "foreigen")
}

//
func GetContryNames(d string) (n []string ) {
	files, _ := ioutil.ReadDir(conf.ContryIPlist)
	for _, f := range files {
		n = append(n,f.Name())
	}
	return
}