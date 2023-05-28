// Code generated with goa v3.11.3, DO NOT EDIT.
//
// todo protocol buffer definition
//
// Command:
// $ goa gen todo/design

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: goadesign_goagen_todo.proto

package todopb

import (
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

type CreateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of task
	Name *string `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_todo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTaskRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type CreateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of task
	Id *string `protobuf:"bytes,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// name of task
	Name *string `protobuf:"bytes,11,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// status of task
	Completed *bool `protobuf:"varint,12,opt,name=completed,proto3,oneof" json:"completed,omitempty"`
}

func (x *CreateTaskResponse) Reset() {
	*x = CreateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskResponse) ProtoMessage() {}

func (x *CreateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskResponse.ProtoReflect.Descriptor instead.
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *CreateTaskResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CreateTaskResponse) GetCompleted() bool {
	if x != nil && x.Completed != nil {
		return *x.Completed
	}
	return false
}

type CompleteTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of task
	Id *string `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *CompleteTaskRequest) Reset() {
	*x = CompleteTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_todo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskRequest) ProtoMessage() {}

func (x *CompleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_todo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskRequest.ProtoReflect.Descriptor instead.
func (*CompleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_todo_proto_rawDescGZIP(), []int{2}
}

func (x *CompleteTaskRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type CompleteTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of task
	Id *string `protobuf:"bytes,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// name of task
	Name *string `protobuf:"bytes,11,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// status of task
	Completed *bool `protobuf:"varint,12,opt,name=completed,proto3,oneof" json:"completed,omitempty"`
}

func (x *CompleteTaskResponse) Reset() {
	*x = CompleteTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_todo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteTaskResponse) ProtoMessage() {}

func (x *CompleteTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_todo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteTaskResponse.ProtoReflect.Descriptor instead.
func (*CompleteTaskResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_todo_proto_rawDescGZIP(), []int{3}
}

func (x *CompleteTaskResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *CompleteTaskResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *CompleteTaskResponse) GetCompleted() bool {
	if x != nil && x.Completed != nil {
		return *x.Completed
	}
	return false
}

type RevertTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of task
	Id *string `protobuf:"bytes,3,opt,name=id,proto3,oneof" json:"id,omitempty"`
}

func (x *RevertTaskRequest) Reset() {
	*x = RevertTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_todo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevertTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertTaskRequest) ProtoMessage() {}

func (x *RevertTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_todo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertTaskRequest.ProtoReflect.Descriptor instead.
func (*RevertTaskRequest) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_todo_proto_rawDescGZIP(), []int{4}
}

func (x *RevertTaskRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type RevertTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id of task
	Id *string `protobuf:"bytes,10,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// name of task
	Name *string `protobuf:"bytes,11,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// status of task
	Completed *bool `protobuf:"varint,12,opt,name=completed,proto3,oneof" json:"completed,omitempty"`
}

func (x *RevertTaskResponse) Reset() {
	*x = RevertTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goadesign_goagen_todo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevertTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevertTaskResponse) ProtoMessage() {}

func (x *RevertTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goadesign_goagen_todo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevertTaskResponse.ProtoReflect.Descriptor instead.
func (*RevertTaskResponse) Descriptor() ([]byte, []int) {
	return file_goadesign_goagen_todo_proto_rawDescGZIP(), []int{5}
}

func (x *RevertTaskResponse) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *RevertTaskResponse) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RevertTaskResponse) GetCompleted() bool {
	if x != nil && x.Completed != nil {
		return *x.Completed
	}
	return false
}

var File_goadesign_goagen_todo_proto protoreflect.FileDescriptor

var file_goadesign_goagen_todo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x6f, 0x61, 0x64, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x67, 0x6f, 0x61, 0x67,
	0x65, 0x6e, 0x5f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x22, 0x35, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x22, 0x31, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03,
	0x5f, 0x69, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x2f, 0x0a, 0x11, 0x52,
	0x65, 0x76, 0x65, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0x83, 0x01, 0x0a,
	0x12, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x21, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x32, 0xcf, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x3f, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x19, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x52, 0x65, 0x76, 0x65, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goadesign_goagen_todo_proto_rawDescOnce sync.Once
	file_goadesign_goagen_todo_proto_rawDescData = file_goadesign_goagen_todo_proto_rawDesc
)

func file_goadesign_goagen_todo_proto_rawDescGZIP() []byte {
	file_goadesign_goagen_todo_proto_rawDescOnce.Do(func() {
		file_goadesign_goagen_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_goadesign_goagen_todo_proto_rawDescData)
	})
	return file_goadesign_goagen_todo_proto_rawDescData
}

var file_goadesign_goagen_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_goadesign_goagen_todo_proto_goTypes = []interface{}{
	(*CreateTaskRequest)(nil),    // 0: todo.CreateTaskRequest
	(*CreateTaskResponse)(nil),   // 1: todo.CreateTaskResponse
	(*CompleteTaskRequest)(nil),  // 2: todo.CompleteTaskRequest
	(*CompleteTaskResponse)(nil), // 3: todo.CompleteTaskResponse
	(*RevertTaskRequest)(nil),    // 4: todo.RevertTaskRequest
	(*RevertTaskResponse)(nil),   // 5: todo.RevertTaskResponse
}
var file_goadesign_goagen_todo_proto_depIdxs = []int32{
	0, // 0: todo.Todo.CreateTask:input_type -> todo.CreateTaskRequest
	2, // 1: todo.Todo.CompleteTask:input_type -> todo.CompleteTaskRequest
	4, // 2: todo.Todo.RevertTask:input_type -> todo.RevertTaskRequest
	1, // 3: todo.Todo.CreateTask:output_type -> todo.CreateTaskResponse
	3, // 4: todo.Todo.CompleteTask:output_type -> todo.CompleteTaskResponse
	5, // 5: todo.Todo.RevertTask:output_type -> todo.RevertTaskResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_goadesign_goagen_todo_proto_init() }
func file_goadesign_goagen_todo_proto_init() {
	if File_goadesign_goagen_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goadesign_goagen_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRequest); i {
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
		file_goadesign_goagen_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskResponse); i {
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
		file_goadesign_goagen_todo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskRequest); i {
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
		file_goadesign_goagen_todo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteTaskResponse); i {
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
		file_goadesign_goagen_todo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevertTaskRequest); i {
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
		file_goadesign_goagen_todo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevertTaskResponse); i {
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
	file_goadesign_goagen_todo_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_goadesign_goagen_todo_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_goadesign_goagen_todo_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_goadesign_goagen_todo_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_goadesign_goagen_todo_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_goadesign_goagen_todo_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_goadesign_goagen_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goadesign_goagen_todo_proto_goTypes,
		DependencyIndexes: file_goadesign_goagen_todo_proto_depIdxs,
		MessageInfos:      file_goadesign_goagen_todo_proto_msgTypes,
	}.Build()
	File_goadesign_goagen_todo_proto = out.File
	file_goadesign_goagen_todo_proto_rawDesc = nil
	file_goadesign_goagen_todo_proto_goTypes = nil
	file_goadesign_goagen_todo_proto_depIdxs = nil
}
