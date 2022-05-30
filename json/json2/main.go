package main

import (
	"encoding/json"
	"fmt"
)



type EventCfg struct {
	Hdr   string       `json:"hdr,omitempty"`   //event header
	Event json.RawMessage `json:"event,omitempty"` //event struct
}
type E struct {
	Header json.RawMessage `json:"header"`
	PfcpSessReportReq json.RawMessage `json:"pfcpSessReportReq"`
}
var str = `
{
  "hdr": "header test",
  "event": {
	"header": {
	  "Version": 1,
	  "MP": false,
	  "S": true,
	  "MessageType": 56,
	  "SEID": 1,
	  "MessagePriority": 0
	},
	"pfcpSessReportReq": {
	  "ReportType": {
		"Usar": true
	  },
	  "UsageReport":[
		{
		  "URRID": {
			"UrrIdValue": 1073741832
		  },
		  "URSEQN": {
			"UrseqnValue": 1
		  },
		  "UsageReportTrigger": {
			"Volqu": true
		  },
		  "MeasurementMethod": {
			"Volum": true
		  },
		  "VolumeQuota": {
			"Tovol": true,
			"TotalVolume": 2200
		  }
		}
	  ]
	}
  }
}
`
func main() {
	var testCfg EventCfg
	var e E
	if err := json.Unmarshal([]byte(str), &testCfg); err != nil {
		fmt.Println(err)
	}
	fmt.Println(testCfg)
	fmt.Println(testCfg.Hdr)
	fmt.Println(string(testCfg.Event))

	if err := json.Unmarshal(testCfg.Event,&e); err != nil {
		fmt.Println(err)
	}
	fmt.Println(e)
	fmt.Println(string(e.Header))
	fmt.Println(string(e.PfcpSessReportReq))

}
