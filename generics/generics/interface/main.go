package main

var (
	grpcPort     = uint16(50051) // In 8.M8.0 is 9051
	gtpcPort     = uint16(50052) // In 8.M8.0 is 9052
	grpcPgwPort  = uint16(50053) // In 8.M8.0 is 9053
	grpc2Port    = uint16(50054) // In 8.M8.0 is 9054
	grpc3Port    = uint16(50055) // In 8.M8.0 is 9055
	grpc4Port    = uint16(50056) // SMF grpc port for gtpc, in 8.M8.0 is 9056
	grpcMgmtPort = uint16(50057) // In 8.M8.0 is 9057
	// 50058 is use for tracing, in 8.M8.0 is 9058
	grpcAAAPort        = uint16(50059) // In 8.M8.0 is 9059
	grpcPcmsPort       = uint16(50060) // aaaClient <--> pcms, in 8.M8.0 is 9060
	grpcCdrmgrPort     = uint16(50061) // smfsm <--> cdrmgr, in 8.M8.0 is 9061
	grpcAAAcliPort     = uint16(50062) // In 8.M8.0 is 9062
	grpcCdrMgrMgmtPort = uint16(9071)  //apiserver <--> cdrmgr

	dbPort   = uint16(6379)
	svcPort  = uint16(80)
	sftpPort = uint(22)

	// service port used by SMF containers
	smfSvcPort       = uint16(8088)
	smfCfgPort       = uint16(8080)
	smfCdrmgrCdfPort = uint16(3386)
	gtpcv2Port       = uint16(2123)
	pfcpPort         = uint16(8805)
	pgwcPort         = uint16(35805)
	gtpuPort         = uint16(2152)
	pfcpUdpPort      = uint16(38805)
	sftpTcpPort      = uint16(2022)
	timerClientPort  = uint16(50063) // timersvc --> smfsm, in 8.M8.0 is 9063

	httpsPort = uint16(443)
	liPort    = uint16(8443)
	x1Port    = uint16(8443) // limgr as a server for x1
	t3Port    = uint16(8444) // limgr as a server for t3
	dhcpv4    = uint16(67)
	dhcpv6    = uint16(547)
	// smfDhcpv4 and smfDhcpv6 port used by SMF containers
	smfDhcpv4Port  = uint16(8068)
	smfDhcpv6Port  = uint16(8546)
	rediusAuthPort = uint16(1812)
	radiusCoaPort  = uint16(3799)
	metricPath     = "/metrics"

	recovery                       = uint8(0)
	recoveryTimeStampForPod        = int64(0)
	SEC_FROM_1900_TO_1970          = int64(2208988800)
	cdrVolumeName                  = "cdr"
	cdrVolumeMountPath             = "/var/cdr/"
	BgpHttpPort             uint16 = 6543
	PcapHttpPort            uint16 = 7778
	RmSvcExternalPort       uint16 = 8443
)

type portNumericType interface {
	~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~int | ~uint
}

// SMF Service gRPC port should not use 50051+ which located in default
// ephemeral ports range [32768,60999]. They will be occupied by
// others TCP connections like redis client.
func grpcPortConverter[T portNumericType](srcPort uint16, newPortEnabled bool) T {
	if !newPortEnabled || srcPort < grpcPort || srcPort > timerClientPort {
		return T(srcPort)
	}
	// Rough calculation. 50051-50063 to 9051-9063.
	return T(srcPort - 41000)
}

func main() {

}
