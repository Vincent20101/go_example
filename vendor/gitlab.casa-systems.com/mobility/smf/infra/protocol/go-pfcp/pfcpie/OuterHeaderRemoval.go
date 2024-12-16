package pfcpie

import (
	"fmt"
)

const (
	OuterHeaderRemovalGtpUUdpIpv4 uint8 = iota
	OuterHeaderRemovalGtpUUdpIpv6
	OuterHeaderRemovalUdpIpv4
	OuterHeaderRemovalUdpIpv6
	OuterHeaderRemovalIpv4
	OuterHeaderRemovalIpv6
	OuterHeaderRemovalGtpUUdpIp
	OuterHeaderRemovalVlanSTag
	OuterHeaderRemovalSTagCTag

	PduSessionContainerToBeDeleted uint8 = 1
)

type OuterHeaderRemoval struct {
	OuterHeaderRemovalDescription uint8 `json:"outerHeaderRemovalDescription,omitempty"`
	GtpUExtensionHeaderDeletion   bool  `json:"gtpUExtensionHeaderDeletion,omitempty"`
}

func (o *OuterHeaderRemoval) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(o.OuterHeaderRemovalDescription))
	// Octet 6
	if o.GtpUExtensionHeaderDeletion {
		data = append(data, byte(PduSessionContainerToBeDeleted))
	}
	return data, nil
}

func (o *OuterHeaderRemoval) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Unmarshal OuterHeaderRemoval octet5 failed: Inadequate TLV length: %d", length)
	}
	o.OuterHeaderRemovalDescription = uint8(data[idx])
	idx = idx + 1
	// Octet 6
	if length > idx {
		if data[idx] == PduSessionContainerToBeDeleted {
			o.GtpUExtensionHeaderDeletion = true
		} else {
			o.GtpUExtensionHeaderDeletion = false
		}
		idx = idx + 1
	}

	return nil
}
