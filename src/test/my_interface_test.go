package test

import (
	"basic/model"
	"bytes"
	"fmt"
	"strconv"
	"testing"
	//"golang.org/x/tour/reader"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	//result := ""
	// 效率更高的字符串拼装方式
	var buffer bytes.Buffer
	for index,currentInt := range ip {
		if 0 != index {
			//result += ":"
			buffer.WriteString(":")
		}
		//result += strconv.Itoa(int(currentInt))
		buffer.WriteString(strconv.Itoa(int(currentInt)))
	}
	return buffer.String()
}


func TestPrintAddr(t *testing.T) {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func TestPrintHomeWork(t *testing.T) {
	w := model.HomeWork{"computer", "Print Homework Message", "写一个实现了String方法的类，并用printf打印"}
	fmt.Printf("作业的信息为: %v\n", w)
}
