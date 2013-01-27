package main

import (
	"os"
	"fmt"
	"encoding/json"
	"strings"
)

func outJSON(src string, YesNo string) []byte{
	prep := src[0:len(src)-1]
	preout := strings.Split(prep, ",")
	out,err := json.Marshal(preout)
	if err != nil{
		panic(err)
	}
	return out
}

func main() {
	userArgs := os.Args
	if userArgs[1] == "todo" {
		todo()
	}
	if userArgs[1] == "blackberryMEP" {
		if userArgs[2] == "unlock" {
			preout := Blackberry(userArgs[3], userArgs[4])
			out := outJSON(preout)
			fmt.Println(string(out))
		}
		if userArgs[2] == "getSupported" {
			out,err := json.Marshal(getSupported("MEP"))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
	if userArgs[1] == "blackberryPRD" {
		if userArgs[2] == "unlock" {
			preout := Blackberry(userArgs[3], userArgs[4])
			out := outJSON(preout)
			fmt.Println(string(out))
		}
		if userArgs[2] == "getSupported" {
			out,err := json.Marshal(getSupported("PRD"))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
	if userArgs[1] == "huaweiOld" {
		if userArgs[2] == "unlock" {
			out,err := json.Marshal(HuaweiOld(userArgs[3]))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
	if userArgs[1] == "zteOld" {
		if userArgs[2] == "unlock" {
			out,err := json.Marshal(zteOld(userArgs[3]))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
	if userArgs[1] == "zteB03" {
		if userArgs[2] == "unlock" {
			out,err := json.Marshal(zteB03(userArgs[3]))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
	if userArgs[1] == "zteB04" {
		if userArgs[2] == "unlock" {
			out,err := json.Marshal(zteB04(userArgs[3]))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
	if userArgs[1] == "alcatelC700" {
		if userArgs[2] == "unlock" {
			out,err := json.Marshal(alcatelC700(userArgs[3], userArgs[4]))
			if err != nil{
				panic(err)
			}
			fmt.Println(string(out))
		}
	}
}
