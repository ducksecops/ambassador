// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/matcher/v3alpha/metadata.proto

package envoy_type_matcher_v3alpha

import (
	fmt "fmt"
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

// MetadataMatcher provides a general interface to check if a given value is matched in
// :ref:`Metadata <envoy_api_msg_api.v3alpha.core.Metadata>`. It uses `filter` and `path` to
// retrieve the value from the Metadata and then check if it's matched to the specified value.
//
// For example, for the following Metadata:
//
// .. code-block:: yaml
//
//    filter_metadata:
//      envoy.filters.http.rbac:
//        fields:
//          a:
//            struct_value:
//              fields:
//                b:
//                  struct_value:
//                    fields:
//                      c:
//                        string_value: pro
//                t:
//                  list_value:
//                    values:
//                      - string_value: m
//                      - string_value: n
//
// The following MetadataMatcher is matched as the path [a, b, c] will retrieve a string value "pro"
// from the Metadata which is matched to the specified prefix match.
//
// .. code-block:: yaml
//
//    filter: envoy.filters.http.rbac
//    path:
//    - key: a
//    - key: b
//    - key: c
//    value:
//      string_match:
//        prefix: pr
//
// The following MetadataMatcher is matched as the code will match one of the string values in the
// list at the path [a, t].
//
// .. code-block:: yaml
//
//    filter: envoy.filters.http.rbac
//    path:
//    - key: a
//    - key: t
//    value:
//      list_match:
//        one_of:
//          string_match:
//            exact: m
//
// An example use of MetadataMatcher is specifying additional metadata in envoy.filters.http.rbac to
// enforce access control based on dynamic metadata in a request. See :ref:`Permission
// <envoy_api_msg_config.rbac.v3alpha.Permission>` and :ref:`Principal
// <envoy_api_msg_config.rbac.v3alpha.Principal>`.
type MetadataMatcher struct {
	// The filter name to retrieve the Struct from the Metadata.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	// The path to retrieve the Value from the Struct.
	Path []*MetadataMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	// The MetadataMatcher is matched if the value retrieved by path is matched to this value.
	Value                *ValueMatcher `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MetadataMatcher) Reset()         { *m = MetadataMatcher{} }
func (m *MetadataMatcher) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher) ProtoMessage()    {}
func (*MetadataMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_d306cd29ee202178, []int{0}
}
func (m *MetadataMatcher) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMatcher.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher.Merge(m, src)
}
func (m *MetadataMatcher) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher proto.InternalMessageInfo

func (m *MetadataMatcher) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *MetadataMatcher) GetPath() []*MetadataMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *MetadataMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

// Specifies the segment in a path to retrieve value from Metadata.
// Note: Currently it's not supported to retrieve a value from a list in Metadata. This means that
// if the segment key refers to a list, it has to be the last segment in a path.
type MetadataMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*MetadataMatcher_PathSegment_Key
	Segment              isMetadataMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *MetadataMatcher_PathSegment) Reset()         { *m = MetadataMatcher_PathSegment{} }
func (m *MetadataMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*MetadataMatcher_PathSegment) ProtoMessage()    {}
func (*MetadataMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_d306cd29ee202178, []int{0, 0}
}
func (m *MetadataMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMatcher_PathSegment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMatcher_PathSegment.Merge(m, src)
}
func (m *MetadataMatcher_PathSegment) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMatcher_PathSegment proto.InternalMessageInfo

type isMetadataMatcher_PathSegment_Segment interface {
	isMetadataMatcher_PathSegment_Segment()
	MarshalTo([]byte) (int, error)
	Size() int
}

type MetadataMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*MetadataMatcher_PathSegment_Key) isMetadataMatcher_PathSegment_Segment() {}

func (m *MetadataMatcher_PathSegment) GetSegment() isMetadataMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *MetadataMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*MetadataMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetadataMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetadataMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*MetadataMatcher)(nil), "envoy.type.matcher.v3alpha.MetadataMatcher")
	proto.RegisterType((*MetadataMatcher_PathSegment)(nil), "envoy.type.matcher.v3alpha.MetadataMatcher.PathSegment")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v3alpha/metadata.proto", fileDescriptor_d306cd29ee202178)
}

var fileDescriptor_d306cd29ee202178 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x02, 0x2b, 0xd5, 0x03, 0x29, 0xd5, 0x83, 0x2a, 0xd5,
	0x83, 0x2a, 0x95, 0x52, 0xc3, 0x63, 0x4c, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0xc4, 0x0c, 0x29, 0xf1,
	0xb2, 0xc4, 0x9c, 0xcc, 0x94, 0xc4, 0x92, 0x54, 0x7d, 0x18, 0x03, 0x22, 0xa1, 0x34, 0x93, 0x89,
	0x8b, 0xdf, 0x17, 0x6a, 0x9f, 0x2f, 0xc4, 0x00, 0x21, 0x45, 0x2e, 0xb6, 0xb4, 0xcc, 0x9c, 0x92,
	0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xce, 0x5d, 0x2f, 0x0f, 0x30, 0xb3, 0x14,
	0x31, 0x29, 0x30, 0x06, 0x41, 0x25, 0x84, 0xc2, 0xb9, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x24, 0x98,
	0x14, 0x98, 0x35, 0xb8, 0x8d, 0xcc, 0xf5, 0x70, 0x3b, 0x51, 0x0f, 0xcd, 0x74, 0xbd, 0x80, 0xc4,
	0x92, 0x8c, 0xe0, 0xd4, 0xf4, 0xdc, 0xd4, 0xbc, 0x12, 0x27, 0x2e, 0x90, 0xc9, 0xac, 0x93, 0x18,
	0x99, 0x38, 0x18, 0x83, 0xc0, 0x06, 0x0a, 0x79, 0x71, 0xb1, 0x82, 0xdd, 0x2d, 0xc1, 0xac, 0xc0,
	0xa8, 0xc1, 0x6d, 0xa4, 0x81, 0xcf, 0xe4, 0x30, 0x90, 0x42, 0xa8, 0xb1, 0x50, 0xa3, 0xba, 0x18,
	0x99, 0x04, 0x18, 0x83, 0x20, 0x46, 0x48, 0xd9, 0x71, 0x71, 0x23, 0x59, 0x26, 0x24, 0xcb, 0xc5,
	0x9c, 0x9d, 0x5a, 0x89, 0xe1, 0x27, 0x0f, 0x86, 0x20, 0x90, 0xb8, 0x93, 0x00, 0x17, 0x7b, 0x31,
	0x54, 0x25, 0xeb, 0x8e, 0x97, 0x07, 0x98, 0x19, 0x9d, 0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0x46, 0x2e, 0x8d, 0xcc, 0x7c, 0x88, 0x63, 0x0a, 0x8a, 0xf2,
	0x2b, 0x2a, 0xf1, 0xb8, 0xcb, 0x89, 0x17, 0xe6, 0xe5, 0x00, 0x50, 0x10, 0x07, 0x30, 0x26, 0xb1,
	0x81, 0xc3, 0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc5, 0xa3, 0x12, 0xf5, 0x01, 0x00,
	0x00,
}

func (m *MetadataMatcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMatcher) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataMatcher) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Value != nil {
		{
			size, err := m.Value.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMetadata(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Path) > 0 {
		for iNdEx := len(m.Path) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Path[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMetadata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Filter) > 0 {
		i -= len(m.Filter)
		copy(dAtA[i:], m.Filter)
		i = encodeVarintMetadata(dAtA, i, uint64(len(m.Filter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetadataMatcher_PathSegment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMatcher_PathSegment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataMatcher_PathSegment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Segment != nil {
		{
			size := m.Segment.Size()
			i -= size
			if _, err := m.Segment.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *MetadataMatcher_PathSegment_Key) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *MetadataMatcher_PathSegment_Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Key)
	copy(dAtA[i:], m.Key)
	i = encodeVarintMetadata(dAtA, i, uint64(len(m.Key)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}
func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetadataMatcher) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Filter)
	if l > 0 {
		n += 1 + l + sovMetadata(uint64(l))
	}
	if len(m.Path) > 0 {
		for _, e := range m.Path {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovMetadata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMatcher_PathSegment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Segment != nil {
		n += m.Segment.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMatcher_PathSegment_Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func sovMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetadataMatcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: MetadataMatcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataMatcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = append(m.Path, &MetadataMatcher_PathSegment{})
			if err := m.Path[len(m.Path)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &ValueMatcher{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func (m *MetadataMatcher_PathSegment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
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
			return fmt.Errorf("proto: PathSegment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PathSegment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
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
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Segment = &MetadataMatcher_PathSegment_Key{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetadata
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
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
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
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
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
				return 0, ErrInvalidLengthMetadata
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetadata
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMetadata
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata   = fmt.Errorf("proto: integer overflow")
)