// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/gorm.proto

package proto

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type AutoTimeUnit int32

const (
	AutoTimeUnit_AutoTimeUnitNone   AutoTimeUnit = 0
	AutoTimeUnit_AutoTimeUnitSecond AutoTimeUnit = 1
	AutoTimeUnit_AutoTimeUnitNano   AutoTimeUnit = 2
	AutoTimeUnit_AutoTimeUnitMilli  AutoTimeUnit = 3
)

// Enum value maps for AutoTimeUnit.
var (
	AutoTimeUnit_name = map[int32]string{
		0: "AutoTimeUnitNone",
		1: "AutoTimeUnitSecond",
		2: "AutoTimeUnitNano",
		3: "AutoTimeUnitMilli",
	}
	AutoTimeUnit_value = map[string]int32{
		"AutoTimeUnitNone":   0,
		"AutoTimeUnitSecond": 1,
		"AutoTimeUnitNano":   2,
		"AutoTimeUnitMilli":  3,
	}
)

func (x AutoTimeUnit) Enum() *AutoTimeUnit {
	p := new(AutoTimeUnit)
	*p = x
	return p
}

func (x AutoTimeUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoTimeUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_gorm_proto_enumTypes[0].Descriptor()
}

func (AutoTimeUnit) Type() protoreflect.EnumType {
	return &file_proto_gorm_proto_enumTypes[0]
}

func (x AutoTimeUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoTimeUnit.Descriptor instead.
func (AutoTimeUnit) EnumDescriptor() ([]byte, []int) {
	return file_proto_gorm_proto_rawDescGZIP(), []int{0}
}

type WritePermission int32

const (
	WritePermission_WritePermissionAll    WritePermission = 0
	WritePermission_WritePermissionCreate WritePermission = 1
	WritePermission_WritePermissionUpdate WritePermission = 2
	WritePermission_WritePermissionNone   WritePermission = 3
)

// Enum value maps for WritePermission.
var (
	WritePermission_name = map[int32]string{
		0: "WritePermissionAll",
		1: "WritePermissionCreate",
		2: "WritePermissionUpdate",
		3: "WritePermissionNone",
	}
	WritePermission_value = map[string]int32{
		"WritePermissionAll":    0,
		"WritePermissionCreate": 1,
		"WritePermissionUpdate": 2,
		"WritePermissionNone":   3,
	}
)

func (x WritePermission) Enum() *WritePermission {
	p := new(WritePermission)
	*p = x
	return p
}

func (x WritePermission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WritePermission) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_gorm_proto_enumTypes[1].Descriptor()
}

func (WritePermission) Type() protoreflect.EnumType {
	return &file_proto_gorm_proto_enumTypes[1]
}

func (x WritePermission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WritePermission.Descriptor instead.
func (WritePermission) EnumDescriptor() ([]byte, []int) {
	return file_proto_gorm_proto_rawDescGZIP(), []int{1}
}

type Model struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *Model) Reset() {
	*x = Model{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gorm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Model) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Model) ProtoMessage() {}

func (x *Model) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gorm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Model.ProtoReflect.Descriptor instead.
func (*Model) Descriptor() ([]byte, []int) {
	return file_proto_gorm_proto_rawDescGZIP(), []int{0}
}

func (x *Model) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

type ExtType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImportPath string `protobuf:"bytes,1,opt,name=import_path,json=importPath,proto3" json:"import_path,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ExtType) Reset() {
	*x = ExtType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gorm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtType) ProtoMessage() {}

func (x *ExtType) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gorm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtType.ProtoReflect.Descriptor instead.
func (*ExtType) Descriptor() ([]byte, []int) {
	return file_proto_gorm_proto_rawDescGZIP(), []int{1}
}

func (x *ExtType) GetImportPath() string {
	if x != nil {
		return x.ImportPath
	}
	return ""
}

func (x *ExtType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column            string          `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Type              string          `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Serializer        string          `protobuf:"bytes,3,opt,name=serializer,proto3" json:"serializer,omitempty"`
	Size              int64           `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Pk                bool            `protobuf:"varint,5,opt,name=pk,proto3" json:"pk,omitempty"`
	Unique            bool            `protobuf:"varint,6,opt,name=unique,proto3" json:"unique,omitempty"`
	DefaultValue      string          `protobuf:"bytes,7,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	Precision         int32           `protobuf:"varint,8,opt,name=precision,proto3" json:"precision,omitempty"`
	Scale             int32           `protobuf:"varint,9,opt,name=scale,proto3" json:"scale,omitempty"`
	NotNull           bool            `protobuf:"varint,10,opt,name=not_null,json=notNull,proto3" json:"not_null,omitempty"`
	AutoIncrement     bool            `protobuf:"varint,11,opt,name=auto_increment,json=autoIncrement,proto3" json:"auto_increment,omitempty"`
	AutoIncrementStep int32           `protobuf:"varint,12,opt,name=auto_increment_step,json=autoIncrementStep,proto3" json:"auto_increment_step,omitempty"`
	Embedded          bool            `protobuf:"varint,13,opt,name=embedded,proto3" json:"embedded,omitempty"`
	EmbeddedPrefix    string          `protobuf:"bytes,14,opt,name=embedded_prefix,json=embeddedPrefix,proto3" json:"embedded_prefix,omitempty"`
	AutoCreateTime    AutoTimeUnit    `protobuf:"varint,15,opt,name=auto_create_time,json=autoCreateTime,proto3,enum=yuez.gorm.AutoTimeUnit" json:"auto_create_time,omitempty"`
	AutoUpdateTime    AutoTimeUnit    `protobuf:"varint,16,opt,name=auto_update_time,json=autoUpdateTime,proto3,enum=yuez.gorm.AutoTimeUnit" json:"auto_update_time,omitempty"`
	Index             string          `protobuf:"bytes,17,opt,name=index,proto3" json:"index,omitempty"`
	UniqIndex         string          `protobuf:"bytes,18,opt,name=uniq_index,json=uniqIndex,proto3" json:"uniq_index,omitempty"`
	Check             string          `protobuf:"bytes,19,opt,name=check,proto3" json:"check,omitempty"`
	WritePerm         WritePermission `protobuf:"varint,20,opt,name=write_perm,json=writePerm,proto3,enum=yuez.gorm.WritePermission" json:"write_perm,omitempty"`
	NoReadPerm        bool            `protobuf:"varint,21,opt,name=no_read_perm,json=noReadPerm,proto3" json:"no_read_perm,omitempty"`
	Ignore            string          `protobuf:"bytes,22,opt,name=ignore,proto3" json:"ignore,omitempty"`
	NullableType      bool            `protobuf:"varint,23,opt,name=nullable_type,json=nullableType,proto3" json:"nullable_type,omitempty"`
	ExtType           *ExtType        `protobuf:"bytes,24,opt,name=ext_type,json=extType,proto3" json:"ext_type,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gorm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gorm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_proto_gorm_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *Field) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Field) GetSerializer() string {
	if x != nil {
		return x.Serializer
	}
	return ""
}

func (x *Field) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Field) GetPk() bool {
	if x != nil {
		return x.Pk
	}
	return false
}

func (x *Field) GetUnique() bool {
	if x != nil {
		return x.Unique
	}
	return false
}

func (x *Field) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

func (x *Field) GetPrecision() int32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *Field) GetScale() int32 {
	if x != nil {
		return x.Scale
	}
	return 0
}

func (x *Field) GetNotNull() bool {
	if x != nil {
		return x.NotNull
	}
	return false
}

func (x *Field) GetAutoIncrement() bool {
	if x != nil {
		return x.AutoIncrement
	}
	return false
}

func (x *Field) GetAutoIncrementStep() int32 {
	if x != nil {
		return x.AutoIncrementStep
	}
	return 0
}

func (x *Field) GetEmbedded() bool {
	if x != nil {
		return x.Embedded
	}
	return false
}

func (x *Field) GetEmbeddedPrefix() string {
	if x != nil {
		return x.EmbeddedPrefix
	}
	return ""
}

func (x *Field) GetAutoCreateTime() AutoTimeUnit {
	if x != nil {
		return x.AutoCreateTime
	}
	return AutoTimeUnit_AutoTimeUnitNone
}

func (x *Field) GetAutoUpdateTime() AutoTimeUnit {
	if x != nil {
		return x.AutoUpdateTime
	}
	return AutoTimeUnit_AutoTimeUnitNone
}

func (x *Field) GetIndex() string {
	if x != nil {
		return x.Index
	}
	return ""
}

func (x *Field) GetUniqIndex() string {
	if x != nil {
		return x.UniqIndex
	}
	return ""
}

func (x *Field) GetCheck() string {
	if x != nil {
		return x.Check
	}
	return ""
}

func (x *Field) GetWritePerm() WritePermission {
	if x != nil {
		return x.WritePerm
	}
	return WritePermission_WritePermissionAll
}

func (x *Field) GetNoReadPerm() bool {
	if x != nil {
		return x.NoReadPerm
	}
	return false
}

func (x *Field) GetIgnore() string {
	if x != nil {
		return x.Ignore
	}
	return ""
}

func (x *Field) GetNullableType() bool {
	if x != nil {
		return x.NullableType
	}
	return false
}

func (x *Field) GetExtType() *ExtType {
	if x != nil {
		return x.ExtType
	}
	return nil
}

var file_proto_gorm_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*Model)(nil),
		Field:         51234,
		Name:          "yuez.gorm.model",
		Tag:           "bytes,51234,opt,name=model",
		Filename:      "proto/gorm.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*Field)(nil),
		Field:         51234,
		Name:          "yuez.gorm.field",
		Tag:           "bytes,51234,opt,name=field",
		Filename:      "proto/gorm.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional yuez.gorm.Model model = 51234;
	E_Model = &file_proto_gorm_proto_extTypes[0]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional yuez.gorm.Field field = 51234;
	E_Field = &file_proto_gorm_proto_extTypes[1]
)

var File_proto_gorm_proto protoreflect.FileDescriptor

var file_proto_gorm_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x79, 0x75, 0x65, 0x7a, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x26, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3e, 0x0a, 0x07, 0x45, 0x78, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xb9, 0x06, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x70,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x69, 0x6e, 0x63,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x62, 0x65, 0x64,
	0x64, 0x65, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x41, 0x0a, 0x10, 0x61, 0x75, 0x74,
	0x6f, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x79, 0x75, 0x65, 0x7a, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e,
	0x41, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x0e, 0x61, 0x75,
	0x74, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x10,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x79, 0x75, 0x65, 0x7a, 0x2e, 0x67, 0x6f,
	0x72, 0x6d, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x52,
	0x0e, 0x61, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x13, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x79, 0x75, 0x65, 0x7a, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x77, 0x72, 0x69, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x6f, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x70, 0x65, 0x72, 0x6d, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6e, 0x6f, 0x52,
	0x65, 0x61, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72,
	0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x75, 0x6c, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x79, 0x75, 0x65, 0x7a, 0x2e, 0x67, 0x6f,
	0x72, 0x6d, 0x2e, 0x45, 0x78, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x65, 0x78, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x2a, 0x69, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x75, 0x74,
	0x6f, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x75, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69,
	0x74, 0x4e, 0x61, 0x6e, 0x6f, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x54,
	0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x10, 0x03, 0x2a, 0x78,
	0x0a, 0x0f, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x6c, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12,
	0x17, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4e, 0x6f, 0x6e, 0x65, 0x10, 0x03, 0x3a, 0x4c, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xa2, 0x90, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x79, 0x75, 0x65,
	0x7a, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x3a, 0x4a, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa2,
	0x90, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x79, 0x75, 0x65, 0x7a, 0x2e, 0x67, 0x6f,
	0x72, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x88,
	0x01, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x7a, 0x67, 0x73, 0x32, 0x32, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gorm_proto_rawDescOnce sync.Once
	file_proto_gorm_proto_rawDescData = file_proto_gorm_proto_rawDesc
)

func file_proto_gorm_proto_rawDescGZIP() []byte {
	file_proto_gorm_proto_rawDescOnce.Do(func() {
		file_proto_gorm_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gorm_proto_rawDescData)
	})
	return file_proto_gorm_proto_rawDescData
}

var file_proto_gorm_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_gorm_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_gorm_proto_goTypes = []interface{}{
	(AutoTimeUnit)(0),                 // 0: yuez.gorm.AutoTimeUnit
	(WritePermission)(0),              // 1: yuez.gorm.WritePermission
	(*Model)(nil),                     // 2: yuez.gorm.Model
	(*ExtType)(nil),                   // 3: yuez.gorm.ExtType
	(*Field)(nil),                     // 4: yuez.gorm.Field
	(*descriptor.MessageOptions)(nil), // 5: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 6: google.protobuf.FieldOptions
}
var file_proto_gorm_proto_depIdxs = []int32{
	0, // 0: yuez.gorm.Field.auto_create_time:type_name -> yuez.gorm.AutoTimeUnit
	0, // 1: yuez.gorm.Field.auto_update_time:type_name -> yuez.gorm.AutoTimeUnit
	1, // 2: yuez.gorm.Field.write_perm:type_name -> yuez.gorm.WritePermission
	3, // 3: yuez.gorm.Field.ext_type:type_name -> yuez.gorm.ExtType
	5, // 4: yuez.gorm.model:extendee -> google.protobuf.MessageOptions
	6, // 5: yuez.gorm.field:extendee -> google.protobuf.FieldOptions
	2, // 6: yuez.gorm.model:type_name -> yuez.gorm.Model
	4, // 7: yuez.gorm.field:type_name -> yuez.gorm.Field
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	4, // [4:6] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_gorm_proto_init() }
func file_proto_gorm_proto_init() {
	if File_proto_gorm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gorm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Model); i {
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
		file_proto_gorm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtType); i {
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
		file_proto_gorm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
			RawDescriptor: file_proto_gorm_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_proto_gorm_proto_goTypes,
		DependencyIndexes: file_proto_gorm_proto_depIdxs,
		EnumInfos:         file_proto_gorm_proto_enumTypes,
		MessageInfos:      file_proto_gorm_proto_msgTypes,
		ExtensionInfos:    file_proto_gorm_proto_extTypes,
	}.Build()
	File_proto_gorm_proto = out.File
	file_proto_gorm_proto_rawDesc = nil
	file_proto_gorm_proto_goTypes = nil
	file_proto_gorm_proto_depIdxs = nil
}
