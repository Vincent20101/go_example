package v1alpha2

import (
	"encoding/json"
	"strconv"
	"strings"
)

// UpfIpPool Configuration
//
//	Purpose:
//	  IpPoolInfo of upf
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under IpPoolInst
type UpfIpPool struct {
	// Upf instance id associated with ip pool
	// Mandatory
	UpfInstId string `mapstructure:"upfInstId" json:"upfInstId,omitempty"`
	// Ppe/upfinfo id associated with ip pool
	// Optional
	UpfInfoKey string `mapstructure:"upfInfokey" json:"upfInfokey,omitempty"`
	// Dnn associated with ip pool
	// Mandatory
	Dnn string `mapstructure:"dnn" json:"dnn,omitempty"`
	// Snssai associated with ip pool
	// Mandatory
	Snssai Snssai `mapstructure:"snssai" json:"snssai,omitempty"`
	// List of ranges of static IPv4 addresses handled by UPF.\n
	// Default value nil.\n
	// Optional.
	Ipv4AddressRanges []Ipv4AddressRange `mapstructure:"ipv4AddressRanges" json:"ipv4AddressRanges,omitempty"`
	// List of ranges of static IPv6 prefixes handled by the UPF.\n
	// Default value nil.\n
	// Optional.
	Ipv6PrefixRanges []Ipv6PrefixRange `mapstructure:"ipv6PrefixRanges" json:"ipv6PrefixRanges,omitempty"`
	// Indicates the current state of the pool.\n
	// If it is true, the pool is waiting to be updated and SMF should not assign ip addresses from this pool. \n
	// Default: false.\n
	// Optional.
	IsBusyOut bool `mapstructure:"isBusyOut" json:"isBusyOut,omitempty"`
	// Ip Pool size, Calculated when reading the ip pool configuration, used in back-end
	Size uint32 `mapstructure:"size" json:"size,omitempty"`
	// Ip Pool usage,Calculated when get pool infos,used in back-end
	Usage float64 `mapstructure:"usage" json:"usage,omitempty"`
	// 5G terminology for VRF, filled by operator, no need to config
	Vrf string `mapstructure:"vrf" json:"vrf,omitempty"`
	// upfIpPoolName, filled by operator, no need to config,
	// viper map key will be lowercase,if upf ip pool name contains uppercase letters,
	// they will be converted to lowercase. use UpfIpPoolName to save the upf ip pool name
	UpfIpPoolName string `mapstructure:"upfIpPoolName" json:"upfIpPoolName,omitempty"`
	// use to select upf, no need to config
	PoolName string `mapstructure:"poolName" json:"poolName,omitempty"`
}

// keyName: IPv4/IPv6-SmfIpPool-<upfInstId>-<poolName>-dnn-<sst-sd>
// or if upfInfokey!=”
// keyName: IPv4/IPv6-SmfIpPool-<upfInstId>-<upfInfoKey>-<poolName>-dnn-<sst-sd>
func (u *UpfIpPool) Key(poolName string, ipv4 bool) string {
	if poolName == "" {
		return u.keyWithoutPoolName(ipv4)
	}
	name := make([]string, 14, 14)
	conjunction := ""
	ppeId := ""
	if u.UpfInfoKey != "" {
		conjunction = "-"
		ppeId = "<" + u.UpfInfoKey + ">"
	}

	name[0] = "-<"
	name[1] = u.UpfInstId
	name[2] = ">-"
	name[3] = ppeId
	name[4] = conjunction
	name[5] = "<"
	name[6] = poolName
	name[7] = ">-"
	name[8] = u.Dnn
	name[9] = "-<"
	name[10] = strconv.Itoa(int(u.Snssai.Sst))
	name[11] = "-"
	name[12] = u.Snssai.Sd
	name[13] = ">"

	if ipv4 {
		return "IPv4-SmfIpPool" + strings.Join(name, "")
	} else {
		return "IPv6-SmfIpPool" + strings.Join(name, "")
	}
}

// keyName: IPv4/IPv6-<upfInstId>-dnn-<sst-sd>
// or if upfInfokey!=”
// keyName: IPv4/IPv6-<upfInstId>-<upfInfoKey>-dnn-<sst-sd>
func (u *UpfIpPool) keyWithoutPoolName(ipv4 bool) string {
	name := make([]string, 11, 11)
	conjunction := ""
	ppeId := ""
	if u.UpfInfoKey != "" {
		conjunction = "-"
		ppeId = "<" + u.UpfInfoKey + ">"
	}

	name[0] = "-<"
	name[1] = u.UpfInstId
	name[2] = ">-"
	name[3] = ppeId
	name[4] = conjunction
	name[5] = u.Dnn
	name[6] = "-<"
	name[7] = strconv.Itoa(int(u.Snssai.Sst))
	name[8] = "-"
	name[9] = u.Snssai.Sd
	name[10] = ">"

	if ipv4 {
		return "IPv4" + strings.Join(name, "")
	} else {
		return "IPv6" + strings.Join(name, "")
	}
}

// keyName: IPv4/IPv6-SmfIpPool-<upfInstId>-<poolName>-dnn-<sst-sd>
// or if upfInfokey!=”
// keyName: IPv4/IPv6-SmfIpPool-<upfInstId>-<upfInfoKey>-<poolName>-dnn-<sst-sd>
func (u *UpfIpPool) GetPoolInfoFromName(poolName string) {
	if !strings.Contains(poolName, "-SmfIpPool-") {
		u.getPoolInfoFromKeyWithoutPoolName(poolName)
		return
	}

	if len(poolName) < 16 {
		return
	}

	instanceIDIndex := strings.Index(poolName, ">-<")
	u.UpfInstId = poolName[16:instanceIDIndex]
	instanceIDIndex += 3 //+ len(">-<")

	keyOrPoolIndex := strings.Index(poolName[instanceIDIndex:], ">-")
	if keyOrPoolIndex == -1 {
		return
	}

	keyOrPoolIndex += instanceIDIndex
	v := poolName[instanceIDIndex:keyOrPoolIndex]
	keyOrPoolIndex += 2 //+ len(">-")
	nextIndex := strings.Index(poolName[keyOrPoolIndex:], ">-")
	if nextIndex == -1 {
		u.PoolName = v
		nextIndex = keyOrPoolIndex
	} else {
		nextIndex += keyOrPoolIndex
		u.UpfInfoKey = v
		u.PoolName = poolName[keyOrPoolIndex+1 : nextIndex]
		nextIndex += 2 //+ len(">-")
	}

	//dnn
	if len(poolName) < nextIndex {
		return
	}
	dnnIndex := strings.Index(poolName[nextIndex:], "-<")
	if dnnIndex == -1 {
		return
	}
	dnnIndex += nextIndex
	u.Dnn = poolName[nextIndex:dnnIndex]

	//slice
	dnnIndex += 2 // +len("-<")
	sliceIndex := strings.Index(poolName[dnnIndex:], ">")
	if sliceIndex == -1 {
		return
	}
	slice := poolName[dnnIndex : sliceIndex+dnnIndex]
	l := strings.Split(slice, "-")
	if len(l) == 2 {
		u.Snssai.Sd = l[1]
		sst, _ := strconv.Atoi(l[0])
		u.Snssai.Sst = int32(sst)
	}
}

// keyName: IPv4/IPv6-<upfInstId>-dnn-<sst-sd>
// or if upfInfokey!=”
// keyName: IPv4/IPv6-<upfInstId>-<upfInfoKey>-dnn-<sst-sd>
func (u *UpfIpPool) getPoolInfoFromKeyWithoutPoolName(key string) {
	if len(key) < 6 {
		return
	}

	instanceIDIndex := strings.Index(key, ">-")
	u.UpfInstId = key[6:instanceIDIndex]
	instanceIDIndex += 2 //+ len(">-")

	var nextIndex int
	keyIndex := strings.Index(key[instanceIDIndex:], ">-")
	if keyIndex != -1 {
		keyIndex += instanceIDIndex
		instanceIDIndex += 1 // + len("<")
		u.UpfInfoKey = key[instanceIDIndex:keyIndex]
		nextIndex = keyIndex + 2 //+ len(">-")
	} else {
		nextIndex = instanceIDIndex
	}

	//dnn
	if len(key) < nextIndex {
		return
	}
	dnnIndex := strings.Index(key[nextIndex:], "-<")
	if dnnIndex == -1 {
		return
	}
	dnnIndex += nextIndex
	u.Dnn = key[nextIndex:dnnIndex]

	//slice
	dnnIndex += 2 // +len("-<")
	sliceIndex := strings.Index(key[dnnIndex:], ">")
	if sliceIndex == -1 {
		return
	}
	slice := key[dnnIndex : sliceIndex+dnnIndex]
	l := strings.Split(slice, "-")
	if len(l) == 2 {
		u.Snssai.Sd = l[1]
		sst, _ := strconv.Atoi(l[0])
		u.Snssai.Sst = int32(sst)
	}
}

// IpPoolInst Configuration
//
//	Purpose:
//	  Configured IpPoolInfo of smf
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under SmfConfig and UpfConfig
type IpPoolInst struct {
	// Indicating it is IPv4 or IPv6 pool \n
	// Default: false \n
	// Optional
	IsIpv6 bool `mapstructure:"isIpv6" json:"isIpv6,omitempty"`
	// Indicating it is public or private IPv4 address,\n
	// “true” for IPv6 or public IPv4 for vzwStatic or Enterprise APNs.\n
	// Default: false.\n
	// Optional.
	IsPublic bool `mapstructure:"isPublic" json:"isPublic,omitempty"`
	// Indicating if SMF needs to do address allocation or not,\n
	// “false” for some enterprise/vzwStatic APNs \n
	// Default: false.\n
	// Deprecated
	// Optional.
	AllocateUeIpBySmf bool `mapstructure:"allocateUeIpBySmf" json:"allocateUeIpBySmf,omitempty"`
	// Indicating if SMF needs to select UPF based ip-pool name,\n
	// “true” only used for enterprise APN \n
	// Default: false.\n
	// Optional.
	ValidatePoolName bool `mapstructure:"validatePoolName" json:"validatePoolName,omitempty"`
	// IPv6 prefix length for IPv6 prefix delegation for enterprise ip pool \n
	// Range: 51 - 64 \n
	// Default value: 64 \n
	// Optional
	Ipv6PrefixLength uint8 `mapstructure:"ipv6PrefixLength" json:"ipv6PrefixLength,omitempty"`
	// List of Subnet \n
	// IPv4 Subnet Prefix Ranges: /13 ~ /32 \n
	// IPv6 Subnet Prefix Ranges: /45 ~ /64 \n
	// Example: 10.0.10.0/24; 10:10:10::/48 \n
	// Default value nil \n
	// Optional.
	Subnets []string `mapstructure:"subnets" json:"subnets,omitempty"`
	// Indicates the current state of the pool.\n
	// If it is true, the pool is waiting to be updated and SMF should not assign ip addresses from this pool. \n
	// Default: false.\n
	// Optional.
	IsBusyOut bool `mapstructure:"isBusyOut" json:"isBusyOut,omitempty"`
	//upf associated with ip pool, filled by operator, no need to config
	UpfIpPools map[string]*UpfIpPool `mapstructure:"upfIpPools" json:"upfIpPools,omitempty"`
	// Pool name, filled by operator, no need to config
	PoolName string `mapstructure:"poolName" json:"poolName,omitempty"`
	// ip address stickiness quarantine time
	// Unit:  second, default 0,range 0-31,536,000(One Year), 0 means no ip address stickiness.
	IpAddressStickinessQuarantineTime uint32 `mapstructure:"ipAddressStickinessQuarantineTime" json:"ipAddressStickinessQuarantineTime,omitempty"`
	// Indicating if ip pool is using hard stickiness.\n
	// Default: false.\n
	// Optional.
	HardStickiness bool `mapstructure:"hardStickiness" json:"hardStickiness,omitempty"`
}

// UpfIPv4PoolRef Configuration
//
//	Purpose:
//	  Configured IPv4 Pool Ref of upf
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under UpfIpPoolRef
type UpfIPv4PoolRef struct {
	// IP pool name \n
	// Mandatory
	Name string `mapstructure:"name" json:"name"`
	// deprecate, replaced by UpfIpPoolRef.Vrf \n
	// 5G terminology for VRF \n
	// Only used for IPv4 private pool for enterprise APN \n
	// Default: "" \n
	// Optional
	Vrf string `mapstructure:"vrf" json:"vrf,omitempty"`
	// CGNAT profile name \n
	// Only used for IPv4 private pool \n
	// SMF will pass CgnatProfile config in PFCP msg for a session \n
	// Default: "" \n
	// Optional
	CgnatProfile string `mapstructure:"cgnatProfile" json:"cgnatProfile,omitempty"`
	// List of ranges of static IPv4 addresses handled by UPF.\n
	// Default value nil.\n
	// Optional.
	Ipv4AddressRanges []Ipv4AddressRange `mapstructure:"ipv4AddressRanges" json:"ipv4AddressRanges,omitempty"`
}

// UpfIPv6PoolRef Configuration
//
//	Purpose:
//	  Configured IPv6 Pool Ref of upf
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under UpfIpPoolRef
type UpfIPv6PoolRef struct {
	// IP pool name \n
	// Mandatory
	Name string `mapstructure:"name" json:"name"`
	// IP pool name for IPv6 prefix delegation \n
	// Default: "" \n
	// Optional
	PrefixDelegationPool string `mapstructure:"prefixDelegationPool" json:"prefixDelegationPool,omitempty"`
	// List of ranges of static IPv6 prefixes handled by the UPF.\n
	// Default value nil.\n
	// Optional.
	Ipv6PrefixRanges []Ipv6PrefixRange `mapstructure:"ipv6PrefixRanges" json:"ipv6PrefixRanges,omitempty"`
}

// UpfIpPoolRef Configuration
//
//	Purpose:
//	  Configured IP Pool Ref of upf
//
//	Data model:
//	  Refer to the description for each attribute below
//
//	Usage:
//	  Configured under DnnUpfInfoItem
type UpfIpPoolRef struct {
	// 5G terminology for VRF \n
	// Default: "" \n
	// Optional
	Vrf string `mapstructure:"vrf" json:"vrf,omitempty"`
	// IPv4 Pool Ref \n
	// Default: nil \n
	// Legacy configuration, deprecated \n
	// Optional
	IPv4Pool *UpfIPv4PoolRef `mapstructure:"ipv4Pool" json:"ipv4Pool,omitempty"`
	// IPv6 Pool Ref \n
	// Default: nil \n
	// Legacy configuration, deprecated \n
	// Optional
	IPv6Pool *UpfIPv6PoolRef `mapstructure:"ipv6Pool" json:"ipv6Pool,omitempty"`
	// To support multiple IP Pools per DNN/APN,multiple Ipv4PoolItem, each Ipv4PoolItem defines a IPv4 pool \n
	// Default: nil \n
	// Optional
	Ipv4PoolList []*UpfIPv4PoolRef `mapstructure:"ipv4PoolList" json:"ipv4PoolList,omitempty"`
	// To support multiple IP Pools per DNN/APN,multiple Ipv6PoolItem, each Ipv6PoolItem defines a IPv6 pool \n
	// Default: nil \n
	// Optional
	Ipv6PoolList []*UpfIPv6PoolRef `mapstructure:"ipv6PoolList" json:"ipv6PoolList,omitempty"`
}

func NewIpPoolInst() *IpPoolInst {
	pool := &IpPoolInst{}
	pool.UpfIpPools = make(map[string]*UpfIpPool)

	return pool
}

// GetSmfIpPoolCfgFromUpfProfile :smf operator use this function to read smf-upf ip pool config
func GetSmfIpPoolCfgFromUpfProfile(upfProfile interface{}) (map[string]*IpPoolInst, error) {
	var upf UpfNfProfile

	data, err := json.Marshal(upfProfile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &upf)
	if err != nil {
		return nil, err
	}

	return getUpfIpPoolInfoFromUpfInfo(&upf), nil
}

func getUpfIpPoolInfoFromUpfInfo(upf *UpfNfProfile) map[string]*IpPoolInst {
	poolMap := map[string]*IpPoolInst{}
	if upf == nil {
		return poolMap
	}

	if len(upf.UpfInfo.SNssaiUpfInfoList) > 0 {
		getIpPoolInfoFromUpfInfo("", upf.NfProfile.NfInstanceId, &upf.UpfInfo, poolMap)
	}

	for k, v := range upf.UpfInfoList {
		getIpPoolInfoFromUpfInfo(k, upf.NfProfile.NfInstanceId, &v, poolMap)
	}

	return poolMap
}

func getIpPoolInfoFromUpfInfo(infoKey, upfInstanceId string, upfInfo *UpfInfo, poolMap map[string]*IpPoolInst) {
	var upfIpPoolInfo UpfIpPool

	upfIpPoolInfo.UpfInstId = upfInstanceId
	for _, v := range upfInfo.SNssaiUpfInfoList {
		for _, dnnInfo := range v.DnnUpfInfoList {
			//ip pool config in dnnInfo.IpPools
			upfIpPoolInfo.UpfInfoKey = infoKey
			upfIpPoolInfo.Snssai = v.SNssai
			upfIpPoolInfo.Dnn = dnnInfo.Dnn

			if dnnInfo.IpPools != nil {
				if dnnInfo.IpPools.Ipv4PoolList != nil { //If ipv4/6poollist is configured and the old ipv4/6pool is also configured, ignore the old configuration
					addIpv4PoolListToPoolMap(upfIpPoolInfo, &dnnInfo, poolMap)
				} else if dnnInfo.IpPools.IPv4Pool != nil && dnnInfo.IpPools.IPv4Pool.Name != "" {
					v4Pool := upfIpPoolInfo
					v4Pool.PoolName = dnnInfo.IpPools.IPv4Pool.Name
					v4Pool.Ipv4AddressRanges = dnnInfo.IpPools.IPv4Pool.Ipv4AddressRanges
					v4Pool.Vrf = dnnInfo.IpPools.Vrf
					addIpv4PoolToPoolMap(&v4Pool, poolMap)
				}

				if dnnInfo.IpPools.Ipv6PoolList != nil { //If ipv4/6poollist is configured and the old ipv4/6pool is also configured, ignore the old configuration
					addIpv6PoolListToPoolMap(upfIpPoolInfo, &dnnInfo, poolMap)
				} else if dnnInfo.IpPools.IPv6Pool != nil && dnnInfo.IpPools.IPv6Pool.Name != "" {
					v6Pool := upfIpPoolInfo
					v6Pool.PoolName = dnnInfo.IpPools.IPv6Pool.Name
					v6Pool.Ipv6PrefixRanges = dnnInfo.IpPools.IPv6Pool.Ipv6PrefixRanges
					v6Pool.Vrf = dnnInfo.IpPools.Vrf
					addIpv6PoolToPoolMap(&v6Pool, poolMap)
				}
			} else {
				//ip range configured under upf.iprange
				var smfPool *IpPoolInst
				var exist bool
				if smfPool, exist = poolMap[""]; !exist {
					smfPool = NewIpPoolInst()
					poolMap[""] = smfPool
				}

				if len(dnnInfo.Ipv4AddressRanges) > 0 {
					v4Pool := upfIpPoolInfo
					v4Pool.Ipv4AddressRanges = dnnInfo.Ipv4AddressRanges
					smfPool.UpfIpPools[v4Pool.Key("", true)] = &v4Pool
				}
				if len(dnnInfo.Ipv6PrefixRanges) > 0 {
					v6Pool := upfIpPoolInfo
					v6Pool.Ipv6PrefixRanges = dnnInfo.Ipv6PrefixRanges
					smfPool.UpfIpPools[v6Pool.Key("", false)] = &v6Pool
				}
			}
		}
	}

}

func addIpv4PoolListToPoolMap(upfIpPoolInfo UpfIpPool, dnnInfo *DnnUpfInfoItem, poolMap map[string]*IpPoolInst) {
	for _, upfIPv4Pool := range dnnInfo.IpPools.Ipv4PoolList {
		if upfIPv4Pool.Name != "" {
			v4Pool := upfIpPoolInfo
			v4Pool.PoolName = upfIPv4Pool.Name
			v4Pool.Ipv4AddressRanges = upfIPv4Pool.Ipv4AddressRanges
			v4Pool.Vrf = dnnInfo.IpPools.Vrf
			addIpv4PoolToPoolMap(&v4Pool, poolMap)
		}
	}
}

func addIpv6PoolListToPoolMap(upfIpPoolInfo UpfIpPool, dnnInfo *DnnUpfInfoItem, poolMap map[string]*IpPoolInst) {
	for _, upfIPv6Pool := range dnnInfo.IpPools.Ipv6PoolList {
		if upfIPv6Pool.Name != "" {
			v6Pool := upfIpPoolInfo
			v6Pool.PoolName = upfIPv6Pool.Name
			v6Pool.Ipv6PrefixRanges = upfIPv6Pool.Ipv6PrefixRanges
			v6Pool.Vrf = dnnInfo.IpPools.Vrf
			addIpv6PoolToPoolMap(&v6Pool, poolMap)
		}
	}
}

func addIpv4PoolToPoolMap(v4Pool *UpfIpPool, poolMap map[string]*IpPoolInst) {

	//ip pool range config under dnnItem.IpPools.IPv4Pool.Ipv4AddressRanges
	var smfPool *IpPoolInst
	var exist bool
	if smfPool, exist = poolMap[v4Pool.PoolName]; !exist {
		smfPool = NewIpPoolInst()
		poolMap[v4Pool.PoolName] = smfPool
	}
	smfPool.UpfIpPools[v4Pool.Key(v4Pool.PoolName, true)] = v4Pool
}

func addIpv6PoolToPoolMap(v6Pool *UpfIpPool, poolMap map[string]*IpPoolInst) {

	//ip pool range config under dnnItem.IpPools.IPv6Pool.Ipv6PrefixRanges
	var smfPool *IpPoolInst
	var exist bool
	if smfPool, exist = poolMap[v6Pool.PoolName]; !exist {
		smfPool = NewIpPoolInst()
		poolMap[v6Pool.PoolName] = smfPool
	}
	smfPool.UpfIpPools[v6Pool.Key(v6Pool.PoolName, false)] = v6Pool
}
