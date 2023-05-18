package main

import (
	"fmt"
	"crypto/sha1"
	"io"
	"encoding/hex"
	"strconv"
	"strings"
)

func alcatelMW43TM(imei string, imp string, perm string, xorn string) string {  //C700 & friends
	var (
		swapimei = "08" //08 21 43 65 87 09 21 43 07
		permimei = "" //perm of swap
		doxor [20]int
		myxor byte
		pre = ""
		doswap [9]int
		err error
		tmpswap [9]string

	)
	simei := imei + "0"
	for i:=0; i < (len(simei)-1); i++{ //swaps imei
		swapimei = swapimei + string(simei[i+1]) + string(simei[i])
		i++
	}
	for i:=0; i<(len(perm)-1); i++ {  //sets up permimei
		doswap[i],err = strconv.Atoi(string(perm[i]))
		if err != nil{
			panic(err)
		}
	}
	j:=0
	for i:=0; i<9; i++{
		tmpswap[i] = string(swapimei[j]) + string(swapimei[j+1])
		j = j + 2
	}
	for i:=0; i<(len(tmpswap)); i++{ //creates permimei
		permimei = permimei + string(tmpswap[(doswap[i])])
	}
	presha1input := imp + swapimei + swapimei + permimei //forms sha1message
	sha1input,err := hex.DecodeString(presha1input)
	if err != nil{
		panic(err)
	}
	digest := sha1.New()
	io.WriteString(digest,string(sha1input))
	hexStr := fmt.Sprintf("%x",digest.Sum(nil))
	bRay,err := hex.DecodeString(string(hexStr))
	if err != nil{
		panic(err)
	}
	xorder,err := hex.DecodeString(xorn) //parses xorder
	if err != nil{
		panic(err)
	}
	for i := 0; i < (len(xorder)-1); i++{  //makes []int of xorder
		doxor[i],err = strconv.Atoi(fmt.Sprintf("%d",xorder[i]))
		if err != nil{
			panic(err)
		}
	}
	for i:= 0; i < (len(doxor)-1); i++{  //xors bytes
		myxor = bRay[(doxor[i])] ^ bRay[(doxor[i+1])] ^ bRay[(doxor[i+2])] ^ bRay[(doxor[i+3])] ^ bRay[(doxor[i+4])]
		pre += fmt.Sprintf("%02x", myxor)  //hex encodes bytes to string
		i = i+4
	}
	out,err := strconv.ParseInt(pre,16,64) //parses hexstring to base 10 code
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d",out)  //makes code a string
}

func alcatelC700(model string, imei string) []string {
	var (
		modata = map[string][]string{
		"duck":[]string{"8F", "BE", "876543210", "110A090201100B0803000F0C0704130E0D060512"}, //MANDARINA DUCK | C820 | C825
		"playboy":[]string{"3C", "E2", "785432106", "0B121307000C110806010D100905020E0F0A0403"}, //C717 | C700 | C701 | EL03 | PLAYBOY
		"misssixty":[]string{"6C", "B9", "456132807", "0503011311040200121007090B0D0F06080A0C0E"}, //MISSSIXTY | S520
		"s215":[]string{"74", "9A", "547682031", "0504010D0F0C06070A0010080B0E031202111309"}, //S215 | S218 | S219 | S320 | S321
		"s853":[]string{"53", "AE", "876543210", "0004080C100105090D1102060A0E1203070B0F13"}} //S853  !!!PERM UNKNOWN!!!
		modeldb = map[string][]string{
		"MandarinaDuck":modata["duck"],
		"C820":modata["duck"],
		"C825":modata["duck"],
		"Playboy":modata["playboy"],
		"C717":modata["playboy"],
		"C700":modata["playboy"],
		"EL03":modata["playboy"],
		"MissSixty":modata["misssixty"],
		"S520":modata["misssixty"],
		"S215":modata["s215"],
		"S218":modata["s215"],
		"S219":modata["s215"],
		"S320":modata["s215"],
		"S321":modata["s215"],
		"S853":modata["duck"]}
	)
	whichmodel := modeldb[model]
	return strings.Fields("NCK:"+ alcatelC700Calc(imei,whichmodel[0],whichmodel[2], whichmodel[3]) + " SPCK:" + alcatelC700Calc(imei,whichmodel[1],whichmodel[2], whichmodel[3]))
}
