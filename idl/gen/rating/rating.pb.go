// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: idl/rating.proto

package rating

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

type RecommendReasonType int32

const (
	RecommendReasonType_RECOMMEND_REASON_TYPE_TAG   RecommendReasonType = 0
	RecommendReasonType_RECOMMEND_REASON_TYPE_MOVIE RecommendReasonType = 1
	RecommendReasonType_RECOMMEND_REASON_TYPE_LOG   RecommendReasonType = 2
	RecommendReasonType_RECOMMEND_REASON_TYPE_TOP_K RecommendReasonType = 3
)

// Enum value maps for RecommendReasonType.
var (
	RecommendReasonType_name = map[int32]string{
		0: "RECOMMEND_REASON_TYPE_TAG",
		1: "RECOMMEND_REASON_TYPE_MOVIE",
		2: "RECOMMEND_REASON_TYPE_LOG",
		3: "RECOMMEND_REASON_TYPE_TOP_K",
	}
	RecommendReasonType_value = map[string]int32{
		"RECOMMEND_REASON_TYPE_TAG":   0,
		"RECOMMEND_REASON_TYPE_MOVIE": 1,
		"RECOMMEND_REASON_TYPE_LOG":   2,
		"RECOMMEND_REASON_TYPE_TOP_K": 3,
	}
)

func (x RecommendReasonType) Enum() *RecommendReasonType {
	p := new(RecommendReasonType)
	*p = x
	return p
}

func (x RecommendReasonType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecommendReasonType) Descriptor() protoreflect.EnumDescriptor {
	return file_idl_rating_proto_enumTypes[0].Descriptor()
}

func (RecommendReasonType) Type() protoreflect.EnumType {
	return &file_idl_rating_proto_enumTypes[0]
}

func (x RecommendReasonType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecommendReasonType.Descriptor instead.
func (RecommendReasonType) EnumDescriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{0}
}

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrNo  int64  `protobuf:"varint,1,opt,name=err_no,json=errNo,proto3" json:"err_no,omitempty"`
	ErrMsg string `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetErrNo() int64 {
	if x != nil {
		return x.ErrNo
	}
	return 0
}

func (x *BaseResp) GetErrMsg() string {
	if x != nil {
		return x.ErrMsg
	}
	return ""
}

type Participant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Character string `protobuf:"bytes,1,opt,name=character,proto3" json:"character,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Participant) Reset() {
	*x = Participant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Participant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Participant) ProtoMessage() {}

func (x *Participant) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Participant.ProtoReflect.Descriptor instead.
func (*Participant) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{1}
}

func (x *Participant) GetCharacter() string {
	if x != nil {
		return x.Character
	}
	return ""
}

func (x *Participant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RecommendReason struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieReason *Movie              `protobuf:"bytes,1,opt,name=movie_reason,json=movieReason,proto3" json:"movie_reason,omitempty"`
	TagReason   string              `protobuf:"bytes,2,opt,name=tag_reason,json=tagReason,proto3" json:"tag_reason,omitempty"`
	ReasonType  RecommendReasonType `protobuf:"varint,3,opt,name=reason_type,json=reasonType,proto3,enum=rating.RecommendReasonType" json:"reason_type,omitempty"`
}

func (x *RecommendReason) Reset() {
	*x = RecommendReason{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendReason) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendReason) ProtoMessage() {}

func (x *RecommendReason) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendReason.ProtoReflect.Descriptor instead.
func (*RecommendReason) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{2}
}

func (x *RecommendReason) GetMovieReason() *Movie {
	if x != nil {
		return x.MovieReason
	}
	return nil
}

func (x *RecommendReason) GetTagReason() string {
	if x != nil {
		return x.TagReason
	}
	return ""
}

func (x *RecommendReason) GetReasonType() RecommendReasonType {
	if x != nil {
		return x.ReasonType
	}
	return RecommendReasonType_RECOMMEND_REASON_TYPE_TAG
}

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string           `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	PicUrl        string           `protobuf:"bytes,3,opt,name=pic_url,json=picUrl,proto3" json:"pic_url,omitempty"`
	Introduction  *string          `protobuf:"bytes,4,opt,name=introduction,proto3,oneof" json:"introduction,omitempty"`
	Participant   []*Participant   `protobuf:"bytes,5,rep,name=participant,proto3" json:"participant,omitempty"`
	ReleaseDate   *string          `protobuf:"bytes,6,opt,name=release_date,json=releaseDate,proto3,oneof" json:"release_date,omitempty"`
	Language      *string          `protobuf:"bytes,7,opt,name=language,proto3,oneof" json:"language,omitempty"`
	Reason        *RecommendReason `protobuf:"bytes,8,opt,name=reason,proto3,oneof" json:"reason,omitempty"`
	AverageRating *float64         `protobuf:"fixed64,9,opt,name=average_rating,json=averageRating,proto3,oneof" json:"average_rating,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[3]
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
	return file_idl_rating_proto_rawDescGZIP(), []int{3}
}

func (x *Movie) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetPicUrl() string {
	if x != nil {
		return x.PicUrl
	}
	return ""
}

func (x *Movie) GetIntroduction() string {
	if x != nil && x.Introduction != nil {
		return *x.Introduction
	}
	return ""
}

func (x *Movie) GetParticipant() []*Participant {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *Movie) GetReleaseDate() string {
	if x != nil && x.ReleaseDate != nil {
		return *x.ReleaseDate
	}
	return ""
}

func (x *Movie) GetLanguage() string {
	if x != nil && x.Language != nil {
		return *x.Language
	}
	return ""
}

func (x *Movie) GetReason() *RecommendReason {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *Movie) GetAverageRating() float64 {
	if x != nil && x.AverageRating != nil {
		return *x.AverageRating
	}
	return 0
}

type RateRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Movie  *Movie  `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	Rating float64 `protobuf:"fixed64,2,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RateRecord) Reset() {
	*x = RateRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRecord) ProtoMessage() {}

func (x *RateRecord) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRecord.ProtoReflect.Descriptor instead.
func (*RateRecord) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{4}
}

func (x *RateRecord) GetMovie() *Movie {
	if x != nil {
		return x.Movie
	}
	return nil
}

func (x *RateRecord) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId string  `protobuf:"bytes,1,opt,name=movie_id,json=movieId,proto3" json:"movie_id,omitempty"`
	UserId  string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rating  float64 `protobuf:"fixed64,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RateReq) Reset() {
	*x = RateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateReq) ProtoMessage() {}

func (x *RateReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateReq.ProtoReflect.Descriptor instead.
func (*RateReq) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{5}
}

func (x *RateReq) GetMovieId() string {
	if x != nil {
		return x.MovieId
	}
	return ""
}

func (x *RateReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RateReq) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type RateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *RateResp) Reset() {
	*x = RateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResp) ProtoMessage() {}

func (x *RateResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResp.ProtoReflect.Descriptor instead.
func (*RateResp) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{6}
}

func (x *RateResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type QueryRateRecordsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *QueryRateRecordsReq) Reset() {
	*x = QueryRateRecordsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRateRecordsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRateRecordsReq) ProtoMessage() {}

func (x *QueryRateRecordsReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRateRecordsReq.ProtoReflect.Descriptor instead.
func (*QueryRateRecordsReq) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{7}
}

func (x *QueryRateRecordsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryRateRecordsReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *QueryRateRecordsReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type QueryRateRecordsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp     `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Records  []*RateRecord `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	NRecords int64         `protobuf:"varint,3,opt,name=n_records,json=nRecords,proto3" json:"n_records,omitempty"`
}

func (x *QueryRateRecordsResp) Reset() {
	*x = QueryRateRecordsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryRateRecordsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryRateRecordsResp) ProtoMessage() {}

func (x *QueryRateRecordsResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryRateRecordsResp.ProtoReflect.Descriptor instead.
func (*QueryRateRecordsResp) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{8}
}

func (x *QueryRateRecordsResp) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *QueryRateRecordsResp) GetRecords() []*RateRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

func (x *QueryRateRecordsResp) GetNRecords() int64 {
	if x != nil {
		return x.NRecords
	}
	return 0
}

type QueryMovieRatingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MovieIdList []string `protobuf:"bytes,2,rep,name=movie_id_list,json=movieIdList,proto3" json:"movie_id_list,omitempty"`
}

func (x *QueryMovieRatingReq) Reset() {
	*x = QueryMovieRatingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMovieRatingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMovieRatingReq) ProtoMessage() {}

func (x *QueryMovieRatingReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMovieRatingReq.ProtoReflect.Descriptor instead.
func (*QueryMovieRatingReq) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{9}
}

func (x *QueryMovieRatingReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *QueryMovieRatingReq) GetMovieIdList() []string {
	if x != nil {
		return x.MovieIdList
	}
	return nil
}

type QueryMovieRatingResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MovieId2PersonalRating map[string]float64 `protobuf:"bytes,1,rep,name=movie_id2_personal_rating,json=movieId2PersonalRating,proto3" json:"movie_id2_personal_rating,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *QueryMovieRatingResp) Reset() {
	*x = QueryMovieRatingResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_rating_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMovieRatingResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMovieRatingResp) ProtoMessage() {}

func (x *QueryMovieRatingResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_rating_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMovieRatingResp.ProtoReflect.Descriptor instead.
func (*QueryMovieRatingResp) Descriptor() ([]byte, []int) {
	return file_idl_rating_proto_rawDescGZIP(), []int{10}
}

func (x *QueryMovieRatingResp) GetMovieId2PersonalRating() map[string]float64 {
	if x != nil {
		return x.MovieId2PersonalRating
	}
	return nil
}

var File_idl_rating_proto protoreflect.FileDescriptor

var file_idl_rating_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x64, 0x6c, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x3a, 0x0a, 0x08, 0x42, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x5f, 0x6e, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x65, 0x72, 0x72, 0x4e, 0x6f, 0x12, 0x17, 0x0a,
	0x07, 0x65, 0x72, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x65, 0x72, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x3f, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x0c, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x52, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x61, 0x67, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x67, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x0b,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x9e, 0x03, 0x0a, 0x05, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x69,
	0x63, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x69, 0x63,
	0x55, 0x72, 0x6c, 0x12, 0x27, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x0b,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0c, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x48, 0x03, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x48, 0x04, 0x52, 0x0d, 0x61, 0x76,
	0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x0f, 0x0a, 0x0d, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x49, 0x0a, 0x0a, 0x52,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x05, 0x6d, 0x6f, 0x76,
	0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x55, 0x0a, 0x07, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x39, 0x0a,
	0x08, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5a, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x14, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x07,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x52, 0x0a, 0x13, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xd6, 0x01, 0x0a, 0x14,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x73, 0x0a, 0x19, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x5f, 0x69, 0x64,
	0x32, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x32, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x16, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x32, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x49, 0x0a, 0x1b, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x49, 0x64, 0x32, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x2a, 0x95, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19,
	0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x52,
	0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x56, 0x49, 0x45, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19,
	0x52, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x52,
	0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x44, 0x5f, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x4f, 0x50, 0x5f, 0x4b, 0x10, 0x03, 0x32, 0xe8, 0x01, 0x0a,
	0x0d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x0f, 0x2e, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x54, 0x0a, 0x15, 0x42, 0x61, 0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x72, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_idl_rating_proto_rawDescOnce sync.Once
	file_idl_rating_proto_rawDescData = file_idl_rating_proto_rawDesc
)

func file_idl_rating_proto_rawDescGZIP() []byte {
	file_idl_rating_proto_rawDescOnce.Do(func() {
		file_idl_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_rating_proto_rawDescData)
	})
	return file_idl_rating_proto_rawDescData
}

var file_idl_rating_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_idl_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_idl_rating_proto_goTypes = []interface{}{
	(RecommendReasonType)(0),     // 0: rating.RecommendReasonType
	(*BaseResp)(nil),             // 1: rating.BaseResp
	(*Participant)(nil),          // 2: rating.Participant
	(*RecommendReason)(nil),      // 3: rating.RecommendReason
	(*Movie)(nil),                // 4: rating.Movie
	(*RateRecord)(nil),           // 5: rating.RateRecord
	(*RateReq)(nil),              // 6: rating.RateReq
	(*RateResp)(nil),             // 7: rating.RateResp
	(*QueryRateRecordsReq)(nil),  // 8: rating.QueryRateRecordsReq
	(*QueryRateRecordsResp)(nil), // 9: rating.QueryRateRecordsResp
	(*QueryMovieRatingReq)(nil),  // 10: rating.QueryMovieRatingReq
	(*QueryMovieRatingResp)(nil), // 11: rating.QueryMovieRatingResp
	nil,                          // 12: rating.QueryMovieRatingResp.MovieId2PersonalRatingEntry
}
var file_idl_rating_proto_depIdxs = []int32{
	4,  // 0: rating.RecommendReason.movie_reason:type_name -> rating.Movie
	0,  // 1: rating.RecommendReason.reason_type:type_name -> rating.RecommendReasonType
	2,  // 2: rating.Movie.participant:type_name -> rating.Participant
	3,  // 3: rating.Movie.reason:type_name -> rating.RecommendReason
	4,  // 4: rating.RateRecord.movie:type_name -> rating.Movie
	1,  // 5: rating.RateResp.base_resp:type_name -> rating.BaseResp
	1,  // 6: rating.QueryRateRecordsResp.base_resp:type_name -> rating.BaseResp
	5,  // 7: rating.QueryRateRecordsResp.records:type_name -> rating.RateRecord
	12, // 8: rating.QueryMovieRatingResp.movie_id2_personal_rating:type_name -> rating.QueryMovieRatingResp.MovieId2PersonalRatingEntry
	6,  // 9: rating.RatingService.RateMovie:input_type -> rating.RateReq
	8,  // 10: rating.RatingService.QueryRateRecords:input_type -> rating.QueryRateRecordsReq
	10, // 11: rating.RatingService.BatchQueryMovieRating:input_type -> rating.QueryMovieRatingReq
	7,  // 12: rating.RatingService.RateMovie:output_type -> rating.RateResp
	9,  // 13: rating.RatingService.QueryRateRecords:output_type -> rating.QueryRateRecordsResp
	11, // 14: rating.RatingService.BatchQueryMovieRating:output_type -> rating.QueryMovieRatingResp
	12, // [12:15] is the sub-list for method output_type
	9,  // [9:12] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_idl_rating_proto_init() }
func file_idl_rating_proto_init() {
	if File_idl_rating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_rating_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_idl_rating_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Participant); i {
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
		file_idl_rating_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendReason); i {
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
		file_idl_rating_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_idl_rating_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRecord); i {
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
		file_idl_rating_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateReq); i {
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
		file_idl_rating_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResp); i {
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
		file_idl_rating_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRateRecordsReq); i {
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
		file_idl_rating_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryRateRecordsResp); i {
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
		file_idl_rating_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMovieRatingReq); i {
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
		file_idl_rating_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMovieRatingResp); i {
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
	file_idl_rating_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_rating_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_rating_proto_goTypes,
		DependencyIndexes: file_idl_rating_proto_depIdxs,
		EnumInfos:         file_idl_rating_proto_enumTypes,
		MessageInfos:      file_idl_rating_proto_msgTypes,
	}.Build()
	File_idl_rating_proto = out.File
	file_idl_rating_proto_rawDesc = nil
	file_idl_rating_proto_goTypes = nil
	file_idl_rating_proto_depIdxs = nil
}
