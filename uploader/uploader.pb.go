// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: uploader/uploader.proto

package uploader

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type FileTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *FileHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Data   []byte      `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *FileTransferRequest) Reset() {
	*x = FileTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTransferRequest) ProtoMessage() {}

func (x *FileTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTransferRequest.ProtoReflect.Descriptor instead.
func (*FileTransferRequest) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_proto_rawDescGZIP(), []int{0}
}

func (x *FileTransferRequest) GetHeader() *FileHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *FileTransferRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type FileStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Contents:
	//	*FileStreamRequest_Header
	//	*FileStreamRequest_Chunk
	Contents isFileStreamRequest_Contents `protobuf_oneof:"contents"`
}

func (x *FileStreamRequest) Reset() {
	*x = FileStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStreamRequest) ProtoMessage() {}

func (x *FileStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStreamRequest.ProtoReflect.Descriptor instead.
func (*FileStreamRequest) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_proto_rawDescGZIP(), []int{1}
}

func (m *FileStreamRequest) GetContents() isFileStreamRequest_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *FileStreamRequest) GetHeader() *FileHeader {
	if x, ok := x.GetContents().(*FileStreamRequest_Header); ok {
		return x.Header
	}
	return nil
}

func (x *FileStreamRequest) GetChunk() []byte {
	if x, ok := x.GetContents().(*FileStreamRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isFileStreamRequest_Contents interface {
	isFileStreamRequest_Contents()
}

type FileStreamRequest_Header struct {
	Header *FileHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type FileStreamRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*FileStreamRequest_Header) isFileStreamRequest_Contents() {}

func (*FileStreamRequest_Chunk) isFileStreamRequest_Contents() {}

type FileHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	FileSize *int64 `protobuf:"varint,2,opt,name=file_size,json=fileSize,proto3,oneof" json:"file_size,omitempty"`
}

func (x *FileHeader) Reset() {
	*x = FileHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uploader_uploader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileHeader) ProtoMessage() {}

func (x *FileHeader) ProtoReflect() protoreflect.Message {
	mi := &file_uploader_uploader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileHeader.ProtoReflect.Descriptor instead.
func (*FileHeader) Descriptor() ([]byte, []int) {
	return file_uploader_uploader_proto_rawDescGZIP(), []int{2}
}

func (x *FileHeader) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileHeader) GetFileSize() int64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

var File_uploader_uploader_proto protoreflect.FileDescriptor

var file_uploader_uploader_proto_rawDesc = []byte{
	0x0a, 0x17, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x57, 0x0a, 0x13, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x67, 0x0a, 0x11, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x50, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x32, 0x96, 0x01, 0x0a, 0x08, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x1b, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x19, 0x5a,
	0x17, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uploader_uploader_proto_rawDescOnce sync.Once
	file_uploader_uploader_proto_rawDescData = file_uploader_uploader_proto_rawDesc
)

func file_uploader_uploader_proto_rawDescGZIP() []byte {
	file_uploader_uploader_proto_rawDescOnce.Do(func() {
		file_uploader_uploader_proto_rawDescData = protoimpl.X.CompressGZIP(file_uploader_uploader_proto_rawDescData)
	})
	return file_uploader_uploader_proto_rawDescData
}

var file_uploader_uploader_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_uploader_uploader_proto_goTypes = []interface{}{
	(*FileTransferRequest)(nil), // 0: uploader.FileTransferRequest
	(*FileStreamRequest)(nil),   // 1: uploader.FileStreamRequest
	(*FileHeader)(nil),          // 2: uploader.FileHeader
	(*empty.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_uploader_uploader_proto_depIdxs = []int32{
	2, // 0: uploader.FileTransferRequest.header:type_name -> uploader.FileHeader
	2, // 1: uploader.FileStreamRequest.header:type_name -> uploader.FileHeader
	0, // 2: uploader.Uploader.Upload:input_type -> uploader.FileTransferRequest
	1, // 3: uploader.Uploader.UploadStream:input_type -> uploader.FileStreamRequest
	3, // 4: uploader.Uploader.Upload:output_type -> google.protobuf.Empty
	3, // 5: uploader.Uploader.UploadStream:output_type -> google.protobuf.Empty
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_uploader_uploader_proto_init() }
func file_uploader_uploader_proto_init() {
	if File_uploader_uploader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uploader_uploader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileTransferRequest); i {
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
		file_uploader_uploader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStreamRequest); i {
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
		file_uploader_uploader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileHeader); i {
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
	file_uploader_uploader_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*FileStreamRequest_Header)(nil),
		(*FileStreamRequest_Chunk)(nil),
	}
	file_uploader_uploader_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_uploader_uploader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uploader_uploader_proto_goTypes,
		DependencyIndexes: file_uploader_uploader_proto_depIdxs,
		MessageInfos:      file_uploader_uploader_proto_msgTypes,
	}.Build()
	File_uploader_uploader_proto = out.File
	file_uploader_uploader_proto_rawDesc = nil
	file_uploader_uploader_proto_goTypes = nil
	file_uploader_uploader_proto_depIdxs = nil
}
