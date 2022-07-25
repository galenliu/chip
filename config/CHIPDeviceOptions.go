package config

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net"
	"os"
	"path"
	"sync"
)

func getDefaultKVS() string {
	dir, _ := os.UserHomeDir()
	return path.Join(dir, "chip.ini")
}

type ConfigFlag struct {
	Key          string
	DefaultValue any
	Usage        string
}

var (
	DeviceOptionVersion = ConfigFlag{
		Key:          "version",
		DefaultValue: 0,
		Usage:        "The version indication provides versioning of the setup payload.\n",
	}

	DeviceOptionVendorID = ConfigFlag{
		"vendor-id",
		0,
		"The Vendor ID is assigned by the Connectivity Standards Alliance.\n",
	}

	DeviceOptionProductID = ConfigFlag{
		"product-id",
		0,
		"The Product ID is specified by vendor.\n",
	}

	DeviceOptionCustomFlow = ConfigFlag{
		"custom-flow",
		0,
		"A 2-bit unsigned enumeration specifying manufacturer-specific custom flow options.\n",
	}

	DeviceOptionCapabilities = ConfigFlag{
		"capabilities",
		0,
		"Discovery Capabilities Bitmask which contains information about Device’s available technologies for device discovery.\n",
	}

	DeviceOptionDiscriminator = ConfigFlag{
		"discriminator",
		0,
		"A 12-bit unsigned integer match the value which a device advertises during commissioning.\n",
	}

	DeviceOptionPasscode = ConfigFlag{
		"passcode",
		0xFFFFFFF,
		"A 27-bit unsigned integer, which serves as proof of possession during commissioning. If not provided to compute a verifier, the --spake2p-verifier-base64 must be provided. \n",
	}

	DeviceOptionSpake2pVerifierBase64 = ConfigFlag{
		"spake2p-verifier-base64",
		0xFFFFF,
		"A raw concatenation of 'W0' and 'L' (67 bytes) as base64 to override the verifier auto-computed from the passcode, if provided.\n",
	}

	DeviceOptionSpake2pSaltBase64 = ConfigFlag{
		"spake2p-salt-base64",
		0,
		"16-32 bytes of salt to use for the PASE verifier, as base64. If omitted, will be generated randomly. If a --spake2p-verifier-base64 is passed, it must match against the salt otherwise failure will arise.\n",
	}

	DeviceOptionSpake2pIterations = ConfigFlag{
		"spake2p-iterations",
		0,
		"Number of PB DF iterations to use. If omitted, will be 1000. If a --spake2p-verifier-base64 is passed, the iteration counts must match that used to generate the verifier otherwise failure will arise.\n",
	}

	DeviceOptionSecuredDevicePort = ConfigFlag{
		"secured-device-port",
		5540,
		"A 16-bit unsigned integer specifying the listen port to use for secure device messages (default is 5540).\n",
	}

	DeviceOptionSecuredCommissionerPort = ConfigFlag{
		"secured-commissioner-port",
		5542,
		"A 16-bit unsigned integer specifying the listen port to use for secure commissioner messages (default is 5542). Only valid when app is both device and commissioner.\n",
	}

	DeviceOptionUnsecuredCommissionerPort = ConfigFlag{
		"unsecured-commissioner-port",
		5550,
		"A 16-bit unsigned integer specifying the port to use for unsecured commissioner messages (default is 5550).\n",
	}

	DeviceOptionCommand = ConfigFlag{
		"command",
		"command",
		"A name for a command to execute during startup.\n"}

	DeviceOptionPics = ConfigFlag{
		"PICS",
		"",
		"A file containing PICS items.\n"}

	DeviceOptionKvs = ConfigFlag{
		"KVS",
		getDefaultKVS(),
		"A file to store Key Value Store items.\n"}

	DeviceOptionInterfaceId = ConfigFlag{
		"interface-id",
		"interface-id",
		"A interface id to advertise on.\n"}
)

type DeviceOptions struct {
	Spake2pIterations         uint32
	Spake2pVerifier           []byte
	Spake2pSalt               []byte
	Discriminator             uint16
	Payload                   PayloadContents
	BleDevice                 uint32
	WiFi                      bool
	Thread                    bool
	SecuredDevicePort         uint16
	SecuredCommissionerPort   uint16
	UnsecuredCommissionerPort uint16
	Command                   string
	PICS                      string
	KVS                       string
	InterfaceId               net.Interface
	TraceStreamDecodeEnabled  bool
	TraceStreamToLogEnabled   bool
	TraceStreamFilename       string
	TestEventTriggerEnableKey []byte
}

var _instance *DeviceOptions
var _once sync.Once

func GetDeviceOptionsInstance() *DeviceOptions {
	_once.Do(func() {
		if _instance == nil {
			_instance = &DeviceOptions{}
		}
	})
	return _instance
}

func SetDeviceOptions(c *cobra.Command) {
	c.Flags().Uint8(DeviceOptionVersion.Key, cast.ToUint8(DeviceOptionVersion.DefaultValue), DeviceOptionVersion.Usage)
	c.Flags().Uint64(DeviceOptionVendorID.Key, cast.ToUint64(DeviceOptionVendorID.DefaultValue), DeviceOptionVendorID.Usage)
	c.Flags().Uint64(DeviceOptionProductID.Key, cast.ToUint64(DeviceOptionProductID.DefaultValue), DeviceOptionProductID.Usage)
	c.Flags().Uint8(DeviceOptionCustomFlow.Key, cast.ToUint8(DeviceOptionCustomFlow.DefaultValue), DeviceOptionCustomFlow.Usage)
	c.Flags().Uint8(DeviceOptionCapabilities.Key, cast.ToUint8(DeviceOptionCapabilities.DefaultValue), DeviceOptionCapabilities.Usage)
	c.Flags().Uint16(DeviceOptionDiscriminator.Key, cast.ToUint16(DeviceOptionDiscriminator.DefaultValue), DeviceOptionDiscriminator.Usage)
	c.Flags().Uint32(DeviceOptionPasscode.Key, cast.ToUint32(DeviceOptionPasscode.DefaultValue), DeviceOptionPasscode.Usage)
	c.Flags().Uint32(DeviceOptionSpake2pVerifierBase64.Key, cast.ToUint32(DeviceOptionSpake2pVerifierBase64.DefaultValue), DeviceOptionSpake2pVerifierBase64.Usage)
	c.Flags().Uint32(DeviceOptionSpake2pSaltBase64.Key, cast.ToUint32(DeviceOptionSpake2pSaltBase64.DefaultValue), DeviceOptionSpake2pSaltBase64.Usage)
	c.Flags().Uint64(DeviceOptionSpake2pIterations.Key, cast.ToUint64(DeviceOptionSpake2pIterations.DefaultValue), DeviceOptionSpake2pIterations.Usage)
	c.Flags().Uint16(DeviceOptionSecuredDevicePort.Key, cast.ToUint16(DeviceOptionSecuredDevicePort.DefaultValue), DeviceOptionSecuredDevicePort.Usage)
	c.Flags().Uint16(DeviceOptionSecuredCommissionerPort.Key, cast.ToUint16(DeviceOptionSecuredCommissionerPort.DefaultValue), DeviceOptionSecuredCommissionerPort.Usage)
	c.Flags().Uint16(DeviceOptionUnsecuredCommissionerPort.Key, cast.ToUint16(DeviceOptionUnsecuredCommissionerPort.DefaultValue), DeviceOptionUnsecuredCommissionerPort.Usage)
	c.Flags().String(DeviceOptionCommand.Key, cast.ToString(DeviceOptionCommand.DefaultValue), DeviceOptionCommand.Usage)
	c.Flags().String(DeviceOptionPics.Key, cast.ToString(DeviceOptionPics.DefaultValue), DeviceOptionPics.Usage)
	c.Flags().String(DeviceOptionKvs.Key, cast.ToString(DeviceOptionKvs.DefaultValue), DeviceOptionKvs.Usage)
	c.Flags().String(DeviceOptionInterfaceId.Key, cast.ToString(DeviceOptionInterfaceId.DefaultValue), DeviceOptionInterfaceId.Usage)
}

func GetDeviceOptions(config *viper.Viper) *DeviceOptions {

	GetDeviceOptionsInstance().Payload.Version = uint8(config.GetUint(DeviceOptionVersion.Key))
	GetDeviceOptionsInstance().Payload.VendorID = uint16(config.GetUint32(DeviceOptionVendorID.Key))
	GetDeviceOptionsInstance().Payload.Discriminator = uint16(config.GetUint32(DeviceOptionDiscriminator.Key))

	GetDeviceOptionsInstance().SecuredDevicePort = uint16(config.GetUint32(DeviceOptionSecuredDevicePort.Key))
	GetDeviceOptionsInstance().SecuredCommissionerPort = uint16(config.GetUint32(DeviceOptionSecuredCommissionerPort.Key))
	GetDeviceOptionsInstance().UnsecuredCommissionerPort = uint16(config.GetUint32(DeviceOptionUnsecuredCommissionerPort.Key))
	GetDeviceOptionsInstance().Command = config.GetString(DeviceOptionCommand.Key)
	GetDeviceOptionsInstance().PICS = config.GetString(DeviceOptionPics.Key)
	GetDeviceOptionsInstance().KVS = config.GetString(DeviceOptionKvs.Key)
	GetDeviceOptionsInstance().InterfaceId = net.Interface{}
	GetDeviceOptionsInstance().TraceStreamDecodeEnabled = false
	GetDeviceOptionsInstance().TraceStreamToLogEnabled = false

	return GetDeviceOptionsInstance()
}
