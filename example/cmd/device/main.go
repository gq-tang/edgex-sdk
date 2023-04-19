/**
@description device文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 14:44
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/gq-tang/edgex-sdk/pkg/service"
	"os"
	"os/signal"
	"time"
)

func main() {
	service.BootStrap("test")
	fmt.Println("after BootStrap")
	ds := service.Service()
	print := func(method string) func(device dtos.Device) error {
		return func(device dtos.Device) error {
			data, _ := json.MarshalIndent(device, "", " ")
			fmt.Printf("device %s execute: %s", method, string(data))
			return nil
		}
	}
	ds.HandleDeviceAdd(print("add"))
	ds.HandleDeviceUpdate(print("update"))
	ds.HandleDeviceDelete(print("delete"))

	time.Sleep(time.Second * 2)
	cli := service.CommandClient()
	params := map[string]string{
		"switch": "0",
	}
	res, err := cli.IssueSetCommandByName(context.Background(), "94805a42b456", "switch", params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	service.Stop()
}
