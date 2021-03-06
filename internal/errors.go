package internal

import "fmt"

var (
	ChipErrorInvalidArgument      = fmt.Errorf("CHIP_ERROR_INVALID_ARGUMENT")
	ChipErrorIncorrectState       = fmt.Errorf("CHIP_ERROR_INCORRECT_STATE")
	ChipErrorNotImplemented       = fmt.Errorf("CHIP_ERROR_NOT_IMPLEMENTED")
	ChipErrorInternal             = fmt.Errorf("CHIP_ERROR_INTERNAL")
	ChipDeviceErrorConfigNotFound = fmt.Errorf("CHIP_DEVICE_ERROR_CONFIG_NOT_FOUND")
)
