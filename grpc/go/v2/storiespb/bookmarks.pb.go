// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: bookmarks.proto

package storiespb

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

// LoadBookmarkType ...
type LoadBookmarkType int32

const (
	LoadBookmarkType_LBT_IGNORE    LoadBookmarkType = 0
	LoadBookmarkType_LBT_BOOKMARK  LoadBookmarkType = 1
	LoadBookmarkType_LBT_READLATER LoadBookmarkType = 2
)

// Enum value maps for LoadBookmarkType.
var (
	LoadBookmarkType_name = map[int32]string{
		0: "LBT_IGNORE",
		1: "LBT_BOOKMARK",
		2: "LBT_READLATER",
	}
	LoadBookmarkType_value = map[string]int32{
		"LBT_IGNORE":    0,
		"LBT_BOOKMARK":  1,
		"LBT_READLATER": 2,
	}
)

func (x LoadBookmarkType) Enum() *LoadBookmarkType {
	p := new(LoadBookmarkType)
	*p = x
	return p
}

func (x LoadBookmarkType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoadBookmarkType) Descriptor() protoreflect.EnumDescriptor {
	return file_bookmarks_proto_enumTypes[0].Descriptor()
}

func (LoadBookmarkType) Type() protoreflect.EnumType {
	return &file_bookmarks_proto_enumTypes[0]
}

func (x LoadBookmarkType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoadBookmarkType.Descriptor instead.
func (LoadBookmarkType) EnumDescriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{0}
}

type BookmarkSort_FieldBy int32

const (
	BookmarkSort_IGNORE            BookmarkSort_FieldBy = 0
	BookmarkSort_STORY_TITLE       BookmarkSort_FieldBy = 1
	BookmarkSort_STORY_WRITER      BookmarkSort_FieldBy = 2
	BookmarkSort_BOOKMARK_MIX_DATE BookmarkSort_FieldBy = 3
)

// Enum value maps for BookmarkSort_FieldBy.
var (
	BookmarkSort_FieldBy_name = map[int32]string{
		0: "IGNORE",
		1: "STORY_TITLE",
		2: "STORY_WRITER",
		3: "BOOKMARK_MIX_DATE",
	}
	BookmarkSort_FieldBy_value = map[string]int32{
		"IGNORE":            0,
		"STORY_TITLE":       1,
		"STORY_WRITER":      2,
		"BOOKMARK_MIX_DATE": 3,
	}
)

func (x BookmarkSort_FieldBy) Enum() *BookmarkSort_FieldBy {
	p := new(BookmarkSort_FieldBy)
	*p = x
	return p
}

func (x BookmarkSort_FieldBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BookmarkSort_FieldBy) Descriptor() protoreflect.EnumDescriptor {
	return file_bookmarks_proto_enumTypes[1].Descriptor()
}

func (BookmarkSort_FieldBy) Type() protoreflect.EnumType {
	return &file_bookmarks_proto_enumTypes[1]
}

func (x BookmarkSort_FieldBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BookmarkSort_FieldBy.Descriptor instead.
func (BookmarkSort_FieldBy) EnumDescriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{6, 0}
}

// RequestCreateBookmark ...
type RequestCreateBookmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID puede ser:
	//
	// ID de una historia el Path tiene q estar vacio.
	// ID de un capitulo el Path tiene q tener algun valor.
	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Abstract string `protobuf:"bytes,2,opt,name=Abstract,proto3" json:"Abstract,omitempty"`
	// Path ruta donde apunta el marcador
	//
	// Path vacio indica q el marcador es para leer mas tarde una historia
	// Path no vacio indica q es un marcador en el texto de un capitulo
	Path string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
}

func (x *RequestCreateBookmark) Reset() {
	*x = RequestCreateBookmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestCreateBookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestCreateBookmark) ProtoMessage() {}

func (x *RequestCreateBookmark) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestCreateBookmark.ProtoReflect.Descriptor instead.
func (*RequestCreateBookmark) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{0}
}

func (x *RequestCreateBookmark) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RequestCreateBookmark) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *RequestCreateBookmark) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// RequestEditBookmark ...
type RequestEditBookmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Abstract string `protobuf:"bytes,2,opt,name=Abstract,proto3" json:"Abstract,omitempty"`
	Path     string `protobuf:"bytes,3,opt,name=Path,proto3" json:"Path,omitempty"`
}

func (x *RequestEditBookmark) Reset() {
	*x = RequestEditBookmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestEditBookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestEditBookmark) ProtoMessage() {}

func (x *RequestEditBookmark) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestEditBookmark.ProtoReflect.Descriptor instead.
func (*RequestEditBookmark) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{1}
}

func (x *RequestEditBookmark) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RequestEditBookmark) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *RequestEditBookmark) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// RequestMyBookmarks ...
type RequestMyBookmarks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Load      LoadBookmarkType   `protobuf:"varint,1,opt,name=Load,proto3,enum=stories.LoadBookmarkType" json:"Load,omitempty"`
	LoadTotal bool               `protobuf:"varint,2,opt,name=LoadTotal,proto3" json:"LoadTotal,omitempty"`
	Filter    *FilterMyBookmarks `protobuf:"bytes,3,opt,name=Filter,proto3" json:"Filter,omitempty"`
	Sort      *BookmarkSort      `protobuf:"bytes,4,opt,name=Sort,proto3" json:"Sort,omitempty"`
	Paginator *Paginator         `protobuf:"bytes,5,opt,name=Paginator,proto3" json:"Paginator,omitempty"`
}

func (x *RequestMyBookmarks) Reset() {
	*x = RequestMyBookmarks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMyBookmarks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMyBookmarks) ProtoMessage() {}

func (x *RequestMyBookmarks) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMyBookmarks.ProtoReflect.Descriptor instead.
func (*RequestMyBookmarks) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{2}
}

func (x *RequestMyBookmarks) GetLoad() LoadBookmarkType {
	if x != nil {
		return x.Load
	}
	return LoadBookmarkType_LBT_IGNORE
}

func (x *RequestMyBookmarks) GetLoadTotal() bool {
	if x != nil {
		return x.LoadTotal
	}
	return false
}

func (x *RequestMyBookmarks) GetFilter() *FilterMyBookmarks {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *RequestMyBookmarks) GetSort() *BookmarkSort {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *RequestMyBookmarks) GetPaginator() *Paginator {
	if x != nil {
		return x.Paginator
	}
	return nil
}

// ResponseMyBookmarks ...
type ResponseMyBookmarks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total     uint64      `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Bookmarks []*Bookmark `protobuf:"bytes,2,rep,name=Bookmarks,proto3" json:"Bookmarks,omitempty"`
}

func (x *ResponseMyBookmarks) Reset() {
	*x = ResponseMyBookmarks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMyBookmarks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMyBookmarks) ProtoMessage() {}

func (x *ResponseMyBookmarks) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMyBookmarks.ProtoReflect.Descriptor instead.
func (*ResponseMyBookmarks) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseMyBookmarks) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ResponseMyBookmarks) GetBookmarks() []*Bookmark {
	if x != nil {
		return x.Bookmarks
	}
	return nil
}

// Bookmark ...
type Bookmark struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string       `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID    string       `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Story     *StorySimple `protobuf:"bytes,3,opt,name=Story,proto3" json:"Story,omitempty"`
	Chapter   *ChapterInfo `protobuf:"bytes,4,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
	Abstract  string       `protobuf:"bytes,6,opt,name=Abstract,proto3" json:"Abstract,omitempty"`
	Path      string       `protobuf:"bytes,7,opt,name=Path,proto3" json:"Path,omitempty"`
	CreatedAt int64        `protobuf:"varint,8,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt int64        `protobuf:"varint,9,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *Bookmark) Reset() {
	*x = Bookmark{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bookmark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bookmark) ProtoMessage() {}

func (x *Bookmark) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bookmark.ProtoReflect.Descriptor instead.
func (*Bookmark) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{4}
}

func (x *Bookmark) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Bookmark) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *Bookmark) GetStory() *StorySimple {
	if x != nil {
		return x.Story
	}
	return nil
}

func (x *Bookmark) GetChapter() *ChapterInfo {
	if x != nil {
		return x.Chapter
	}
	return nil
}

func (x *Bookmark) GetAbstract() string {
	if x != nil {
		return x.Abstract
	}
	return ""
}

func (x *Bookmark) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Bookmark) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Bookmark) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// FilterMyBookmarks ...
type FilterMyBookmarks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoryWriterID string `protobuf:"bytes,1,opt,name=StoryWriterID,proto3" json:"StoryWriterID,omitempty"`
	// WordsOnTitle palabras q contiene el titulo
	StoryWordsOnTitle     string     `protobuf:"bytes,2,opt,name=StoryWordsOnTitle,proto3" json:"StoryWordsOnTitle,omitempty"`
	StoryCategoriesID     []string   `protobuf:"bytes,3,rep,name=StoryCategoriesID,proto3" json:"StoryCategoriesID,omitempty"`
	StoryStartPublishedAt int64      `protobuf:"varint,4,opt,name=StoryStartPublishedAt,proto3" json:"StoryStartPublishedAt,omitempty"`
	StoryEndPublishedAt   int64      `protobuf:"varint,5,opt,name=StoryEndPublishedAt,proto3" json:"StoryEndPublishedAt,omitempty"`
	StoryIsFinished       BoolOption `protobuf:"varint,6,opt,name=StoryIsFinished,proto3,enum=stories.BoolOption" json:"StoryIsFinished,omitempty"`
}

func (x *FilterMyBookmarks) Reset() {
	*x = FilterMyBookmarks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterMyBookmarks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterMyBookmarks) ProtoMessage() {}

func (x *FilterMyBookmarks) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterMyBookmarks.ProtoReflect.Descriptor instead.
func (*FilterMyBookmarks) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{5}
}

func (x *FilterMyBookmarks) GetStoryWriterID() string {
	if x != nil {
		return x.StoryWriterID
	}
	return ""
}

func (x *FilterMyBookmarks) GetStoryWordsOnTitle() string {
	if x != nil {
		return x.StoryWordsOnTitle
	}
	return ""
}

func (x *FilterMyBookmarks) GetStoryCategoriesID() []string {
	if x != nil {
		return x.StoryCategoriesID
	}
	return nil
}

func (x *FilterMyBookmarks) GetStoryStartPublishedAt() int64 {
	if x != nil {
		return x.StoryStartPublishedAt
	}
	return 0
}

func (x *FilterMyBookmarks) GetStoryEndPublishedAt() int64 {
	if x != nil {
		return x.StoryEndPublishedAt
	}
	return 0
}

func (x *FilterMyBookmarks) GetStoryIsFinished() BoolOption {
	if x != nil {
		return x.StoryIsFinished
	}
	return BoolOption_BO_IGNORE
}

// BookmarkSort ...
type BookmarkSort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	By         BookmarkSort_FieldBy `protobuf:"varint,1,opt,name=By,proto3,enum=stories.BookmarkSort_FieldBy" json:"By,omitempty"`
	Descendent bool                 `protobuf:"varint,2,opt,name=Descendent,proto3" json:"Descendent,omitempty"`
}

func (x *BookmarkSort) Reset() {
	*x = BookmarkSort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmarks_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookmarkSort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkSort) ProtoMessage() {}

func (x *BookmarkSort) ProtoReflect() protoreflect.Message {
	mi := &file_bookmarks_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkSort.ProtoReflect.Descriptor instead.
func (*BookmarkSort) Descriptor() ([]byte, []int) {
	return file_bookmarks_proto_rawDescGZIP(), []int{6}
}

func (x *BookmarkSort) GetBy() BookmarkSort_FieldBy {
	if x != nil {
		return x.By
	}
	return BookmarkSort_IGNORE
}

func (x *BookmarkSort) GetDescendent() bool {
	if x != nil {
		return x.Descendent
	}
	return false
}

var File_bookmarks_proto protoreflect.FileDescriptor

var file_bookmarks_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x15, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x22, 0x55, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x64, 0x69,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x22, 0xf2, 0x01, 0x0a, 0x12, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x4c, 0x6f, 0x61, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4d, 0x79,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x29, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x09,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x09, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x5c,
	0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x09, 0x42,
	0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72,
	0x6b, 0x52, 0x09, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x22, 0xfa, 0x01, 0x0a,
	0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x2a, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x2e, 0x0a,
	0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x41, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xbc, 0x02, 0x0a, 0x11, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x4d, 0x79, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12,
	0x24, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x57, 0x6f,
	0x72, 0x64, 0x73, 0x4f, 0x6e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x4f, 0x6e, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49, 0x44, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x49,
	0x44, 0x12, 0x34, 0x0a, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x15, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x45, 0x6e, 0x64, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x6e, 0x64, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3d, 0x0a, 0x0f, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x49, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0xae, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x2d, 0x0a, 0x02, 0x42, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x6f, 0x72, 0x74, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x42, 0x79, 0x52, 0x02, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x63,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x44, 0x65,
	0x73, 0x63, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x07, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x42, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x4f, 0x4f, 0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x5f, 0x4d,
	0x49, 0x58, 0x5f, 0x44, 0x41, 0x54, 0x45, 0x10, 0x03, 0x2a, 0x47, 0x0a, 0x10, 0x4c, 0x6f, 0x61,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a,
	0x0a, 0x4c, 0x42, 0x54, 0x5f, 0x49, 0x47, 0x4e, 0x4f, 0x52, 0x45, 0x10, 0x00, 0x12, 0x10, 0x0a,
	0x0c, 0x4c, 0x42, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x4d, 0x41, 0x52, 0x4b, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x4c, 0x42, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x4c, 0x41, 0x54, 0x45, 0x52,
	0x10, 0x02, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x45, 0x72, 0x6e, 0x65, 0x73, 0x74, 0x6f, 0x47, 0x46, 0x2f, 0x62, 0x6f, 0x75, 0x6b, 0x6b,
	0x65, 0x72, 0x2d, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookmarks_proto_rawDescOnce sync.Once
	file_bookmarks_proto_rawDescData = file_bookmarks_proto_rawDesc
)

func file_bookmarks_proto_rawDescGZIP() []byte {
	file_bookmarks_proto_rawDescOnce.Do(func() {
		file_bookmarks_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookmarks_proto_rawDescData)
	})
	return file_bookmarks_proto_rawDescData
}

var file_bookmarks_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_bookmarks_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bookmarks_proto_goTypes = []interface{}{
	(LoadBookmarkType)(0),         // 0: stories.LoadBookmarkType
	(BookmarkSort_FieldBy)(0),     // 1: stories.BookmarkSort.FieldBy
	(*RequestCreateBookmark)(nil), // 2: stories.RequestCreateBookmark
	(*RequestEditBookmark)(nil),   // 3: stories.RequestEditBookmark
	(*RequestMyBookmarks)(nil),    // 4: stories.RequestMyBookmarks
	(*ResponseMyBookmarks)(nil),   // 5: stories.ResponseMyBookmarks
	(*Bookmark)(nil),              // 6: stories.Bookmark
	(*FilterMyBookmarks)(nil),     // 7: stories.FilterMyBookmarks
	(*BookmarkSort)(nil),          // 8: stories.BookmarkSort
	(*Paginator)(nil),             // 9: stories.Paginator
	(*StorySimple)(nil),           // 10: stories.StorySimple
	(*ChapterInfo)(nil),           // 11: stories.ChapterInfo
	(BoolOption)(0),               // 12: stories.BoolOption
}
var file_bookmarks_proto_depIdxs = []int32{
	0,  // 0: stories.RequestMyBookmarks.Load:type_name -> stories.LoadBookmarkType
	7,  // 1: stories.RequestMyBookmarks.Filter:type_name -> stories.FilterMyBookmarks
	8,  // 2: stories.RequestMyBookmarks.Sort:type_name -> stories.BookmarkSort
	9,  // 3: stories.RequestMyBookmarks.Paginator:type_name -> stories.Paginator
	6,  // 4: stories.ResponseMyBookmarks.Bookmarks:type_name -> stories.Bookmark
	10, // 5: stories.Bookmark.Story:type_name -> stories.StorySimple
	11, // 6: stories.Bookmark.Chapter:type_name -> stories.ChapterInfo
	12, // 7: stories.FilterMyBookmarks.StoryIsFinished:type_name -> stories.BoolOption
	1,  // 8: stories.BookmarkSort.By:type_name -> stories.BookmarkSort.FieldBy
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_bookmarks_proto_init() }
func file_bookmarks_proto_init() {
	if File_bookmarks_proto != nil {
		return
	}
	file_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bookmarks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestCreateBookmark); i {
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
		file_bookmarks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestEditBookmark); i {
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
		file_bookmarks_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMyBookmarks); i {
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
		file_bookmarks_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMyBookmarks); i {
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
		file_bookmarks_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bookmark); i {
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
		file_bookmarks_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterMyBookmarks); i {
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
		file_bookmarks_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookmarkSort); i {
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
			RawDescriptor: file_bookmarks_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bookmarks_proto_goTypes,
		DependencyIndexes: file_bookmarks_proto_depIdxs,
		EnumInfos:         file_bookmarks_proto_enumTypes,
		MessageInfos:      file_bookmarks_proto_msgTypes,
	}.Build()
	File_bookmarks_proto = out.File
	file_bookmarks_proto_rawDesc = nil
	file_bookmarks_proto_goTypes = nil
	file_bookmarks_proto_depIdxs = nil
}
