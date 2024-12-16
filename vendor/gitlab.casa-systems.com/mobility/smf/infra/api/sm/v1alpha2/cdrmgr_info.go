package v1alpha2

// CdrMgr Instance configuration
//
//	Purpose:
//	  This is default service configuration to be used
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Used as value of CdrMgrSvcCfg
type CdrMgrInstanceInfo struct {
	// The controller will sync the cdrMgrSvcInstanceId \n
	// Optional
	CdrMgrSvcInstanceId string `mapstructure:"cdrMgrSvcInstanceId" json:"cdrMgrSvcInstanceId,omitempty"`
	// The controller will sync the cdrMgrSvcFqdn \n
	// Optional
	CdrMgrSvcFqdn string `mapstructure:"cdrMgrSvcFqdn" json:"cdrMgrSvcFqdn,omitempty"`
	// The service port for the cdrMgr \n
	// Default value 8088 \n
	// Optional
	CdrMgrSvcPort uint16 `mapstructure:"cdrMgrSvcPort" json:"cdrMgrSvcPort,omitempty"`
	// The port for GRPC services \n
	// Default value 50061 \n
	// Optional
	CdrMgrGrpcPort uint16 `mapstructure:"cdrMgrGrpcSvcPort" json:"cdrMgrGrpcSvcPort,omitempty"`
}
