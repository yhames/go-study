// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: grpc/proto/auth.proto

package auth

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

type ResponseType int32

const (
	ResponseType_SUCCESS ResponseType = 0
	ResponseType_FAILURE ResponseType = 1
	ResponseType_EXPIRED ResponseType = 2
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILURE",
		2: "EXPIRED",
	}
	ResponseType_value = map[string]int32{
		"SUCCESS": 0,
		"FAILURE": 1,
		"EXPIRED": 2,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_proto_auth_proto_enumTypes[0].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_grpc_proto_auth_proto_enumTypes[0]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{0}
}

type AuthData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Token         string                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt     int64                  `protobuf:"varint,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	mi := &file_grpc_proto_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AuthData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AuthData) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AuthData) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type Verify struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        ResponseType           `protobuf:"varint,1,opt,name=status,proto3,enum=ResponseType" json:"status,omitempty"`
	AuthData      *AuthData              `protobuf:"bytes,2,opt,name=auth_data,json=authData,proto3" json:"auth_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Verify) Reset() {
	*x = Verify{}
	mi := &file_grpc_proto_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Verify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Verify) ProtoMessage() {}

func (x *Verify) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Verify.ProtoReflect.Descriptor instead.
func (*Verify) Descriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *Verify) GetStatus() ResponseType {
	if x != nil {
		return x.Status
	}
	return ResponseType_SUCCESS
}

func (x *Verify) GetAuthData() *AuthData {
	if x != nil {
		return x.AuthData
	}
	return nil
}

type CreateTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthData      *AuthData              `protobuf:"bytes,1,opt,name=auth_data,json=authData,proto3" json:"auth_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	mi := &file_grpc_proto_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTokenRequest) GetAuthData() *AuthData {
	if x != nil {
		return x.AuthData
	}
	return nil
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AuthData      *AuthData              `protobuf:"bytes,1,opt,name=auth_data,json=authData,proto3" json:"auth_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	mi := &file_grpc_proto_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTokenResponse) GetAuthData() *AuthData {
	if x != nil {
		return x.AuthData
	}
	return nil
}

type VerifyTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	mi := &file_grpc_proto_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type VerifyTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Verify        *Verify                `protobuf:"bytes,1,opt,name=verify,proto3" json:"verify,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VerifyTokenResponse) Reset() {
	*x = VerifyTokenResponse{}
	mi := &file_grpc_proto_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VerifyTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenResponse) ProtoMessage() {}

func (x *VerifyTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyTokenResponse) GetVerify() *Verify {
	if x != nil {
		return x.Verify
	}
	return nil
}

var File_grpc_proto_auth_proto protoreflect.FileDescriptor

const file_grpc_proto_auth_proto_rawDesc = "" +
	"\n" +
	"\x15grpc/proto/auth.proto\"r\n" +
	"\bAuthData\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05token\x18\x02 \x01(\tR\x05token\x12\x1d\n" +
	"\n" +
	"created_at\x18\x03 \x01(\x03R\tcreatedAt\x12\x1d\n" +
	"\n" +
	"expires_at\x18\x04 \x01(\x03R\texpiresAt\"W\n" +
	"\x06Verify\x12%\n" +
	"\x06status\x18\x01 \x01(\x0e2\r.ResponseTypeR\x06status\x12&\n" +
	"\tauth_data\x18\x02 \x01(\v2\t.AuthDataR\bauthData\"<\n" +
	"\x12CreateTokenRequest\x12&\n" +
	"\tauth_data\x18\x01 \x01(\v2\t.AuthDataR\bauthData\"=\n" +
	"\x13CreateTokenResponse\x12&\n" +
	"\tauth_data\x18\x01 \x01(\v2\t.AuthDataR\bauthData\"*\n" +
	"\x12VerifyTokenRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token\"6\n" +
	"\x13VerifyTokenResponse\x12\x1f\n" +
	"\x06verify\x18\x01 \x01(\v2\a.VerifyR\x06verify*5\n" +
	"\fResponseType\x12\v\n" +
	"\aSUCCESS\x10\x00\x12\v\n" +
	"\aFAILURE\x10\x01\x12\v\n" +
	"\aEXPIRED\x10\x022\x81\x01\n" +
	"\vAuthService\x128\n" +
	"\vCreateToken\x12\x13.CreateTokenRequest\x1a\x14.CreateTokenResponse\x128\n" +
	"\vVerifyToken\x12\x13.VerifyTokenRequest\x1a\x14.VerifyTokenResponseB\aZ\x05/authb\x06proto3"

var (
	file_grpc_proto_auth_proto_rawDescOnce sync.Once
	file_grpc_proto_auth_proto_rawDescData []byte
)

func file_grpc_proto_auth_proto_rawDescGZIP() []byte {
	file_grpc_proto_auth_proto_rawDescOnce.Do(func() {
		file_grpc_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_grpc_proto_auth_proto_rawDesc), len(file_grpc_proto_auth_proto_rawDesc)))
	})
	return file_grpc_proto_auth_proto_rawDescData
}

var file_grpc_proto_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_proto_auth_proto_goTypes = []any{
	(ResponseType)(0),           // 0: ResponseType
	(*AuthData)(nil),            // 1: AuthData
	(*Verify)(nil),              // 2: Verify
	(*CreateTokenRequest)(nil),  // 3: CreateTokenRequest
	(*CreateTokenResponse)(nil), // 4: CreateTokenResponse
	(*VerifyTokenRequest)(nil),  // 5: VerifyTokenRequest
	(*VerifyTokenResponse)(nil), // 6: VerifyTokenResponse
}
var file_grpc_proto_auth_proto_depIdxs = []int32{
	0, // 0: Verify.status:type_name -> ResponseType
	1, // 1: Verify.auth_data:type_name -> AuthData
	1, // 2: CreateTokenRequest.auth_data:type_name -> AuthData
	1, // 3: CreateTokenResponse.auth_data:type_name -> AuthData
	2, // 4: VerifyTokenResponse.verify:type_name -> Verify
	3, // 5: AuthService.CreateToken:input_type -> CreateTokenRequest
	5, // 6: AuthService.VerifyToken:input_type -> VerifyTokenRequest
	4, // 7: AuthService.CreateToken:output_type -> CreateTokenResponse
	6, // 8: AuthService.VerifyToken:output_type -> VerifyTokenResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpc_proto_auth_proto_init() }
func file_grpc_proto_auth_proto_init() {
	if File_grpc_proto_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_grpc_proto_auth_proto_rawDesc), len(file_grpc_proto_auth_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_auth_proto_goTypes,
		DependencyIndexes: file_grpc_proto_auth_proto_depIdxs,
		EnumInfos:         file_grpc_proto_auth_proto_enumTypes,
		MessageInfos:      file_grpc_proto_auth_proto_msgTypes,
	}.Build()
	File_grpc_proto_auth_proto = out.File
	file_grpc_proto_auth_proto_goTypes = nil
	file_grpc_proto_auth_proto_depIdxs = nil
}
