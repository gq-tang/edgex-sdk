/**
@description service文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/19 13:53
*/

package service

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/container"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
)

func (s *service) CommandClient() interfaces.CommandClient {
	return container.CommandClientFrom(s.dic.Get)
}
