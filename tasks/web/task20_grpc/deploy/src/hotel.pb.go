// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: hotel.proto

package hotel

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

// The request message containing the room type.
type AvailabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomType     string `protobuf:"bytes,1,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	CheckInDate  string `protobuf:"bytes,2,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate string `protobuf:"bytes,3,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
}

func (x *AvailabilityRequest) Reset() {
	*x = AvailabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailabilityRequest) ProtoMessage() {}

func (x *AvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailabilityRequest.ProtoReflect.Descriptor instead.
func (*AvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{0}
}

func (x *AvailabilityRequest) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *AvailabilityRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *AvailabilityRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

// The response message containing the availability status.
type AvailabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Available bool   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AvailabilityResponse) Reset() {
	*x = AvailabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailabilityResponse) ProtoMessage() {}

func (x *AvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailabilityResponse.ProtoReflect.Descriptor instead.
func (*AvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{1}
}

func (x *AvailabilityResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *AvailabilityResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The request message for booking a room.
type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomType     string `protobuf:"bytes,1,opt,name=room_type,json=roomType,proto3" json:"room_type,omitempty"`
	GuestName    string `protobuf:"bytes,2,opt,name=guest_name,json=guestName,proto3" json:"guest_name,omitempty"`
	CheckInDate  string `protobuf:"bytes,3,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`
	CheckOutDate string `protobuf:"bytes,4,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{2}
}

func (x *BookingRequest) GetRoomType() string {
	if x != nil {
		return x.RoomType
	}
	return ""
}

func (x *BookingRequest) GetGuestName() string {
	if x != nil {
		return x.GuestName
	}
	return ""
}

func (x *BookingRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *BookingRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

// The response message for a booking request.
type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success            bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ConfirmationNumber string `protobuf:"bytes,2,opt,name=confirmation_number,json=confirmationNumber,proto3" json:"confirmation_number,omitempty"`
	Message            string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{3}
}

func (x *BookingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BookingResponse) GetConfirmationNumber() string {
	if x != nil {
		return x.ConfirmationNumber
	}
	return ""
}

func (x *BookingResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The request message for canceling a booking.
type CancellationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfirmationNumber string `protobuf:"bytes,1,opt,name=confirmation_number,json=confirmationNumber,proto3" json:"confirmation_number,omitempty"`
}

func (x *CancellationRequest) Reset() {
	*x = CancellationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancellationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancellationRequest) ProtoMessage() {}

func (x *CancellationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancellationRequest.ProtoReflect.Descriptor instead.
func (*CancellationRequest) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{4}
}

func (x *CancellationRequest) GetConfirmationNumber() string {
	if x != nil {
		return x.ConfirmationNumber
	}
	return ""
}

// The response message for a cancellation request.
type CancellationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CancellationResponse) Reset() {
	*x = CancellationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hotel_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancellationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancellationResponse) ProtoMessage() {}

func (x *CancellationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hotel_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancellationResponse.ProtoReflect.Descriptor instead.
func (*CancellationResponse) Descriptor() ([]byte, []int) {
	return file_hotel_proto_rawDescGZIP(), []int{5}
}

func (x *CancellationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CancellationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_hotel_proto protoreflect.FileDescriptor

var file_hotel_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x6f, 0x74, 0x65, 0x6c, 0x22, 0x7c, 0x0a, 0x13, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x4e, 0x0a, 0x14, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x6f,
	0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x76, 0x0a, 0x0f, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x13, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x14, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xe1, 0x01, 0x0a, 0x0c, 0x48, 0x6f, 0x74, 0x65,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x2e,
	0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x15, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x6f, 0x74, 0x65,
	0x6c, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x74, 0x72, 0x6f,
	0x76, 0x6f, 0x6b, 0x2f, 0x68, 0x6f, 0x74, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_hotel_proto_rawDescOnce sync.Once
	file_hotel_proto_rawDescData = file_hotel_proto_rawDesc
)

func file_hotel_proto_rawDescGZIP() []byte {
	file_hotel_proto_rawDescOnce.Do(func() {
		file_hotel_proto_rawDescData = protoimpl.X.CompressGZIP(file_hotel_proto_rawDescData)
	})
	return file_hotel_proto_rawDescData
}

var file_hotel_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_hotel_proto_goTypes = []any{
	(*AvailabilityRequest)(nil),  // 0: hotel.AvailabilityRequest
	(*AvailabilityResponse)(nil), // 1: hotel.AvailabilityResponse
	(*BookingRequest)(nil),       // 2: hotel.BookingRequest
	(*BookingResponse)(nil),      // 3: hotel.BookingResponse
	(*CancellationRequest)(nil),  // 4: hotel.CancellationRequest
	(*CancellationResponse)(nil), // 5: hotel.CancellationResponse
}
var file_hotel_proto_depIdxs = []int32{
	0, // 0: hotel.HotelService.CheckAvailability:input_type -> hotel.AvailabilityRequest
	2, // 1: hotel.HotelService.BookRoom:input_type -> hotel.BookingRequest
	4, // 2: hotel.HotelService.CancelBooking:input_type -> hotel.CancellationRequest
	1, // 3: hotel.HotelService.CheckAvailability:output_type -> hotel.AvailabilityResponse
	3, // 4: hotel.HotelService.BookRoom:output_type -> hotel.BookingResponse
	5, // 5: hotel.HotelService.CancelBooking:output_type -> hotel.CancellationResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hotel_proto_init() }
func file_hotel_proto_init() {
	if File_hotel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hotel_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AvailabilityRequest); i {
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
		file_hotel_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AvailabilityResponse); i {
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
		file_hotel_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*BookingRequest); i {
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
		file_hotel_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*BookingResponse); i {
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
		file_hotel_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CancellationRequest); i {
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
		file_hotel_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CancellationResponse); i {
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
			RawDescriptor: file_hotel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hotel_proto_goTypes,
		DependencyIndexes: file_hotel_proto_depIdxs,
		MessageInfos:      file_hotel_proto_msgTypes,
	}.Build()
	File_hotel_proto = out.File
	file_hotel_proto_rawDesc = nil
	file_hotel_proto_goTypes = nil
	file_hotel_proto_depIdxs = nil
}
