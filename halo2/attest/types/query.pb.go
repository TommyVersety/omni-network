// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: halo2/attest/types/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ApprovedFromRequest queries halo for approved aggregate attestations for the given chain_id
// and from the given height (inclusive). The response will contain at most max attestations sequentially
// following from_height.
type ApprovedFromRequest struct {
	ChainId    uint64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	FromHeight uint64 `protobuf:"varint,2,opt,name=from_height,json=fromHeight,proto3" json:"from_height,omitempty"`
}

func (m *ApprovedFromRequest) Reset()         { *m = ApprovedFromRequest{} }
func (m *ApprovedFromRequest) String() string { return proto.CompactTextString(m) }
func (*ApprovedFromRequest) ProtoMessage()    {}
func (*ApprovedFromRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_95af3dbf1e28ff30, []int{0}
}
func (m *ApprovedFromRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApprovedFromRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApprovedFromRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApprovedFromRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedFromRequest.Merge(m, src)
}
func (m *ApprovedFromRequest) XXX_Size() int {
	return m.Size()
}
func (m *ApprovedFromRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedFromRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedFromRequest proto.InternalMessageInfo

func (m *ApprovedFromRequest) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ApprovedFromRequest) GetFromHeight() uint64 {
	if m != nil {
		return m.FromHeight
	}
	return 0
}

// ApprovedFromRequest queries halo for approved aggregate attestations for the given chain_id
// and from the given height (inclusive). The response will contain at most max attestations.
type ApprovedFromResponse struct {
	Aggregates []*MsgAggAttestation `protobuf:"bytes,1,rep,name=aggregates,proto3" json:"aggregates,omitempty"`
}

func (m *ApprovedFromResponse) Reset()         { *m = ApprovedFromResponse{} }
func (m *ApprovedFromResponse) String() string { return proto.CompactTextString(m) }
func (*ApprovedFromResponse) ProtoMessage()    {}
func (*ApprovedFromResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_95af3dbf1e28ff30, []int{1}
}
func (m *ApprovedFromResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApprovedFromResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApprovedFromResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApprovedFromResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedFromResponse.Merge(m, src)
}
func (m *ApprovedFromResponse) XXX_Size() int {
	return m.Size()
}
func (m *ApprovedFromResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedFromResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedFromResponse proto.InternalMessageInfo

func (m *ApprovedFromResponse) GetAggregates() []*MsgAggAttestation {
	if m != nil {
		return m.Aggregates
	}
	return nil
}

func init() {
	proto.RegisterType((*ApprovedFromRequest)(nil), "halo2.attest.types.ApprovedFromRequest")
	proto.RegisterType((*ApprovedFromResponse)(nil), "halo2.attest.types.ApprovedFromResponse")
}

func init() { proto.RegisterFile("halo2/attest/types/query.proto", fileDescriptor_95af3dbf1e28ff30) }

var fileDescriptor_95af3dbf1e28ff30 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x48, 0xcc, 0xc9,
	0x37, 0xd2, 0x4f, 0x2c, 0x29, 0x49, 0x2d, 0x2e, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f,
	0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x02, 0xcb, 0xeb, 0x41,
	0xe4, 0xf5, 0xc0, 0xf2, 0x52, 0xd2, 0x58, 0xf4, 0x94, 0x54, 0x40, 0x34, 0x28, 0x05, 0x72, 0x09,
	0x3b, 0x16, 0x14, 0x14, 0xe5, 0x97, 0xa5, 0xa6, 0xb8, 0x15, 0xe5, 0xe7, 0x06, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x08, 0x49, 0x72, 0x71, 0x24, 0x67, 0x24, 0x66, 0xe6, 0xc5, 0x67, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0xb1, 0x83, 0xf9, 0x9e, 0x29, 0x42, 0xf2, 0x5c, 0xdc, 0x69,
	0x45, 0xf9, 0xb9, 0xf1, 0x19, 0xa9, 0x99, 0xe9, 0x19, 0x25, 0x12, 0x4c, 0x60, 0x59, 0x2e, 0x90,
	0x90, 0x07, 0x58, 0x44, 0x29, 0x96, 0x4b, 0x04, 0xd5, 0xc8, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x57, 0x2e, 0xae, 0xc4, 0xf4, 0xf4, 0xa2, 0xd4, 0xf4, 0xc4, 0x92, 0xd4, 0x62, 0x09, 0x46,
	0x05, 0x66, 0x0d, 0x6e, 0x23, 0x55, 0x3d, 0x4c, 0x07, 0xeb, 0xf9, 0x16, 0xa7, 0x3b, 0xa6, 0xa7,
	0x3b, 0x82, 0x85, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x90, 0x34, 0x1a, 0xe5, 0x70, 0xb1, 0x06,
	0x82, 0x7c, 0x2c, 0x94, 0xcc, 0xc5, 0x83, 0x6c, 0x8f, 0x90, 0x3a, 0x36, 0xb3, 0xb0, 0x78, 0x4e,
	0x4a, 0x83, 0xb0, 0x42, 0x88, 0x93, 0x95, 0x18, 0x9c, 0x74, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x08, 0x33, 0x58, 0x93, 0xd8, 0xc0, 0x81, 0x6a, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xa0, 0x11, 0x3a, 0x63, 0xa7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ApprovedFrom queries halo for approved aggregate attestations for the given chain_id
	// and from the given height (inclusive). The response will contain at most max attestations sequentially
	// following from_height.
	ApprovedFrom(ctx context.Context, in *ApprovedFromRequest, opts ...grpc.CallOption) (*ApprovedFromResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ApprovedFrom(ctx context.Context, in *ApprovedFromRequest, opts ...grpc.CallOption) (*ApprovedFromResponse, error) {
	out := new(ApprovedFromResponse)
	err := c.cc.Invoke(ctx, "/halo2.attest.types.Query/ApprovedFrom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ApprovedFrom queries halo for approved aggregate attestations for the given chain_id
	// and from the given height (inclusive). The response will contain at most max attestations sequentially
	// following from_height.
	ApprovedFrom(context.Context, *ApprovedFromRequest) (*ApprovedFromResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ApprovedFrom(ctx context.Context, req *ApprovedFromRequest) (*ApprovedFromResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApprovedFrom not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ApprovedFrom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApprovedFromRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApprovedFrom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/halo2.attest.types.Query/ApprovedFrom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApprovedFrom(ctx, req.(*ApprovedFromRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "halo2.attest.types.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApprovedFrom",
			Handler:    _Query_ApprovedFrom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "halo2/attest/types/query.proto",
}

func (m *ApprovedFromRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApprovedFromRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApprovedFromRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FromHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.FromHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ApprovedFromResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApprovedFromResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApprovedFromResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Aggregates) > 0 {
		for iNdEx := len(m.Aggregates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Aggregates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ApprovedFromRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovQuery(uint64(m.ChainId))
	}
	if m.FromHeight != 0 {
		n += 1 + sovQuery(uint64(m.FromHeight))
	}
	return n
}

func (m *ApprovedFromResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Aggregates) > 0 {
		for _, e := range m.Aggregates {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ApprovedFromRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ApprovedFromRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApprovedFromRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromHeight", wireType)
			}
			m.FromHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FromHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ApprovedFromResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ApprovedFromResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApprovedFromResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aggregates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Aggregates = append(m.Aggregates, &MsgAggAttestation{})
			if err := m.Aggregates[len(m.Aggregates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)