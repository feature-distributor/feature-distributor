// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: common/protobuf/req_group.proto

package pb

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

type ReqGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Key         string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Options     []*ReqGroupOption `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *ReqGroup) Reset() {
	*x = ReqGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGroup) ProtoMessage() {}

func (x *ReqGroup) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGroup.ProtoReflect.Descriptor instead.
func (*ReqGroup) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{0}
}

func (x *ReqGroup) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ReqGroup) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReqGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ReqGroup) GetOptions() []*ReqGroupOption {
	if x != nil {
		return x.Options
	}
	return nil
}

type ReqGroupOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index         int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	AttrType      string `protobuf:"bytes,2,opt,name=attrType,proto3" json:"attrType,omitempty"`
	AttrName      string `protobuf:"bytes,3,opt,name=attrName,proto3" json:"attrName,omitempty"`
	OperationType string `protobuf:"bytes,4,opt,name=operationType,proto3" json:"operationType,omitempty"`
	AttrValue     string `protobuf:"bytes,5,opt,name=attrValue,proto3" json:"attrValue,omitempty"`
}

func (x *ReqGroupOption) Reset() {
	*x = ReqGroupOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGroupOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGroupOption) ProtoMessage() {}

func (x *ReqGroupOption) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGroupOption.ProtoReflect.Descriptor instead.
func (*ReqGroupOption) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{1}
}

func (x *ReqGroupOption) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ReqGroupOption) GetAttrType() string {
	if x != nil {
		return x.AttrType
	}
	return ""
}

func (x *ReqGroupOption) GetAttrName() string {
	if x != nil {
		return x.AttrName
	}
	return ""
}

func (x *ReqGroupOption) GetOperationType() string {
	if x != nil {
		return x.OperationType
	}
	return ""
}

func (x *ReqGroupOption) GetAttrValue() string {
	if x != nil {
		return x.AttrValue
	}
	return ""
}

type ReqGroupOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *ReqGroupOperationResponse) Reset() {
	*x = ReqGroupOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqGroupOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGroupOperationResponse) ProtoMessage() {}

func (x *ReqGroupOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGroupOperationResponse.ProtoReflect.Descriptor instead.
func (*ReqGroupOperationResponse) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{2}
}

func (x *ReqGroupOperationResponse) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ListReqGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index    int64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size     int64   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Key      *string `protobuf:"bytes,3,opt,name=key,proto3,oneof" json:"key,omitempty"`
	Keywords *string `protobuf:"bytes,4,opt,name=keywords,proto3,oneof" json:"keywords,omitempty"`
}

func (x *ListReqGroupRequest) Reset() {
	*x = ListReqGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReqGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReqGroupRequest) ProtoMessage() {}

func (x *ListReqGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReqGroupRequest.ProtoReflect.Descriptor instead.
func (*ListReqGroupRequest) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{3}
}

func (x *ListReqGroupRequest) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ListReqGroupRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListReqGroupRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *ListReqGroupRequest) GetKeywords() string {
	if x != nil && x.Keywords != nil {
		return *x.Keywords
	}
	return ""
}

type ListReqGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int64               `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Groups []*ListItemReqGroup `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *ListReqGroupResponse) Reset() {
	*x = ListReqGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReqGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReqGroupResponse) ProtoMessage() {}

func (x *ListReqGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReqGroupResponse.ProtoReflect.Descriptor instead.
func (*ListReqGroupResponse) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{4}
}

func (x *ListReqGroupResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListReqGroupResponse) GetGroups() []*ListItemReqGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

type UpdateReqGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     int64  `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateReqGroupRequest) Reset() {
	*x = UpdateReqGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReqGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReqGroupRequest) ProtoMessage() {}

func (x *UpdateReqGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReqGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateReqGroupRequest) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateReqGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *UpdateReqGroupRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateReqGroupRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GetReqGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId int64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *GetReqGroupRequest) Reset() {
	*x = GetReqGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReqGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReqGroupRequest) ProtoMessage() {}

func (x *GetReqGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReqGroupRequest.ProtoReflect.Descriptor instead.
func (*GetReqGroupRequest) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{6}
}

func (x *GetReqGroupRequest) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

type ListItemReqGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     int64  `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Key         string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime  int64  `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime  int64  `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *ListItemReqGroup) Reset() {
	*x = ListItemReqGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_protobuf_req_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItemReqGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItemReqGroup) ProtoMessage() {}

func (x *ListItemReqGroup) ProtoReflect() protoreflect.Message {
	mi := &file_common_protobuf_req_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItemReqGroup.ProtoReflect.Descriptor instead.
func (*ListItemReqGroup) Descriptor() ([]byte, []int) {
	return file_common_protobuf_req_group_proto_rawDescGZIP(), []int{7}
}

func (x *ListItemReqGroup) GetGroupId() int64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *ListItemReqGroup) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ListItemReqGroup) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListItemReqGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListItemReqGroup) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ListItemReqGroup) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_common_protobuf_req_group_proto protoreflect.FileDescriptor

var file_common_protobuf_req_group_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x72, 0x65, 0x71, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7f, 0x0a, 0x08, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x74, 0x74, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x74, 0x74, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x74, 0x74, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x74,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74,
	0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x35, 0x0a, 0x19, 0x52, 0x65, 0x71, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x8c,
	0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x57, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x06, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x69, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x2e, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0xb6, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xbf, 0x02, 0x0a, 0x0f, 0x52,
	0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x09, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x37, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x09, 0x2e, 0x52,
	0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x1a, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x13, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_protobuf_req_group_proto_rawDescOnce sync.Once
	file_common_protobuf_req_group_proto_rawDescData = file_common_protobuf_req_group_proto_rawDesc
)

func file_common_protobuf_req_group_proto_rawDescGZIP() []byte {
	file_common_protobuf_req_group_proto_rawDescOnce.Do(func() {
		file_common_protobuf_req_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_protobuf_req_group_proto_rawDescData)
	})
	return file_common_protobuf_req_group_proto_rawDescData
}

var file_common_protobuf_req_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_protobuf_req_group_proto_goTypes = []interface{}{
	(*ReqGroup)(nil),                  // 0: ReqGroup
	(*ReqGroupOption)(nil),            // 1: ReqGroupOption
	(*ReqGroupOperationResponse)(nil), // 2: ReqGroupOperationResponse
	(*ListReqGroupRequest)(nil),       // 3: ListReqGroupRequest
	(*ListReqGroupResponse)(nil),      // 4: ListReqGroupResponse
	(*UpdateReqGroupRequest)(nil),     // 5: UpdateReqGroupRequest
	(*GetReqGroupRequest)(nil),        // 6: GetReqGroupRequest
	(*ListItemReqGroup)(nil),          // 7: ListItemReqGroup
}
var file_common_protobuf_req_group_proto_depIdxs = []int32{
	1, // 0: ReqGroup.options:type_name -> ReqGroupOption
	7, // 1: ListReqGroupResponse.groups:type_name -> ListItemReqGroup
	3, // 2: ReqGroupService.ListReqGroup:input_type -> ListReqGroupRequest
	6, // 3: ReqGroupService.GetReqGroup:input_type -> GetReqGroupRequest
	0, // 4: ReqGroupService.CreateReqGroup:input_type -> ReqGroup
	5, // 5: ReqGroupService.UpdateReqGroup:input_type -> UpdateReqGroupRequest
	6, // 6: ReqGroupService.DeleteReqGroup:input_type -> GetReqGroupRequest
	4, // 7: ReqGroupService.ListReqGroup:output_type -> ListReqGroupResponse
	0, // 8: ReqGroupService.GetReqGroup:output_type -> ReqGroup
	2, // 9: ReqGroupService.CreateReqGroup:output_type -> ReqGroupOperationResponse
	2, // 10: ReqGroupService.UpdateReqGroup:output_type -> ReqGroupOperationResponse
	2, // 11: ReqGroupService.DeleteReqGroup:output_type -> ReqGroupOperationResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_protobuf_req_group_proto_init() }
func file_common_protobuf_req_group_proto_init() {
	if File_common_protobuf_req_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_protobuf_req_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGroup); i {
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
		file_common_protobuf_req_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGroupOption); i {
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
		file_common_protobuf_req_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqGroupOperationResponse); i {
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
		file_common_protobuf_req_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReqGroupRequest); i {
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
		file_common_protobuf_req_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReqGroupResponse); i {
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
		file_common_protobuf_req_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReqGroupRequest); i {
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
		file_common_protobuf_req_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReqGroupRequest); i {
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
		file_common_protobuf_req_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItemReqGroup); i {
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
	file_common_protobuf_req_group_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_protobuf_req_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_protobuf_req_group_proto_goTypes,
		DependencyIndexes: file_common_protobuf_req_group_proto_depIdxs,
		MessageInfos:      file_common_protobuf_req_group_proto_msgTypes,
	}.Build()
	File_common_protobuf_req_group_proto = out.File
	file_common_protobuf_req_group_proto_rawDesc = nil
	file_common_protobuf_req_group_proto_goTypes = nil
	file_common_protobuf_req_group_proto_depIdxs = nil
}