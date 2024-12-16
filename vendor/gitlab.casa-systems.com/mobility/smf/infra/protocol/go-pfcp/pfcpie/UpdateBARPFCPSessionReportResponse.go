package pfcpie

type UpdateBARPFCPSessionReportResponse struct {
	BARID                           *BARID                           `tlv:"88" json:"bARID,omitempty"`
	DownlinkDataNotificationDelay   *DownlinkDataNotificationDelay   `tlv:"46" json:"downlinkDataNotificationDelay,omitempty"`
	DLBufferingDuration             *DLBufferingDuration             `tlv:"47" json:"dLBufferingDuration,omitempty"`
	DLBufferingSuggestedPacketCount *DLBufferingSuggestedPacketCount `tlv:"48" json:"dLBufferingSuggestedPacketCount,omitempty"`
	SuggestedBufferingPacketsCount  *SuggestedBufferingPacketsCount  `tlv:"140" json:"suggestedBufferingPacketsCount,omitempty"`
}
