// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/datawire/ambassador/pkg/api/envoy/config/common/dynamic_forward_proxy/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for the dynamic forward proxy HTTP filter. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.filters.http.dynamic_forward_proxy]
type FilterConfig struct {
	// The DNS cache configuration that the filter will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_api_field_config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

// Per route Configuration for the dynamic forward proxy HTTP filter.
type PerRouteConfig struct {
	// Types that are valid to be assigned to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewrite
	//	*PerRouteConfig_AutoHostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{1}
}
func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return m.Size()
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
	MarshalTo([]byte) (int, error)
	Size() int
}

type PerRouteConfig_HostRewrite struct {
	HostRewrite string `protobuf:"bytes,1,opt,name=host_rewrite,json=hostRewrite,proto3,oneof" json:"host_rewrite,omitempty"`
}
type PerRouteConfig_AutoHostRewriteHeader struct {
	AutoHostRewriteHeader string `protobuf:"bytes,2,opt,name=auto_host_rewrite_header,json=autoHostRewriteHeader,proto3,oneof" json:"auto_host_rewrite_header,omitempty"`
}

func (*PerRouteConfig_HostRewrite) isPerRouteConfig_HostRewriteSpecifier()           {}
func (*PerRouteConfig_AutoHostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (m *PerRouteConfig) GetHostRewrite() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewrite); ok {
		return x.HostRewrite
	}
	return ""
}

func (m *PerRouteConfig) GetAutoHostRewriteHeader() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_AutoHostRewriteHeader); ok {
		return x.AutoHostRewriteHeader
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerRouteConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerRouteConfig_HostRewrite)(nil),
		(*PerRouteConfig_AutoHostRewriteHeader)(nil),
	}
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto", fileDescriptor_85a2356b260c47da)
}

var fileDescriptor_85a2356b260c47da = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x8b, 0xd4, 0x40,
	0x14, 0xc7, 0x77, 0x16, 0xee, 0xd0, 0xb9, 0xe3, 0x58, 0xa2, 0x77, 0x86, 0x53, 0x83, 0x5c, 0x65,
	0x95, 0x81, 0x5b, 0x10, 0xdb, 0xcd, 0x2e, 0x4b, 0xca, 0x90, 0xc2, 0x36, 0x8c, 0xc9, 0x64, 0x33,
	0x90, 0xcc, 0x84, 0x99, 0x97, 0xfd, 0xd1, 0xd9, 0xd8, 0xd8, 0xd8, 0xfa, 0x87, 0x58, 0xd9, 0x0b,
	0x5b, 0xea, 0x7f, 0x20, 0xfb, 0x27, 0x58, 0x5a, 0x88, 0x64, 0x26, 0xab, 0x06, 0x97, 0x15, 0xae,
	0x0b, 0xef, 0xf3, 0xe6, 0x33, 0xef, 0x65, 0xbe, 0x38, 0x66, 0x62, 0x29, 0x37, 0x24, 0x95, 0x22,
	0xe7, 0x0b, 0x92, 0xf3, 0x12, 0x98, 0x22, 0x05, 0x40, 0x4d, 0xb2, 0x8d, 0xa0, 0x15, 0x4f, 0x93,
	0x5c, 0xaa, 0x15, 0x55, 0x59, 0x52, 0x2b, 0xb9, 0xde, 0x90, 0xe5, 0x2d, 0x2d, 0xeb, 0x82, 0x1e,
	0xa6, 0x7e, 0xad, 0x24, 0x48, 0xe7, 0x85, 0x71, 0xfa, 0xd6, 0xe9, 0x5b, 0xa7, 0xdf, 0x3a, 0xfd,
	0xc3, 0xa7, 0x3a, 0xe7, 0xf5, 0xa4, 0x37, 0x4b, 0x2a, 0xab, 0x4a, 0x8a, 0xff, 0x8d, 0x21, 0x74,
	0x92, 0xd2, 0xb4, 0x60, 0xf6, 0xea, 0x6b, 0xaf, 0xc9, 0x6a, 0x4a, 0xa8, 0x10, 0x12, 0x28, 0x70,
	0x29, 0x34, 0xa9, 0xf8, 0x42, 0x51, 0xd8, 0xf3, 0xa7, 0xff, 0x70, 0x0d, 0x14, 0x1a, 0xdd, 0xe1,
	0x47, 0x4b, 0x5a, 0xf2, 0x8c, 0x02, 0x23, 0xfb, 0x0f, 0x0b, 0x6e, 0xde, 0x22, 0x7c, 0x3e, 0x37,
	0x8b, 0x4c, 0xcd, 0x74, 0x4e, 0x83, 0x47, 0xbf, 0xef, 0x4e, 0xec, 0xc4, 0x2e, 0x7a, 0x86, 0x9e,
	0x9f, 0xdd, 0x4e, 0xfc, 0xde, 0xfa, 0x76, 0x8d, 0xe3, 0x9b, 0xfb, 0x33, 0xa1, 0xa7, 0xad, 0xc9,
	0xca, 0x83, 0x7b, 0x3f, 0x82, 0x93, 0x77, 0x68, 0x38, 0x42, 0xf1, 0x45, 0xd6, 0x23, 0x37, 0x9f,
	0x11, 0xbe, 0x88, 0x98, 0x8a, 0x65, 0x03, 0x5d, 0xc9, 0x99, 0xe0, 0xf3, 0x42, 0x6a, 0x48, 0x14,
	0x5b, 0x29, 0x0e, 0xcc, 0x4c, 0x71, 0x3f, 0x78, 0xf2, 0xfd, 0xc3, 0xcf, 0xf7, 0x27, 0x57, 0xf8,
	0xe1, 0xdf, 0x2c, 0x29, 0x39, 0x30, 0x45, 0xcb, 0x70, 0x10, 0x9f, 0xb5, 0xf5, 0xd8, 0x96, 0x9d,
	0x57, 0xd8, 0xa5, 0x0d, 0xc8, 0xa4, 0xd7, 0x5b, 0x30, 0x9a, 0x31, 0xe5, 0x0e, 0x8d, 0xee, 0xb1,
	0xd1, 0x5d, 0xe2, 0x07, 0x07, 0x5a, 0xc2, 0x41, 0x7c, 0xd9, 0x1e, 0x0f, 0xff, 0x18, 0x43, 0x03,
	0x02, 0x17, 0x5f, 0xf5, 0xfa, 0x75, 0xcd, 0x52, 0x9e, 0x73, 0xa6, 0x82, 0x8f, 0x68, 0xbb, 0xf3,
	0xd0, 0x97, 0x9d, 0x87, 0xbe, 0xed, 0x3c, 0x64, 0xe4, 0x2f, 0xf7, 0xb9, 0x61, 0x6b, 0x60, 0x42,
	0xb7, 0x8f, 0xd3, 0x65, 0x47, 0x1f, 0x0d, 0xcf, 0xf8, 0xd3, 0x9b, 0xed, 0xd7, 0xd3, 0xe1, 0x68,
	0x80, 0x67, 0x5c, 0xda, 0x7f, 0x6f, 0xc9, 0xdd, 0x52, 0x18, 0xb8, 0x33, 0x8b, 0xe7, 0x96, 0x46,
	0x2d, 0x8c, 0xda, 0x10, 0x44, 0xe8, 0xf5, 0xa9, 0x49, 0xc3, 0xf8, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0xcb, 0x45, 0x51, 0x36, 0x03, 0x00, 0x00,
}

func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DnsCacheConfig != nil {
		{
			size, err := m.DnsCacheConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PerRouteConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerRouteConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerRouteConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.HostRewriteSpecifier != nil {
		{
			size := m.HostRewriteSpecifier.Size()
			i -= size
			if _, err := m.HostRewriteSpecifier.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *PerRouteConfig_HostRewrite) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerRouteConfig_HostRewrite) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.HostRewrite)
	copy(dAtA[i:], m.HostRewrite)
	i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(len(m.HostRewrite)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func (m *PerRouteConfig_AutoHostRewriteHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerRouteConfig_AutoHostRewriteHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.AutoHostRewriteHeader)
	copy(dAtA[i:], m.AutoHostRewriteHeader)
	i = encodeVarintDynamicForwardProxy(dAtA, i, uint64(len(m.AutoHostRewriteHeader)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintDynamicForwardProxy(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicForwardProxy(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DnsCacheConfig != nil {
		l = m.DnsCacheConfig.Size()
		n += 1 + l + sovDynamicForwardProxy(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PerRouteConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HostRewriteSpecifier != nil {
		n += m.HostRewriteSpecifier.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PerRouteConfig_HostRewrite) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HostRewrite)
	n += 1 + l + sovDynamicForwardProxy(uint64(l))
	return n
}
func (m *PerRouteConfig_AutoHostRewriteHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AutoHostRewriteHeader)
	n += 1 + l + sovDynamicForwardProxy(uint64(l))
	return n
}

func sovDynamicForwardProxy(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicForwardProxy(x uint64) (n int) {
	return sovDynamicForwardProxy(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicForwardProxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsCacheConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsCacheConfig == nil {
				m.DnsCacheConfig = &v2alpha.DnsCacheConfig{}
			}
			if err := m.DnsCacheConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicForwardProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PerRouteConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicForwardProxy
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PerRouteConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerRouteConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostRewrite", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostRewriteSpecifier = &PerRouteConfig_HostRewrite{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AutoHostRewriteHeader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicForwardProxy
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostRewriteSpecifier = &PerRouteConfig_AutoHostRewriteHeader{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicForwardProxy(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDynamicForwardProxy
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDynamicForwardProxy(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicForwardProxy
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDynamicForwardProxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDynamicForwardProxy
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDynamicForwardProxy
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicForwardProxy
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicForwardProxy
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicForwardProxy        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicForwardProxy          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicForwardProxy = fmt.Errorf("proto: unexpected end of group")
)
