package pfcpie

// User Plane Path Failure Report IE . Refer to [TS29244-g40 7.4.5.1.2	User Plane Path Failure Report].
type UserPlanePathFailureReport struct {
	RemoteGTPUPeer []*RemoteGTPUPeer `tlv:"103" json:"remoteGTPUPeer,omitempty"`
}

// Marshal and Unmarshal functions not needed since all members of struct are TLV encoded.
