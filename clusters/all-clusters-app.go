package clusters

const (
	BootreasontypeKunspecified = iota
	BootReasonType_kPowerOnReboot
	BootReasonType_kBrownOutReset
	BootReasonType_kSoftwareWatchdogReset
	BootReasonType_kHardwareWatchdogReset
	BootReasonType_kSoftwareUpdateCompleted
	BootReasonType_kSoftwareReset

	RegulatorylocationtypeKindoor = iota
	RegulatoryLocationType_kOutdoor
	RegulatorylocationtypeKindooroutdoor
)
