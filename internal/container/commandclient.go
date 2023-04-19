/**
@description container文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/19 9:13
*/

package container

import (
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
)

// CommandClientName contains the name of the CommandClient's implementation in the DIC.
var CommandClientName = di.TypeInstanceToName((*interfaces.CommandClient)(nil))

func CommandClientFrom(get di.Get) interfaces.CommandClient {
	return get(CommandClientName).(interfaces.CommandClient)
}
