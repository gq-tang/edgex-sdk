/**
@description device文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 14:44
*/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/gq-tang/edgex-sdk/pkg/service"
)

func main() {
	ds, _ := service.NewDeviceService("test")
	print := func(device dtos.Device) error {
		data, _ := json.MarshalIndent(device, "", " ")
		fmt.Println("device add execute: ", string(data))
		return nil
	}
	ds.HandleDeviceAdd(print)
	ds.HandleDeviceUpdate(print)
	ds.HandleDeviceDelete(print)
	if err := ds.Run(); err != nil {
		fmt.Println("err:", err)
		return
	}

}
