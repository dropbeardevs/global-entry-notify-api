// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: global_entry_notify_api.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type UpdateNotificationTokenRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateNotificationTokenRq) Reset() {
	*x = UpdateNotificationTokenRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNotificationTokenRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotificationTokenRq) ProtoMessage() {}

func (x *UpdateNotificationTokenRq) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotificationTokenRq.ProtoReflect.Descriptor instead.
func (*UpdateNotificationTokenRq) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateNotificationTokenRq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateNotificationTokenRq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UpdateNotificationDetailsRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId              string               `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NotificationDetails *NotificationDetails `protobuf:"bytes,2,opt,name=notificationDetails,proto3" json:"notificationDetails,omitempty"`
}

func (x *UpdateNotificationDetailsRq) Reset() {
	*x = UpdateNotificationDetailsRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNotificationDetailsRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNotificationDetailsRq) ProtoMessage() {}

func (x *UpdateNotificationDetailsRq) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNotificationDetailsRq.ProtoReflect.Descriptor instead.
func (*UpdateNotificationDetailsRq) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateNotificationDetailsRq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateNotificationDetailsRq) GetNotificationDetails() *NotificationDetails {
	if x != nil {
		return x.NotificationDetails
	}
	return nil
}

type DeleteNotificationDetailsRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	LocationId int32  `protobuf:"varint,2,opt,name=locationId,proto3" json:"locationId,omitempty"`
}

func (x *DeleteNotificationDetailsRq) Reset() {
	*x = DeleteNotificationDetailsRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNotificationDetailsRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNotificationDetailsRq) ProtoMessage() {}

func (x *DeleteNotificationDetailsRq) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNotificationDetailsRq.ProtoReflect.Descriptor instead.
func (*DeleteNotificationDetailsRq) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNotificationDetailsRq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteNotificationDetailsRq) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

type NotificationDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationId       int32                  `protobuf:"varint,1,opt,name=locationId,proto3" json:"locationId,omitempty"`
	TargetDate       *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=targetDate,proto3" json:"targetDate,omitempty"`
	AppointmentDate  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=appointmentDate,proto3" json:"appointmentDate,omitempty"`
	LastNotifiedDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=lastNotifiedDate,proto3" json:"lastNotifiedDate,omitempty"`
}

func (x *NotificationDetails) Reset() {
	*x = NotificationDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationDetails) ProtoMessage() {}

func (x *NotificationDetails) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationDetails.ProtoReflect.Descriptor instead.
func (*NotificationDetails) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{4}
}

func (x *NotificationDetails) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *NotificationDetails) GetTargetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.TargetDate
	}
	return nil
}

func (x *NotificationDetails) GetAppointmentDate() *timestamppb.Timestamp {
	if x != nil {
		return x.AppointmentDate
	}
	return nil
}

func (x *NotificationDetails) GetLastNotifiedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastNotifiedDate
	}
	return nil
}

type LocationsRp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocationId        int32  `protobuf:"varint,1,opt,name=locationId,proto3" json:"locationId,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address           string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	AddressAdditional string `protobuf:"bytes,4,opt,name=addressAdditional,proto3" json:"addressAdditional,omitempty"`
	City              string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	State             string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	PostalCode        string `protobuf:"bytes,7,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	CountryCode       string `protobuf:"bytes,8,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	TimezoneData      string `protobuf:"bytes,9,opt,name=timezoneData,proto3" json:"timezoneData,omitempty"`
	PhoneNumber       string `protobuf:"bytes,10,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	PhoneAreaCode     string `protobuf:"bytes,11,opt,name=phoneAreaCode,proto3" json:"phoneAreaCode,omitempty"`
	PhoneCountryCode  string `protobuf:"bytes,12,opt,name=phoneCountryCode,proto3" json:"phoneCountryCode,omitempty"`
	Directions        string `protobuf:"bytes,13,opt,name=Directions,proto3" json:"Directions,omitempty"`
	Notes             string `protobuf:"bytes,14,opt,name=Notes,proto3" json:"Notes,omitempty"`
}

func (x *LocationsRp) Reset() {
	*x = LocationsRp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationsRp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationsRp) ProtoMessage() {}

func (x *LocationsRp) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationsRp.ProtoReflect.Descriptor instead.
func (*LocationsRp) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{5}
}

func (x *LocationsRp) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *LocationsRp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LocationsRp) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LocationsRp) GetAddressAdditional() string {
	if x != nil {
		return x.AddressAdditional
	}
	return ""
}

func (x *LocationsRp) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *LocationsRp) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *LocationsRp) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *LocationsRp) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *LocationsRp) GetTimezoneData() string {
	if x != nil {
		return x.TimezoneData
	}
	return ""
}

func (x *LocationsRp) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *LocationsRp) GetPhoneAreaCode() string {
	if x != nil {
		return x.PhoneAreaCode
	}
	return ""
}

func (x *LocationsRp) GetPhoneCountryCode() string {
	if x != nil {
		return x.PhoneCountryCode
	}
	return ""
}

func (x *LocationsRp) GetDirections() string {
	if x != nil {
		return x.Directions
	}
	return ""
}

func (x *LocationsRp) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type GetNotificationDetailsRq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetNotificationDetailsRq) Reset() {
	*x = GetNotificationDetailsRq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationDetailsRq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationDetailsRq) ProtoMessage() {}

func (x *GetNotificationDetailsRq) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationDetailsRq.ProtoReflect.Descriptor instead.
func (*GetNotificationDetailsRq) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{6}
}

func (x *GetNotificationDetailsRq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetNotificationDetailsRp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NotificationDetails []*NotificationDetails `protobuf:"bytes,1,rep,name=notificationDetails,proto3" json:"notificationDetails,omitempty"`
}

func (x *GetNotificationDetailsRp) Reset() {
	*x = GetNotificationDetailsRp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_global_entry_notify_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNotificationDetailsRp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNotificationDetailsRp) ProtoMessage() {}

func (x *GetNotificationDetailsRp) ProtoReflect() protoreflect.Message {
	mi := &file_global_entry_notify_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNotificationDetailsRp.ProtoReflect.Descriptor instead.
func (*GetNotificationDetailsRp) Descriptor() ([]byte, []int) {
	return file_global_entry_notify_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetNotificationDetailsRp) GetNotificationDetails() []*NotificationDetails {
	if x != nil {
		return x.NotificationDetails
	}
	return nil
}

var File_global_entry_notify_api_proto protoreflect.FileDescriptor

var file_global_entry_notify_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x49, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x95, 0x01, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x71,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x5e, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x55, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0xff, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x44, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x22, 0xc3, 0x03, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2c, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x52, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x70, 0x12, 0x5e, 0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x32, 0xdc, 0x04, 0x0a, 0x18, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x1b, 0x47, 0x72, 0x70, 0x63, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x32, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x71, 0x1a, 0x1e, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x1d, 0x47, 0x72, 0x70,
	0x63, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x34, 0x2e, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x71,
	0x1a, 0x1e, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x00, 0x12, 0x77, 0x0a, 0x1d, 0x47, 0x72, 0x70, 0x63, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x12, 0x34, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x71, 0x1a, 0x1e, 0x2e, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x10, 0x47,
	0x72, 0x70, 0x63, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x24, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x70, 0x22, 0x00, 0x12,
	0x84, 0x01, 0x0a, 0x1a, 0x47, 0x72, 0x70, 0x63, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x31,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52,
	0x71, 0x1a, 0x31, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x70, 0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x64, 0x72, 0x6f, 0x70, 0x62, 0x65, 0x61, 0x72,
	0x64, 0x65, 0x76, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2d, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_global_entry_notify_api_proto_rawDescOnce sync.Once
	file_global_entry_notify_api_proto_rawDescData = file_global_entry_notify_api_proto_rawDesc
)

func file_global_entry_notify_api_proto_rawDescGZIP() []byte {
	file_global_entry_notify_api_proto_rawDescOnce.Do(func() {
		file_global_entry_notify_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_global_entry_notify_api_proto_rawDescData)
	})
	return file_global_entry_notify_api_proto_rawDescData
}

var file_global_entry_notify_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_global_entry_notify_api_proto_goTypes = []interface{}{
	(*Error)(nil),                       // 0: global_entry_notify_api.Error
	(*UpdateNotificationTokenRq)(nil),   // 1: global_entry_notify_api.UpdateNotificationTokenRq
	(*UpdateNotificationDetailsRq)(nil), // 2: global_entry_notify_api.UpdateNotificationDetailsRq
	(*DeleteNotificationDetailsRq)(nil), // 3: global_entry_notify_api.DeleteNotificationDetailsRq
	(*NotificationDetails)(nil),         // 4: global_entry_notify_api.NotificationDetails
	(*LocationsRp)(nil),                 // 5: global_entry_notify_api.LocationsRp
	(*GetNotificationDetailsRq)(nil),    // 6: global_entry_notify_api.GetNotificationDetailsRq
	(*GetNotificationDetailsRp)(nil),    // 7: global_entry_notify_api.GetNotificationDetailsRp
	(*timestamppb.Timestamp)(nil),       // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_global_entry_notify_api_proto_depIdxs = []int32{
	4,  // 0: global_entry_notify_api.UpdateNotificationDetailsRq.notificationDetails:type_name -> global_entry_notify_api.NotificationDetails
	8,  // 1: global_entry_notify_api.NotificationDetails.targetDate:type_name -> google.protobuf.Timestamp
	8,  // 2: global_entry_notify_api.NotificationDetails.appointmentDate:type_name -> google.protobuf.Timestamp
	8,  // 3: global_entry_notify_api.NotificationDetails.lastNotifiedDate:type_name -> google.protobuf.Timestamp
	4,  // 4: global_entry_notify_api.GetNotificationDetailsRp.notificationDetails:type_name -> global_entry_notify_api.NotificationDetails
	1,  // 5: global_entry_notify_api.GlobalEntryNotifyService.GrpcUpdateNotificationToken:input_type -> global_entry_notify_api.UpdateNotificationTokenRq
	2,  // 6: global_entry_notify_api.GlobalEntryNotifyService.GrpcUpdateNotificationDetails:input_type -> global_entry_notify_api.UpdateNotificationDetailsRq
	3,  // 7: global_entry_notify_api.GlobalEntryNotifyService.GrpcDeleteNotificationDetails:input_type -> global_entry_notify_api.DeleteNotificationDetailsRq
	9,  // 8: global_entry_notify_api.GlobalEntryNotifyService.GrpcGetLocations:input_type -> google.protobuf.Empty
	6,  // 9: global_entry_notify_api.GlobalEntryNotifyService.GrpcGetNotificationDetails:input_type -> global_entry_notify_api.GetNotificationDetailsRq
	0,  // 10: global_entry_notify_api.GlobalEntryNotifyService.GrpcUpdateNotificationToken:output_type -> global_entry_notify_api.Error
	0,  // 11: global_entry_notify_api.GlobalEntryNotifyService.GrpcUpdateNotificationDetails:output_type -> global_entry_notify_api.Error
	0,  // 12: global_entry_notify_api.GlobalEntryNotifyService.GrpcDeleteNotificationDetails:output_type -> global_entry_notify_api.Error
	5,  // 13: global_entry_notify_api.GlobalEntryNotifyService.GrpcGetLocations:output_type -> global_entry_notify_api.LocationsRp
	7,  // 14: global_entry_notify_api.GlobalEntryNotifyService.GrpcGetNotificationDetails:output_type -> global_entry_notify_api.GetNotificationDetailsRp
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_global_entry_notify_api_proto_init() }
func file_global_entry_notify_api_proto_init() {
	if File_global_entry_notify_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_global_entry_notify_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_global_entry_notify_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNotificationTokenRq); i {
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
		file_global_entry_notify_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNotificationDetailsRq); i {
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
		file_global_entry_notify_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNotificationDetailsRq); i {
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
		file_global_entry_notify_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationDetails); i {
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
		file_global_entry_notify_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationsRp); i {
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
		file_global_entry_notify_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationDetailsRq); i {
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
		file_global_entry_notify_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNotificationDetailsRp); i {
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
			RawDescriptor: file_global_entry_notify_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_global_entry_notify_api_proto_goTypes,
		DependencyIndexes: file_global_entry_notify_api_proto_depIdxs,
		MessageInfos:      file_global_entry_notify_api_proto_msgTypes,
	}.Build()
	File_global_entry_notify_api_proto = out.File
	file_global_entry_notify_api_proto_rawDesc = nil
	file_global_entry_notify_api_proto_goTypes = nil
	file_global_entry_notify_api_proto_depIdxs = nil
}
