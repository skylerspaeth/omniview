// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/settings.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type SettingType int32

const (
	SettingType_TEXT     SettingType = 0
	SettingType_INTEGER  SettingType = 1
	SettingType_FLOAT    SettingType = 2
	SettingType_TOGGLE   SettingType = 3
	SettingType_COLOR    SettingType = 4
	SettingType_DATETIME SettingType = 5
	SettingType_PASSWORD SettingType = 6
)

// Enum value maps for SettingType.
var (
	SettingType_name = map[int32]string{
		0: "TEXT",
		1: "INTEGER",
		2: "FLOAT",
		3: "TOGGLE",
		4: "COLOR",
		5: "DATETIME",
		6: "PASSWORD",
	}
	SettingType_value = map[string]int32{
		"TEXT":     0,
		"INTEGER":  1,
		"FLOAT":    2,
		"TOGGLE":   3,
		"COLOR":    4,
		"DATETIME": 5,
		"PASSWORD": 6,
	}
)

func (x SettingType) Enum() *SettingType {
	p := new(SettingType)
	*p = x
	return p
}

func (x SettingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettingType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_settings_proto_enumTypes[0].Descriptor()
}

func (SettingType) Type() protoreflect.EnumType {
	return &file_proto_settings_proto_enumTypes[0]
}

func (x SettingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettingType.Descriptor instead.
func (SettingType) EnumDescriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{0}
}

type VisualResourceComponent_ComponentType int32

const (
	VisualResourceComponent_SIDEBAR    VisualResourceComponent_ComponentType = 0
	VisualResourceComponent_TABLE      VisualResourceComponent_ComponentType = 1
	VisualResourceComponent_TABLE_ROW  VisualResourceComponent_ComponentType = 2
	VisualResourceComponent_TABLE_CELL VisualResourceComponent_ComponentType = 3
)

// Enum value maps for VisualResourceComponent_ComponentType.
var (
	VisualResourceComponent_ComponentType_name = map[int32]string{
		0: "SIDEBAR",
		1: "TABLE",
		2: "TABLE_ROW",
		3: "TABLE_CELL",
	}
	VisualResourceComponent_ComponentType_value = map[string]int32{
		"SIDEBAR":    0,
		"TABLE":      1,
		"TABLE_ROW":  2,
		"TABLE_CELL": 3,
	}
)

func (x VisualResourceComponent_ComponentType) Enum() *VisualResourceComponent_ComponentType {
	p := new(VisualResourceComponent_ComponentType)
	*p = x
	return p
}

func (x VisualResourceComponent_ComponentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VisualResourceComponent_ComponentType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_settings_proto_enumTypes[1].Descriptor()
}

func (VisualResourceComponent_ComponentType) Type() protoreflect.EnumType {
	return &file_proto_settings_proto_enumTypes[1]
}

func (x VisualResourceComponent_ComponentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VisualResourceComponent_ComponentType.Descriptor instead.
func (VisualResourceComponent_ComponentType) EnumDescriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{11, 0}
}

type SettingsCategory struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label         string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Icon          string                 `protobuf:"bytes,4,opt,name=icon,proto3" json:"icon,omitempty"`
	Settings      map[string]*Setting    `protobuf:"bytes,5,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SettingsCategory) Reset() {
	*x = SettingsCategory{}
	mi := &file_proto_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SettingsCategory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsCategory) ProtoMessage() {}

func (x *SettingsCategory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsCategory.ProtoReflect.Descriptor instead.
func (*SettingsCategory) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsCategory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SettingsCategory) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *SettingsCategory) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SettingsCategory) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *SettingsCategory) GetSettings() map[string]*Setting {
	if x != nil {
		return x.Settings
	}
	return nil
}

type Setting struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label         string                 `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Type          SettingType            `protobuf:"varint,4,opt,name=type,proto3,enum=com.omniview.pluginsdk.SettingType" json:"type,omitempty"`
	Value         *anypb.Any             `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Default       *anypb.Any             `protobuf:"bytes,6,opt,name=default,proto3" json:"default,omitempty"`
	Options       []*SettingOption       `protobuf:"bytes,7,rep,name=options,proto3" json:"options,omitempty"`
	FileSelection *SettingFileSelection  `protobuf:"bytes,8,opt,name=file_selection,json=fileSelection,proto3,oneof" json:"file_selection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Setting) Reset() {
	*x = Setting{}
	mi := &file_proto_settings_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{1}
}

func (x *Setting) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Setting) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Setting) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Setting) GetType() SettingType {
	if x != nil {
		return x.Type
	}
	return SettingType_TEXT
}

func (x *Setting) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Setting) GetDefault() *anypb.Any {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *Setting) GetOptions() []*SettingOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *Setting) GetFileSelection() *SettingFileSelection {
	if x != nil {
		return x.FileSelection
	}
	return nil
}

type SettingOption struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Label         string                 `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Value         *anypb.Any             `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SettingOption) Reset() {
	*x = SettingOption{}
	mi := &file_proto_settings_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SettingOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingOption) ProtoMessage() {}

func (x *SettingOption) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingOption.ProtoReflect.Descriptor instead.
func (*SettingOption) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{2}
}

func (x *SettingOption) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *SettingOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SettingOption) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SettingFileSelection struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enabled       bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	AllowFolders  bool                   `protobuf:"varint,2,opt,name=allow_folders,json=allowFolders,proto3" json:"allow_folders,omitempty"`
	Extensions    []string               `protobuf:"bytes,3,rep,name=extensions,proto3" json:"extensions,omitempty"`
	Multiple      bool                   `protobuf:"varint,4,opt,name=multiple,proto3" json:"multiple,omitempty"`
	Relative      bool                   `protobuf:"varint,5,opt,name=relative,proto3" json:"relative,omitempty"`
	DefaultPath   string                 `protobuf:"bytes,6,opt,name=default_path,json=defaultPath,proto3" json:"default_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SettingFileSelection) Reset() {
	*x = SettingFileSelection{}
	mi := &file_proto_settings_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SettingFileSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingFileSelection) ProtoMessage() {}

func (x *SettingFileSelection) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingFileSelection.ProtoReflect.Descriptor instead.
func (*SettingFileSelection) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{3}
}

func (x *SettingFileSelection) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *SettingFileSelection) GetAllowFolders() bool {
	if x != nil {
		return x.AllowFolders
	}
	return false
}

func (x *SettingFileSelection) GetExtensions() []string {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *SettingFileSelection) GetMultiple() bool {
	if x != nil {
		return x.Multiple
	}
	return false
}

func (x *SettingFileSelection) GetRelative() bool {
	if x != nil {
		return x.Relative
	}
	return false
}

func (x *SettingFileSelection) GetDefaultPath() string {
	if x != nil {
		return x.DefaultPath
	}
	return ""
}

type GetSettingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSettingRequest) Reset() {
	*x = GetSettingRequest{}
	mi := &file_proto_settings_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingRequest) ProtoMessage() {}

func (x *GetSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingRequest.ProtoReflect.Descriptor instead.
func (*GetSettingRequest) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{4}
}

func (x *GetSettingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSettingValueRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSettingValueRequest) Reset() {
	*x = GetSettingValueRequest{}
	mi := &file_proto_settings_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSettingValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingValueRequest) ProtoMessage() {}

func (x *GetSettingValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingValueRequest.ProtoReflect.Descriptor instead.
func (*GetSettingValueRequest) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{5}
}

func (x *GetSettingValueRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSettingValueResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         *anypb.Any             `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSettingValueResponse) Reset() {
	*x = GetSettingValueResponse{}
	mi := &file_proto_settings_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSettingValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingValueResponse) ProtoMessage() {}

func (x *GetSettingValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingValueResponse.ProtoReflect.Descriptor instead.
func (*GetSettingValueResponse) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{6}
}

func (x *GetSettingValueResponse) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetSettingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value         *anypb.Any             `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetSettingRequest) Reset() {
	*x = SetSettingRequest{}
	mi := &file_proto_settings_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSettingRequest) ProtoMessage() {}

func (x *SetSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSettingRequest.ProtoReflect.Descriptor instead.
func (*SetSettingRequest) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{7}
}

func (x *SetSettingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetSettingRequest) GetValue() *anypb.Any {
	if x != nil {
		return x.Value
	}
	return nil
}

type SetSettingsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Settings      map[string]*anypb.Any  `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetSettingsRequest) Reset() {
	*x = SetSettingsRequest{}
	mi := &file_proto_settings_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSettingsRequest) ProtoMessage() {}

func (x *SetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSettingsRequest.ProtoReflect.Descriptor instead.
func (*SetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{8}
}

func (x *SetSettingsRequest) GetSettings() map[string]*anypb.Any {
	if x != nil {
		return x.Settings
	}
	return nil
}

type ListSettingsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Settings      map[string]*Setting    `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSettingsResponse) Reset() {
	*x = ListSettingsResponse{}
	mi := &file_proto_settings_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSettingsResponse) ProtoMessage() {}

func (x *ListSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSettingsResponse.ProtoReflect.Descriptor instead.
func (*ListSettingsResponse) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{9}
}

func (x *ListSettingsResponse) GetSettings() map[string]*Setting {
	if x != nil {
		return x.Settings
	}
	return nil
}

// VisualComponentList is a list of all the visual components provided
// by the plugin, grouped by the section they are applicable to.
type VisualComponentList struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Resource      []*VisualResourceComponent `protobuf:"bytes,1,rep,name=resource,proto3" json:"resource,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisualComponentList) Reset() {
	*x = VisualComponentList{}
	mi := &file_proto_settings_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisualComponentList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisualComponentList) ProtoMessage() {}

func (x *VisualComponentList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisualComponentList.ProtoReflect.Descriptor instead.
func (*VisualComponentList) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{10}
}

func (x *VisualComponentList) GetResource() []*VisualResourceComponent {
	if x != nil {
		return x.Resource
	}
	return nil
}

// VisualResourceComponent is a visual component that can be displayed
// in the UI within a resource.
type VisualResourceComponent struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The unique identifier of the component.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The plugin for which this component is applicable.
	Plugin string `protobuf:"bytes,2,opt,name=plugin,proto3" json:"plugin,omitempty"`
	// The type of the component.
	Type VisualResourceComponent_ComponentType `protobuf:"varint,3,opt,name=type,proto3,enum=com.omniview.pluginsdk.VisualResourceComponent_ComponentType" json:"type,omitempty"`
	// The resources for which this component is applicable.
	Resources     []string `protobuf:"bytes,4,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VisualResourceComponent) Reset() {
	*x = VisualResourceComponent{}
	mi := &file_proto_settings_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VisualResourceComponent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VisualResourceComponent) ProtoMessage() {}

func (x *VisualResourceComponent) ProtoReflect() protoreflect.Message {
	mi := &file_proto_settings_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VisualResourceComponent.ProtoReflect.Descriptor instead.
func (*VisualResourceComponent) Descriptor() ([]byte, []int) {
	return file_proto_settings_proto_rawDescGZIP(), []int{11}
}

func (x *VisualResourceComponent) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VisualResourceComponent) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *VisualResourceComponent) GetType() VisualResourceComponent_ComponentType {
	if x != nil {
		return x.Type
	}
	return VisualResourceComponent_SIDEBAR
}

func (x *VisualResourceComponent) GetResources() []string {
	if x != nil {
		return x.Resources
	}
	return nil
}

var File_proto_settings_proto protoreflect.FileDescriptor

const file_proto_settings_proto_rawDesc = "" +
	"\n" +
	"\x14proto/settings.proto\x12\x16com.omniview.pluginsdk\x1a\x1bgoogle/protobuf/empty.proto\x1a\x19google/protobuf/any.proto\"\xa0\x02\n" +
	"\x10SettingsCategory\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05label\x18\x02 \x01(\tR\x05label\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x12\n" +
	"\x04icon\x18\x04 \x01(\tR\x04icon\x12R\n" +
	"\bsettings\x18\x05 \x03(\v26.com.omniview.pluginsdk.SettingsCategory.SettingsEntryR\bsettings\x1a\\\n" +
	"\rSettingsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x125\n" +
	"\x05value\x18\x02 \x01(\v2\x1f.com.omniview.pluginsdk.SettingR\x05value:\x028\x01\"\x94\x03\n" +
	"\aSetting\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05label\x18\x02 \x01(\tR\x05label\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x127\n" +
	"\x04type\x18\x04 \x01(\x0e2#.com.omniview.pluginsdk.SettingTypeR\x04type\x12*\n" +
	"\x05value\x18\x05 \x01(\v2\x14.google.protobuf.AnyR\x05value\x12.\n" +
	"\adefault\x18\x06 \x01(\v2\x14.google.protobuf.AnyR\adefault\x12?\n" +
	"\aoptions\x18\a \x03(\v2%.com.omniview.pluginsdk.SettingOptionR\aoptions\x12X\n" +
	"\x0efile_selection\x18\b \x01(\v2,.com.omniview.pluginsdk.SettingFileSelectionH\x00R\rfileSelection\x88\x01\x01B\x11\n" +
	"\x0f_file_selection\"s\n" +
	"\rSettingOption\x12\x14\n" +
	"\x05label\x18\x01 \x01(\tR\x05label\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12*\n" +
	"\x05value\x18\x03 \x01(\v2\x14.google.protobuf.AnyR\x05value\"\xd0\x01\n" +
	"\x14SettingFileSelection\x12\x18\n" +
	"\aenabled\x18\x01 \x01(\bR\aenabled\x12#\n" +
	"\rallow_folders\x18\x02 \x01(\bR\fallowFolders\x12\x1e\n" +
	"\n" +
	"extensions\x18\x03 \x03(\tR\n" +
	"extensions\x12\x1a\n" +
	"\bmultiple\x18\x04 \x01(\bR\bmultiple\x12\x1a\n" +
	"\brelative\x18\x05 \x01(\bR\brelative\x12!\n" +
	"\fdefault_path\x18\x06 \x01(\tR\vdefaultPath\"#\n" +
	"\x11GetSettingRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"(\n" +
	"\x16GetSettingValueRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"E\n" +
	"\x17GetSettingValueResponse\x12*\n" +
	"\x05value\x18\x01 \x01(\v2\x14.google.protobuf.AnyR\x05value\"O\n" +
	"\x11SetSettingRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05value\"\xbd\x01\n" +
	"\x12SetSettingsRequest\x12T\n" +
	"\bsettings\x18\x01 \x03(\v28.com.omniview.pluginsdk.SetSettingsRequest.SettingsEntryR\bsettings\x1aQ\n" +
	"\rSettingsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12*\n" +
	"\x05value\x18\x02 \x01(\v2\x14.google.protobuf.AnyR\x05value:\x028\x01\"\xcc\x01\n" +
	"\x14ListSettingsResponse\x12V\n" +
	"\bsettings\x18\x01 \x03(\v2:.com.omniview.pluginsdk.ListSettingsResponse.SettingsEntryR\bsettings\x1a\\\n" +
	"\rSettingsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x125\n" +
	"\x05value\x18\x02 \x01(\v2\x1f.com.omniview.pluginsdk.SettingR\x05value:\x028\x01\"b\n" +
	"\x13VisualComponentList\x12K\n" +
	"\bresource\x18\x01 \x03(\v2/.com.omniview.pluginsdk.VisualResourceComponentR\bresource\"\xfa\x01\n" +
	"\x17VisualResourceComponent\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x16\n" +
	"\x06plugin\x18\x02 \x01(\tR\x06plugin\x12Q\n" +
	"\x04type\x18\x03 \x01(\x0e2=.com.omniview.pluginsdk.VisualResourceComponent.ComponentTypeR\x04type\x12\x1c\n" +
	"\tresources\x18\x04 \x03(\tR\tresources\"F\n" +
	"\rComponentType\x12\v\n" +
	"\aSIDEBAR\x10\x00\x12\t\n" +
	"\x05TABLE\x10\x01\x12\r\n" +
	"\tTABLE_ROW\x10\x02\x12\x0e\n" +
	"\n" +
	"TABLE_CELL\x10\x03*b\n" +
	"\vSettingType\x12\b\n" +
	"\x04TEXT\x10\x00\x12\v\n" +
	"\aINTEGER\x10\x01\x12\t\n" +
	"\x05FLOAT\x10\x02\x12\n" +
	"\n" +
	"\x06TOGGLE\x10\x03\x12\t\n" +
	"\x05COLOR\x10\x04\x12\f\n" +
	"\bDATETIME\x10\x05\x12\f\n" +
	"\bPASSWORD\x10\x062\xd8\x03\n" +
	"\x0eSettingsPlugin\x12T\n" +
	"\fListSettings\x12\x16.google.protobuf.Empty\x1a,.com.omniview.pluginsdk.ListSettingsResponse\x12X\n" +
	"\n" +
	"GetSetting\x12).com.omniview.pluginsdk.GetSettingRequest\x1a\x1f.com.omniview.pluginsdk.Setting\x12r\n" +
	"\x0fGetSettingValue\x12..com.omniview.pluginsdk.GetSettingValueRequest\x1a/.com.omniview.pluginsdk.GetSettingValueResponse\x12O\n" +
	"\n" +
	"SetSetting\x12).com.omniview.pluginsdk.SetSettingRequest\x1a\x16.google.protobuf.Empty\x12Q\n" +
	"\vSetSettings\x12*.com.omniview.pluginsdk.SetSettingsRequest\x1a\x16.google.protobuf.EmptyB\tZ\a./protob\x06proto3"

var (
	file_proto_settings_proto_rawDescOnce sync.Once
	file_proto_settings_proto_rawDescData []byte
)

func file_proto_settings_proto_rawDescGZIP() []byte {
	file_proto_settings_proto_rawDescOnce.Do(func() {
		file_proto_settings_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_settings_proto_rawDesc), len(file_proto_settings_proto_rawDesc)))
	})
	return file_proto_settings_proto_rawDescData
}

var file_proto_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_proto_settings_proto_goTypes = []any{
	(SettingType)(0), // 0: com.omniview.pluginsdk.SettingType
	(VisualResourceComponent_ComponentType)(0), // 1: com.omniview.pluginsdk.VisualResourceComponent.ComponentType
	(*SettingsCategory)(nil),                   // 2: com.omniview.pluginsdk.SettingsCategory
	(*Setting)(nil),                            // 3: com.omniview.pluginsdk.Setting
	(*SettingOption)(nil),                      // 4: com.omniview.pluginsdk.SettingOption
	(*SettingFileSelection)(nil),               // 5: com.omniview.pluginsdk.SettingFileSelection
	(*GetSettingRequest)(nil),                  // 6: com.omniview.pluginsdk.GetSettingRequest
	(*GetSettingValueRequest)(nil),             // 7: com.omniview.pluginsdk.GetSettingValueRequest
	(*GetSettingValueResponse)(nil),            // 8: com.omniview.pluginsdk.GetSettingValueResponse
	(*SetSettingRequest)(nil),                  // 9: com.omniview.pluginsdk.SetSettingRequest
	(*SetSettingsRequest)(nil),                 // 10: com.omniview.pluginsdk.SetSettingsRequest
	(*ListSettingsResponse)(nil),               // 11: com.omniview.pluginsdk.ListSettingsResponse
	(*VisualComponentList)(nil),                // 12: com.omniview.pluginsdk.VisualComponentList
	(*VisualResourceComponent)(nil),            // 13: com.omniview.pluginsdk.VisualResourceComponent
	nil,                                        // 14: com.omniview.pluginsdk.SettingsCategory.SettingsEntry
	nil,                                        // 15: com.omniview.pluginsdk.SetSettingsRequest.SettingsEntry
	nil,                                        // 16: com.omniview.pluginsdk.ListSettingsResponse.SettingsEntry
	(*anypb.Any)(nil),                          // 17: google.protobuf.Any
	(*emptypb.Empty)(nil),                      // 18: google.protobuf.Empty
}
var file_proto_settings_proto_depIdxs = []int32{
	14, // 0: com.omniview.pluginsdk.SettingsCategory.settings:type_name -> com.omniview.pluginsdk.SettingsCategory.SettingsEntry
	0,  // 1: com.omniview.pluginsdk.Setting.type:type_name -> com.omniview.pluginsdk.SettingType
	17, // 2: com.omniview.pluginsdk.Setting.value:type_name -> google.protobuf.Any
	17, // 3: com.omniview.pluginsdk.Setting.default:type_name -> google.protobuf.Any
	4,  // 4: com.omniview.pluginsdk.Setting.options:type_name -> com.omniview.pluginsdk.SettingOption
	5,  // 5: com.omniview.pluginsdk.Setting.file_selection:type_name -> com.omniview.pluginsdk.SettingFileSelection
	17, // 6: com.omniview.pluginsdk.SettingOption.value:type_name -> google.protobuf.Any
	17, // 7: com.omniview.pluginsdk.GetSettingValueResponse.value:type_name -> google.protobuf.Any
	17, // 8: com.omniview.pluginsdk.SetSettingRequest.value:type_name -> google.protobuf.Any
	15, // 9: com.omniview.pluginsdk.SetSettingsRequest.settings:type_name -> com.omniview.pluginsdk.SetSettingsRequest.SettingsEntry
	16, // 10: com.omniview.pluginsdk.ListSettingsResponse.settings:type_name -> com.omniview.pluginsdk.ListSettingsResponse.SettingsEntry
	13, // 11: com.omniview.pluginsdk.VisualComponentList.resource:type_name -> com.omniview.pluginsdk.VisualResourceComponent
	1,  // 12: com.omniview.pluginsdk.VisualResourceComponent.type:type_name -> com.omniview.pluginsdk.VisualResourceComponent.ComponentType
	3,  // 13: com.omniview.pluginsdk.SettingsCategory.SettingsEntry.value:type_name -> com.omniview.pluginsdk.Setting
	17, // 14: com.omniview.pluginsdk.SetSettingsRequest.SettingsEntry.value:type_name -> google.protobuf.Any
	3,  // 15: com.omniview.pluginsdk.ListSettingsResponse.SettingsEntry.value:type_name -> com.omniview.pluginsdk.Setting
	18, // 16: com.omniview.pluginsdk.SettingsPlugin.ListSettings:input_type -> google.protobuf.Empty
	6,  // 17: com.omniview.pluginsdk.SettingsPlugin.GetSetting:input_type -> com.omniview.pluginsdk.GetSettingRequest
	7,  // 18: com.omniview.pluginsdk.SettingsPlugin.GetSettingValue:input_type -> com.omniview.pluginsdk.GetSettingValueRequest
	9,  // 19: com.omniview.pluginsdk.SettingsPlugin.SetSetting:input_type -> com.omniview.pluginsdk.SetSettingRequest
	10, // 20: com.omniview.pluginsdk.SettingsPlugin.SetSettings:input_type -> com.omniview.pluginsdk.SetSettingsRequest
	11, // 21: com.omniview.pluginsdk.SettingsPlugin.ListSettings:output_type -> com.omniview.pluginsdk.ListSettingsResponse
	3,  // 22: com.omniview.pluginsdk.SettingsPlugin.GetSetting:output_type -> com.omniview.pluginsdk.Setting
	8,  // 23: com.omniview.pluginsdk.SettingsPlugin.GetSettingValue:output_type -> com.omniview.pluginsdk.GetSettingValueResponse
	18, // 24: com.omniview.pluginsdk.SettingsPlugin.SetSetting:output_type -> google.protobuf.Empty
	18, // 25: com.omniview.pluginsdk.SettingsPlugin.SetSettings:output_type -> google.protobuf.Empty
	21, // [21:26] is the sub-list for method output_type
	16, // [16:21] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_proto_settings_proto_init() }
func file_proto_settings_proto_init() {
	if File_proto_settings_proto != nil {
		return
	}
	file_proto_settings_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_settings_proto_rawDesc), len(file_proto_settings_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_settings_proto_goTypes,
		DependencyIndexes: file_proto_settings_proto_depIdxs,
		EnumInfos:         file_proto_settings_proto_enumTypes,
		MessageInfos:      file_proto_settings_proto_msgTypes,
	}.Build()
	File_proto_settings_proto = out.File
	file_proto_settings_proto_goTypes = nil
	file_proto_settings_proto_depIdxs = nil
}
