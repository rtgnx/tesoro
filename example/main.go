package main

import (
	"fmt"
	"log"
	"strings"

	"encoding/json"

	"github.com/chzyer/readline"
	"github.com/conejoninja/trezor"
	"github.com/zserge/hid"
)

func main() {

	var c trezor.TrezorClient

	numberDevices := 0
	hid.UsbWalk(func(device hid.Device) {
		info := device.Info()
		// 0x534c : 21324 vendor
		// 0x0001 : 1     product
		if info.Vendor == 21324 && info.Product == 1 {
			numberDevices++
			c.SetTransport(device)
		}
	})
	if numberDevices == 0 {
		fmt.Println("No TREZOR devices found, make sure your TREZOR device is connected")
	} else {
		fmt.Printf("Found %d TREZOR devices connected\n", numberDevices)
		shell(c)
		defer c.CloseTransport()
	}
}

func shell(c trezor.TrezorClient) {
	var str string
	var msgType uint16
	rl, err := readline.NewEx(&readline.Config{
		Prompt: ">",
	})
	if err != nil {
		panic(err)
	}

	defer rl.Close()
	log.SetOutput(rl.Stderr())

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("ERR", err)
			break
		}
		args := strings.Split(strings.ToLower(line), " ")

		switch args[0] {
		case "ping":
			if len(args) < 2 {
				fmt.Println("Missing parameters")
			} else {
				str, msgType = c.Call(c.Ping(strings.Join(args[1:], " ")))
			}
			break
		case "signmessage":
			if len(args) < 2 {
				fmt.Println("Missing parameters")
			} else {
				msg := strings.Join(args[1:], " ")
				str, msgType = c.Call(c.SignMessage([]byte(msg)))
				if msgType == 26 {
					str, msgType = c.Call(c.ButtonAck())
					var sm trezor.SignMessage
					err = json.Unmarshal([]byte(str), &sm)
					if err == nil {
						sm.Message = msg
					}
					smJSON, _ := json.Marshal(sm)
					str = string(smJSON)
				}
			}
			break
		case "getaddress":
			str, msgType = c.Call(c.GetAddress())
			break
		default:
			if msgType == 18 { // PIN INPUT
				str, msgType = c.Call(c.PinMatrixAck(line))
			} else {
				fmt.Println("Unknown command")
				str = line
				msgType = 999
			}
			break
		}
		fmt.Println(str, msgType)
	}
}