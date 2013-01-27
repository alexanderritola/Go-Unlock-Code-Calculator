package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func createPrivatePass(MEP [16]int) [64]int {
	arrayStatic := [16]int{0x16, 0x27, 0x99, 0xC2, 0xC8, 0x99, 0xB8, 0xC9, 0xDE, 0xED, 0x77, 0xA2, 0x62, 0xD2, 0x66, 0x5E}
	Pass := [64]int{}
	for i := 0; i < 16; i++ {
		Pass[i] = arrayStatic[i] ^ MEP[i]
	}
	return Pass
}

func makeSHA1(strIMEI string, pPass [64]int, mepNumber int) string { //generates SHA1
	var (
		imei      = fmt.Sprintf("%02X", strIMEI) + "0" + strconv.Itoa(mepNumber)
		o_key     = [64]string{}
		i_key     = [64]string{}
		o_key_pad = ""
		i_key_pad = ""
	)
	for i := 0; i < 64; i++ {
		o_key[i] = fmt.Sprintf("%02X", pPass[i]^0x5C)
		i_key[i] = fmt.Sprintf("%02X", pPass[i]^0x36)
		o_key_pad += o_key[i]
		i_key_pad += i_key[i]
	}
	p1inputSha1 := i_key_pad + imei
	p1digest := sha1digest(p1inputSha1)
	p2inputSha1 := o_key_pad + p1digest
	p2digest := sha1digest(p2inputSha1)
	return p2digest
}

func sha1digest(src string) string { //digest hex to sha1hex
	p, _ := hex.DecodeString(src)
	h := sha1.New()
	h.Write(p)
	s := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(s))
}

func makeCode(hash string, mep string) string { //creates unlock code
	var (
		out  = ""
		pos  int
		size int
	)

	const mep8digits = "MEP-23361-001,MEP-30218-002,MEP-15326-002,MEP-04626-002,MEP-27501-003,MEP-31845-001,MEP-22793-001,MEP-04103-001"
	pos = strings.Count(mep8digits, mep)
	bytes, err := hex.DecodeString(hash)
	if err != nil {
		panic(err)
	}
	if pos == 0 {
		size = 16
	}
	if pos != 0 {
		size = 8
	}
	for i := 0; i < size; i++ {
		preout := bytes[i]
		preout1 := strconv.FormatInt(int64(preout), 10)
		lenpreout := len(preout1)
		preout2 := preout1[lenpreout-1]
		out += string(preout2)
	}
	return out
}

func Blackberry(myMEP string, strIMEI string) string{
	var (
		out = ""
		inMEP [16]int
	)
	if myMEP[0:3] == "MEP" {
		inMEP = findMEP(myMEP)
	}
	if myMEP[0:3] == "PRD" {
		inMEP = findPRD(myMEP)
	}
	for i := 1; i <= 5; i++ {
		precode := makeCode(makeSHA1(strIMEI, createPrivatePass(inMEP), i), myMEP)
		out+="MEP" + strconv.Itoa(i) + ":" + precode+","
	}
	return out
}
