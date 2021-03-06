// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: api/booksearch.proto

package kvadotest

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

// Books search request is either by author, or by content or by title
type SearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*SearchRequest_ByAuthor
	//	*SearchRequest_ByContent
	//	*SearchRequest_ByTitle
	Request isSearchRequest_Request `protobuf_oneof:"request"`
}

func (x *SearchRequest) Reset() {
	*x = SearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_booksearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchRequest) ProtoMessage() {}

func (x *SearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_booksearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchRequest.ProtoReflect.Descriptor instead.
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return file_api_booksearch_proto_rawDescGZIP(), []int{0}
}

func (m *SearchRequest) GetRequest() isSearchRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *SearchRequest) GetByAuthor() string {
	if x, ok := x.GetRequest().(*SearchRequest_ByAuthor); ok {
		return x.ByAuthor
	}
	return ""
}

func (x *SearchRequest) GetByContent() string {
	if x, ok := x.GetRequest().(*SearchRequest_ByContent); ok {
		return x.ByContent
	}
	return ""
}

func (x *SearchRequest) GetByTitle() string {
	if x, ok := x.GetRequest().(*SearchRequest_ByTitle); ok {
		return x.ByTitle
	}
	return ""
}

type isSearchRequest_Request interface {
	isSearchRequest_Request()
}

type SearchRequest_ByAuthor struct {
	ByAuthor string `protobuf:"bytes,1,opt,name=by_author,json=byAuthor,proto3,oneof"`
}

type SearchRequest_ByContent struct {
	ByContent string `protobuf:"bytes,2,opt,name=by_content,json=byContent,proto3,oneof"`
}

type SearchRequest_ByTitle struct {
	ByTitle string `protobuf:"bytes,3,opt,name=by_title,json=byTitle,proto3,oneof"`
}

func (*SearchRequest_ByAuthor) isSearchRequest_Request() {}

func (*SearchRequest_ByContent) isSearchRequest_Request() {}

func (*SearchRequest_ByTitle) isSearchRequest_Request() {}

// Books list response
type SearchReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *SearchReply) Reset() {
	*x = SearchReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_booksearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchReply) ProtoMessage() {}

func (x *SearchReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_booksearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchReply.ProtoReflect.Descriptor instead.
func (*SearchReply) Descriptor() ([]byte, []int) {
	return file_api_booksearch_proto_rawDescGZIP(), []int{1}
}

func (x *SearchReply) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

// Book description
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []string `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_booksearch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_api_booksearch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_api_booksearch_proto_rawDescGZIP(), []int{2}
}

func (x *Book) GetAuthors() []string {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_api_booksearch_proto protoreflect.FileDescriptor

var file_api_booksearch_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x22, 0x77, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x79, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x62, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x62, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x62, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x0b, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x22, 0x50, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x32, 0x4a, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x3c, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x70, 0x72, 0x74, 0x72, 0x32, 0x35, 0x38, 0x2f, 0x6b, 0x76, 0x61, 0x64, 0x6f, 0x74, 0x65, 0x73,
	0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_booksearch_proto_rawDescOnce sync.Once
	file_api_booksearch_proto_rawDescData = file_api_booksearch_proto_rawDesc
)

func file_api_booksearch_proto_rawDescGZIP() []byte {
	file_api_booksearch_proto_rawDescOnce.Do(func() {
		file_api_booksearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_booksearch_proto_rawDescData)
	})
	return file_api_booksearch_proto_rawDescData
}

var file_api_booksearch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_booksearch_proto_goTypes = []interface{}{
	(*SearchRequest)(nil), // 0: booksearch.SearchRequest
	(*SearchReply)(nil),   // 1: booksearch.SearchReply
	(*Book)(nil),          // 2: booksearch.Book
}
var file_api_booksearch_proto_depIdxs = []int32{
	2, // 0: booksearch.SearchReply.books:type_name -> booksearch.Book
	0, // 1: booksearch.BookSearch.Search:input_type -> booksearch.SearchRequest
	1, // 2: booksearch.BookSearch.Search:output_type -> booksearch.SearchReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_booksearch_proto_init() }
func file_api_booksearch_proto_init() {
	if File_api_booksearch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_booksearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchRequest); i {
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
		file_api_booksearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchReply); i {
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
		file_api_booksearch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
	file_api_booksearch_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SearchRequest_ByAuthor)(nil),
		(*SearchRequest_ByContent)(nil),
		(*SearchRequest_ByTitle)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_booksearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_booksearch_proto_goTypes,
		DependencyIndexes: file_api_booksearch_proto_depIdxs,
		MessageInfos:      file_api_booksearch_proto_msgTypes,
	}.Build()
	File_api_booksearch_proto = out.File
	file_api_booksearch_proto_rawDesc = nil
	file_api_booksearch_proto_goTypes = nil
	file_api_booksearch_proto_depIdxs = nil
}
