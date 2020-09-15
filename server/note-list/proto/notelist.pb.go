// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: noteList.proto

package lovers_srv_noteList

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type NoteListUpReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	NoteListLevel string `protobuf:"bytes,2,opt,name=NoteListLevel,proto3" json:"NoteListLevel,omitempty"`
	Timestamp     string `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	NoeListData   []byte `protobuf:"bytes,4,opt,name=NoeListData,proto3" json:"NoeListData,omitempty"`
}

func (x *NoteListUpReq) Reset() {
	*x = NoteListUpReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noteList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListUpReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListUpReq) ProtoMessage() {}

func (x *NoteListUpReq) ProtoReflect() protoreflect.Message {
	mi := &file_noteList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListUpReq.ProtoReflect.Descriptor instead.
func (*NoteListUpReq) Descriptor() ([]byte, []int) {
	return file_noteList_proto_rawDescGZIP(), []int{0}
}

func (x *NoteListUpReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *NoteListUpReq) GetNoteListLevel() string {
	if x != nil {
		return x.NoteListLevel
	}
	return ""
}

func (x *NoteListUpReq) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *NoteListUpReq) GetNoeListData() []byte {
	if x != nil {
		return x.NoeListData
	}
	return nil
}

type NoteListUpResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID           string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	NoteListUpResult string `protobuf:"bytes,2,opt,name=NoteListUpResult,proto3" json:"NoteListUpResult,omitempty"`
	Err              string `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *NoteListUpResp) Reset() {
	*x = NoteListUpResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noteList_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListUpResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListUpResp) ProtoMessage() {}

func (x *NoteListUpResp) ProtoReflect() protoreflect.Message {
	mi := &file_noteList_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListUpResp.ProtoReflect.Descriptor instead.
func (*NoteListUpResp) Descriptor() ([]byte, []int) {
	return file_noteList_proto_rawDescGZIP(), []int{1}
}

func (x *NoteListUpResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *NoteListUpResp) GetNoteListUpResult() string {
	if x != nil {
		return x.NoteListUpResult
	}
	return ""
}

func (x *NoteListUpResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type NoteListDownReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID       string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Timestamp    string `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	TimestampEnd string `protobuf:"bytes,3,opt,name=TimestampEnd,proto3" json:"TimestampEnd,omitempty"`
}

func (x *NoteListDownReq) Reset() {
	*x = NoteListDownReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noteList_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListDownReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListDownReq) ProtoMessage() {}

func (x *NoteListDownReq) ProtoReflect() protoreflect.Message {
	mi := &file_noteList_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListDownReq.ProtoReflect.Descriptor instead.
func (*NoteListDownReq) Descriptor() ([]byte, []int) {
	return file_noteList_proto_rawDescGZIP(), []int{2}
}

func (x *NoteListDownReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *NoteListDownReq) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *NoteListDownReq) GetTimestampEnd() string {
	if x != nil {
		return x.TimestampEnd
	}
	return ""
}

type NoteListDownResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteListDownRet string `protobuf:"bytes,1,opt,name=NoteListDownRet,proto3" json:"NoteListDownRet,omitempty"`
	UserID          string `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Timestamp       string `protobuf:"bytes,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	NoteListLevel   string `protobuf:"bytes,4,opt,name=NoteListLevel,proto3" json:"NoteListLevel,omitempty"`
	NoteListData    []byte `protobuf:"bytes,5,opt,name=NoteListData,proto3" json:"NoteListData,omitempty"`
	Err             string `protobuf:"bytes,6,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *NoteListDownResp) Reset() {
	*x = NoteListDownResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noteList_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListDownResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListDownResp) ProtoMessage() {}

func (x *NoteListDownResp) ProtoReflect() protoreflect.Message {
	mi := &file_noteList_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListDownResp.ProtoReflect.Descriptor instead.
func (*NoteListDownResp) Descriptor() ([]byte, []int) {
	return file_noteList_proto_rawDescGZIP(), []int{3}
}

func (x *NoteListDownResp) GetNoteListDownRet() string {
	if x != nil {
		return x.NoteListDownRet
	}
	return ""
}

func (x *NoteListDownResp) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *NoteListDownResp) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *NoteListDownResp) GetNoteListLevel() string {
	if x != nil {
		return x.NoteListLevel
	}
	return ""
}

func (x *NoteListDownResp) GetNoteListData() []byte {
	if x != nil {
		return x.NoteListData
	}
	return nil
}

func (x *NoteListDownResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type NoteListDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
}

func (x *NoteListDelReq) Reset() {
	*x = NoteListDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noteList_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListDelReq) ProtoMessage() {}

func (x *NoteListDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_noteList_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListDelReq.ProtoReflect.Descriptor instead.
func (*NoteListDelReq) Descriptor() ([]byte, []int) {
	return file_noteList_proto_rawDescGZIP(), []int{4}
}

func (x *NoteListDelReq) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *NoteListDelReq) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

type NoteListDelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteListDelRet string `protobuf:"bytes,1,opt,name=NoteListDelRet,proto3" json:"NoteListDelRet,omitempty"`
	Err            string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *NoteListDelResp) Reset() {
	*x = NoteListDelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_noteList_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteListDelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteListDelResp) ProtoMessage() {}

func (x *NoteListDelResp) ProtoReflect() protoreflect.Message {
	mi := &file_noteList_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteListDelResp.ProtoReflect.Descriptor instead.
func (*NoteListDelResp) Descriptor() ([]byte, []int) {
	return file_noteList_proto_rawDescGZIP(), []int{5}
}

func (x *NoteListDelResp) GetNoteListDelRet() string {
	if x != nil {
		return x.NoteListDelRet
	}
	return ""
}

func (x *NoteListDelResp) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_noteList_proto protoreflect.FileDescriptor

var file_noteList_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x8d, 0x01, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x55, 0x70, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x24, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x6f, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x4e, 0x6f, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x66, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x2a, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4e, 0x6f, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x55, 0x70, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x6b, 0x0a,
	0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x45, 0x6e, 0x64, 0x22, 0xce, 0x01, 0x0a, 0x10, 0x4e,
	0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x28, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x24, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4e, 0x6f, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x46, 0x0a, 0x0e, 0x4e,
	0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x4b, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72,
	0x32, 0x9e, 0x02, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x57, 0x0a,
	0x0a, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x12, 0x22, 0x2e, 0x6c, 0x6f,
	0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x24, 0x2e, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x6c,
	0x6f, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x44, 0x65, 0x6c, 0x12, 0x23, 0x2e, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x73, 0x72,
	0x76, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x24, 0x2e, 0x6c, 0x6f, 0x76, 0x65,
	0x72, 0x73, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x4e, 0x6f, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_noteList_proto_rawDescOnce sync.Once
	file_noteList_proto_rawDescData = file_noteList_proto_rawDesc
)

func file_noteList_proto_rawDescGZIP() []byte {
	file_noteList_proto_rawDescOnce.Do(func() {
		file_noteList_proto_rawDescData = protoimpl.X.CompressGZIP(file_noteList_proto_rawDescData)
	})
	return file_noteList_proto_rawDescData
}

var file_noteList_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_noteList_proto_goTypes = []interface{}{
	(*NoteListUpReq)(nil),    // 0: lovers.srv.noteList.NoteListUpReq
	(*NoteListUpResp)(nil),   // 1: lovers.srv.noteList.NoteListUpResp
	(*NoteListDownReq)(nil),  // 2: lovers.srv.noteList.NoteListDownReq
	(*NoteListDownResp)(nil), // 3: lovers.srv.noteList.NoteListDownResp
	(*NoteListDelReq)(nil),   // 4: lovers.srv.noteList.NoteListDelReq
	(*NoteListDelResp)(nil),  // 5: lovers.srv.noteList.NoteListDelResp
}
var file_noteList_proto_depIdxs = []int32{
	0, // 0: lovers.srv.noteList.NoteList.NoteListUp:input_type -> lovers.srv.noteList.NoteListUpReq
	2, // 1: lovers.srv.noteList.NoteList.NoteListDown:input_type -> lovers.srv.noteList.NoteListDownReq
	4, // 2: lovers.srv.noteList.NoteList.NoteListDel:input_type -> lovers.srv.noteList.NoteListDelReq
	1, // 3: lovers.srv.noteList.NoteList.NoteListUp:output_type -> lovers.srv.noteList.NoteListUpResp
	3, // 4: lovers.srv.noteList.NoteList.NoteListDown:output_type -> lovers.srv.noteList.NoteListDownResp
	5, // 5: lovers.srv.noteList.NoteList.NoteListDel:output_type -> lovers.srv.noteList.NoteListDelResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_noteList_proto_init() }
func file_noteList_proto_init() {
	if File_noteList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_noteList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListUpReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noteList_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListUpResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noteList_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListDownReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noteList_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListDownResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noteList_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListDelReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_noteList_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoteListDelResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_noteList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_noteList_proto_goTypes,
		DependencyIndexes: file_noteList_proto_depIdxs,
		MessageInfos:      file_noteList_proto_msgTypes,
	}.Build()
	File_noteList_proto = out.File
	file_noteList_proto_rawDesc = nil
	file_noteList_proto_goTypes = nil
	file_noteList_proto_depIdxs = nil
}
