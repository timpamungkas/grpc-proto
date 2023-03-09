// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.12.4
// source: proto/transaction/cart.proto

package transaction

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

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemUuid  string `protobuf:"bytes,1,opt,name=item_uuid,json=item_id,proto3" json:"item_uuid,omitempty"`
	Quantity  uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ItemPrice uint32 `protobuf:"varint,3,opt,name=item_price,proto3" json:"item_price,omitempty"`
	Taxable   bool   `protobuf:"varint,4,opt,name=taxable,proto3" json:"taxable,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_proto_transaction_cart_proto_rawDescGZIP(), []int{0}
}

func (x *CartItem) GetItemUuid() string {
	if x != nil {
		return x.ItemUuid
	}
	return ""
}

func (x *CartItem) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetItemPrice() uint32 {
	if x != nil {
		return x.ItemPrice
	}
	return 0
}

func (x *CartItem) GetTaxable() bool {
	if x != nil {
		return x.Taxable
	}
	return false
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartUuid string      `protobuf:"bytes,1,opt,name=cart_uuid,proto3" json:"cart_uuid,omitempty"`
	Items    []*CartItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_transaction_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_proto_transaction_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_proto_transaction_cart_proto_rawDescGZIP(), []int{1}
}

func (x *Cart) GetCartUuid() string {
	if x != nil {
		return x.CartUuid
	}
	return ""
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_proto_transaction_cart_proto protoreflect.FileDescriptor

var file_proto_transaction_cart_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x08, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x74, 0x61, 0x78, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x51, 0x0a, 0x04, 0x43, 0x61, 0x72,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x2b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x6d, 0x70, 0x61,
	0x6d, 0x75, 0x6e, 0x67, 0x6b, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_transaction_cart_proto_rawDescOnce sync.Once
	file_proto_transaction_cart_proto_rawDescData = file_proto_transaction_cart_proto_rawDesc
)

func file_proto_transaction_cart_proto_rawDescGZIP() []byte {
	file_proto_transaction_cart_proto_rawDescOnce.Do(func() {
		file_proto_transaction_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_transaction_cart_proto_rawDescData)
	})
	return file_proto_transaction_cart_proto_rawDescData
}

var file_proto_transaction_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_transaction_cart_proto_goTypes = []interface{}{
	(*CartItem)(nil), // 0: transaction.CartItem
	(*Cart)(nil),     // 1: transaction.Cart
}
var file_proto_transaction_cart_proto_depIdxs = []int32{
	0, // 0: transaction.Cart.items:type_name -> transaction.CartItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_transaction_cart_proto_init() }
func file_proto_transaction_cart_proto_init() {
	if File_proto_transaction_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_transaction_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_proto_transaction_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
			RawDescriptor: file_proto_transaction_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_transaction_cart_proto_goTypes,
		DependencyIndexes: file_proto_transaction_cart_proto_depIdxs,
		MessageInfos:      file_proto_transaction_cart_proto_msgTypes,
	}.Build()
	File_proto_transaction_cart_proto = out.File
	file_proto_transaction_cart_proto_rawDesc = nil
	file_proto_transaction_cart_proto_goTypes = nil
	file_proto_transaction_cart_proto_depIdxs = nil
}
