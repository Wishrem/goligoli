// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: video.proto

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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Liked       int64  `protobuf:"varint,4,opt,name=liked,proto3" json:"liked,omitempty"`
	Shared      int64  `protobuf:"varint,5,opt,name=shared,proto3" json:"shared,omitempty"`
	VideoId     int64  `protobuf:"varint,6,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	VideoUrl    string `protobuf:"bytes,7,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
	Passed      bool   `protobuf:"varint,8,opt,name=passed,proto3" json:"passed,omitempty"`
	Reason      string `protobuf:"bytes,9,opt,name=reason,proto3" json:"reason,omitempty"`
	Year        int64  `protobuf:"varint,10,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Video) GetLiked() int64 {
	if x != nil {
		return x.Liked
	}
	return 0
}

func (x *Video) GetShared() int64 {
	if x != nil {
		return x.Shared
	}
	return 0
}

func (x *Video) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *Video) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

func (x *Video) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *Video) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Video) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

type UploadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Video       []byte `protobuf:"bytes,2,opt,name=video,proto3" json:"video,omitempty"`
	Title       string `protobuf:"bytes,3,opt,name=title,proto3" json:"title" form:"title" binding:"required"`             // @gotags: json:"title" form:"title" binding:"required"
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description" form:"description" binding:"required"` // @gotags: json:"description" form:"description" binding:"required"
	Year        int64  `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *UploadReq) Reset() {
	*x = UploadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadReq) ProtoMessage() {}

func (x *UploadReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadReq.ProtoReflect.Descriptor instead.
func (*UploadReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *UploadReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UploadReq) GetVideo() []byte {
	if x != nil {
		return x.Video
	}
	return nil
}

func (x *UploadReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UploadReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UploadReq) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

type UploadResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video *Video `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *UploadResp) Reset() {
	*x = UploadResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadResp) ProtoMessage() {}

func (x *UploadResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadResp.ProtoReflect.Descriptor instead.
func (*UploadResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *UploadResp) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type LikeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id" form:"video_id" binding:"required"` // @gotags: json:"video_id" form:"video_id" binding:"required"
}

func (x *LikeReq) Reset() {
	*x = LikeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeReq) ProtoMessage() {}

func (x *LikeReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeReq.ProtoReflect.Descriptor instead.
func (*LikeReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *LikeReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LikeReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type LikeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LikeResp) Reset() {
	*x = LikeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeResp) ProtoMessage() {}

func (x *LikeResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeResp.ProtoReflect.Descriptor instead.
func (*LikeResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

type ShareReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	VideoId int64 `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id" form:"video_id" binding:"required"` // @gotags: json:"video_id" form:"video_id" binding:"required"
}

func (x *ShareReq) Reset() {
	*x = ShareReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareReq) ProtoMessage() {}

func (x *ShareReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareReq.ProtoReflect.Descriptor instead.
func (*ShareReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{5}
}

func (x *ShareReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ShareReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type ShareResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShareUrl string `protobuf:"bytes,1,opt,name=share_url,json=shareUrl,proto3" json:"share_url,omitempty"`
}

func (x *ShareResp) Reset() {
	*x = ShareResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareResp) ProtoMessage() {}

func (x *ShareResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareResp.ProtoReflect.Descriptor instead.
func (*ShareResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{6}
}

func (x *ShareResp) GetShareUrl() string {
	if x != nil {
		return x.ShareUrl
	}
	return ""
}

type JudgeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId int64  `protobuf:"varint,1,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	VideoId int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id" form:"video_id" binding:"required"` // @gotags: json:"video_id" form:"video_id" binding:"required"
	Passed  bool   `protobuf:"varint,3,opt,name=passed,proto3" json:"passed" form:"passed" binding:"required"`                  // @gotags: json:"passed" form:"passed" binding:"required"
	Reason  string `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason" form:"reason"`                   // @gotags: json:"reason" form:"reason"
}

func (x *JudgeReq) Reset() {
	*x = JudgeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeReq) ProtoMessage() {}

func (x *JudgeReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeReq.ProtoReflect.Descriptor instead.
func (*JudgeReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{7}
}

func (x *JudgeReq) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *JudgeReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *JudgeReq) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *JudgeReq) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type JudgeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Video *Video `protobuf:"bytes,1,opt,name=video,proto3" json:"video,omitempty"`
}

func (x *JudgeResp) Reset() {
	*x = JudgeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeResp) ProtoMessage() {}

func (x *JudgeResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeResp.ProtoReflect.Descriptor instead.
func (*JudgeResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{8}
}

func (x *JudgeResp) GetVideo() *Video {
	if x != nil {
		return x.Video
	}
	return nil
}

type ViewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *ViewReq) Reset() {
	*x = ViewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewReq) ProtoMessage() {}

func (x *ViewReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewReq.ProtoReflect.Descriptor instead.
func (*ViewReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{9}
}

func (x *ViewReq) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type ViewResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ViewResp) Reset() {
	*x = ViewResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewResp) ProtoMessage() {}

func (x *ViewResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewResp.ProtoReflect.Descriptor instead.
func (*ViewResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{10}
}

type GetVideosReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year   int64  `protobuf:"varint,1,opt,name=year,proto3" json:"year" form:"year"`     // @gotags: json:"year" form:"year"
	Liked  int64  `protobuf:"varint,2,opt,name=liked,proto3" json:"liked" form:"liked"`   // @gotags: json:"liked" form:"liked"
	Shared int64  `protobuf:"varint,3,opt,name=shared,proto3" json:"shared" form:"shared"` // @gotags: json:"shared" form:"shared"
	Title  string `protobuf:"bytes,4,opt,name=title,proto3" json:"title" form:"title"`    // @gotags: json:"title" form:"title"
}

func (x *GetVideosReq) Reset() {
	*x = GetVideosReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideosReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideosReq) ProtoMessage() {}

func (x *GetVideosReq) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideosReq.ProtoReflect.Descriptor instead.
func (*GetVideosReq) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{11}
}

func (x *GetVideosReq) GetYear() int64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *GetVideosReq) GetLiked() int64 {
	if x != nil {
		return x.Liked
	}
	return 0
}

func (x *GetVideosReq) GetShared() int64 {
	if x != nil {
		return x.Shared
	}
	return 0
}

func (x *GetVideosReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type GetVideosResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Videos []*Video `protobuf:"bytes,1,rep,name=videos,proto3" json:"videos,omitempty"`
}

func (x *GetVideosResp) Reset() {
	*x = GetVideosResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideosResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideosResp) ProtoMessage() {}

func (x *GetVideosResp) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideosResp.ProtoReflect.Descriptor instead.
func (*GetVideosResp) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{12}
}

func (x *GetVideosResp) GetVideos() []*Video {
	if x != nil {
		return x.Videos
	}
	return nil
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x82, 0x02, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b,
	0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x86, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22,
	0x2d, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x3d,
	0x0a, 0x07, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x0a, 0x0a,
	0x08, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x3e, 0x0a, 0x08, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x09, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61, 0x72, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x55, 0x72, 0x6c, 0x22, 0x70, 0x0a, 0x08, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x2c, 0x0a, 0x09, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x05, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x22, 0x25, 0x0a, 0x07, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x56, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x22, 0x66, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x32,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x21, 0x0a, 0x06, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x06, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04,
	0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x24, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x65, 0x12, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x05, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x12, 0x0c,
	0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70,
	0x62, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x04, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x12, 0x30,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x12, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_video_proto_goTypes = []interface{}{
	(*Video)(nil),         // 0: pb.Video
	(*UploadReq)(nil),     // 1: pb.UploadReq
	(*UploadResp)(nil),    // 2: pb.UploadResp
	(*LikeReq)(nil),       // 3: pb.LikeReq
	(*LikeResp)(nil),      // 4: pb.LikeResp
	(*ShareReq)(nil),      // 5: pb.ShareReq
	(*ShareResp)(nil),     // 6: pb.ShareResp
	(*JudgeReq)(nil),      // 7: pb.JudgeReq
	(*JudgeResp)(nil),     // 8: pb.JudgeResp
	(*ViewReq)(nil),       // 9: pb.ViewReq
	(*ViewResp)(nil),      // 10: pb.ViewResp
	(*GetVideosReq)(nil),  // 11: pb.GetVideosReq
	(*GetVideosResp)(nil), // 12: pb.GetVideosResp
}
var file_video_proto_depIdxs = []int32{
	0,  // 0: pb.UploadResp.video:type_name -> pb.Video
	0,  // 1: pb.JudgeResp.video:type_name -> pb.Video
	0,  // 2: pb.GetVideosResp.videos:type_name -> pb.Video
	1,  // 3: pb.VideoService.Upload:input_type -> pb.UploadReq
	3,  // 4: pb.VideoService.Like:input_type -> pb.LikeReq
	5,  // 5: pb.VideoService.Share:input_type -> pb.ShareReq
	7,  // 6: pb.VideoService.Judge:input_type -> pb.JudgeReq
	9,  // 7: pb.VideoService.View:input_type -> pb.ViewReq
	11, // 8: pb.VideoService.GetVideos:input_type -> pb.GetVideosReq
	2,  // 9: pb.VideoService.Upload:output_type -> pb.UploadResp
	4,  // 10: pb.VideoService.Like:output_type -> pb.LikeResp
	6,  // 11: pb.VideoService.Share:output_type -> pb.ShareResp
	8,  // 12: pb.VideoService.Judge:output_type -> pb.JudgeResp
	10, // 13: pb.VideoService.View:output_type -> pb.ViewResp
	12, // 14: pb.VideoService.GetVideos:output_type -> pb.GetVideosResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadReq); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadResp); i {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeReq); i {
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
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeResp); i {
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
		file_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareReq); i {
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
		file_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareResp); i {
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
		file_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeReq); i {
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
		file_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeResp); i {
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
		file_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewReq); i {
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
		file_video_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewResp); i {
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
		file_video_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideosReq); i {
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
		file_video_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideosResp); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
