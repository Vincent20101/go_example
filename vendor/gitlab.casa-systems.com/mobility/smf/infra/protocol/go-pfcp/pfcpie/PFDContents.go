package pfcpie

import (
	"fmt"
)

// Refer to [TS29244 8.2.39 PFD Contents]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 61 (decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |  ADNP |  AURL |  AFD  |  DNP  |  CP   |   DN  |  URL  |  FD    |
//	         |----------------------------------------------------------------|
//	  6      |                         Spare                                  |
//	         |----------------------------------------------------------------|
//
// m to (m+1)|                 Length of Flow Description                     |
//
//	|----------------------------------------------------------------|
//
// (m+2) to p|                  Flow Description                              |
//
//	|----------------------------------------------------------------|
//
// q to (q+1)|                       Length of URL                            |
//
//	|----------------------------------------------------------------|
//
// (q+2) to r|                            URL                                 |
//
//	|----------------------------------------------------------------|
//
// s to (s+1)|                    Length of Domain Name                       |
//
//	|----------------------------------------------------------------|
//
// (s+2) to t|                         Domain Name                            |
//
//	|----------------------------------------------------------------|
//
// u to (u+1)|                  Length of Custom PFD Content                  |
//
//	|----------------------------------------------------------------|
//
// (u+2) to v|                        Custom PFD Content                      |
//
//	|----------------------------------------------------------------|
//
// w to (w+1)|                   Length of Domain Name Protocol               |
//
//	|----------------------------------------------------------------|
//
// (w+2) to x|                       Domain Name Protocol                     |
//
//	|----------------------------------------------------------------|
//
// y to (y+1)|                Length of Additional Flow Description           |
//
//	|----------------------------------------------------------------|
//
// (y+2) to z |                     Additional Flow Description                |
//
//	|----------------------------------------------------------------|
//
// a to (a+1) |                     Length of Additional URL                   |
//
//	|----------------------------------------------------------------|
//
// (a+2) to b |                        Additional URL                          |
//
//	|----------------------------------------------------------------|
//
// c to (c+1) |    Length of Additional Domain Name and Domain Name Protocol   |
//
//	|----------------------------------------------------------------|
//
// (c+2) to d |        Additional Domain Name and Domain Name Protocol         |
//
//	|----------------------------------------------------------------|
//
// e to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5 in the Figure 8.2.39-1:
//
//   - Bit 1 – FD (Flow Description):
//     If this bit is set to "1", then the Length of Flow Description and the Flow
//     Description fields shall be present, otherwise they shall not be present.
//
//   - Bit 2 – URL (URL):
//     If this bit is set to "1", then the Length of URL and the URL fields shall
//     be present, otherwise they shall not be present.
//
//   - Bit 3 – DN (Domain Name):
//     If this bit is set to "1", then the Length of Domain Name and the Domain
//     Name fields shall be present, otherwise they shall not be present.
//
//   - Bit 4 – CP (Custom PFD Content):
//     If this bit is set to "1", then the Length of Custom PFD Content and the
//     Custom PFD Content fields shall be present, otherwise they shall not be
//     present.
//
//   - Bit 5 – DNP (Domain Name Protocol):
//     If this bit is set to "1", then the Length of Domain Name Protocol and the
//     Domain Name Protocol shall be present, otherwise they shall not be present;
//     and if this bit is set to "1", the Length of Domain Name and the Domain
//     Name fields shall also be present.
//
//   - Bit 6 – AFD (Additional Flow Description):
//     If this bit is set to "1", the Length of Additional Flow Description and
//     the Additional Flow Description field shall be present, otherwise they
//     shall not be present.
//
//   - Bit 7 – AURL (Additional URL):
//     If this bit is set to "1", the Length of Additional URL and the Additional
//     URL field shall be present, otherwise they shall not be present.
//
//   - Bit 8 – ADNP (Additional Domain Name and Domain Name Protocol):
//     If this bit is set to "1", the Length of Additional Domain Name and Domain
//     Name Protocol, and the Additional Domain Name and Domain Name Protocol
//     field shall be present, otherwise they shall not be present.
//
//     TOD0: The encoding/decoding ways and the format of CustompFDContent field
//     was not found in protocols from [TS29244-gue] to [TS29244-g60]. we assume
//     the format of this field is the same as URL,DomainName，which is a octet
//     string.
type PFDContents struct {
	Adnp bool `json:"adnp,omitempty"`
	Aurl bool `json:"aurl,omitempty"`
	Afd  bool `json:"afd,omitempty"`
	Dnp  bool `json:"dnp,omitempty"`
	Cp   bool `json:"cp,omitempty"`
	Dn   bool `json:"dn,omitempty"`
	Url  bool `json:"url,omitempty"`
	Fd   bool `json:"fd,omitempty"`

	FlowDescription                           string                                      `json:"flowDescription,omitempty"`
	URL                                       string                                      `json:"URL,omitempty"`
	DomainName                                string                                      `json:"domainName,omitempty"`
	CustomPFDContents                         CustomPFDContent                            `json:"customPFDContents,omitempty"`
	DomainNameProtocol                        string                                      `json:"domainNameProtocol,omitempty"`
	AdditionalFlowDescription                 []string                                    `json:"additionalFlowDescription,omitempty"`
	AdditionalURL                             []string                                    `json:"additionalURL,omitempty"`
	AdditionalDomainNameAndDomainNameProtocol []AdditionalDomainNameAndDomainNameProtocol `json:"additionalDomainNameAndDomainNameProtocol,omitempty"`
}

type AdditionalDomainNameAndDomainNameProtocol struct {
	DomainName         string `json:"domainName,omitempty"`
	DomainNameProtocol string `json:"domainNameProtocol,omitempty"`
}

func (p *PFDContents) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(p.Adnp)<<7 |
		btou(p.Aurl)<<6 |
		btou(p.Afd)<<5 |
		btou(p.Dnp)<<4 |
		btou(p.Cp)<<3 |
		btou(p.Dn)<<2 |
		btou(p.Url)<<1 |
		btou(p.Fd)
	if tmpUint8&0xff == 0 {
		return []byte(""), fmt.Errorf("at least one of FD, URL, DN, CP, DNP, AFD, AURL, ADNP shall be set")
	}
	data = append([]byte(""), byte(tmpUint8))

	// Octet 6 (spare)
	data = append(data, byte(0))

	// FlowDescription
	if p.Fd {
		data = AppendUint16(data, uint16(len(p.FlowDescription)))
		data = append(data, []byte(p.FlowDescription)...)
	}

	// URL
	if p.Url {
		data = AppendUint16(data, uint16(len(p.URL)))
		data = append(data, []byte(p.URL)...)
	}

	// Domain Name
	if p.Dn {
		data = AppendUint16(data, uint16(len(p.DomainName)))
		data = append(data, []byte(p.DomainName)...)
	}

	// Custom PFD Contents
	if p.Cp {
		buf, err := p.CustomPFDContents.MarshalBinary()
		if err != nil {
			return nil, err
		}
		data = AppendUint16(data, uint16(len(buf)))
		data = append(data, buf...)
	}

	// Domain Name Protocols
	if p.Dnp {
		if !p.Dn {
			return []byte(""), fmt.Errorf("when Adnp is set,the Domain Name fields shall also be present")
		}
		data = AppendUint16(data, uint16(len(p.DomainNameProtocol)))
		data = append(data, []byte(p.DomainNameProtocol)...)
	}

	// Additional Flow Description
	if p.Afd {
		t := []byte("")
		for _, v := range p.AdditionalFlowDescription {
			t = AppendUint16(t, uint16(len(v)))
			t = append(t, []byte(v)...)
		}
		data = AppendUint16(data, uint16(len(t)))
		data = append(data, t...)

	}

	// Aditional URL
	if p.Aurl {
		t := []byte("")
		for _, v := range p.AdditionalURL {
			t = AppendUint16(t, uint16(len(v)))
			t = append(t, []byte(v)...)
		}
		data = AppendUint16(data, uint16(len(t)))
		data = append(data, t...)
	}

	// Additional Domain Name and Domain Name Protocols
	if p.Adnp {
		t := []byte("")
		for _, v := range p.AdditionalDomainNameAndDomainNameProtocol {
			t = AppendUint16(t, uint16(len(v.DomainName)))
			t = append(t, []byte(v.DomainName)...)

			t = AppendUint16(t, uint16(len(v.DomainNameProtocol)))
			t = append(t, []byte(v.DomainNameProtocol)...)
		}
		data = AppendUint16(data, uint16(len(t)))
		data = append(data, t...)
	}
	return

}

func (p *PFDContents) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)
	octet5, e := byteReader.ReadUint8()
	if e != nil {
		return fmt.Errorf("inadequate TLV length for octet5: %d", e)
	}

	// Octet 5
	p.Adnp = utob(octet5 & BitMask8)
	p.Aurl = utob(octet5 & BitMask7)
	p.Afd = utob(octet5 & BitMask6)
	p.Dnp = utob(octet5 & BitMask5)
	p.Cp = utob(octet5 & BitMask4)
	p.Dn = utob(octet5 & BitMask3)
	p.Url = utob(octet5 & BitMask2)
	p.Fd = utob(octet5 & BitMask1)

	// Octet 6 (spare)
	_, _ = byteReader.ReadUint8()

	// Flow Description
	if p.Fd {
		fdLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for fdLen: %v", e)
		}

		flowDescription, e := byteReader.ReadBytes(int(fdLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for FlowDescription: %v", e)
		}
		p.FlowDescription = string(flowDescription)
	}

	// URL
	if p.Url {
		urlLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for urlLen: %v", e)
		}

		url, e := byteReader.ReadBytes(int(urlLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for URL: %v", e)
		}
		p.URL = string(url)
	}

	// Domain Name
	if p.Dn {
		dnLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for dnLen: %v", e)
		}

		domainName, e := byteReader.ReadBytes(int(dnLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for DomainName: %v", e)
		}
		p.DomainName = string(domainName)
	}

	// Custom PFD Contents
	if p.Cp {
		cpLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for cpLen: %v", e)
		}
		cpBuf, _ := byteReader.ReadBytes(int(cpLen))
		var cpContent CustomPFDContent
		e = cpContent.UnmarshalBinary(cpBuf)
		if e != nil {
			return fmt.Errorf("Unmarshal CustomPFDContent failed: %v", e)
		}
		p.CustomPFDContents = cpContent
	}

	// Domain Name Protocols
	if p.Dnp {
		if !p.Dn {
			return fmt.Errorf("When Adnp is set,the Domain Name fields shall also be present in IE PFDContents ")
		}
		dnpLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for dnpLen: %v", e)
		}

		domainNameProtocols, e := byteReader.ReadBytes(int(dnpLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for DomainNameProtocols: %v", e)
		}

		p.DomainNameProtocol = string(domainNameProtocols)
	}

	// Additional Flow Description
	if p.Afd {
		afdLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for afdLen: %v", e)
		}
		afd, e := byteReader.ReadBytes(int(afdLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for AdditionalFlowDescription: %v", e)
		}
		afdReader := NewByteReader(afd)

		for {
			//len of FlowDescription
			fdLen, e := afdReader.ReadUint16()
			// Break when adfReader is drained.
			if fdLen == 0 {
				break
			}
			if e != nil {
				return fmt.Errorf("unmarshal lenght of Afd.Fd Inadequate TLV length: %v", e)
			}

			//FlowDescription
			fd, e := afdReader.ReadBytes(int(fdLen))
			if e != nil {
				return fmt.Errorf("unmarshal Afd.Fd Inadequate TLV length: %v", e)
			}
			p.AdditionalFlowDescription = append(p.AdditionalFlowDescription, string(fd))
		}
	}

	// Additional URL
	if p.Aurl {
		aurlLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for aurlLen: %v", e)
		}
		aurl, e := byteReader.ReadBytes(int(aurlLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for AdditionalFlowDescription: %v", e)
		}
		aurlReader := NewByteReader(aurl)
		for {
			urlLen, e := aurlReader.ReadUint16()
			// Break when aurlReader is drained.
			if urlLen == 0 {
				break
			}
			if e != nil {
				return fmt.Errorf("unmarshal length of Aurl.url Inadequate TLV length: %v", e)
			}

			//FlowDescription
			url, e := aurlReader.ReadBytes(int(urlLen))
			if e != nil {
				return fmt.Errorf("unmarshal Aurl.url Inadequate TLV length: %v", e)
			}
			p.AdditionalURL = append(p.AdditionalURL, string(url))
		}

	}

	// Additional Domain Name and Domain Name Protocols
	if p.Adnp {

		adnpLen, e := byteReader.ReadUint16()
		if e != nil {
			return fmt.Errorf("inadequate TLV length for adnpLen: %v", e)
		}
		adnp, e := byteReader.ReadBytes(int(adnpLen))
		if e != nil {
			return fmt.Errorf("inadequate TLV length for AdditionalFlowDescription: %v", e)
		}
		adnpReader := NewByteReader(adnp)
		//for !adnpReader.IsDrained() {
		for {
			dnAndDnp := AdditionalDomainNameAndDomainNameProtocol{}
			// Domain Name
			dnLen, e := adnpReader.ReadUint16()
			// Reader Buffer is Drained.
			if dnLen == 0 {
				break
			}
			if e != nil {
				return fmt.Errorf("unmarshal length of adnp.dn Inadequate TLV length: %v", e)
			}
			dn, _ := adnpReader.ReadBytes(int(dnLen))
			dnAndDnp.DomainName = string(dn)

			// Domain Name Protocol
			dnpLen, e := adnpReader.ReadUint16()
			if e != nil {
				return fmt.Errorf("unmarshal length of adnp.dnp Inadequate TLV length: %v", e)
			}
			dnp, e := adnpReader.ReadBytes(int(dnpLen))
			dnAndDnp.DomainNameProtocol = string(dnp)
			if e != nil {
				return fmt.Errorf("unmarshal adnp.dnp Inadequate TLV length: %v", e)
			}
			p.AdditionalDomainNameAndDomainNameProtocol = append(p.AdditionalDomainNameAndDomainNameProtocol, dnAndDnp)
		}
	}

	return nil
}
