/**
@description service文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 10:20
*/

package service

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/gq-tang/edgex-sdk/config"
)

type deviceService struct {
	lc     logger.LoggingClient
	config *config.ConfigurationStruct
	dic    *di.Container
}

func NewDeviceService() (*deviceService, error) {
	return &deviceService{
		lc:  nil,
		dic: nil,
	}, nil
}
