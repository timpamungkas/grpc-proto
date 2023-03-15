// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.12.4
// source: proto/bank/type/transaction.proto

package bank

import (
	date "google.golang.org/genproto/googleapis/type/date"
	datetime "google.golang.org/genproto/googleapis/type/datetime"
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

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_TRANSACTION_TYPE_IN          TransactionType = 1
	TransactionType_TRANSACTION_TYPE_OUT         TransactionType = 2
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "TRANSACTION_TYPE_UNSPECIFIED",
		1: "TRANSACTION_TYPE_IN",
		2: "TRANSACTION_TYPE_OUT",
	}
	TransactionType_value = map[string]int32{
		"TRANSACTION_TYPE_UNSPECIFIED": 0,
		"TRANSACTION_TYPE_IN":          1,
		"TRANSACTION_TYPE_OUT":         2,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_bank_type_transaction_proto_enumTypes[0].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_proto_bank_type_transaction_proto_enumTypes[0]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_proto_bank_type_transaction_proto_rawDescGZIP(), []int{0}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber string             `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
	Type          TransactionType    `protobuf:"varint,2,opt,name=type,proto3,enum=bank.TransactionType" json:"type,omitempty"`
	Amount        float64            `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Timestamp     *datetime.DateTime `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Notes         string             `protobuf:"bytes,16,opt,name=notes,proto3" json:"notes,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_type_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *Transaction) GetType() TransactionType {
	if x != nil {
		return x.Type
	}
	return TransactionType_TRANSACTION_TYPE_UNSPECIFIED
}

func (x *Transaction) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Transaction) GetTimestamp() *datetime.DateTime {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Transaction) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

type TransactionSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountNumber   string     `protobuf:"bytes,1,opt,name=account_number,proto3" json:"account_number,omitempty"`
	SumAmountIn     float64    `protobuf:"fixed64,2,opt,name=sum_amount_in,proto3" json:"sum_amount_in,omitempty"`
	SumAmountOut    float64    `protobuf:"fixed64,3,opt,name=sum_amount_out,proto3" json:"sum_amount_out,omitempty"`
	SumTotal        float64    `protobuf:"fixed64,4,opt,name=sum_total,proto3" json:"sum_total,omitempty"`
	TransactionDate *date.Date `protobuf:"bytes,5,opt,name=transaction_date,proto3" json:"transaction_date,omitempty"`
}

func (x *TransactionSummary) Reset() {
	*x = TransactionSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_bank_type_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionSummary) ProtoMessage() {}

func (x *TransactionSummary) ProtoReflect() protoreflect.Message {
	mi := &file_proto_bank_type_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionSummary.ProtoReflect.Descriptor instead.
func (*TransactionSummary) Descriptor() ([]byte, []int) {
	return file_proto_bank_type_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionSummary) GetAccountNumber() string {
	if x != nil {
		return x.AccountNumber
	}
	return ""
}

func (x *TransactionSummary) GetSumAmountIn() float64 {
	if x != nil {
		return x.SumAmountIn
	}
	return 0
}

func (x *TransactionSummary) GetSumAmountOut() float64 {
	if x != nil {
		return x.SumAmountOut
	}
	return 0
}

func (x *TransactionSummary) GetSumTotal() float64 {
	if x != nil {
		return x.SumTotal
	}
	return 0
}

func (x *TransactionSummary) GetTransactionDate() *date.Date {
	if x != nil {
		return x.TransactionDate
	}
	return nil
}

var File_proto_bank_type_transaction_proto protoreflect.FileDescriptor

var file_proto_bank_type_transaction_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x33, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x22,
	0xe7, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x75, 0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x75,
	0x6d, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x73, 0x75, 0x6d, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3d, 0x0a, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x2a, 0x66, 0x0a, 0x0f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x1c,
	0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17,
	0x0a, 0x13, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x54, 0x52, 0x41, 0x4e, 0x53,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x6d, 0x70, 0x61, 0x6d, 0x75, 0x6e, 0x67, 0x6b, 0x61, 0x73, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_bank_type_transaction_proto_rawDescOnce sync.Once
	file_proto_bank_type_transaction_proto_rawDescData = file_proto_bank_type_transaction_proto_rawDesc
)

func file_proto_bank_type_transaction_proto_rawDescGZIP() []byte {
	file_proto_bank_type_transaction_proto_rawDescOnce.Do(func() {
		file_proto_bank_type_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_bank_type_transaction_proto_rawDescData)
	})
	return file_proto_bank_type_transaction_proto_rawDescData
}

var file_proto_bank_type_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_bank_type_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_bank_type_transaction_proto_goTypes = []interface{}{
	(TransactionType)(0),       // 0: bank.TransactionType
	(*Transaction)(nil),        // 1: bank.Transaction
	(*TransactionSummary)(nil), // 2: bank.TransactionSummary
	(*datetime.DateTime)(nil),  // 3: google.type.DateTime
	(*date.Date)(nil),          // 4: google.type.Date
}
var file_proto_bank_type_transaction_proto_depIdxs = []int32{
	0, // 0: bank.Transaction.type:type_name -> bank.TransactionType
	3, // 1: bank.Transaction.timestamp:type_name -> google.type.DateTime
	4, // 2: bank.TransactionSummary.transaction_date:type_name -> google.type.Date
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_bank_type_transaction_proto_init() }
func file_proto_bank_type_transaction_proto_init() {
	if File_proto_bank_type_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_bank_type_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_proto_bank_type_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionSummary); i {
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
			RawDescriptor: file_proto_bank_type_transaction_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_bank_type_transaction_proto_goTypes,
		DependencyIndexes: file_proto_bank_type_transaction_proto_depIdxs,
		EnumInfos:         file_proto_bank_type_transaction_proto_enumTypes,
		MessageInfos:      file_proto_bank_type_transaction_proto_msgTypes,
	}.Build()
	File_proto_bank_type_transaction_proto = out.File
	file_proto_bank_type_transaction_proto_rawDesc = nil
	file_proto_bank_type_transaction_proto_goTypes = nil
	file_proto_bank_type_transaction_proto_depIdxs = nil
}
