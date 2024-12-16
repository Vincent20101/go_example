package v1alpha2

// LoadAndOverloadControl Configuration
//
//	Purpose:
//	  Configure SMF load and overload control information.
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfConfig
type LoadAndOverloadControl struct {
	//current load used by smfsm to send LCI
	Load uint8 `mapstructure:"load" json:"load,omitempty"`
	//operator tells smfsm it is overload now or not
	Enable bool `mapstructure:"enable" json:"enable,omitempty"`
	//hard code by controller as 30 seconds for now, when
	//GCS-12580 support configurable interval for cpu/memory metric for overload control is implemented
	//take the configured value for interval which operator calls metric API.
	PeriodOfValidity uint32 `mapstructure:"periodOfValidity" json:"periodOfValidity,omitempty"`
	//current overload level used by smfsm to trigger overload alarms
	OverloadLevel string `mapstructure:"overloadLevel" json:"overloadLevel,omitempty"`
	// What percentage of the sessions' attaching request will be rejected by SMF.
	// Default value: 0 \n
	// Unit: percentage(%)
	SessRejectPercentage uint32 `mapstructure:"sessRejectPercentage" json:"sessRejectPercentage"`
	// Control whether log sbi error responses
	// Default value: false
	DisableSbiErrRspLog bool `mapstructure:"disableSbiErrRspLog" json:"disableSbiErrRspLog,omitempty"`
}
