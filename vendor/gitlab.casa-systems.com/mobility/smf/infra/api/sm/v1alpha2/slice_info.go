package v1alpha2

type SlicePolicyProfile struct {
	// Ref to Cc
	DefCcRef string `mapstructure:"defCcRef" json:"defCcRef"`
	// Ref to PccProfile
	DefStaticPccRef string `mapstructure:"defStaticPccRef" json:"defStaticPccRef"`
	// User Plane security polices applied to a PDU session at the establishment.\n
	// Default value nil.\n
	// optional
	UpSecurityCfg *UpSecurityPolicy `mapstructure:"upSecurityCfg" json:"upSecurityCfg,omitempty"`
	// Transport level marking, mapping from QoS Flow basis to DSCP value
	MessageMarkingMappingCfg []MessageMarkingMapping `mapstructure:"messageMarkingMappingCfg" json:"messageMarkingMappingCfg,omitempty"`
	// Mapping from NF application error to 5GSM cause values.\n
	// Default value nil.\n
	// optional
	CauseCodeMappingCfg *CauseCodeMappingProfile `mapstructure:"causeCodeMappingCfg" json:"causeCodeMappingCfg,omitempty"`
	// during congestion SMF is to configured to drop certain percentage
	// of new sessions. How priority session (i.e. MPS, MCS, mission critical service)
	// require priority treatment. this configuration configures percentage threshold for certian priority
	// session to be dropped if drop percentage is above threshold
	// i.e. priority 0 drop percentage threshold 99, which means priority 0 session can not be dropped in congestion unless drop percentage is 100%
	// i.e. priority 1 drop percentage threshold 90, which means priority 1 session ca be dropped when drop percentage above 90%
	// if priority level is not configured with drop threshold, which can be dropped during congestion.
	// key should be priority level allowed value: 0-31
	// value should be congestion drop percentage allowed value: 0-99. 100 is not allowed as it means dropped all new sessions
	PriorityDropThreshold map[uint8]uint8 `mapstructure:"priorityDropThreshold" json:"priorityDropThreshold,omitempty"`
	//  Local configure fail handle while pcf data request fail
	PcfClientPolicy *PcfClientPolicy `mapstructure:"pcfClientPolicy" json:"pcfClientPolicy,omitempty"`
	//Access Network Charging Identifier: Indicates the access network charging identifier for default QoS flow or whole PDU session.
	AccNetChId *AccNetChId `mapstructure:"accNetChId" json:"accNetChId"`
	// Ultra Reliable Low Latency Communication policy for supporting redundant transmission
	// Optional
	UrllcPolicy *UrllcPolicy `mapstructure:"urllcPolicy" json:"urllcPolicy,omitempty"`
	// TODO: add other slice policies according to frontend Policy resource yaml
}
