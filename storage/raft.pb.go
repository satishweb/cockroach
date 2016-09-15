// Code generated by protoc-gen-gogo.
// source: cockroach/storage/raft.proto
// DO NOT EDIT!

package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb3 "github.com/cockroachdb/cockroach/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"
import raftpb "github.com/coreos/etcd/raft/raftpb"

import github_com_cockroachdb_cockroach_roachpb "github.com/cockroachdb/cockroach/roachpb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// RaftMessageRequest is the request used to send raft messages using our
// protobuf-based RPC codec.
type RaftMessageRequest struct {
	RangeID     github_com_cockroachdb_cockroach_roachpb.RangeID `protobuf:"varint,1,opt,name=range_id,json=rangeId,casttype=github.com/cockroachdb/cockroach/roachpb.RangeID" json:"range_id"`
	FromReplica cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,2,opt,name=from_replica,json=fromReplica" json:"from_replica"`
	ToReplica   cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,3,opt,name=to_replica,json=toReplica" json:"to_replica"`
	Message     raftpb.Message                                   `protobuf:"bytes,4,opt,name=message" json:"message"`
	// Is this a quiesce request? A quiesce request is a MsgHeartbeat
	// which is requesting the recipient to stop ticking its local
	// replica as long as the current Raft state matches the heartbeat
	// Term/Commit. If the Term/Commit match, the recipient is marked as
	// quiescent. If they don't match, the message is passed along to
	// Raft which will generate a MsgHeartbeatResp that will unquiesce
	// the sender.
	Quiesce bool `protobuf:"varint,5,opt,name=quiesce" json:"quiesce"`
}

func (m *RaftMessageRequest) Reset()                    { *m = RaftMessageRequest{} }
func (m *RaftMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*RaftMessageRequest) ProtoMessage()               {}
func (*RaftMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{0} }

type RaftMessageResponseUnion struct {
	Error *cockroach_roachpb3.Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *RaftMessageResponseUnion) Reset()                    { *m = RaftMessageResponseUnion{} }
func (m *RaftMessageResponseUnion) String() string            { return proto.CompactTextString(m) }
func (*RaftMessageResponseUnion) ProtoMessage()               {}
func (*RaftMessageResponseUnion) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{1} }

// RaftMessageResponse may be sent to the sender of a
// RaftMessageRequest. RaftMessage does not use the usual
// request/response pattern; it is primarily modeled as a one-way
// stream of requests. Normal 'responses' are usually sent as new
// requests on a separate stream in the other direction.
// RaftMessageResponse is not sent for every RaftMessageRequest, but
// may be used for certain error conditions.
type RaftMessageResponse struct {
	RangeID     github_com_cockroachdb_cockroach_roachpb.RangeID `protobuf:"varint,1,opt,name=range_id,json=rangeId,casttype=github.com/cockroachdb/cockroach/roachpb.RangeID" json:"range_id"`
	FromReplica cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,2,opt,name=from_replica,json=fromReplica" json:"from_replica"`
	ToReplica   cockroach_roachpb.ReplicaDescriptor              `protobuf:"bytes,3,opt,name=to_replica,json=toReplica" json:"to_replica"`
	Union       RaftMessageResponseUnion                         `protobuf:"bytes,4,opt,name=union" json:"union"`
}

func (m *RaftMessageResponse) Reset()                    { *m = RaftMessageResponse{} }
func (m *RaftMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*RaftMessageResponse) ProtoMessage()               {}
func (*RaftMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{2} }

// ConfChangeContext is encoded in the raftpb.ConfChange.Context field.
type ConfChangeContext struct {
	CommandID string `protobuf:"bytes,1,opt,name=command_id,json=commandId" json:"command_id"`
	// Payload is the application-level command (i.e. an encoded
	// roachpb.EndTransactionRequest).
	Payload []byte `protobuf:"bytes,2,opt,name=payload" json:"payload,omitempty"`
	// Replica contains full details about the replica being added or removed.
	Replica cockroach_roachpb.ReplicaDescriptor `protobuf:"bytes,3,opt,name=replica" json:"replica"`
}

func (m *ConfChangeContext) Reset()                    { *m = ConfChangeContext{} }
func (m *ConfChangeContext) String() string            { return proto.CompactTextString(m) }
func (*ConfChangeContext) ProtoMessage()               {}
func (*ConfChangeContext) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{3} }

func init() {
	proto.RegisterType((*RaftMessageRequest)(nil), "cockroach.storage.RaftMessageRequest")
	proto.RegisterType((*RaftMessageResponseUnion)(nil), "cockroach.storage.RaftMessageResponseUnion")
	proto.RegisterType((*RaftMessageResponse)(nil), "cockroach.storage.RaftMessageResponse")
	proto.RegisterType((*ConfChangeContext)(nil), "cockroach.storage.ConfChangeContext")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MultiRaft service

type MultiRaftClient interface {
	RaftMessage(ctx context.Context, opts ...grpc.CallOption) (MultiRaft_RaftMessageClient, error)
	RaftMessageSync(ctx context.Context, in *RaftMessageRequest, opts ...grpc.CallOption) (*RaftMessageResponse, error)
}

type multiRaftClient struct {
	cc *grpc.ClientConn
}

func NewMultiRaftClient(cc *grpc.ClientConn) MultiRaftClient {
	return &multiRaftClient{cc}
}

func (c *multiRaftClient) RaftMessage(ctx context.Context, opts ...grpc.CallOption) (MultiRaft_RaftMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_MultiRaft_serviceDesc.Streams[0], c.cc, "/cockroach.storage.MultiRaft/RaftMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiRaftRaftMessageClient{stream}
	return x, nil
}

type MultiRaft_RaftMessageClient interface {
	Send(*RaftMessageRequest) error
	Recv() (*RaftMessageResponse, error)
	grpc.ClientStream
}

type multiRaftRaftMessageClient struct {
	grpc.ClientStream
}

func (x *multiRaftRaftMessageClient) Send(m *RaftMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiRaftRaftMessageClient) Recv() (*RaftMessageResponse, error) {
	m := new(RaftMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *multiRaftClient) RaftMessageSync(ctx context.Context, in *RaftMessageRequest, opts ...grpc.CallOption) (*RaftMessageResponse, error) {
	out := new(RaftMessageResponse)
	err := grpc.Invoke(ctx, "/cockroach.storage.MultiRaft/RaftMessageSync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MultiRaft service

type MultiRaftServer interface {
	RaftMessage(MultiRaft_RaftMessageServer) error
	RaftMessageSync(context.Context, *RaftMessageRequest) (*RaftMessageResponse, error)
}

func RegisterMultiRaftServer(s *grpc.Server, srv MultiRaftServer) {
	s.RegisterService(&_MultiRaft_serviceDesc, srv)
}

func _MultiRaft_RaftMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiRaftServer).RaftMessage(&multiRaftRaftMessageServer{stream})
}

type MultiRaft_RaftMessageServer interface {
	Send(*RaftMessageResponse) error
	Recv() (*RaftMessageRequest, error)
	grpc.ServerStream
}

type multiRaftRaftMessageServer struct {
	grpc.ServerStream
}

func (x *multiRaftRaftMessageServer) Send(m *RaftMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiRaftRaftMessageServer) Recv() (*RaftMessageRequest, error) {
	m := new(RaftMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MultiRaft_RaftMessageSync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RaftMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiRaftServer).RaftMessageSync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.storage.MultiRaft/RaftMessageSync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiRaftServer).RaftMessageSync(ctx, req.(*RaftMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MultiRaft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.storage.MultiRaft",
	HandlerType: (*MultiRaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RaftMessageSync",
			Handler:    _MultiRaft_RaftMessageSync_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RaftMessage",
			Handler:       _MultiRaft_RaftMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptorRaft,
}

func (m *RaftMessageRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRaft(data, i, uint64(m.RangeID))
	data[i] = 0x12
	i++
	i = encodeVarintRaft(data, i, uint64(m.FromReplica.Size()))
	n1, err := m.FromReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.ToReplica.Size()))
	n2, err := m.ToReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x22
	i++
	i = encodeVarintRaft(data, i, uint64(m.Message.Size()))
	n3, err := m.Message.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	data[i] = 0x28
	i++
	if m.Quiesce {
		data[i] = 1
	} else {
		data[i] = 0
	}
	i++
	return i, nil
}

func (m *RaftMessageResponseUnion) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageResponseUnion) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintRaft(data, i, uint64(m.Error.Size()))
		n4, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *RaftMessageResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *RaftMessageResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintRaft(data, i, uint64(m.RangeID))
	data[i] = 0x12
	i++
	i = encodeVarintRaft(data, i, uint64(m.FromReplica.Size()))
	n5, err := m.FromReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.ToReplica.Size()))
	n6, err := m.ToReplica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	data[i] = 0x22
	i++
	i = encodeVarintRaft(data, i, uint64(m.Union.Size()))
	n7, err := m.Union.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	return i, nil
}

func (m *ConfChangeContext) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ConfChangeContext) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintRaft(data, i, uint64(len(m.CommandID)))
	i += copy(data[i:], m.CommandID)
	if m.Payload != nil {
		data[i] = 0x12
		i++
		i = encodeVarintRaft(data, i, uint64(len(m.Payload)))
		i += copy(data[i:], m.Payload)
	}
	data[i] = 0x1a
	i++
	i = encodeVarintRaft(data, i, uint64(m.Replica.Size()))
	n8, err := m.Replica.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeFixed64Raft(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Raft(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaft(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *RaftMessageRequest) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.RangeID))
	l = m.FromReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.ToReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovRaft(uint64(l))
	n += 2
	return n
}

func (m *RaftMessageResponseUnion) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func (m *RaftMessageResponse) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovRaft(uint64(m.RangeID))
	l = m.FromReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.ToReplica.Size()
	n += 1 + l + sovRaft(uint64(l))
	l = m.Union.Size()
	n += 1 + l + sovRaft(uint64(l))
	return n
}

func (m *ConfChangeContext) Size() (n int) {
	var l int
	_ = l
	l = len(m.CommandID)
	n += 1 + l + sovRaft(uint64(l))
	if m.Payload != nil {
		l = len(m.Payload)
		n += 1 + l + sovRaft(uint64(l))
	}
	l = m.Replica.Size()
	n += 1 + l + sovRaft(uint64(l))
	return n
}

func sovRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RaftMessageResponseUnion) GetValue() interface{} {
	if this.Error != nil {
		return this.Error
	}
	return nil
}

func (this *RaftMessageResponseUnion) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *cockroach_roachpb3.Error:
		this.Error = vt
	default:
		return false
	}
	return true
}
func (m *RaftMessageRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeID |= (github_com_cockroachdb_cockroach_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quiesce", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Quiesce = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *RaftMessageResponseUnion) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMessageResponseUnion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageResponseUnion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &cockroach_roachpb3.Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *RaftMessageResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.RangeID |= (github_com_cockroachdb_cockroach_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToReplica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Union", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Union.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func (m *ConfChangeContext) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfChangeContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfChangeContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommandID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CommandID = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], data[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Replica.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
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
func skipRaft(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipRaft(data[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/storage/raft.proto", fileDescriptorRaft) }

var fileDescriptorRaft = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe4, 0x53, 0xcf, 0xaa, 0xd3, 0x4e,
	0x14, 0xee, 0xf4, 0xd7, 0x92, 0xdb, 0xd3, 0x0b, 0x97, 0xce, 0xcf, 0x45, 0xa8, 0x92, 0xd6, 0xa2,
	0x52, 0x10, 0x26, 0xa5, 0x4b, 0x97, 0x6d, 0x45, 0xba, 0x28, 0x48, 0xc4, 0x8d, 0x0b, 0x2f, 0x93,
	0xc9, 0x34, 0x0d, 0x36, 0x99, 0xdc, 0xc9, 0x14, 0xbc, 0x6f, 0xe1, 0x23, 0xb8, 0xf1, 0x5d, 0xba,
	0x14, 0xdc, 0xdc, 0x55, 0xd5, 0xf8, 0x16, 0xae, 0x24, 0x93, 0x49, 0x6f, 0xa5, 0x57, 0xfc, 0x83,
	0x3b, 0x37, 0xc3, 0xcc, 0xf9, 0xce, 0xf7, 0x9d, 0x73, 0xbe, 0xc3, 0xc0, 0x1d, 0x26, 0xd8, 0x2b,
	0x29, 0x28, 0x5b, 0xb9, 0x99, 0x12, 0x92, 0x86, 0xdc, 0x95, 0x74, 0xa9, 0x48, 0x2a, 0x85, 0x12,
	0xb8, 0xb3, 0x47, 0x89, 0x41, 0xbb, 0xce, 0x35, 0x41, 0x9f, 0xa9, 0xef, 0x72, 0x29, 0x85, 0xcc,
	0x4a, 0x4a, 0xb7, 0x7f, 0x8c, 0xc7, 0x5c, 0xd1, 0x80, 0x2a, 0x6a, 0x32, 0x6e, 0x73, 0xc5, 0x02,
	0x5d, 0x45, 0x1f, 0xa9, 0x7f, 0x50, 0xb1, 0x7b, 0x2b, 0x14, 0xa1, 0xd0, 0x57, 0xb7, 0xb8, 0x95,
	0xd1, 0xc1, 0xc7, 0x3a, 0x60, 0x8f, 0x2e, 0xd5, 0x82, 0x67, 0x19, 0x0d, 0xb9, 0xc7, 0x2f, 0x36,
	0x3c, 0x53, 0xf8, 0x25, 0x9c, 0x48, 0x9a, 0x84, 0xfc, 0x3c, 0x0a, 0x6c, 0xd4, 0x47, 0xc3, 0xc6,
	0x64, 0xba, 0xdd, 0xf5, 0x6a, 0xf9, 0xae, 0x67, 0x79, 0x45, 0x7c, 0x3e, 0xfb, 0xba, 0xeb, 0x8d,
	0xc2, 0x48, 0xad, 0x36, 0x3e, 0x61, 0x22, 0x76, 0xf7, 0xbd, 0x05, 0xbe, 0x7b, 0xd4, 0x27, 0x31,
	0x1c, 0xcf, 0xd2, 0xa2, 0xf3, 0x00, 0x2f, 0xe0, 0x74, 0x29, 0x45, 0x7c, 0x2e, 0x79, 0xba, 0x8e,
	0x18, 0xb5, 0xeb, 0x7d, 0x34, 0x6c, 0x8f, 0xef, 0x91, 0x6b, 0x57, 0xf6, 0xd4, 0x32, 0x63, 0xc6,
	0x33, 0x26, 0xa3, 0x54, 0x09, 0x39, 0x69, 0x14, 0x9d, 0x78, 0xed, 0x82, 0x6f, 0x40, 0x3c, 0x07,
	0x50, 0x62, 0x2f, 0xf6, 0xdf, 0x6f, 0x8b, 0xb5, 0x94, 0xa8, 0xa4, 0x5c, 0xb0, 0xe2, 0xd2, 0x0b,
	0xbb, 0xa1, 0x75, 0xce, 0x48, 0xe9, 0x25, 0x31, 0x16, 0x19, 0x4a, 0x95, 0x85, 0x1d, 0xb0, 0x2e,
	0x36, 0x11, 0xcf, 0x18, 0xb7, 0x9b, 0x7d, 0x34, 0x3c, 0xa9, 0x70, 0x13, 0x1c, 0x3c, 0x05, 0xfb,
	0x3b, 0x83, 0xb3, 0x54, 0x24, 0x19, 0x7f, 0x9e, 0x44, 0x22, 0xc1, 0x04, 0x9a, 0x7a, 0xc5, 0xda,
	0xe3, 0xf6, 0xd8, 0xbe, 0xa1, 0xe5, 0xc7, 0x05, 0xee, 0x95, 0x69, 0x8f, 0x1a, 0xdb, 0xb7, 0x3d,
	0x34, 0xb8, 0xaa, 0xc3, 0xff, 0x37, 0x48, 0xfe, 0xc3, 0x4b, 0x7b, 0x02, 0xcd, 0x4d, 0x61, 0xa8,
	0x59, 0xd9, 0x43, 0x72, 0xf4, 0xbb, 0xc8, 0x8f, 0x76, 0x60, 0xc4, 0x4a, 0xfe, 0xe0, 0x1d, 0x82,
	0xce, 0x54, 0x24, 0xcb, 0xe9, 0xaa, 0x98, 0x79, 0x2a, 0x12, 0xc5, 0x5f, 0x2b, 0x3c, 0x02, 0x60,
	0x22, 0x8e, 0x69, 0x12, 0x54, 0xd6, 0xb6, 0x26, 0x1d, 0x63, 0x6d, 0x6b, 0x5a, 0x22, 0xf3, 0x99,
	0xd7, 0x32, 0x49, 0xf3, 0x00, 0xdb, 0x60, 0xa5, 0xf4, 0x72, 0x2d, 0x68, 0xa0, 0x5d, 0x3a, 0xf5,
	0xaa, 0x27, 0x9e, 0x81, 0xf5, 0xe7, 0x23, 0x57, 0xd4, 0xf1, 0x07, 0x04, 0xad, 0xc5, 0x66, 0xad,
	0xa2, 0x62, 0x2c, 0xec, 0x43, 0xfb, 0x60, 0x3c, 0x7c, 0xff, 0x67, 0xe3, 0xeb, 0x3f, 0xde, 0x7d,
	0xf0, 0x6b, 0x2e, 0x0d, 0x6a, 0x43, 0x34, 0x42, 0xd8, 0x87, 0xb3, 0x03, 0xf0, 0xd9, 0x65, 0xc2,
	0xfe, 0x7a, 0x9d, 0xc9, 0xdd, 0xed, 0x67, 0xa7, 0xb6, 0xcd, 0x1d, 0xf4, 0x3e, 0x77, 0xd0, 0x55,
	0xee, 0xa0, 0x4f, 0xb9, 0x83, 0xde, 0x7c, 0x71, 0x6a, 0x2f, 0x2c, 0x43, 0xfc, 0x16, 0x00, 0x00,
	0xff, 0xff, 0x20, 0xa3, 0x0d, 0x1b, 0x56, 0x05, 0x00, 0x00,
}
