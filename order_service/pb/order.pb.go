// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: order.proto

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

type RequestOrderData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     uint32       `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Address    *Address     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	OrderItems []*OrderItem `protobuf:"bytes,3,rep,name=order_items,json=orderItems,proto3" json:"order_items,omitempty"`
}

func (x *RequestOrderData) Reset() {
	*x = RequestOrderData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestOrderData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestOrderData) ProtoMessage() {}

func (x *RequestOrderData) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestOrderData.ProtoReflect.Descriptor instead.
func (*RequestOrderData) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *RequestOrderData) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *RequestOrderData) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RequestOrderData) GetOrderItems() []*OrderItem {
	if x != nil {
		return x.OrderItems
	}
	return nil
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MenuId string `protobuf:"bytes,1,opt,name=menu_id,json=menuId,proto3" json:"menu_id,omitempty"`
	Qty    uint32 `protobuf:"varint,2,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *OrderItem) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

func (x *OrderItem) GetQty() uint32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type RestaurantId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RestaurantId) Reset() {
	*x = RestaurantId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantId) ProtoMessage() {}

func (x *RestaurantId) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantId.ProtoReflect.Descriptor instead.
func (*RestaurantId) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *RestaurantId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DriverId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DriverId) Reset() {
	*x = DriverId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverId) ProtoMessage() {}

func (x *DriverId) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverId.ProtoReflect.Descriptor instead.
func (*DriverId) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *DriverId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *UserId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type OrderId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *OrderId) Reset() {
	*x = OrderId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderId) ProtoMessage() {}

func (x *OrderId) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderId.ProtoReflect.Descriptor instead.
func (*OrderId) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *OrderId) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *Orders) GetOrders() []*Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurant  *Restaurant  `protobuf:"bytes,1,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	OrderDetail *OrderDetail `protobuf:"bytes,2,opt,name=order_detail,json=orderDetail,proto3" json:"order_detail,omitempty"`
	User        *User        `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Driver      *Driver      `protobuf:"bytes,4,opt,name=driver,proto3" json:"driver,omitempty"`
	Payment     *Payment     `protobuf:"bytes,5,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *Order) GetRestaurant() *Restaurant {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *Order) GetOrderDetail() *OrderDetail {
	if x != nil {
		return x.OrderDetail
	}
	return nil
}

func (x *Order) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Order) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *Order) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{8}
}

func (x *Driver) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Driver) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address *Address `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{9}
}

func (x *User) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{10}
}

func (x *Address) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Address) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Restaurant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AdminId uint32   `protobuf:"varint,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	Name    string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address *Address `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Status  string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Restaurant) Reset() {
	*x = Restaurant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurant) ProtoMessage() {}

func (x *Restaurant) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurant.ProtoReflect.Descriptor instead.
func (*Restaurant) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{11}
}

func (x *Restaurant) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Restaurant) GetAdminId() uint32 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *Restaurant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Restaurant) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Restaurant) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Quantity uint32  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Subtotal float32 `protobuf:"fixed32,3,opt,name=subtotal,proto3" json:"subtotal,omitempty"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{12}
}

func (x *Menu) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Menu) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Menu) GetSubtotal() float32 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

type OrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu  []*Menu `protobuf:"bytes,1,rep,name=menu,proto3" json:"menu,omitempty"`
	Total float32 `protobuf:"fixed32,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *OrderDetail) Reset() {
	*x = OrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderDetail) ProtoMessage() {}

func (x *OrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderDetail.ProtoReflect.Descriptor instead.
func (*OrderDetail) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{13}
}

func (x *OrderDetail) GetMenu() []*Menu {
	if x != nil {
		return x.Menu
	}
	return nil
}

func (x *OrderDetail) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InvoiceId string  `protobuf:"bytes,2,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	Total     float32 `protobuf:"fixed32,3,opt,name=total,proto3" json:"total,omitempty"`
	Method    string  `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Status    string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{14}
}

func (x *Payment) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Payment) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *Payment) GetTotal() float32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Payment) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Payment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x36, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x6d, 0x65, 0x6e,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6e, 0x75,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x71, 0x74, 0x79, 0x22, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x08, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x19, 0x0a, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12,
	0x23, 0x0a, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x22, 0xde, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x30,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x12, 0x34, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x53, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x8c, 0x01,
	0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4e, 0x0a, 0x04,
	0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x0b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x6d,
	0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x22, 0x7e, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x81, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x30, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x1a,
	0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x0c,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2d, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_order_proto_goTypes = []interface{}{
	(*RequestOrderData)(nil), // 0: user.RequestOrderData
	(*OrderItem)(nil),        // 1: user.OrderItem
	(*RestaurantId)(nil),     // 2: user.RestaurantId
	(*DriverId)(nil),         // 3: user.DriverId
	(*UserId)(nil),           // 4: user.UserId
	(*OrderId)(nil),          // 5: user.OrderId
	(*Orders)(nil),           // 6: user.Orders
	(*Order)(nil),            // 7: user.Order
	(*Driver)(nil),           // 8: user.Driver
	(*User)(nil),             // 9: user.User
	(*Address)(nil),          // 10: user.Address
	(*Restaurant)(nil),       // 11: user.Restaurant
	(*Menu)(nil),             // 12: user.Menu
	(*OrderDetail)(nil),      // 13: user.OrderDetail
	(*Payment)(nil),          // 14: user.Payment
}
var file_order_proto_depIdxs = []int32{
	10, // 0: user.RequestOrderData.address:type_name -> user.Address
	1,  // 1: user.RequestOrderData.order_items:type_name -> user.OrderItem
	7,  // 2: user.Orders.orders:type_name -> user.Order
	11, // 3: user.Order.restaurant:type_name -> user.Restaurant
	13, // 4: user.Order.order_detail:type_name -> user.OrderDetail
	9,  // 5: user.Order.user:type_name -> user.User
	8,  // 6: user.Order.driver:type_name -> user.Driver
	14, // 7: user.Order.payment:type_name -> user.Payment
	10, // 8: user.User.address:type_name -> user.Address
	10, // 9: user.Restaurant.address:type_name -> user.Address
	12, // 10: user.OrderDetail.menu:type_name -> user.Menu
	0,  // 11: user.OrderService.PostOrder:input_type -> user.RequestOrderData
	2,  // 12: user.OrderService.GetRestaurantOrders:input_type -> user.RestaurantId
	4,  // 13: user.OrderService.GetUserOrders:input_type -> user.UserId
	3,  // 14: user.OrderService.GetDriverOrder:input_type -> user.DriverId
	5,  // 15: user.OrderService.GetOrderById:input_type -> user.OrderId
	7,  // 16: user.OrderService.PostOrder:output_type -> user.Order
	6,  // 17: user.OrderService.GetRestaurantOrders:output_type -> user.Orders
	6,  // 18: user.OrderService.GetUserOrders:output_type -> user.Orders
	7,  // 19: user.OrderService.GetDriverOrder:output_type -> user.Order
	7,  // 20: user.OrderService.GetOrderById:output_type -> user.Order
	16, // [16:21] is the sub-list for method output_type
	11, // [11:16] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestOrderData); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantId); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriverId); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderId); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
		file_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
		file_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurant); i {
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
		file_order_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
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
		file_order_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderDetail); i {
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
		file_order_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
