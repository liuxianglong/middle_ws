// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: api_ws.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WSRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// @inject_tag: dc:"下一页"
	Body          []byte         `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty" dc:"下一页"`
	Head          *WSRequestHead `protobuf:"bytes,2,opt,name=head,proto3" json:"head,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WSRequest) Reset() {
	*x = WSRequest{}
	mi := &file_api_ws_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSRequest) ProtoMessage() {}

func (x *WSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ws_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSRequest.ProtoReflect.Descriptor instead.
func (*WSRequest) Descriptor() ([]byte, []int) {
	return file_api_ws_proto_rawDescGZIP(), []int{0}
}

func (x *WSRequest) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *WSRequest) GetHead() *WSRequestHead {
	if x != nil {
		return x.Head
	}
	return nil
}

type WSRequestHead struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	HeadPad       []byte                 `protobuf:"bytes,1,opt,name=head_pad,json=headPad,proto3" json:"head_pad,omitempty"`
	Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`                             //消息唯一id
	Uid           int64                  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`                           //用户数字id
	SendTime      int64                  `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"` //消息发送时间，精确到毫秒
	Size          int32                  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`                         //消息内容长度
	Check         []byte                 `protobuf:"bytes,6,opt,name=check,proto3" json:"check,omitempty"`                        //校验位
	EndPad        []byte                 `protobuf:"bytes,7,opt,name=end_pad,json=endPad,proto3" json:"end_pad,omitempty"`        //消息头结束标识
	Cmd           string                 `protobuf:"bytes,8,opt,name=cmd,proto3" json:"cmd,omitempty"`                            //请求路由
	Tm            []byte                 `protobuf:"bytes,9,opt,name=tm,proto3" json:"tm,omitempty"`                              //消息内容格式 0二进制 1json串
	Rt            []byte                 `protobuf:"bytes,10,opt,name=rt,proto3" json:"rt,omitempty"`                             //请求类型 0push 1request
	Free1         []byte                 `protobuf:"bytes,11,opt,name=free1,proto3" json:"free1,omitempty"`                       //保留字段
	Free2         []byte                 `protobuf:"bytes,12,opt,name=free2,proto3" json:"free2,omitempty"`                       //保留字段
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WSRequestHead) Reset() {
	*x = WSRequestHead{}
	mi := &file_api_ws_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WSRequestHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WSRequestHead) ProtoMessage() {}

func (x *WSRequestHead) ProtoReflect() protoreflect.Message {
	mi := &file_api_ws_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WSRequestHead.ProtoReflect.Descriptor instead.
func (*WSRequestHead) Descriptor() ([]byte, []int) {
	return file_api_ws_proto_rawDescGZIP(), []int{1}
}

func (x *WSRequestHead) GetHeadPad() []byte {
	if x != nil {
		return x.HeadPad
	}
	return nil
}

func (x *WSRequestHead) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WSRequestHead) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *WSRequestHead) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *WSRequestHead) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *WSRequestHead) GetCheck() []byte {
	if x != nil {
		return x.Check
	}
	return nil
}

func (x *WSRequestHead) GetEndPad() []byte {
	if x != nil {
		return x.EndPad
	}
	return nil
}

func (x *WSRequestHead) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *WSRequestHead) GetTm() []byte {
	if x != nil {
		return x.Tm
	}
	return nil
}

func (x *WSRequestHead) GetRt() []byte {
	if x != nil {
		return x.Rt
	}
	return nil
}

func (x *WSRequestHead) GetFree1() []byte {
	if x != nil {
		return x.Free1
	}
	return nil
}

func (x *WSRequestHead) GetFree2() []byte {
	if x != nil {
		return x.Free2
	}
	return nil
}

var File_api_ws_proto protoreflect.FileDescriptor

const file_api_ws_proto_rawDesc = "" +
	"\n" +
	"\fapi_ws.proto\x12\x02pb\"F\n" +
	"\tWSRequest\x12\x12\n" +
	"\x04body\x18\x01 \x01(\fR\x04body\x12%\n" +
	"\x04head\x18\x02 \x01(\v2\x11.pb.WSRequestHeadR\x04head\"\x8a\x02\n" +
	"\rWSRequestHead\x12\x19\n" +
	"\bhead_pad\x18\x01 \x01(\fR\aheadPad\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\x03R\x02id\x12\x10\n" +
	"\x03uid\x18\x03 \x01(\x03R\x03uid\x12\x1b\n" +
	"\tsend_time\x18\x04 \x01(\x03R\bsendTime\x12\x12\n" +
	"\x04size\x18\x05 \x01(\x05R\x04size\x12\x14\n" +
	"\x05check\x18\x06 \x01(\fR\x05check\x12\x17\n" +
	"\aend_pad\x18\a \x01(\fR\x06endPad\x12\x10\n" +
	"\x03cmd\x18\b \x01(\tR\x03cmd\x12\x0e\n" +
	"\x02tm\x18\t \x01(\fR\x02tm\x12\x0e\n" +
	"\x02rt\x18\n" +
	" \x01(\fR\x02rt\x12\x14\n" +
	"\x05free1\x18\v \x01(\fR\x05free1\x12\x14\n" +
	"\x05free2\x18\f \x01(\fR\x05free2B\aZ\x05../pbb\x06proto3"

var (
	file_api_ws_proto_rawDescOnce sync.Once
	file_api_ws_proto_rawDescData []byte
)

func file_api_ws_proto_rawDescGZIP() []byte {
	file_api_ws_proto_rawDescOnce.Do(func() {
		file_api_ws_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_ws_proto_rawDesc), len(file_api_ws_proto_rawDesc)))
	})
	return file_api_ws_proto_rawDescData
}

var file_api_ws_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_ws_proto_goTypes = []any{
	(*WSRequest)(nil),     // 0: pb.WSRequest
	(*WSRequestHead)(nil), // 1: pb.WSRequestHead
}
var file_api_ws_proto_depIdxs = []int32{
	1, // 0: pb.WSRequest.head:type_name -> pb.WSRequestHead
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ws_proto_init() }
func file_api_ws_proto_init() {
	if File_api_ws_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_ws_proto_rawDesc), len(file_api_ws_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_ws_proto_goTypes,
		DependencyIndexes: file_api_ws_proto_depIdxs,
		MessageInfos:      file_api_ws_proto_msgTypes,
	}.Build()
	File_api_ws_proto = out.File
	file_api_ws_proto_goTypes = nil
	file_api_ws_proto_depIdxs = nil
}
