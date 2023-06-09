// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package container

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"

	"github.com/gq-tang/edgex-sdk/internal/config"
)

// ConfigurationName contains the name of device service's ConfigurationStruct implementation in the DIC.
var ConfigurationName = di.TypeInstanceToName(config.ConfigurationStruct{})

// ConfigurationFrom helper function queries the DIC and returns device service's ConfigurationStruct implementation.
func ConfigurationFrom(get di.Get) *config.ConfigurationStruct {
	return get(ConfigurationName).(*config.ConfigurationStruct)
}
