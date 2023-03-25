package main

import "fmt"

func main() {
	fmt.Println((0 ^ UPFMGR_HTTP) & 0)
	fmt.Println((70 ^ 15) & 15)
	fmt.Println(70 ^ 15&15)
	fmt.Printf("%08b\n", 15)
	fmt.Printf("%08b\n", 6)
	fmt.Printf("%08b\n", 6^15)
	fmt.Printf("%08b\n", 9&15)
}

// HealthType The health value of each interface
type HealthType uint64

const (
	DB HealthType = 1 << iota
	//aaaclient server
	AAA_GRPC
	//cdrmgr server
	CDRMGR_GRPC
	//gtpc server
	GTPC_GRPC
	GTPC_UDP
	//pfcp server
	PFCP_GRPC_SMFSM
	PFCP_GRPC_UPFMGR
	PFCP_UDP
	PFCP_GTPU
	//limgr server
	LI_GRPC
	LI_X1_HTTP
	LI_X2_T3_HTTP
	//pcms server
	PCMS_GRPC
	PCMS_HTTP
	//smfsm server
	SMFSM_GRPC_PFCP
	SMFSM_GRPC_UPFMGR
	SMFSM_HTTP
	SMFSM_GRPC_TRACE
	SMFSM_GRPC_GTPC
	SMFSM_GRPC_AAA
	SMFSM_GRPC_LI
	//upfmgr server
	UPFMGR_GRPC_PFCP
	UPFMGR_HTTP
)
