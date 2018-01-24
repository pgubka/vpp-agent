// Code generated by protoc-gen-gogo.
// source: nat.proto
// DO NOT EDIT!

/*
Package nat is a generated protocol buffer package.

It is generated from these files:
	nat.proto

It has these top-level messages:
	Nat44Global
	Nat44SNat
	Nat44DNat
*/
package nat

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Nat44DNat_DNatConfig_Mapping_Protocol int32

const (
	Nat44DNat_DNatConfig_Mapping_TCP  Nat44DNat_DNatConfig_Mapping_Protocol = 0
	Nat44DNat_DNatConfig_Mapping_UDP  Nat44DNat_DNatConfig_Mapping_Protocol = 1
	Nat44DNat_DNatConfig_Mapping_ICMP Nat44DNat_DNatConfig_Mapping_Protocol = 2
)

var Nat44DNat_DNatConfig_Mapping_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
	2: "ICMP",
}
var Nat44DNat_DNatConfig_Mapping_Protocol_value = map[string]int32{
	"TCP":  0,
	"UDP":  1,
	"ICMP": 2,
}

func (x Nat44DNat_DNatConfig_Mapping_Protocol) String() string {
	return proto.EnumName(Nat44DNat_DNatConfig_Mapping_Protocol_name, int32(x))
}

// NAT44 global config
type Nat44Global struct {
	VrfId       uint32                     `protobuf:"varint,1,opt,name=vrfId,proto3" json:"vrfId,omitempty"`
	Forwarding  bool                       `protobuf:"varint,2,opt,name=forwarding,proto3" json:"forwarding,omitempty"`
	In          []string                   `protobuf:"bytes,3,rep,name=in" json:"in,omitempty"`
	Out         []string                   `protobuf:"bytes,4,rep,name=out" json:"out,omitempty"`
	AddressPool []*Nat44Global_AddressPool `protobuf:"bytes,5,rep,name=address_pool" json:"address_pool,omitempty"`
}

func (m *Nat44Global) Reset()         { *m = Nat44Global{} }
func (m *Nat44Global) String() string { return proto.CompactTextString(m) }
func (*Nat44Global) ProtoMessage()    {}

func (m *Nat44Global) GetAddressPool() []*Nat44Global_AddressPool {
	if m != nil {
		return m.AddressPool
	}
	return nil
}

type Nat44Global_AddressPool struct {
	FirstSrcAddress string `protobuf:"bytes,1,opt,proto3" json:"FirstSrcAddress,omitempty"`
	LastSrcAddres   string `protobuf:"bytes,2,opt,proto3" json:"LastSrcAddres,omitempty"`
	TwiceNat        bool   `protobuf:"varint,3,opt,name=twiceNat,proto3" json:"twiceNat,omitempty"`
}

func (m *Nat44Global_AddressPool) Reset()         { *m = Nat44Global_AddressPool{} }
func (m *Nat44Global_AddressPool) String() string { return proto.CompactTextString(m) }
func (*Nat44Global_AddressPool) ProtoMessage()    {}

// Many-to-one (SNAT) setup
type Nat44SNat struct {
	SnatConfig []*Nat44SNat_SNatConfig `protobuf:"bytes,1,rep,name=snat_config" json:"snat_config,omitempty"`
}

func (m *Nat44SNat) Reset()         { *m = Nat44SNat{} }
func (m *Nat44SNat) String() string { return proto.CompactTextString(m) }
func (*Nat44SNat) ProtoMessage()    {}

func (m *Nat44SNat) GetSnatConfig() []*Nat44SNat_SNatConfig {
	if m != nil {
		return m.SnatConfig
	}
	return nil
}

type Nat44SNat_SNatConfig struct {
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
}

func (m *Nat44SNat_SNatConfig) Reset()         { *m = Nat44SNat_SNatConfig{} }
func (m *Nat44SNat_SNatConfig) String() string { return proto.CompactTextString(m) }
func (*Nat44SNat_SNatConfig) ProtoMessage()    {}

// One-to-many (DNAT) setup
type Nat44DNat struct {
	DnatConfig []*Nat44DNat_DNatConfig `protobuf:"bytes,1,rep,name=dnat_config" json:"dnat_config,omitempty"`
}

func (m *Nat44DNat) Reset()         { *m = Nat44DNat{} }
func (m *Nat44DNat) String() string { return proto.CompactTextString(m) }
func (*Nat44DNat) ProtoMessage()    {}

func (m *Nat44DNat) GetDnatConfig() []*Nat44DNat_DNatConfig {
	if m != nil {
		return m.DnatConfig
	}
	return nil
}

type Nat44DNat_DNatConfig struct {
	Label       string                          `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	VrfId       uint32                          `protobuf:"varint,2,opt,name=vrfId,proto3" json:"vrfId,omitempty"`
	SNatEnabled bool                            `protobuf:"varint,3,opt,proto3" json:"SNatEnabled,omitempty"`
	Mapping     []*Nat44DNat_DNatConfig_Mapping `protobuf:"bytes,4,rep,name=mapping" json:"mapping,omitempty"`
}

func (m *Nat44DNat_DNatConfig) Reset()         { *m = Nat44DNat_DNatConfig{} }
func (m *Nat44DNat_DNatConfig) String() string { return proto.CompactTextString(m) }
func (*Nat44DNat_DNatConfig) ProtoMessage()    {}

func (m *Nat44DNat_DNatConfig) GetMapping() []*Nat44DNat_DNatConfig_Mapping {
	if m != nil {
		return m.Mapping
	}
	return nil
}

type Nat44DNat_DNatConfig_Mapping struct {
	ExternalInterface string                                  `protobuf:"bytes,1,opt,name=externalInterface,proto3" json:"externalInterface,omitempty"`
	ExternalIP        string                                  `protobuf:"bytes,2,opt,name=externalIP,proto3" json:"externalIP,omitempty"`
	LocalIp           []*Nat44DNat_DNatConfig_Mapping_LocalIP `protobuf:"bytes,3,rep,name=local_ip" json:"local_ip,omitempty"`
	Protocol          Nat44DNat_DNatConfig_Mapping_Protocol   `protobuf:"varint,4,opt,name=protocol,proto3,enum=nat.Nat44DNat_DNatConfig_Mapping_Protocol" json:"protocol,omitempty"`
}

func (m *Nat44DNat_DNatConfig_Mapping) Reset()         { *m = Nat44DNat_DNatConfig_Mapping{} }
func (m *Nat44DNat_DNatConfig_Mapping) String() string { return proto.CompactTextString(m) }
func (*Nat44DNat_DNatConfig_Mapping) ProtoMessage()    {}

func (m *Nat44DNat_DNatConfig_Mapping) GetLocalIp() []*Nat44DNat_DNatConfig_Mapping_LocalIP {
	if m != nil {
		return m.LocalIp
	}
	return nil
}

type Nat44DNat_DNatConfig_Mapping_LocalIP struct {
	LocalIP     string `protobuf:"bytes,1,opt,name=localIP,proto3" json:"localIP,omitempty"`
	Probability uint32 `protobuf:"varint,2,opt,name=probability,proto3" json:"probability,omitempty"`
}

func (m *Nat44DNat_DNatConfig_Mapping_LocalIP) Reset()         { *m = Nat44DNat_DNatConfig_Mapping_LocalIP{} }
func (m *Nat44DNat_DNatConfig_Mapping_LocalIP) String() string { return proto.CompactTextString(m) }
func (*Nat44DNat_DNatConfig_Mapping_LocalIP) ProtoMessage()    {}

func init() {
	proto.RegisterEnum("nat.Nat44DNat_DNatConfig_Mapping_Protocol", Nat44DNat_DNatConfig_Mapping_Protocol_name, Nat44DNat_DNatConfig_Mapping_Protocol_value)
}
