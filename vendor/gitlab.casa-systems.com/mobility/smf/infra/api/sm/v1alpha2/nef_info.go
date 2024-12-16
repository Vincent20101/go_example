package v1alpha2

// Nef NfProfile
//
//	Purpose:
//	  Specially save the NfProfile information of nef.
//
//	Data model:
//	  Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.2 and 6.1.6.2.48 for details.
//
//	Usage:
//	  It is not currently used in the controller, only in nefmgr.
type NefNfProfile struct {
	// NfProfile information.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.2.\n
	// Mandatory.
	NfProfile NfProfile `mapstructure:"nfProfile" json:"nfProfile"`
	// NefInfo information.\n
	// Please refer https://www.3gpp.org/ftp/Specs/archive/29_series/29.510/29510-g30.zip section 6.1.6.2.48.\n
	// Mandatory.
	NefInfo NefInfo `mapstructure:"nefInfo" json:"nefInfo"`
}

type NefInfo struct {
	// Shall be present and contain the NEF ID of the NEF if NIDD service is supported.\n
	// Default value "".\n
	// Optional.
	NefId string `mapstructure:"nefId" json:"nefId,omitempty"`
	// PFD data.\n
	// PFD data. The NRF shall return the NEF profiles that have at least one nnef-pfdmanagement service.\n
	// matching the application identifiers and/or application function identifiers in the corresponding identifier list.\n
	// If not included, the NRF shall return all the application identifiers and/or application function identifiers registered in the NEF profile.\n
	// Default value nil.\n
	// Optional.
	PfdData *PfdData `mapstructure:"pfdData" json:"pfdData,omitempty"`
	// The AF provided event exposure data. The NEF registers such information in the NRF on behave of the AF.\n
	// Default value nil.\n
	// Optional.
	AfEeData *AfEventExposureData `mapstructure:"afEeData" json:"afEeData,omitempty"`
}

type PfdData struct {
	// List of internal application identifiers of the managed PFDs.\n
	// Default value nil.\n
	// Optional.
	AppIds []string `mapstructure:"appIds" json:"appIds,omitempty"`
	// List of application function identifiers of the managed PFDs.\n
	// Default value nil.\n
	// Optional.
	AfIds []string `mapstructure:"afIds" json:"afIds,omitempty"`
}

type AfEventExposureData struct {
	// AF Event(s) exposed by the NEF after registration of the AF(s) at the NEF.\n
	// Value can be "SVC_EXPERIENCE", "UE_MOBILITY" or "UE_COMM" or "EXCEPTIONS".\n
	// Default value "".\n
	// Mandatory.
	AfEvents []AfEvent `mapstructure:"afEvents" json:"afEvents"`
	// Associated AF identifications to the AfEvents. The absence of this attribute indicate that the NEF can be selected for any AF.\n
	// Default value nil.\n
	// Optional.
	AfIds []string `mapstructure:"afIds" json:"afIds,omitempty"`
	// The list of Application ID(s) the AF(s) connected to the NEF supports. The absence of this attribute indicate that the NEF can be selected for any Application.\n
	// Default value nil.\n
	// Optional.
	AppIds []string `mapstructure:"appIds" json:"appIds,omitempty"`
}

type AfEvent string

const (
	// Indicates that the event subscribed is service experience data for an application.
	AF_EVENT_SVC_EXPERIENCE AfEvent = "SVC_EXPERIENCE"
	// Indicates that the event subscribed is UE mobility information.
	AF_EVENT_UE_MOBILITY AfEvent = "UE_MOBILITY"
	// Indicates that the event subscribed is UE communication information.
	AF_EVENT_UE_COMM AfEvent = "UE_COMM"
	// Indicates that the event subscribed is exceptions information.
	AF_EVENT_EXCEPTIONS AfEvent = "EXCEPTIONS"
)
