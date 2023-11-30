// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: grpc/proto/manager.proto

package proto

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerName   string `protobuf:"bytes,1,opt,name=ser_name,json=serName,proto3" json:"ser_name,omitempty"`
	Active    string `protobuf:"bytes,2,opt,name=active,proto3" json:"active,omitempty"` // start,stop,restart,status
	Directory string `protobuf:"bytes,3,opt,name=directory,proto3" json:"directory,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetSerName() string {
	if x != nil {
		return x.SerName
	}
	return ""
}

func (x *Request) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

func (x *Request) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EmqxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerName string `protobuf:"bytes,1,opt,name=ser_name,json=serName,proto3" json:"ser_name,omitempty"`
	Active  string `protobuf:"bytes,2,opt,name=active,proto3" json:"active,omitempty"` // start,stop,restart,status
}

func (x *EmqxRequest) Reset() {
	*x = EmqxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmqxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmqxRequest) ProtoMessage() {}

func (x *EmqxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmqxRequest.ProtoReflect.Descriptor instead.
func (*EmqxRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{2}
}

func (x *EmqxRequest) GetSerName() string {
	if x != nil {
		return x.SerName
	}
	return ""
}

func (x *EmqxRequest) GetActive() string {
	if x != nil {
		return x.Active
	}
	return ""
}

type EmqxResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Data    []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EmqxResponse) Reset() {
	*x = EmqxResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmqxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmqxResponse) ProtoMessage() {}

func (x *EmqxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmqxResponse.ProtoReflect.Descriptor instead.
func (*EmqxResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{3}
}

func (x *EmqxResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *EmqxResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *EmqxResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk     []byte `protobuf:"bytes,1,opt,name=file_chunk,json=fileChunk,proto3" json:"file_chunk,omitempty"`
	FileName      string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileType      string `protobuf:"bytes,3,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	SerName       string `protobuf:"bytes,4,opt,name=ser_name,json=serName,proto3" json:"ser_name,omitempty"`
	DirectionType string `protobuf:"bytes,5,opt,name=direction_type,json=directionType,proto3" json:"direction_type,omitempty"`
	FileMd5       string `protobuf:"bytes,6,opt,name=file_md5,json=fileMd5,proto3" json:"file_md5,omitempty"`
	SerPath       string `protobuf:"bytes,7,opt,name=ser_path,json=serPath,proto3" json:"ser_path,omitempty"`
}

func (x *UploadRequest) Reset() {
	*x = UploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadRequest) ProtoMessage() {}

func (x *UploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadRequest.ProtoReflect.Descriptor instead.
func (*UploadRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{4}
}

func (x *UploadRequest) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *UploadRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *UploadRequest) GetSerName() string {
	if x != nil {
		return x.SerName
	}
	return ""
}

func (x *UploadRequest) GetDirectionType() string {
	if x != nil {
		return x.DirectionType
	}
	return ""
}

func (x *UploadRequest) GetFileMd5() string {
	if x != nil {
		return x.FileMd5
	}
	return ""
}

func (x *UploadRequest) GetSerPath() string {
	if x != nil {
		return x.SerPath
	}
	return ""
}

type UploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UploadResponse) Reset() {
	*x = UploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResponse) ProtoMessage() {}

func (x *UploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResponse.ProtoReflect.Descriptor instead.
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{5}
}

func (x *UploadResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UploadResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{6}
}

type ServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32                      `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message  string                     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Services []*ServiceResponse_Service `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{7}
}

func (x *ServiceResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ServiceResponse) GetServices() []*ServiceResponse_Service {
	if x != nil {
		return x.Services
	}
	return nil
}

type ServiceResponse_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SerName        string `protobuf:"bytes,1,opt,name=ser_name,json=serName,proto3" json:"ser_name,omitempty"`
	SerUrl         string `protobuf:"bytes,2,opt,name=ser_url,json=serUrl,proto3" json:"ser_url,omitempty"`
	SerNameDisplay string `protobuf:"bytes,3,opt,name=ser_name_display,json=serNameDisplay,proto3" json:"ser_name_display,omitempty"`
	SerStatus      string `protobuf:"bytes,4,opt,name=ser_status,json=serStatus,proto3" json:"ser_status,omitempty"`
	Directory      string `protobuf:"bytes,5,opt,name=directory,proto3" json:"directory,omitempty"`
	Enable         bool   `protobuf:"varint,6,opt,name=enable,proto3" json:"enable,omitempty"`
	Upgrade        bool   `protobuf:"varint,7,opt,name=upgrade,proto3" json:"upgrade,omitempty"`
}

func (x *ServiceResponse_Service) Reset() {
	*x = ServiceResponse_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceResponse_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse_Service) ProtoMessage() {}

func (x *ServiceResponse_Service) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse_Service.ProtoReflect.Descriptor instead.
func (*ServiceResponse_Service) Descriptor() ([]byte, []int) {
	return file_grpc_proto_manager_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ServiceResponse_Service) GetSerName() string {
	if x != nil {
		return x.SerName
	}
	return ""
}

func (x *ServiceResponse_Service) GetSerUrl() string {
	if x != nil {
		return x.SerUrl
	}
	return ""
}

func (x *ServiceResponse_Service) GetSerNameDisplay() string {
	if x != nil {
		return x.SerNameDisplay
	}
	return ""
}

func (x *ServiceResponse_Service) GetSerStatus() string {
	if x != nil {
		return x.SerStatus
	}
	return ""
}

func (x *ServiceResponse_Service) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *ServiceResponse_Service) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *ServiceResponse_Service) GetUpgrade() bool {
	if x != nil {
		return x.Upgrade
	}
	return false
}

var File_grpc_proto_manager_proto protoreflect.FileDescriptor

var file_grpc_proto_manager_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x50, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x40, 0x0a, 0x0b, 0x45, 0x6d, 0x71, 0x78,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x54, 0x0a, 0x0c, 0x45, 0x6d,
	0x71, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0xe0, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e,
	0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x64, 0x35, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x74, 0x68, 0x22, 0x42, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xd2, 0x02, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x34, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0xd6, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x32, 0xcc,
	0x01, 0x0a, 0x0b, 0x47, 0x72, 0x70, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x29,
	0x0a, 0x12, 0x45, 0x78, 0x65, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x12, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x0f, 0x45, 0x78, 0x65,
	0x63, 0x45, 0x6d, 0x71, 0x78, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x0c, 0x2e, 0x45,
	0x6d, 0x71, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x45, 0x6d, 0x71,
	0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x2f, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_manager_proto_rawDescOnce sync.Once
	file_grpc_proto_manager_proto_rawDescData = file_grpc_proto_manager_proto_rawDesc
)

func file_grpc_proto_manager_proto_rawDescGZIP() []byte {
	file_grpc_proto_manager_proto_rawDescOnce.Do(func() {
		file_grpc_proto_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_manager_proto_rawDescData)
	})
	return file_grpc_proto_manager_proto_rawDescData
}

var file_grpc_proto_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grpc_proto_manager_proto_goTypes = []interface{}{
	(*Request)(nil),                 // 0: Request
	(*Response)(nil),                // 1: Response
	(*EmqxRequest)(nil),             // 2: EmqxRequest
	(*EmqxResponse)(nil),            // 3: EmqxResponse
	(*UploadRequest)(nil),           // 4: UploadRequest
	(*UploadResponse)(nil),          // 5: UploadResponse
	(*ServiceRequest)(nil),          // 6: ServiceRequest
	(*ServiceResponse)(nil),         // 7: ServiceResponse
	(*ServiceResponse_Service)(nil), // 8: ServiceResponse.Service
}
var file_grpc_proto_manager_proto_depIdxs = []int32{
	8, // 0: ServiceResponse.services:type_name -> ServiceResponse.Service
	0, // 1: GrpcManager.ExecServiceControl:input_type -> Request
	2, // 2: GrpcManager.ExecEmqxControl:input_type -> EmqxRequest
	4, // 3: GrpcManager.UploadFile:input_type -> UploadRequest
	6, // 4: GrpcManager.GetService:input_type -> ServiceRequest
	1, // 5: GrpcManager.ExecServiceControl:output_type -> Response
	3, // 6: GrpcManager.ExecEmqxControl:output_type -> EmqxResponse
	5, // 7: GrpcManager.UploadFile:output_type -> UploadResponse
	7, // 8: GrpcManager.GetService:output_type -> ServiceResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_manager_proto_init() }
func file_grpc_proto_manager_proto_init() {
	if File_grpc_proto_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_grpc_proto_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_grpc_proto_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmqxRequest); i {
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
		file_grpc_proto_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmqxResponse); i {
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
		file_grpc_proto_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadRequest); i {
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
		file_grpc_proto_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResponse); i {
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
		file_grpc_proto_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceRequest); i {
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
		file_grpc_proto_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse); i {
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
		file_grpc_proto_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceResponse_Service); i {
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
			RawDescriptor: file_grpc_proto_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_manager_proto_goTypes,
		DependencyIndexes: file_grpc_proto_manager_proto_depIdxs,
		MessageInfos:      file_grpc_proto_manager_proto_msgTypes,
	}.Build()
	File_grpc_proto_manager_proto = out.File
	file_grpc_proto_manager_proto_rawDesc = nil
	file_grpc_proto_manager_proto_goTypes = nil
	file_grpc_proto_manager_proto_depIdxs = nil
}
