// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020-2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

// DeviceServiceName contains the name of device service struct in the DIC.
var DeviceServiceName = di.TypeInstanceToName(models.DeviceService{})

// DeviceServiceFrom helper function queries the DIC and returns device service struct.
func DeviceServiceFrom(get di.Get) *models.DeviceService {
	return get(DeviceServiceName).(*models.DeviceService)
}
