package v1alpha2

// Object configuration
//
//	Purpose:
//	  Pass Object-related information to application code
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Object configuration is use config Object relate info, set by operator
type ObjectCfg struct {
	// object's name, e.g smf1
	// Optional.
	Name string `mapstructure:"name" json:"name,omitempty"`
	// object's nameSpace, e.g default
	// Optional.
	NameSpace string `mapstructure:"nameSpace" json:"nameSpace,omitempty"`
	// object's groupVersionResource
	// Optional.
	Gvr *GroupVersionResource `mapstructure:"groupVersionResource" json:"groupVersionResource,omitempty"`
}

// GroupVersionResource Configuration
//
//	Purpose:
//	  Specifies the GroupVersionResource of the object
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under ObjectCfg
type GroupVersionResource struct {
	// group is the API group of the defined custom resource.
	// Mandatory.
	Group string `mapstructure:"group" json:"group"`
	// version is the API version of the defined custom resource.
	// Mandatory.
	Version string `mapstructure:"version" json:"version"`
	// resource is the type of defined custom resource.
	// Mandatory.
	Resource string `mapstructure:"resource" json:"resource"`
}
