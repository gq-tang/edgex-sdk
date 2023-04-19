/**
@description service文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 13:51
*/

package service

import (
	"context"
	"encoding/json"
	"fmt"

	bootstrapContainer "github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
	"github.com/edgexfoundry/go-mod-messaging/v3/pkg/types"
	"github.com/gq-tang/edgex-sdk/internal/container"
)

type DeviceAction func(device dtos.Device) error

type ProfileAction func(profile dtos.DeviceProfile) error

func (s *service) HandleDeviceAdd(fn DeviceAction) {
	s.addDevCallback = append(s.addDevCallback, fn)
}

func (s *service) HandleDeviceUpdate(fn DeviceAction) {
	s.updateDevCallback = append(s.updateDevCallback, fn)
}

func (s *service) HandleDeviceDelete(fn DeviceAction) {
	s.deleteDevCallback = append(s.deleteDevCallback, fn)
}

func (s *service) HandleProfileUpdate(fn ProfileAction) {
	s.updateProfileCallback = append(s.updateProfileCallback, fn)
}

func (s *service) metadataSystemEventsCallback(ctx context.Context, dic *di.Container) errors.EdgeX {
	lc := bootstrapContainer.LoggingClientFrom(dic.Get)
	messageBusInfo := container.ConfigurationFrom(dic.Get).MessageBus

	metadataSystemEventTopic := common.BuildTopic(messageBusInfo.GetBaseTopicPrefix(),
		common.MetadataSystemEventSubscribeTopic, "#")

	lc.Infof("Subscribing to System Events on topic: %s", metadataSystemEventTopic)

	messages := make(chan types.MessageEnvelope)
	messageErrors := make(chan error)
	topics := []types.TopicChannel{
		{
			Topic:    metadataSystemEventTopic,
			Messages: messages,
		},
	}

	messageBus := bootstrapContainer.MessagingClientFrom(dic.Get)
	err := messageBus.Subscribe(topics, messageErrors)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				lc.Infof("Exiting waiting for MessageBus '%s' topic messages", metadataSystemEventTopic)
				return
			case err = <-messageErrors:
				lc.Error(err.Error())
			case msgEnvelope := <-messages:
				lc.Debugf("System event received on message queue. Topic: %s, Correlation-id: %s", msgEnvelope.ReceivedTopic, msgEnvelope.CorrelationID)

				var systemEvent dtos.SystemEvent
				err := json.Unmarshal(msgEnvelope.Payload, &systemEvent)
				if err != nil {
					lc.Errorf("failed to JSON decoding system event: %s", err.Error())
					continue
				}

				switch systemEvent.Type {
				case common.DeviceSystemEventType:
					err = s.deviceSystemEventAction(systemEvent, dic)
					if err != nil {
						lc.Error(err.Error(), common.CorrelationHeader, msgEnvelope.CorrelationID)
					}
				case common.DeviceProfileSystemEventType:
					err = s.deviceProfileSystemEventAction(systemEvent, dic)
					if err != nil {
						lc.Error(err.Error(), common.CorrelationHeader, msgEnvelope.CorrelationID)
					}
				case common.ProvisionWatcherSystemEventType:
					err = s.provisionWatcherSystemEventAction(systemEvent, dic)
					if err != nil {
						lc.Error(err.Error(), common.CorrelationHeader, msgEnvelope.CorrelationID)
					}
				case common.DeviceServiceSystemEventType:
					err = s.deviceServiceSystemEventAction(systemEvent, dic)
					if err != nil {
						lc.Error(err.Error(), common.CorrelationHeader, msgEnvelope.CorrelationID)
					}
				default:
					lc.Errorf("unknown system event type %s", systemEvent.Type)
					continue
				}
			}
		}
	}()

	return nil
}

func (s *service) deviceSystemEventAction(systemEvent dtos.SystemEvent, dic *di.Container) error {
	var device dtos.Device
	err := systemEvent.DecodeDetails(&device)
	if err != nil {
		return fmt.Errorf("failed to decode %s system event details: %s", systemEvent.Type, err.Error())
	}

	switch systemEvent.Action {
	case common.SystemEventActionAdd:
		for _, callback := range s.addDevCallback {
			if err := callback(device); err != nil {
				s.lc.Errorf("device add callback error: %s", err)
				continue
			}
		}
	case common.SystemEventActionUpdate:
		for _, callback := range s.updateDevCallback {
			if err := callback(device); err != nil {
				s.lc.Errorf("device add callback error: %s", err)
				continue
			}
		}
	case common.SystemEventActionDelete:
		for _, callback := range s.deleteDevCallback {
			if err := callback(device); err != nil {
				s.lc.Errorf("device add callback error: %s", err)
				continue
			}
		}
	default:
		return fmt.Errorf("unknown %s system event action %s", systemEvent.Type, systemEvent.Action)
	}

	return err
}

func (s *service) deviceProfileSystemEventAction(systemEvent dtos.SystemEvent, dic *di.Container) error {
	var deviceProfile dtos.DeviceProfile
	err := systemEvent.DecodeDetails(&deviceProfile)
	if err != nil {
		return fmt.Errorf("failed to decode %s system event details: %s", systemEvent.Type, err.Error())
	}

	switch systemEvent.Action {
	case common.SystemEventActionUpdate:
		for _, callback := range s.updateProfileCallback {
			if err := callback(deviceProfile); err != nil {
				s.lc.Errorf("device add callback error: %s", err)
				continue
			}
		}
	case common.SystemEventActionAdd, common.SystemEventActionDelete:
		break
	default:
		return fmt.Errorf("unknown %s system event action %s", systemEvent.Type, systemEvent.Action)
	}

	return err
}

func (s *service) provisionWatcherSystemEventAction(systemEvent dtos.SystemEvent, dic *di.Container) error {
	var pw dtos.ProvisionWatcher
	err := systemEvent.DecodeDetails(&pw)
	if err != nil {
		return fmt.Errorf("failed to decode %s system event details: %s", systemEvent.Type, err.Error())
	}

	switch systemEvent.Action {
	case common.SystemEventActionAdd:

	case common.SystemEventActionUpdate:

	case common.SystemEventActionDelete:

	default:
		return fmt.Errorf("unknown %s system event action %s", systemEvent.Type, systemEvent.Action)
	}

	return err
}

func (s *service) deviceServiceSystemEventAction(systemEvent dtos.SystemEvent, dic *di.Container) error {
	var deviceService dtos.DeviceService
	err := systemEvent.DecodeDetails(&deviceService)
	if err != nil {
		return fmt.Errorf("failed to decode %s system event details: %s", systemEvent.Type, err.Error())
	}

	switch systemEvent.Action {
	case common.SystemEventActionUpdate:

	case common.SystemEventActionAdd, common.SystemEventActionDelete:
		break
	default:
		return fmt.Errorf("unknown %s system event action %s", systemEvent.Type, systemEvent.Action)
	}

	return err
}
