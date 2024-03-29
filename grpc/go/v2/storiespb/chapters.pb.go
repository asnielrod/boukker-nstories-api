// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: chapters.proto

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

// RequestWriteChapter ...
type RequestWriteChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string     `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Content    string     `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	TokenCover string     `protobuf:"bytes,3,opt,name=TokenCover,proto3" json:"TokenCover,omitempty"`
	Status     StatusType `protobuf:"varint,4,opt,name=Status,proto3,enum=stories.StatusType" json:"Status,omitempty"`
	StoryID    string     `protobuf:"bytes,5,opt,name=StoryID,proto3" json:"StoryID,omitempty"`
}

func (x *RequestWriteChapter) Reset() {
	*x = RequestWriteChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestWriteChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestWriteChapter) ProtoMessage() {}

func (x *RequestWriteChapter) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestWriteChapter.ProtoReflect.Descriptor instead.
func (*RequestWriteChapter) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{0}
}

func (x *RequestWriteChapter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestWriteChapter) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *RequestWriteChapter) GetTokenCover() string {
	if x != nil {
		return x.TokenCover
	}
	return ""
}

func (x *RequestWriteChapter) GetStatus() StatusType {
	if x != nil {
		return x.Status
	}
	return StatusType_ST_IGNORE
}

func (x *RequestWriteChapter) GetStoryID() string {
	if x != nil {
		return x.StoryID
	}
	return ""
}

// RequestEditChapter ...
type RequestEditChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *RequestEditChapter) Reset() {
	*x = RequestEditChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestEditChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestEditChapter) ProtoMessage() {}

func (x *RequestEditChapter) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestEditChapter.ProtoReflect.Descriptor instead.
func (*RequestEditChapter) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{1}
}

func (x *RequestEditChapter) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *RequestEditChapter) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RequestEditChapter) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// RequestReorderChapters ...
type RequestReorderChapters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoryID string   `protobuf:"bytes,1,opt,name=StoryID,proto3" json:"StoryID,omitempty"`
	Items   []*Order `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (x *RequestReorderChapters) Reset() {
	*x = RequestReorderChapters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestReorderChapters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestReorderChapters) ProtoMessage() {}

func (x *RequestReorderChapters) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestReorderChapters.ProtoReflect.Descriptor instead.
func (*RequestReorderChapters) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{2}
}

func (x *RequestReorderChapters) GetStoryID() string {
	if x != nil {
		return x.StoryID
	}
	return ""
}

func (x *RequestReorderChapters) GetItems() []*Order {
	if x != nil {
		return x.Items
	}
	return nil
}

// RequestRemoveChapter ...
type RequestRemoveChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RequestRemoveChapter) Reset() {
	*x = RequestRemoveChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRemoveChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRemoveChapter) ProtoMessage() {}

func (x *RequestRemoveChapter) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRemoveChapter.ProtoReflect.Descriptor instead.
func (*RequestRemoveChapter) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{3}
}

func (x *RequestRemoveChapter) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// RequestRemoveChapterCover ...
type RequestRemoveChapterCover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *RequestRemoveChapterCover) Reset() {
	*x = RequestRemoveChapterCover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRemoveChapterCover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRemoveChapterCover) ProtoMessage() {}

func (x *RequestRemoveChapterCover) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRemoveChapterCover.ProtoReflect.Descriptor instead.
func (*RequestRemoveChapterCover) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{4}
}

func (x *RequestRemoveChapterCover) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// ResponseRetrieveChapter ...
type ResponseRetrieveChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapter *Chapter `protobuf:"bytes,1,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
}

func (x *ResponseRetrieveChapter) Reset() {
	*x = ResponseRetrieveChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRetrieveChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRetrieveChapter) ProtoMessage() {}

func (x *ResponseRetrieveChapter) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRetrieveChapter.ProtoReflect.Descriptor instead.
func (*ResponseRetrieveChapter) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseRetrieveChapter) GetChapter() *Chapter {
	if x != nil {
		return x.Chapter
	}
	return nil
}

// ResponseReadChapter ...
type ResponseReadChapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Story         *Story   `protobuf:"bytes,1,opt,name=Story,proto3" json:"Story,omitempty"`
	Chapter       *Chapter `protobuf:"bytes,2,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
	WriterSummary string   `protobuf:"bytes,3,opt,name=WriterSummary,proto3" json:"WriterSummary,omitempty"`
	// WriterStories otras historias publicadas por el escritor
	WriterStories      []*StoryInfo  `protobuf:"bytes,4,rep,name=WriterStories,proto3" json:"WriterStories,omitempty"`
	RecommendedStories []*StoryInfo  `protobuf:"bytes,5,rep,name=RecommendedStories,proto3" json:"RecommendedStories,omitempty"`
	Bookmark           *BookmarkInfo `protobuf:"bytes,8,opt,name=Bookmark,proto3" json:"Bookmark,omitempty"`
	ScrollDepth        float32       `protobuf:"fixed32,9,opt,name=ScrollDepth,proto3" json:"ScrollDepth,omitempty"`
}

func (x *ResponseReadChapter) Reset() {
	*x = ResponseReadChapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chapters_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseReadChapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseReadChapter) ProtoMessage() {}

func (x *ResponseReadChapter) ProtoReflect() protoreflect.Message {
	mi := &file_chapters_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseReadChapter.ProtoReflect.Descriptor instead.
func (*ResponseReadChapter) Descriptor() ([]byte, []int) {
	return file_chapters_proto_rawDescGZIP(), []int{6}
}

func (x *ResponseReadChapter) GetStory() *Story {
	if x != nil {
		return x.Story
	}
	return nil
}

func (x *ResponseReadChapter) GetChapter() *Chapter {
	if x != nil {
		return x.Chapter
	}
	return nil
}

func (x *ResponseReadChapter) GetWriterSummary() string {
	if x != nil {
		return x.WriterSummary
	}
	return ""
}

func (x *ResponseReadChapter) GetWriterStories() []*StoryInfo {
	if x != nil {
		return x.WriterStories
	}
	return nil
}

func (x *ResponseReadChapter) GetRecommendedStories() []*StoryInfo {
	if x != nil {
		return x.RecommendedStories
	}
	return nil
}

func (x *ResponseReadChapter) GetBookmark() *BookmarkInfo {
	if x != nil {
		return x.Bookmark
	}
	return nil
}

func (x *ResponseReadChapter) GetScrollDepth() float32 {
	if x != nil {
		return x.ScrollDepth
	}
	return 0
}

var File_chapters_proto protoreflect.FileDescriptor

var file_chapters_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x12, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22, 0x54, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x45, 0x64, 0x69, 0x74, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x58, 0x0a,
	0x16, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49,
	0x44, 0x12, 0x24, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x2b, 0x0a, 0x19, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x17,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x07, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x22, 0xe0, 0x02, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x53,
	0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a,
	0x12, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x31, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x42, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x63, 0x72, 0x6f, 0x6c, 0x6c, 0x44, 0x65,
	0x70, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x53, 0x63, 0x72, 0x6f, 0x6c,
	0x6c, 0x44, 0x65, 0x70, 0x74, 0x68, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x72, 0x6e, 0x65, 0x73, 0x74, 0x6f, 0x47, 0x46, 0x2f, 0x62,
	0x6f, 0x75, 0x6b, 0x6b, 0x65, 0x72, 0x2d, 0x6e, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chapters_proto_rawDescOnce sync.Once
	file_chapters_proto_rawDescData = file_chapters_proto_rawDesc
)

func file_chapters_proto_rawDescGZIP() []byte {
	file_chapters_proto_rawDescOnce.Do(func() {
		file_chapters_proto_rawDescData = protoimpl.X.CompressGZIP(file_chapters_proto_rawDescData)
	})
	return file_chapters_proto_rawDescData
}

var file_chapters_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_chapters_proto_goTypes = []interface{}{
	(*RequestWriteChapter)(nil),       // 0: stories.RequestWriteChapter
	(*RequestEditChapter)(nil),        // 1: stories.RequestEditChapter
	(*RequestReorderChapters)(nil),    // 2: stories.RequestReorderChapters
	(*RequestRemoveChapter)(nil),      // 3: stories.RequestRemoveChapter
	(*RequestRemoveChapterCover)(nil), // 4: stories.RequestRemoveChapterCover
	(*ResponseRetrieveChapter)(nil),   // 5: stories.ResponseRetrieveChapter
	(*ResponseReadChapter)(nil),       // 6: stories.ResponseReadChapter
	(StatusType)(0),                   // 7: stories.StatusType
	(*Order)(nil),                     // 8: stories.Order
	(*Chapter)(nil),                   // 9: stories.Chapter
	(*Story)(nil),                     // 10: stories.Story
	(*StoryInfo)(nil),                 // 11: stories.StoryInfo
	(*BookmarkInfo)(nil),              // 12: stories.BookmarkInfo
}
var file_chapters_proto_depIdxs = []int32{
	7,  // 0: stories.RequestWriteChapter.Status:type_name -> stories.StatusType
	8,  // 1: stories.RequestReorderChapters.Items:type_name -> stories.Order
	9,  // 2: stories.ResponseRetrieveChapter.Chapter:type_name -> stories.Chapter
	10, // 3: stories.ResponseReadChapter.Story:type_name -> stories.Story
	9,  // 4: stories.ResponseReadChapter.Chapter:type_name -> stories.Chapter
	11, // 5: stories.ResponseReadChapter.WriterStories:type_name -> stories.StoryInfo
	11, // 6: stories.ResponseReadChapter.RecommendedStories:type_name -> stories.StoryInfo
	12, // 7: stories.ResponseReadChapter.Bookmark:type_name -> stories.BookmarkInfo
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_chapters_proto_init() }
func file_chapters_proto_init() {
	if File_chapters_proto != nil {
		return
	}
	file_commons_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chapters_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestWriteChapter); i {
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
		file_chapters_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestEditChapter); i {
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
		file_chapters_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestReorderChapters); i {
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
		file_chapters_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRemoveChapter); i {
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
		file_chapters_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRemoveChapterCover); i {
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
		file_chapters_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRetrieveChapter); i {
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
		file_chapters_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseReadChapter); i {
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
			RawDescriptor: file_chapters_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chapters_proto_goTypes,
		DependencyIndexes: file_chapters_proto_depIdxs,
		MessageInfos:      file_chapters_proto_msgTypes,
	}.Build()
	File_chapters_proto = out.File
	file_chapters_proto_rawDesc = nil
	file_chapters_proto_goTypes = nil
	file_chapters_proto_depIdxs = nil
}
