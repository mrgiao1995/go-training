// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: customer.proto

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

type Customer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DateOfBirth *Date  `protobuf:"bytes,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
	Address     string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Email       string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *Customer) Reset() {
	*x = Customer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Customer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Customer) ProtoMessage() {}

func (x *Customer) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Customer.ProtoReflect.Descriptor instead.
func (*Customer) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Customer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Customer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Customer) GetDateOfBirth() *Date {
	if x != nil {
		return x.DateOfBirth
	}
	return nil
}

func (x *Customer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Customer) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Customer) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangeCustomerPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OldPassword string `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword string `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
}

func (x *ChangeCustomerPasswordRequest) Reset() {
	*x = ChangeCustomerPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeCustomerPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeCustomerPasswordRequest) ProtoMessage() {}

func (x *ChangeCustomerPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeCustomerPasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangeCustomerPasswordRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{1}
}

func (x *ChangeCustomerPasswordRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeCustomerPasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangeCustomerPasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type FindCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindCustomerRequest) Reset() {
	*x = FindCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindCustomerRequest) ProtoMessage() {}

func (x *FindCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindCustomerRequest.ProtoReflect.Descriptor instead.
func (*FindCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{2}
}

func (x *FindCustomerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ViewCustomerBookingHistoriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ViewCustomerBookingHistoriesRequest) Reset() {
	*x = ViewCustomerBookingHistoriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCustomerBookingHistoriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCustomerBookingHistoriesRequest) ProtoMessage() {}

func (x *ViewCustomerBookingHistoriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCustomerBookingHistoriesRequest.ProtoReflect.Descriptor instead.
func (*ViewCustomerBookingHistoriesRequest) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{3}
}

func (x *ViewCustomerBookingHistoriesRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ViewBooking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingCode string  `protobuf:"bytes,1,opt,name=booking_code,json=bookingCode,proto3" json:"booking_code,omitempty"`
	BookingDate *Date   `protobuf:"bytes,2,opt,name=booking_date,json=bookingDate,proto3" json:"booking_date,omitempty"`
	Flight      *Flight `protobuf:"bytes,3,opt,name=flight,proto3" json:"flight,omitempty"`
}

func (x *ViewBooking) Reset() {
	*x = ViewBooking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewBooking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewBooking) ProtoMessage() {}

func (x *ViewBooking) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewBooking.ProtoReflect.Descriptor instead.
func (*ViewBooking) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{4}
}

func (x *ViewBooking) GetBookingCode() string {
	if x != nil {
		return x.BookingCode
	}
	return ""
}

func (x *ViewBooking) GetBookingDate() *Date {
	if x != nil {
		return x.BookingDate
	}
	return nil
}

func (x *ViewBooking) GetFlight() *Flight {
	if x != nil {
		return x.Flight
	}
	return nil
}

type ViewCustomerBookingHistoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*ViewBooking `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
}

func (x *ViewCustomerBookingHistoriesResponse) Reset() {
	*x = ViewCustomerBookingHistoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewCustomerBookingHistoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewCustomerBookingHistoriesResponse) ProtoMessage() {}

func (x *ViewCustomerBookingHistoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewCustomerBookingHistoriesResponse.ProtoReflect.Descriptor instead.
func (*ViewCustomerBookingHistoriesResponse) Descriptor() ([]byte, []int) {
	return file_customer_proto_rawDescGZIP(), []int{5}
}

func (x *ViewCustomerBookingHistoriesResponse) GetBookings() []*ViewBooking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

var File_customer_proto protoreflect.FileDescriptor

var file_customer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x01, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x75, 0x0a, 0x1d,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x23, 0x56, 0x69,
	0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x56, 0x69, 0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x52, 0x06, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x22, 0x56, 0x0a, 0x24, 0x56,
	0x69, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69,
	0x65, 0x77, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x73, 0x32, 0xfb, 0x02, 0x0a, 0x0a, 0x4d, 0x79, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x12, 0x32, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x0f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x16, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x77, 0x0a, 0x1c, 0x56, 0x69, 0x65, 0x77,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x56, 0x69, 0x65, 0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x69, 0x65,
	0x77, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_customer_proto_rawDescOnce sync.Once
	file_customer_proto_rawDescData = file_customer_proto_rawDesc
)

func file_customer_proto_rawDescGZIP() []byte {
	file_customer_proto_rawDescOnce.Do(func() {
		file_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_proto_rawDescData)
	})
	return file_customer_proto_rawDescData
}

var file_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_customer_proto_goTypes = []interface{}{
	(*Customer)(nil),                             // 0: proto.Customer
	(*ChangeCustomerPasswordRequest)(nil),        // 1: proto.ChangeCustomerPasswordRequest
	(*FindCustomerRequest)(nil),                  // 2: proto.FindCustomerRequest
	(*ViewCustomerBookingHistoriesRequest)(nil),  // 3: proto.ViewCustomerBookingHistoriesRequest
	(*ViewBooking)(nil),                          // 4: proto.ViewBooking
	(*ViewCustomerBookingHistoriesResponse)(nil), // 5: proto.ViewCustomerBookingHistoriesResponse
	(*Date)(nil),                                 // 6: proto.Date
	(*Flight)(nil),                               // 7: proto.Flight
	(*Empty)(nil),                                // 8: proto.Empty
}
var file_customer_proto_depIdxs = []int32{
	6, // 0: proto.Customer.date_of_birth:type_name -> proto.Date
	6, // 1: proto.ViewBooking.booking_date:type_name -> proto.Date
	7, // 2: proto.ViewBooking.flight:type_name -> proto.Flight
	4, // 3: proto.ViewCustomerBookingHistoriesResponse.bookings:type_name -> proto.ViewBooking
	0, // 4: proto.MyCustomer.CreateCustomer:input_type -> proto.Customer
	0, // 5: proto.MyCustomer.UpdateCustomer:input_type -> proto.Customer
	2, // 6: proto.MyCustomer.CustomerDetails:input_type -> proto.FindCustomerRequest
	1, // 7: proto.MyCustomer.ChangeCustomerPassword:input_type -> proto.ChangeCustomerPasswordRequest
	3, // 8: proto.MyCustomer.ViewCustomerBookingHistories:input_type -> proto.ViewCustomerBookingHistoriesRequest
	0, // 9: proto.MyCustomer.CreateCustomer:output_type -> proto.Customer
	0, // 10: proto.MyCustomer.UpdateCustomer:output_type -> proto.Customer
	0, // 11: proto.MyCustomer.CustomerDetails:output_type -> proto.Customer
	8, // 12: proto.MyCustomer.ChangeCustomerPassword:output_type -> proto.Empty
	5, // 13: proto.MyCustomer.ViewCustomerBookingHistories:output_type -> proto.ViewCustomerBookingHistoriesResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_customer_proto_init() }
func file_customer_proto_init() {
	if File_customer_proto != nil {
		return
	}
	file_empty_proto_init()
	file_flight_proto_init()
	file_date_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Customer); i {
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
		file_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeCustomerPasswordRequest); i {
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
		file_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindCustomerRequest); i {
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
		file_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCustomerBookingHistoriesRequest); i {
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
		file_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewBooking); i {
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
		file_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewCustomerBookingHistoriesResponse); i {
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
			RawDescriptor: file_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_proto_goTypes,
		DependencyIndexes: file_customer_proto_depIdxs,
		MessageInfos:      file_customer_proto_msgTypes,
	}.Build()
	File_customer_proto = out.File
	file_customer_proto_rawDesc = nil
	file_customer_proto_goTypes = nil
	file_customer_proto_depIdxs = nil
}
