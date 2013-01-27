package main

import (
	"strings"
	"strconv"
	"fmt"
	"crypto/md5"
	"io"
)

//ZTE SFR 114
//ZTE SFR 231
//ZTE SFR 232
//ZTE SFR 241
//ZTE SFR 251
//ZTE SFR 251MSN
//ZTE SFR 341
//ZTE SFR 342
//ZTE SFR 343
//ZTE ORANGE Vegas
//ZTE ORANGE MIAMI
//ZTE VODAFONE INDIE
//ZTE X760
//ZTE X761
//ZTE X960
//ZTE GX760
//ZTE GX761
//ZTE A261+
//ZTE GR230
//ZTE X990
//ZTE T-Mobile Vairy Touch
//ZTE N261 = N281
//ZTE X990
//ZTE X991
//Orange Rio

func zteOld(imei string) []string{ //zteB02?
	var (
		magic = [12]int{6, 8, 8, 9, 5, 0, 0, 0, 0, 0, 0, 0}
		digits []int
		nck = ""
		spck = ""
		crosssum = 0
	)
	for i,_ := range imei[3:15]{
		imeiAsInt, err := strconv.Atoi(string(imei[i+3]))
		if err != nil{
			panic(err)
		}
		digits = append(digits, imeiAsInt)
	}
	for i,_ := range digits{
		crosssum += digits[i]
	}
	for i,_ := range(digits){
		code := (digits[i]*crosssum + digits[11-i]*8 + magic[i]) % 10
		nck += strconv.Itoa(code)
		spck += strconv.Itoa((code + digits[11-i]) % 10)
	}
	return strings.Fields("NCK:"+nck + " SPCK:" + spck)
}

func zteB03(imei string) string { //zteB03
	var (
		key [8]int64
		out = ""
		pre [16]int64
		err error
	)
	mySum := md5.New()
	io.WriteString(mySum, imei)
	imeiMD5 := mySum.Sum(nil)
	for i,_ := range imeiMD5{
		pre[i],err=strconv.ParseInt(fmt.Sprintf("%x", imeiMD5[i]),16,0)
		if err != nil{
			panic(err)
		}
	}
	for i:=0; i <=7; i++{
		key[i] = (((pre[i] + pre[i+8]) & int64(0xFF)*int64(0x09))/int64(0xFF))
		out += fmt.Sprintf("%x", key[i])
	}
	return "NCK:"+out
}

func zteB04(imei string) string {  //zteB04
	var (
		key [8]int64
		out = ""
		pre [16]int64
		err error
	)
	mySum := md5.New()
	io.WriteString(mySum, imei)
	imeiMD5 := mySum.Sum(nil)
	for i,_ := range imeiMD5{
		pre[i],err=strconv.ParseInt(fmt.Sprintf("%x", imeiMD5[i]),16,0)
		if err != nil{
			panic(err)
		}
	}
	for i:=0; i <=7; i++{
		key[i] = (((pre[i] + pre[i+8] + pre[i+4] + 40) & int64(0xFF)*int64(0x09))/int64(0xFF))
		out += fmt.Sprintf("%x", key[i])
	}
	return "NCK:"+out
}
