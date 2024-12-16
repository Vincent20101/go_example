package v1alpha2

// Redundancy Sequence Number which differentiates the PDU Sessions that are handled redundantly
type Rsn string

const (
	Rsn_V1 Rsn = "V1"
	Rsn_V2 Rsn = "V2"
)

// Redundant transmission type which is applied for supporting URLLC
type RedTransType string

const (
	// Redundant transmission on N3/N9 interfaces
	RED_TRANS_TYPE_N3N9_INTERFACE RedTransType = "N3N9_INTERFACE"
	// Redundant transmission at transport layer
	RED_TRANS_TYPE_TRANSPORT_LAYER RedTransType = "TRANSPORT_LAYER"
)

// Redundant PDU session fail handling
type RedFailureHandling string

const (
	// Accept the establishment of a PDU session without redundancy handling
	RedFailureHandling_ACCEPT RedFailureHandling = "ACCEPT"
	// Reject the establishment of the PDU Session
	RedFailureHandling_REJECT RedFailureHandling = "REJECT"
)

type UrllcPolicy struct {
	// Local policy for URLLC service when SMF handles redundant PDU sessions
	// Optional
	RedSessPolicy *RedSessPolicy `mapstructure:"redSessPolicy" json:"redSessPolicy,omitempty"`
	// Preferred redundant transmission type when selected UPF support both methods for URLLC services.\n
	// -"N3N9_INTERFACE": support redundant transmission on N3/N9 interfaces.
	// -"TRANSPORT_LAYER": support redundant transmission at transport layer.
	// Default value "TRANSPORT_LAYER"
	// Optional.
	RedTransTypePreference RedTransType `mapstructure:"redTransTypePreference" json:"redTransTypePreference,omitempty"`
}

type RedSessPolicy struct {
	// RSN value for redundant PDU Session.
	// - "V1"
	// - "V2"
	// Default value ""
	// Mandatory
	Rsn Rsn `mapstructure:"rsn" json:"rsn"`
	// Local configure fail handle when SMF determines that redundant handling is not allowed or not possible for the given PDU Session
	// - "ACCEPT": accept the establishment of a PDU session without redundancy handling
	// - "REJECT": reject the establishment of the PDU Session
	// Default value "ACCEPT"
	// Optional.
	RedFailureHandling RedFailureHandling `mapstructure:"redFailureHandling" json:"redFailureHandling,omitempty"`
}
