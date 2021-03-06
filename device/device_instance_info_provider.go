package device

import (
	"fmt"
	"github.com/galenliu/chip/config"
	"log"
	"sync"
	"time"
)

type DeviceInstanceInfoProvider interface {
	GetVendorName() (string, error)
	GetVendorId() (uint16, error)

	GetProductName() (string, error)
	GetProductId() (uint16, error)

	GetSerialNumber() (string, error)

	GetManufacturingDate() (time.Time, error)

	GetHardwareVersion() (uint16, error)
	GetHardwareVersionString() (string, error)

	GetRotatingDeviceIdUniqueId() ([]byte, error)
}

type DeviceInstanceInfoImpl struct {
	mConfigManager config.ConfigurationManager
}

func NewDeviceInstanceInfoImpl() *DeviceInstanceInfoImpl {
	return GetDeviceInstanceInfoProvider()
}

var _deviceInstanceInfo *DeviceInstanceInfoImpl
var _deviceInstanceInfoOnce sync.Once

func GetDeviceInstanceInfoProvider() *DeviceInstanceInfoImpl {
	_deviceInstanceInfoOnce.Do(func() {
		if _deviceInstanceInfo == nil {
			_deviceInstanceInfo = &DeviceInstanceInfoImpl{}
		}
	})
	return _deviceInstanceInfo
}

func (c *DeviceInstanceInfoImpl) Init(configMgr config.ConfigurationManager) (*DeviceInstanceInfoImpl, error) {
	c.mConfigManager = configMgr
	return c, nil
}

func NewDeviceInstanceInfo() *DeviceInstanceInfoImpl {
	return GetDeviceInstanceInfoProvider()
}

func (d DeviceInstanceInfoImpl) GetVendorId() (uint16, error) {
	return config.ChipDeviceConfigDeviceVendorId, nil
}

func (d DeviceInstanceInfoImpl) GetProductId() (uint16, error) {

	return config.ChipDeviceConfigDeviceProductId, nil

}

func (d DeviceInstanceInfoImpl) GetProductName() (string, error) {
	return config.ChipDeviceConfigDeviceProductName, nil
}

func (d DeviceInstanceInfoImpl) GetVendorName() (string, error) {
	return config.ChipDeviceConfigDeviceVendorName, nil
}

func (d *DeviceInstanceInfoImpl) GetSerialNumber() (string, error) {
	sn, err := d.mConfigManager.ReadConfigValueStr(config.KConfigKey_SerialNum)
	if sn == "" || err != nil {
		return config.ChipDeviceConfigTestSerialNumber, nil
	}
	return sn, nil
}

func (d *DeviceInstanceInfoImpl) GetManufacturingDate() (time.Time, error) {
	data, err := d.mConfigManager.ReadConfigValueStr(config.KConfigKey_ManufacturingDate)
	if err != nil {
		log.Panicf("invalid manufacturing date: %s", err.Error())
	}
	t, err := time.Parse("2006-Jan-02", data)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid manufacturing date: %s", err.Error())
	}
	return t, nil
}

func (d DeviceInstanceInfoImpl) GetHardwareVersion() (uint16, error) {
	version, err := d.mConfigManager.ReadConfigValueUint16(config.KConfigKey_HardwareVersion)
	if err != nil {
		return config.ChipDeviceConfigDefaultDeviceHardwareVersion, nil
	}
	return version, nil
}

func (d DeviceInstanceInfoImpl) GetHardwareVersionString() (string, error) {
	return config.ChipDeviceConfigDefaultDeviceHardwareVersionString, nil
}

func (d DeviceInstanceInfoImpl) GetRotatingDeviceIdUniqueId() ([]byte, error) {
	return config.ChipDeviceConfigRotatingDeviceIdUniqueId, nil
}
