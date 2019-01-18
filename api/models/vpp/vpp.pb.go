// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vpp/vpp.proto

package vpp // import "github.com/ligato/vpp-agent/api/models/vpp"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import acl "github.com/ligato/vpp-agent/api/models/vpp/acl"
import interfaces "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
import ipsec "github.com/ligato/vpp-agent/api/models/vpp/ipsec"
import l2 "github.com/ligato/vpp-agent/api/models/vpp/l2"
import l3 "github.com/ligato/vpp-agent/api/models/vpp/l3"
import nat "github.com/ligato/vpp-agent/api/models/vpp/nat"
import punt "github.com/ligato/vpp-agent/api/models/vpp/punt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Data struct {
	Interfaces           []*interfaces.Interface         `protobuf:"bytes,10,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	ACLs                 []*acl.Acl                      `protobuf:"bytes,20,rep,name=acls,proto3" json:"acls,omitempty"`
	BridgeDomains        []*l2.BridgeDomain              `protobuf:"bytes,30,rep,name=bridge_domains,json=bridgeDomains,proto3" json:"bridge_domains,omitempty"`
	FIBs                 []*l2.FIBEntry                  `protobuf:"bytes,31,rep,name=fibs,proto3" json:"fibs,omitempty"`
	XConnectPairs        []*l2.XConnectPair              `protobuf:"bytes,32,rep,name=xconnect_pairs,json=xconnectPairs,proto3" json:"xconnect_pairs,omitempty"`
	Routes               []*l3.Route                     `protobuf:"bytes,40,rep,name=routes,proto3" json:"routes,omitempty"`
	ARPs                 []*l3.ARPEntry                  `protobuf:"bytes,41,rep,name=arps,proto3" json:"arps,omitempty"`
	ProxyARP             *l3.ProxyARP                    `protobuf:"bytes,42,opt,name=proxy_arp,json=proxyArp,proto3" json:"proxy_arp,omitempty"`
	IPScanNeighbor       *l3.IPScanNeighbor              `protobuf:"bytes,43,opt,name=ipscan_neighbor,json=ipscanNeighbor,proto3" json:"ipscan_neighbor,omitempty"`
	NAT44Global          *nat.Nat44Global                `protobuf:"bytes,50,opt,name=nat44_global,json=nat44Global,proto3" json:"nat44_global,omitempty"`
	DNAT44s              []*nat.DNat44                   `protobuf:"bytes,51,rep,name=dnat44s,proto3" json:"dnat44s,omitempty"`
	IPSecSPDs            []*ipsec.SecurityPolicyDatabase `protobuf:"bytes,60,rep,name=ipsec_spds,json=ipsecSpds,proto3" json:"ipsec_spds,omitempty"`
	IPSecSAs             []*ipsec.SecurityAssociation    `protobuf:"bytes,61,rep,name=ipsec_sas,json=ipsecSas,proto3" json:"ipsec_sas,omitempty"`
	PuntIPRedirects      []*punt.IpRedirect              `protobuf:"bytes,70,rep,name=punt_ipredirects,json=puntIpredirects,proto3" json:"punt_ipredirects,omitempty"`
	PuntToHosts          []*punt.ToHost                  `protobuf:"bytes,71,rep,name=punt_tohosts,json=puntTohosts,proto3" json:"punt_tohosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpp_f7c4ff53b245e3b0, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetInterfaces() []*interfaces.Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *Data) GetACLs() []*acl.Acl {
	if m != nil {
		return m.ACLs
	}
	return nil
}

func (m *Data) GetBridgeDomains() []*l2.BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

func (m *Data) GetFIBs() []*l2.FIBEntry {
	if m != nil {
		return m.FIBs
	}
	return nil
}

func (m *Data) GetXConnectPairs() []*l2.XConnectPair {
	if m != nil {
		return m.XConnectPairs
	}
	return nil
}

func (m *Data) GetRoutes() []*l3.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *Data) GetARPs() []*l3.ARPEntry {
	if m != nil {
		return m.ARPs
	}
	return nil
}

func (m *Data) GetProxyARP() *l3.ProxyARP {
	if m != nil {
		return m.ProxyARP
	}
	return nil
}

func (m *Data) GetIPScanNeighbor() *l3.IPScanNeighbor {
	if m != nil {
		return m.IPScanNeighbor
	}
	return nil
}

func (m *Data) GetNAT44Global() *nat.Nat44Global {
	if m != nil {
		return m.NAT44Global
	}
	return nil
}

func (m *Data) GetDNAT44s() []*nat.DNat44 {
	if m != nil {
		return m.DNAT44s
	}
	return nil
}

func (m *Data) GetIPSecSPDs() []*ipsec.SecurityPolicyDatabase {
	if m != nil {
		return m.IPSecSPDs
	}
	return nil
}

func (m *Data) GetIPSecSAs() []*ipsec.SecurityAssociation {
	if m != nil {
		return m.IPSecSAs
	}
	return nil
}

func (m *Data) GetPuntIPRedirects() []*punt.IpRedirect {
	if m != nil {
		return m.PuntIPRedirects
	}
	return nil
}

func (m *Data) GetPuntToHosts() []*punt.ToHost {
	if m != nil {
		return m.PuntToHosts
	}
	return nil
}

func (*Data) XXX_MessageName() string {
	return "vpp.Data"
}
func init() {
	proto.RegisterType((*Data)(nil), "vpp.Data")
}

func init() { proto.RegisterFile("vpp/vpp.proto", fileDescriptor_vpp_f7c4ff53b245e3b0) }

var fileDescriptor_vpp_f7c4ff53b245e3b0 = []byte{
	// 729 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xd1, 0x6a, 0xe3, 0x46,
	0x14, 0x65, 0xd9, 0xb0, 0x49, 0xc6, 0x71, 0x9c, 0x9d, 0xdd, 0x16, 0x35, 0x0f, 0x51, 0x5a, 0x28,
	0x6c, 0xb7, 0x44, 0x02, 0x3b, 0x14, 0xca, 0xb6, 0x0f, 0xd2, 0xba, 0xd9, 0x15, 0x14, 0x57, 0x8c,
	0xf3, 0x50, 0xfa, 0x62, 0x46, 0xb2, 0xac, 0x0c, 0x28, 0x9a, 0x61, 0xee, 0xd8, 0xc4, 0x7f, 0xd8,
	0x97, 0xfe, 0x82, 0x1e, 0xd4, 0x1f, 0x29, 0x73, 0x47, 0xaa, 0x15, 0xd8, 0x07, 0x09, 0xdd, 0x73,
	0xee, 0x39, 0x9a, 0x39, 0xdc, 0x19, 0x32, 0xde, 0x29, 0x15, 0xee, 0x94, 0x0a, 0x94, 0x96, 0x46,
	0xd2, 0x97, 0x3b, 0xa5, 0x2e, 0x6f, 0x4a, 0x61, 0x1e, 0xb6, 0x59, 0x90, 0xcb, 0xc7, 0xb0, 0x94,
	0xa5, 0x0c, 0x91, 0xcb, 0xb6, 0x1b, 0xac, 0xb0, 0xc0, 0x2f, 0xa7, 0xb9, 0x7c, 0x6d, 0x2d, 0x78,
	0x5e, 0xd9, 0xa7, 0x83, 0xae, 0x2c, 0x24, 0x6a, 0x53, 0xe8, 0x0d, 0xcf, 0x0b, 0x38, 0x7c, 0x76,
	0xfc, 0x57, 0xc8, 0x2b, 0x28, 0x72, 0xf7, 0xee, 0xe0, 0x4b, 0x0b, 0x57, 0xd3, 0x30, 0xd3, 0x62,
	0x5d, 0x16, 0x37, 0x6b, 0xf9, 0xc8, 0x45, 0xdd, 0x71, 0x17, 0x1d, 0xb7, 0x11, 0xd9, 0xd0, 0xa4,
	0x9a, 0x86, 0x4f, 0xb9, 0xac, 0xeb, 0x22, 0x37, 0xcf, 0x1a, 0x67, 0x21, 0xd7, 0xdd, 0xa6, 0x2e,
	0x27, 0x1d, 0x52, 0xcd, 0x3a, 0x80, 0x76, 0x80, 0x96, 0x5b, 0x53, 0x0c, 0x77, 0x51, 0x73, 0x63,
	0x9f, 0x0e, 0x7a, 0x63, 0x21, 0xb5, 0xad, 0x0d, 0xbe, 0x1c, 0xf8, 0xdd, 0x3f, 0xc7, 0xe4, 0x68,
	0xce, 0x0d, 0xa7, 0x3f, 0x13, 0x72, 0xd8, 0xa1, 0x47, 0xae, 0x5f, 0xbe, 0x1b, 0x4d, 0xbf, 0x09,
	0x6c, 0x94, 0x07, 0x38, 0x48, 0xfa, 0x4f, 0x36, 0x68, 0xa6, 0xef, 0xc9, 0x11, 0xcf, 0x2b, 0xf0,
	0xde, 0xa2, 0xe8, 0x0c, 0x45, 0x36, 0xbc, 0x28, 0xaf, 0xe2, 0x93, 0xb6, 0xf1, 0x8f, 0xa2, 0x8f,
	0xbf, 0x03, 0xc3, 0x1e, 0xfa, 0x81, 0x9c, 0xbb, 0x38, 0x56, 0x2e, 0x0e, 0xf0, 0xae, 0x50, 0xf5,
	0x16, 0x55, 0xd5, 0x34, 0x88, 0x91, 0x9d, 0x23, 0xc9, 0xc6, 0xd9, 0xa0, 0x02, 0x1a, 0x90, 0xa3,
	0x8d, 0xc8, 0xc0, 0xf3, 0x51, 0x72, 0xd1, 0x4b, 0xee, 0x92, 0xf8, 0xb7, 0xda, 0xe8, 0xbd, 0xfb,
	0xd9, 0x5d, 0x12, 0x03, 0xc3, 0x3e, 0xba, 0x20, 0xe7, 0x7d, 0x9a, 0x2b, 0xc5, 0x85, 0x06, 0xef,
	0xfa, 0xf9, 0xcf, 0xfe, 0xfc, 0xe8, 0xd8, 0x94, 0x0b, 0x1d, 0xbf, 0x6e, 0x1b, 0x7f, 0x3c, 0x44,
	0x80, 0x8d, 0x7b, 0x39, 0x96, 0xf4, 0x7b, 0xf2, 0x0a, 0x33, 0x06, 0xef, 0x1d, 0xfa, 0x8c, 0x9d,
	0xcf, 0x2c, 0x60, 0x16, 0x65, 0x1d, 0x69, 0x97, 0xc9, 0xb5, 0x02, 0xef, 0x87, 0xe1, 0x32, 0x67,
	0x41, 0xc4, 0xd2, 0xc1, 0x32, 0x23, 0x96, 0xda, 0x4c, 0xb4, 0xb2, 0x99, 0x9c, 0x2a, 0x2d, 0x9f,
	0xf6, 0x2b, 0xae, 0x95, 0xf7, 0xfe, 0xfa, 0xc5, 0x50, 0x94, 0x5a, 0x22, 0x62, 0x69, 0x7c, 0xd6,
	0x36, 0xfe, 0x49, 0x5f, 0xb1, 0x13, 0x14, 0x44, 0x5a, 0xd1, 0x25, 0x99, 0x08, 0x05, 0x39, 0xaf,
	0x57, 0x75, 0x21, 0xca, 0x87, 0x4c, 0x6a, 0xef, 0x47, 0xb4, 0xf8, 0xba, 0xb7, 0x48, 0xd2, 0x65,
	0xce, 0xeb, 0x45, 0xc7, 0xc6, 0xb4, 0x6d, 0xfc, 0xf3, 0xe7, 0x18, 0x3b, 0x77, 0x16, 0x7d, 0x4d,
	0x3f, 0x93, 0xb3, 0x9a, 0x9b, 0xdb, 0xdb, 0x55, 0x59, 0xc9, 0x8c, 0x57, 0xde, 0x14, 0x1d, 0x5d,
	0x6c, 0x76, 0xa0, 0x16, 0x96, 0xfc, 0x84, 0x5c, 0x3c, 0x69, 0x1b, 0x7f, 0xb4, 0x88, 0xee, 0x7b,
	0x80, 0x8d, 0xea, 0x03, 0x4b, 0x7f, 0x22, 0xc7, 0x6b, 0xac, 0xc1, 0x9b, 0x61, 0x1c, 0x93, 0xff,
	0x4d, 0xe6, 0xe8, 0x12, 0x8f, 0xda, 0xc6, 0x3f, 0x9e, 0xa3, 0x01, 0xb0, 0xbe, 0x99, 0xfe, 0x41,
	0x08, 0x1e, 0xa5, 0x15, 0xa8, 0x35, 0x78, 0xbf, 0xa0, 0xf4, 0x5b, 0x37, 0x8e, 0x78, 0xc2, 0x96,
	0x45, 0xbe, 0xd5, 0xc2, 0xec, 0x53, 0x59, 0x89, 0x7c, 0x6f, 0x27, 0x38, 0xe3, 0x50, 0xc4, 0xe3,
	0xb6, 0xf1, 0x4f, 0x93, 0x74, 0x59, 0xe4, 0xcb, 0x74, 0x0e, 0xec, 0x14, 0x9b, 0x97, 0x6a, 0x0d,
	0x34, 0x21, 0xa7, 0x9d, 0x21, 0x07, 0xef, 0x57, 0xf4, 0xbb, 0xfa, 0x82, 0x5f, 0x04, 0x20, 0x73,
	0xc1, 0x8d, 0x90, 0xb5, 0x8b, 0xdc, 0x99, 0x45, 0xc0, 0x4e, 0x9c, 0x17, 0x07, 0x7a, 0x4f, 0x2e,
	0xec, 0x09, 0x5a, 0x09, 0xa5, 0x8b, 0xb5, 0xd0, 0x45, 0x6e, 0xc0, 0xbb, 0x1b, 0x0c, 0x16, 0x1e,
	0xaf, 0x44, 0xb1, 0x8e, 0x8c, 0xdf, 0xb4, 0x8d, 0x3f, 0x49, 0xb7, 0xb5, 0x49, 0xd2, 0x1e, 0x03,
	0x36, 0xb1, 0x5d, 0xc9, 0xc1, 0x81, 0xce, 0xc9, 0x19, 0xba, 0x1a, 0xf9, 0x20, 0xc1, 0x80, 0xf7,
	0x69, 0x30, 0x3d, 0xe8, 0x78, 0x2f, 0x3f, 0x4b, 0x30, 0x2e, 0x6f, 0xeb, 0xe6, 0x6a, 0x60, 0x23,
	0x85, 0x05, 0xaa, 0xe2, 0xdb, 0xbf, 0xff, 0xbd, 0x7a, 0xf1, 0x57, 0x30, 0xb8, 0xf2, 0x2a, 0x51,
	0x72, 0x23, 0xed, 0xa5, 0x78, 0xc3, 0xcb, 0xa2, 0x36, 0x21, 0x57, 0x22, 0x7c, 0x94, 0xeb, 0xa2,
	0x02, 0x0b, 0x7e, 0xd8, 0x29, 0x95, 0xbd, 0xc2, 0xcb, 0x60, 0xf6, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0x18, 0xf7, 0x56, 0x3f, 0x05, 0x00, 0x00,
}
