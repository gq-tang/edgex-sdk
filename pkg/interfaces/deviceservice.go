/**
@description interfaces文件

@copyright    Copyright 2023 seva
@version      1.0.0
@author       tgq <tangguangqiang@rollingstoneiot.com>
@datetime     2023/4/18 10:13
*/

package interfaces

import (
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"
)

type DeviceServiceSDK interface {
	// AddDevice adds a new Device to the Device Service and Core Metadata
	// Returns new Device id or non-nil error.
	AddDevice(device models.Device) (string, error)
	// Devices return all managed Devices from cache
	Devices() []models.Device
	// GetDeviceByName returns the Device by its name if it exists in the cache, or returns an error.
	GetDeviceByName(name string) (models.Device, error)
	// UpdateDevice updates the Device in the cache and ensures that the
	// copy in Core Metadata is also updated.
	UpdateDevice(device models.Device) error
	// RemoveDeviceByName removes the specified Device by name from the cache and ensures that the
	// instance in Core Metadata is also removed.
	RemoveDeviceByName(name string) error
	// AddDeviceProfile adds a new DeviceProfile to the Device Service and Core Metadata
	// Returns new DeviceProfile id or non-nil error.
	AddDeviceProfile(profile models.DeviceProfile) (string, error)
	// DeviceProfiles return all managed DeviceProfiles from cache
	DeviceProfiles() []models.DeviceProfile
	// GetProfileByName returns the Profile by its name if it exists in the cache, or returns an error.
	GetProfileByName(name string) (models.DeviceProfile, error)
	// UpdateDeviceProfile updates the DeviceProfile in the cache and ensures that the
	// copy in Core Metadata is also updated.
	UpdateDeviceProfile(profile models.DeviceProfile) error
	// RemoveDeviceProfileByName removes the specified DeviceProfile by name from the cache and ensures that the
	// instance in Core Metadata is also removed.
	RemoveDeviceProfileByName(name string) error
	// AddProvisionWatcher adds a new Watcher to the cache and Core Metadata
	// Returns new Watcher id or non-nil error.
	AddProvisionWatcher(watcher models.ProvisionWatcher) (string, error)
	// ProvisionWatchers return all managed Watchers from cache
	ProvisionWatchers() []models.ProvisionWatcher
	// GetProvisionWatcherByName returns the Watcher by its name if it exists in the cache, or returns an error.
	GetProvisionWatcherByName(name string) (models.ProvisionWatcher, error)
	// UpdateProvisionWatcher updates the Watcher in the cache and ensures that the
	// copy in Core Metadata is also updated.
	UpdateProvisionWatcher(watcher models.ProvisionWatcher) error
	// RemoveProvisionWatcher removes the specified Watcher by name from the cache and ensures that the
	// instance in Core Metadata is also removed.
	RemoveProvisionWatcher(name string) error
	// DeviceResource retrieves the specific DeviceResource instance from cache according to
	// the Device name and Device Resource name
	DeviceResource(deviceName string, deviceResource string) (models.DeviceResource, bool)
	// DeviceCommand retrieves the specific DeviceCommand instance from cache according to
	// the Device name and Command name
	DeviceCommand(deviceName string, commandName string) (models.DeviceCommand, bool)
	// AddDeviceAutoEvent adds a new AutoEvent to the Device with given name
	AddDeviceAutoEvent(deviceName string, event models.AutoEvent) error
	// RemoveDeviceAutoEvent removes an AutoEvent from the Device with given name
	RemoveDeviceAutoEvent(deviceName string, event models.AutoEvent) error
	// SetDeviceOpState sets the operating state of device
	SetDeviceOpState(name string, state models.OperatingState) error
	// UpdateDeviceOperatingState updates the Device's OperatingState with given name
	// in Core Metadata and device service cache.
	UpdateDeviceOperatingState(deviceName string, state string) error
}
