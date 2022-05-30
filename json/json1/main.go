package main

import (
	"encoding/json"
	"fmt"
)
type EventHdr struct {
	EvntType string `json:"evntType"` //unique event type for req/rsp event
	//sequence number indicating which step in the test case for request or response event
	//multiple events in testCfg could use same TestSeqNo, i.e. receive N1N2Transfer could have
	//3 events one for namf n1n2Transfer event, one for nas event, one for ngapEvnt
	//StepSeqNo value starts with 1, value 0 means it is not used in the step validation i.e. this event has
	//non-zero UnsolicateReqRspPairId
	StepSeqNo uint16 `json:"stepSeqNo"`
	//to handle unsolicited req/rsp which apples to i.e. nnrf_discovery
	//(where SMF discovery validity timeout, so SMF can send discovery to NF anything during a test procedure
	// or when handle multiple test cases one by one, SMF registration and NF discovery has been handled by first
	//test cases, in subsequent test cases, there is no need to enforce StepSeqNo
	//One Req/Rsp event forms a pair with the same UnsolicateReqRspPairId with value greater than 0
	UnsolicateReqRspPairId uint16        `json:"unsolicateReqRspPairId"`
	Direction              string `json:"direction,omitempty"` //direction of event, send or receive
	// deprecate
	ValidatorName      string `json:"validatorName,omitempty"` //event validator function name
	ValidateIEs        []string                  `json:"validateIEs,omitempty"`   //list of msg top-level IEs that will be validated
	SrcNfName          string                    `json:"srcNfName,omitempty"`     //NF instance name which this event is sent
	DstNfName          string                    `json:"dstNfName,omitempty"`     //NF instance name which this event is receive
	FteidNfName        string                    `json:"fteidNfName,omitempty"`   //NF instance name which this event is receive
	ReturnStatusCode   int                       `json:"returnStatusCode"`        //http return status code, i.e. http response status code 200
	HttpHeader         map[string]string         `json:"httpHeader,omitempty"`
	TimeoutMillisecond int                       `json:"timeoutMillisecond"`   //waiting time before handling this event, this can be used to simulator smf timeout handling, retransmission, or call flow to delay sending req/rsp in step
	LinkedTest         string                    `json:"linkedTest,omitempty"` //linked test which is defined in ../tests/common/ i.e sess_estb.json,
	//when linkedTest is provided, only SepSeqNo is required in EventHdr.
	LocRef string `json:"locRef,omitempty"` // refer to location values in the env param cfg
}
type TestCaseCfg struct {
	Events []EventCfg `json:"events"`
}

type EventCfg struct {
	//when LinkedTest is provided, Hdr and Event will not needed
	Hdr   EventHdr        `json:"hdr,omitempty"`   //event header
	Event json.RawMessage `json:"event,omitempty"` //event struct
}
var str = `
{
  "events": [
    {
      "hdr": {
        "evntType": "Pfcp_SessionReportReq",
        "stepSeqNo": 1,
        "direction": "send",
        "srcNfName": "upf-ytl",
        "dstNfName": "smf1"
      },
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
    },
    {
      "hdr": {
        "evntType": "Diam_GyCCR",
        "stepSeqNo": 2,
        "direction": "receive",
        "validateIEs": [
          "SessionId",
          "OriginHost",
          "OriginRealm",
          "DestinationRealm",
          "AuthApplicationId",
          "ServiceContextId",
          "CCRequestType",
          "DestinationHost",
          "UserName",
          "SubscriptionId",
          "MultipleServicesCreditControl",
          "UserEquipmentInfo"
        ],
        "srcNfName": "smf1",
        "dstNfName": "ocs1"
      },
      "event": {
        "Header": {
          "CommandCode":272,
          "CommandFlags":128,
          "Version":1,
          "ApplicationId":4,
          "HopByHopId":1,
          "EndToEndId":1
        },
        "CCR": {
          "sessionId": "lab.casa.com;383205776;1955181538;450051230000000-69;yesnet.mnc152.mcc502.gprs;1;http:__<HOSTPORT>_nchf-convergedcharging_v1_notify_imsi-450051230000000_69",
          "authApplicationId": 4,
          "originHost": "lab.casa.com",
          "originRealm": "casa.com",
          "destinationRealm": "ytlcomms.net",
          "destinationHost": "kl0sae-c01.ytlcomms.net",
          "serviceContextId": "12.32251@3gpp.org",
          "ccRequestType": 2,
          "ccRequestNumber": 1,
          "userName": "450051230000000@yesnet",
          "userEquipmentInfo": {
            "UserEquipmentInfoValue": "450051230000000F"
          },
          "subscriptionId": [
            {
              "SubscriptionIdType": 1,
              "SubscriptionIdData": "450051230000000"
            },
            {
              "SubscriptionIdType": 0,
              "SubscriptionIdData": "450051230000000"
            }
          ],
          "multipleServicesCreditControl": [
            {
              "RequestedServiceUnit": {
                "CCTime": 60,
                "CCTotalOctets": 5000,
                "CCOutputOctets": 1000
              },
              "RatingGroup": 111
            }
          ]
        }
      }
    },
    {
      "hdr": {
        "evntType": "Diam_GyCCA",
        "stepSeqNo": 3,
        "direction": "send",
        "srcNfName": "ocs1",
        "dstNfName": "smf1"
      },
      "event": {
        "header": {
          "Version": 1,
          "CommandFlags": 0,
          "CommandCode": 272,
          "ApplicationID": 4,
          "HopByHopID": 1,
          "EndToEndID": 1
        },
        "CCA": {
          "sessionId": "lab.casa.com;383205776;1955181538;450051230000000-69;yesnet.mnc152.mcc502.gprs;1;http:__<HOSTPORT>_nchf-convergedcharging_v1_notify_imsi-450051230000000_69",
          "resultCode": 4010,
          "authApplicationId": 4,
          "originRealm": "ytlcomms.net",
          "originHost": "kl0sae-c01.ytlcomms.net",
          "ccRequestType": 2,
          "ccRequestNumber": 1
        }
      }
    },
    {
      "hdr": {
        "evntType": "Pfcp_SessionReportRsp",
        "stepSeqNo": 1,
        "unsolicateReqRspPairId": 1,
        "direction": "receive",
        "validateIEs": ["Cause"],
        "dstNfName": "upf-ytl"
      },
      "event": {
        "header": {
          "Version": 1,
          "MP": false,
          "S": true,
          "MessageType": 57,
          "SEID": 197530862419748,
          "MessagePriority": 0
        },
        "pfcpSessReportRsp": {
          "Cause": {
            "CauseValue": 1
          }
        }
      }
    }
  ]
}
`
func main() {
	var testCfg TestCaseCfg
	if err := json.Unmarshal([]byte(str), &testCfg); err != nil {
		fmt.Println(err)
	}
	fmt.Println(testCfg)
	for _, v := range testCfg.Events {
		fmt.Printf("%#v\n", string(v.Event))
	}
}
