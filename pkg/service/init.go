/**
@description service文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 14:06
*/

package service

import (
	"context"
	"fmt"
	"os"
	"sync"

	bootstrapContainer "github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/handlers"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/startup"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/interfaces"
)

var (
	defaultService *service
	once           sync.Once
)

func BootStrap(serviceKey string) {
	once.Do(func() {
		defaultService, _ = NewService(serviceKey)

		if err := defaultService.Run(); err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	})
}

func Stop() {
	defaultService.Stop()
}

func Service() *service {
	return defaultService
}

func CommandClient() interfaces.CommandClient {
	return defaultService.CommandClient()
}

func (s *service) messageBusBootstrapHandler(ctx context.Context, wg *sync.WaitGroup, startupTimer startup.Timer, dic *di.Container) bool {
	if !handlers.MessagingBootstrapHandler(ctx, wg, startupTimer, dic) {
		return false
	}

	lc := bootstrapContainer.LoggingClientFrom(dic.Get)
	err := s.metadataSystemEventsCallback(ctx, dic)
	if err != nil {
		lc.Errorf("Failed to subscribe Metadata system events: %v", err)
		return false
	}

	return true
}
