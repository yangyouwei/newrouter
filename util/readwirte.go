package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAlLFromFile(p string) (string, error) {
	fi, err := os.Open(p)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return "",err
	}
	defer fi.Close()
	a,err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println(err)
	}
	return string(a),nil
}

func Delfile(s string)  {
	//删除
	err1 := os.Remove(s)
	if err1 != nil {
		// 删除失败
		fmt.Println(err1)
	} else {
		fmt.Println("delete file: ",s)
	}
}

func IfFileIsExist(s string) bool {
	_, err := os.Stat(s)
	if err == nil {
	    return true
	}
	if os.IsExist(err) {
		return true
	}
	return false
}