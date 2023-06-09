/**
@description service文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 10:20
*/

package service

import (
	"context"
	"errors"
	"os"
	"sync"

	"github.com/gq-tang/edgex-sdk/container"
	sdkCommon "github.com/gq-tang/edgex-sdk/internal/common"
	"github.com/gq-tang/edgex-sdk/internal/config"

	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/flags"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/handlers"
	bootstrapInterfaces "github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/interfaces"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/startup"
	bootstrapTypes "github.com/edgexfoundry/go-mod-bootstrap/v3/config"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
)

const EnvInstanceName = "EDGEX_INSTANCE_NAME"

type service struct {
	serviceKey string

	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup

	lc     logger.LoggingClient
	config *config.ConfigurationStruct
	flags  *flags.Default
	dic    *di.Container

	addDevCallback    []DeviceAction
	updateDevCallback []DeviceAction
	deleteDevCallback []DeviceAction

	updateProfileCallback []ProfileAction
}

func NewService(serviceKey string) (*service, error) {
	var service service

	if serviceKey == "" {
		return nil, errors.New("please specify device service name")
	}
	ctx, cancel := context.WithCancel(context.Background())
	service.ctx = ctx
	service.cancel = cancel

	service.serviceKey = serviceKey
	service.config = &config.ConfigurationStruct{}

	return &service, nil
}

func (s *service) Run() error {
	var instanceName string
	startupTimer := startup.NewStartUpTimer(s.serviceKey)

	additionalUsage :=
		"    -i, --instance                  Provides a service name suffix which allows unique instance to be created\n" +
			"                                    If the option is provided, service name will be replaced with \"<name>_<instance>\"\n"
	s.flags = flags.NewWithUsage(additionalUsage)
	s.flags.FlagSet.StringVar(&instanceName, "instance", "", "")
	s.flags.FlagSet.StringVar(&instanceName, "i", "", "")
	s.flags.Parse(os.Args[1:])
	s.setServiceName(instanceName)

	s.dic = di.NewContainer(di.ServiceConstructorMap{
		container.ConfigurationName: func(get di.Get) interface{} {
			return s.config
		},
	})

	wg, deferred, successful := bootstrap.RunAndReturnWaitGroup(
		s.ctx,
		s.cancel,
		s.flags,
		s.serviceKey,
		common.ConfigStemDevice,
		s.config,
		nil,
		startupTimer,
		s.dic,
		true,
		bootstrapTypes.ServiceTypeDevice,
		[]bootstrapInterfaces.BootstrapHandler{
			handlers.NewClientsBootstrap().BootstrapHandler,
			s.messageBusBootstrapHandler,
			handlers.NewServiceMetrics(s.serviceKey).BootstrapHandler, // Must be after Messaging
			handlers.NewStartMessage(s.serviceKey, sdkCommon.ServiceVersion).BootstrapHandler,
		})

	defer func() {
		deferred()
	}()

	if !successful {
		s.cancel()
		return errors.New("bootstrapping failed")
	}

	s.wg = wg
	return nil
}

func (s *service) Stop() {
	if s.wg != nil {
		s.cancel()
		s.wg.Wait()
	}
}

func (s *service) setServiceName(instanceName string) {
	envValue := os.Getenv(EnvInstanceName)
	if len(envValue) > 0 {
		instanceName = envValue
	}

	if len(instanceName) > 0 {
		s.serviceKey = s.serviceKey + "_" + instanceName
	}
}
