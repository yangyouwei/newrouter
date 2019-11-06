package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadAlLFromFile(p string) string {
	fi, err := os.Open(p)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	defer fi.Close()
	a,err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println(err)
	}
	return string(a)
}
