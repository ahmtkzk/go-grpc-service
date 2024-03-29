// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: api/MovieService.proto

package api

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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OriginalTitle    string `protobuf:"bytes,2,opt,name=original_title,json=originalTitle,proto3" json:"original_title,omitempty"`
	Overview         string `protobuf:"bytes,3,opt,name=overview,proto3" json:"overview,omitempty"`
	Title            string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	OriginalLanguage string `protobuf:"bytes,5,opt,name=original_language,json=originalLanguage,proto3" json:"original_language,omitempty"`
	Adult            bool   `protobuf:"varint,6,opt,name=adult,proto3" json:"adult,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Movie) GetOriginalTitle() string {
	if x != nil {
		return x.OriginalTitle
	}
	return ""
}

func (x *Movie) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetOriginalLanguage() string {
	if x != nil {
		return x.OriginalLanguage
	}
	return ""
}

func (x *Movie) GetAdult() bool {
	if x != nil {
		return x.Adult
	}
	return false
}

type Keyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Keyword) Reset() {
	*x = Keyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keyword) ProtoMessage() {}

func (x *Keyword) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keyword.ProtoReflect.Descriptor instead.
func (*Keyword) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{1}
}

func (x *Keyword) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Keyword) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Cast struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adult              bool   `protobuf:"varint,1,opt,name=adult,proto3" json:"adult,omitempty"`
	Gender             int32  `protobuf:"varint,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Id                 int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	KnownForDepartment string `protobuf:"bytes,4,opt,name=known_for_department,json=knownForDepartment,proto3" json:"known_for_department,omitempty"`
	Name               string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	OriginalName       string `protobuf:"bytes,6,opt,name=original_name,json=originalName,proto3" json:"original_name,omitempty"`
	ProfilePath        string `protobuf:"bytes,7,opt,name=profile_path,json=profilePath,proto3" json:"profile_path,omitempty"`
	CastId             int32  `protobuf:"varint,8,opt,name=cast_id,json=castId,proto3" json:"cast_id,omitempty"`
	Character          string `protobuf:"bytes,9,opt,name=character,proto3" json:"character,omitempty"`
	CreditId           string `protobuf:"bytes,10,opt,name=credit_id,json=creditId,proto3" json:"credit_id,omitempty"`
	Order              int32  `protobuf:"varint,11,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *Cast) Reset() {
	*x = Cast{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cast) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cast) ProtoMessage() {}

func (x *Cast) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cast.ProtoReflect.Descriptor instead.
func (*Cast) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{2}
}

func (x *Cast) GetAdult() bool {
	if x != nil {
		return x.Adult
	}
	return false
}

func (x *Cast) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Cast) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cast) GetKnownForDepartment() string {
	if x != nil {
		return x.KnownForDepartment
	}
	return ""
}

func (x *Cast) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cast) GetOriginalName() string {
	if x != nil {
		return x.OriginalName
	}
	return ""
}

func (x *Cast) GetProfilePath() string {
	if x != nil {
		return x.ProfilePath
	}
	return ""
}

func (x *Cast) GetCastId() int32 {
	if x != nil {
		return x.CastId
	}
	return 0
}

func (x *Cast) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

func (x *Cast) GetCreditId() string {
	if x != nil {
		return x.CreditId
	}
	return ""
}

func (x *Cast) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type GetMovieListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMovieListRequest) Reset() {
	*x = GetMovieListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieListRequest) ProtoMessage() {}

func (x *GetMovieListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieListRequest.ProtoReflect.Descriptor instead.
func (*GetMovieListRequest) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{3}
}

type GetMovieListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieList []*Movie `protobuf:"bytes,1,rep,name=movieList,proto3" json:"movieList,omitempty"`
}

func (x *GetMovieListResponse) Reset() {
	*x = GetMovieListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieListResponse) ProtoMessage() {}

func (x *GetMovieListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieListResponse.ProtoReflect.Descriptor instead.
func (*GetMovieListResponse) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{4}
}

func (x *GetMovieListResponse) GetMovieList() []*Movie {
	if x != nil {
		return x.MovieList
	}
	return nil
}

type GetMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMovieRequest) Reset() {
	*x = GetMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieRequest) ProtoMessage() {}

func (x *GetMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieRequest.ProtoReflect.Descriptor instead.
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{5}
}

func (x *GetMovieRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie *Movie `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
}

func (x *GetMovieResponse) Reset() {
	*x = GetMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieResponse) ProtoMessage() {}

func (x *GetMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieResponse.ProtoReflect.Descriptor instead.
func (*GetMovieResponse) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{6}
}

func (x *GetMovieResponse) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

type GetMovieKeywordListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMovieKeywordListRequest) Reset() {
	*x = GetMovieKeywordListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieKeywordListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieKeywordListRequest) ProtoMessage() {}

func (x *GetMovieKeywordListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieKeywordListRequest.ProtoReflect.Descriptor instead.
func (*GetMovieKeywordListRequest) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{7}
}

func (x *GetMovieKeywordListRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMovieKeywordListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeywordList []*Keyword `protobuf:"bytes,1,rep,name=keywordList,proto3" json:"keywordList,omitempty"`
}

func (x *GetMovieKeywordListResponse) Reset() {
	*x = GetMovieKeywordListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieKeywordListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieKeywordListResponse) ProtoMessage() {}

func (x *GetMovieKeywordListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieKeywordListResponse.ProtoReflect.Descriptor instead.
func (*GetMovieKeywordListResponse) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{8}
}

func (x *GetMovieKeywordListResponse) GetKeywordList() []*Keyword {
	if x != nil {
		return x.KeywordList
	}
	return nil
}

type GetMovieCreditsListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *GetMovieCreditsListRequest) Reset() {
	*x = GetMovieCreditsListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieCreditsListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieCreditsListRequest) ProtoMessage() {}

func (x *GetMovieCreditsListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieCreditsListRequest.ProtoReflect.Descriptor instead.
func (*GetMovieCreditsListRequest) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{9}
}

func (x *GetMovieCreditsListRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetMovieCreditsListRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type GetMovieCreditsListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CastList []*Cast `protobuf:"bytes,1,rep,name=castList,proto3" json:"castList,omitempty"`
}

func (x *GetMovieCreditsListResponse) Reset() {
	*x = GetMovieCreditsListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_MovieService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieCreditsListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieCreditsListResponse) ProtoMessage() {}

func (x *GetMovieCreditsListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_MovieService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieCreditsListResponse.ProtoReflect.Descriptor instead.
func (*GetMovieCreditsListResponse) Descriptor() ([]byte, []int) {
	return file_api_MovieService_proto_rawDescGZIP(), []int{10}
}

func (x *GetMovieCreditsListResponse) GetCastList() []*Cast {
	if x != nil {
		return x.CastList
	}
	return nil
}

var File_api_MovieService_proto protoreflect.FileDescriptor

var file_api_MovieService_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xb3,
	0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x64, 0x75, 0x6c, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61,
	0x64, 0x75, 0x6c, 0x74, 0x22, 0x2d, 0x0a, 0x07, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xbc, 0x02, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x64, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64, 0x75,
	0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x46,
	0x6f, 0x72, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x61, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x52, 0x09, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x22, 0x2c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x4e, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x48, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x45,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x08, 0x63, 0x61, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x73, 0x74, 0x52, 0x08, 0x63, 0x61, 0x73,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xc8, 0x02, 0x0a, 0x0c, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x4b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x43, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_MovieService_proto_rawDescOnce sync.Once
	file_api_MovieService_proto_rawDescData = file_api_MovieService_proto_rawDesc
)

func file_api_MovieService_proto_rawDescGZIP() []byte {
	file_api_MovieService_proto_rawDescOnce.Do(func() {
		file_api_MovieService_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_MovieService_proto_rawDescData)
	})
	return file_api_MovieService_proto_rawDescData
}

var file_api_MovieService_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_MovieService_proto_goTypes = []interface{}{
	(*Movie)(nil),                       // 0: main.Movie
	(*Keyword)(nil),                     // 1: main.Keyword
	(*Cast)(nil),                        // 2: main.Cast
	(*GetMovieListRequest)(nil),         // 3: main.GetMovieListRequest
	(*GetMovieListResponse)(nil),        // 4: main.GetMovieListResponse
	(*GetMovieRequest)(nil),             // 5: main.GetMovieRequest
	(*GetMovieResponse)(nil),            // 6: main.GetMovieResponse
	(*GetMovieKeywordListRequest)(nil),  // 7: main.GetMovieKeywordListRequest
	(*GetMovieKeywordListResponse)(nil), // 8: main.GetMovieKeywordListResponse
	(*GetMovieCreditsListRequest)(nil),  // 9: main.GetMovieCreditsListRequest
	(*GetMovieCreditsListResponse)(nil), // 10: main.GetMovieCreditsListResponse
}
var file_api_MovieService_proto_depIdxs = []int32{
	0,  // 0: main.GetMovieListResponse.movieList:type_name -> main.Movie
	0,  // 1: main.GetMovieResponse.movie:type_name -> main.Movie
	1,  // 2: main.GetMovieKeywordListResponse.keywordList:type_name -> main.Keyword
	2,  // 3: main.GetMovieCreditsListResponse.castList:type_name -> main.Cast
	3,  // 4: main.MovieService.GetMovieList:input_type -> main.GetMovieListRequest
	5,  // 5: main.MovieService.GetMovie:input_type -> main.GetMovieRequest
	7,  // 6: main.MovieService.GetMovieKeywordList:input_type -> main.GetMovieKeywordListRequest
	9,  // 7: main.MovieService.GetMovieCreditsList:input_type -> main.GetMovieCreditsListRequest
	4,  // 8: main.MovieService.GetMovieList:output_type -> main.GetMovieListResponse
	6,  // 9: main.MovieService.GetMovie:output_type -> main.GetMovieResponse
	8,  // 10: main.MovieService.GetMovieKeywordList:output_type -> main.GetMovieKeywordListResponse
	10, // 11: main.MovieService.GetMovieCreditsList:output_type -> main.GetMovieCreditsListResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_MovieService_proto_init() }
func file_api_MovieService_proto_init() {
	if File_api_MovieService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_MovieService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_api_MovieService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Keyword); i {
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
		file_api_MovieService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cast); i {
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
		file_api_MovieService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieListRequest); i {
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
		file_api_MovieService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieListResponse); i {
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
		file_api_MovieService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieRequest); i {
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
		file_api_MovieService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieResponse); i {
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
		file_api_MovieService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieKeywordListRequest); i {
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
		file_api_MovieService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieKeywordListResponse); i {
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
		file_api_MovieService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieCreditsListRequest); i {
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
		file_api_MovieService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieCreditsListResponse); i {
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
			RawDescriptor: file_api_MovieService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_MovieService_proto_goTypes,
		DependencyIndexes: file_api_MovieService_proto_depIdxs,
		MessageInfos:      file_api_MovieService_proto_msgTypes,
	}.Build()
	File_api_MovieService_proto = out.File
	file_api_MovieService_proto_rawDesc = nil
	file_api_MovieService_proto_goTypes = nil
	file_api_MovieService_proto_depIdxs = nil
}
