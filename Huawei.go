package main

import (
	"crypto/md5"
	"io"
	"fmt"
	"strconv"
	"strings"
)

func HuaweiOld(imei string) []string {
	preUCSalt := md5.New()
	preFCSalt := md5.New()
	io.WriteString(preUCSalt,"hwe620datacard")
	io.WriteString(preFCSalt,"e630upgrade")
	pre2UCSalt := fmt.Sprintf("%x", preUCSalt.Sum(nil))
	pre2FCSalt := fmt.Sprintf("%x", preFCSalt.Sum(nil))
	fmt.Println(pre2UCSalt + " && " + pre2FCSalt)
	out := "Unlock: " + HuaweiCode(imei, pre2UCSalt[8:24]) + " Flash:" + HuaweiCode(imei, pre2FCSalt[8:24])
	return strings.Fields(out)
}

func HuaweiCode(imei string,salt string) string {
	var (
		//imei = "123456789012347"
		appByte []byte
	)
	preIn := imei+salt//"5e8dd316726b0335"
	mySum := md5.New()
	io.WriteString(mySum, preIn)
	myByte := mySum.Sum(nil)
	appByte = append(appByte,(myByte[0] ^ myByte[4] ^ myByte[8] ^ myByte[12]),(myByte[1] ^ myByte[5] ^ myByte[9] ^ myByte[13]),(myByte[2] ^ myByte[6] ^ myByte[10] ^ myByte[14]),(myByte[3] ^ myByte[7] ^ myByte[11] ^ myByte[15]))
	part1 := []byte{0x1, 0xff, 0xff, 0xff}
	part2 := []byte{0x2, 0x00, 0x00, 00}
	appByte = []byte{(appByte[0] & part1[0]), (appByte[1] & part1[1]),(appByte[2] & part1[2]),(appByte[3] & part1[3])}
	appByte = []byte{(appByte[0] | part2[0]), (appByte[1] | part2[1]),(appByte[2] | part2[2]),(appByte[3] | part2[3])}
	preout := fmt.Sprintf("%x",appByte)
	out,err := strconv.ParseUint(preout,16,64)
	if err != nil{
		panic(err)
	}
	return strconv.FormatUint(out,10)
}
