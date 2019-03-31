// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1p1beta1/web_detection.proto

package vision // import "google.golang.org/genproto/googleapis/cloud/vision/v1p1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Relevant information for the image from the Internet.
type WebDetection struct {
	// Deduced entities from similar images on the Internet.
	WebEntities []*WebDetection_WebEntity `protobuf:"bytes,1,rep,name=web_entities,json=webEntities,proto3" json:"web_entities,omitempty"`
	// Fully matching images from the Internet.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,2,rep,name=full_matching_images,json=fullMatchingImages,proto3" json:"full_matching_images,omitempty"`
	// Partial matching images from the Internet.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,3,rep,name=partial_matching_images,json=partialMatchingImages,proto3" json:"partial_matching_images,omitempty"`
	// Web pages containing the matching images from the Internet.
	PagesWithMatchingImages []*WebDetection_WebPage `protobuf:"bytes,4,rep,name=pages_with_matching_images,json=pagesWithMatchingImages,proto3" json:"pages_with_matching_images,omitempty"`
	// The visually similar image results.
	VisuallySimilarImages []*WebDetection_WebImage `protobuf:"bytes,6,rep,name=visually_similar_images,json=visuallySimilarImages,proto3" json:"visually_similar_images,omitempty"`
	// Best guess text labels for the request image.
	BestGuessLabels      []*WebDetection_WebLabel `protobuf:"bytes,8,rep,name=best_guess_labels,json=bestGuessLabels,proto3" json:"best_guess_labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *WebDetection) Reset()         { *m = WebDetection{} }
func (m *WebDetection) String() string { return proto.CompactTextString(m) }
func (*WebDetection) ProtoMessage()    {}
func (*WebDetection) Descriptor() ([]byte, []int) {
	return fileDescriptor_web_detection_1e42c8eaa5ae855c, []int{0}
}
func (m *WebDetection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection.Unmarshal(m, b)
}
func (m *WebDetection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection.Marshal(b, m, deterministic)
}
func (dst *WebDetection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection.Merge(dst, src)
}
func (m *WebDetection) XXX_Size() int {
	return xxx_messageInfo_WebDetection.Size(m)
}
func (m *WebDetection) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection proto.InternalMessageInfo

func (m *WebDetection) GetWebEntities() []*WebDetection_WebEntity {
	if m != nil {
		return m.WebEntities
	}
	return nil
}

func (m *WebDetection) GetFullMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.FullMatchingImages
	}
	return nil
}

func (m *WebDetection) GetPartialMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.PartialMatchingImages
	}
	return nil
}

func (m *WebDetection) GetPagesWithMatchingImages() []*WebDetection_WebPage {
	if m != nil {
		return m.PagesWithMatchingImages
	}
	return nil
}

func (m *WebDetection) GetVisuallySimilarImages() []*WebDetection_WebImage {
	if m != nil {
		return m.VisuallySimilarImages
	}
	return nil
}

func (m *WebDetection) GetBestGuessLabels() []*WebDetection_WebLabel {
	if m != nil {
		return m.BestGuessLabels
	}
	return nil
}

// Entity deduced from similar images on the Internet.
type WebDetection_WebEntity struct {
	// Opaque entity ID.
	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Overall relevancy score for the entity.
	// Not normalized and not comparable across different image queries.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Canonical description of the entity, in English.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebDetection_WebEntity) Reset()         { *m = WebDetection_WebEntity{} }
func (m *WebDetection_WebEntity) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebEntity) ProtoMessage()    {}
func (*WebDetection_WebEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_web_detection_1e42c8eaa5ae855c, []int{0, 0}
}
func (m *WebDetection_WebEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebEntity.Unmarshal(m, b)
}
func (m *WebDetection_WebEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebEntity.Marshal(b, m, deterministic)
}
func (dst *WebDetection_WebEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebEntity.Merge(dst, src)
}
func (m *WebDetection_WebEntity) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebEntity.Size(m)
}
func (m *WebDetection_WebEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebEntity.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebEntity proto.InternalMessageInfo

func (m *WebDetection_WebEntity) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *WebDetection_WebEntity) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WebDetection_WebEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Metadata for online images.
type WebDetection_WebImage struct {
	// The result image URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the image.
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebDetection_WebImage) Reset()         { *m = WebDetection_WebImage{} }
func (m *WebDetection_WebImage) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebImage) ProtoMessage()    {}
func (*WebDetection_WebImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_web_detection_1e42c8eaa5ae855c, []int{0, 1}
}
func (m *WebDetection_WebImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebImage.Unmarshal(m, b)
}
func (m *WebDetection_WebImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebImage.Marshal(b, m, deterministic)
}
func (dst *WebDetection_WebImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebImage.Merge(dst, src)
}
func (m *WebDetection_WebImage) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebImage.Size(m)
}
func (m *WebDetection_WebImage) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebImage.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebImage proto.InternalMessageInfo

func (m *WebDetection_WebImage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebDetection_WebImage) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Metadata for web pages.
type WebDetection_WebPage struct {
	// The result web page URL.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// (Deprecated) Overall relevancy score for the web page.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// Title for the web page, may contain HTML markups.
	PageTitle string `protobuf:"bytes,3,opt,name=page_title,json=pageTitle,proto3" json:"page_title,omitempty"`
	// Fully matching images on the page.
	// Can include resized copies of the query image.
	FullMatchingImages []*WebDetection_WebImage `protobuf:"bytes,4,rep,name=full_matching_images,json=fullMatchingImages,proto3" json:"full_matching_images,omitempty"`
	// Partial matching images on the page.
	// Those images are similar enough to share some key-point features. For
	// example an original image will likely have partial matching for its
	// crops.
	PartialMatchingImages []*WebDetection_WebImage `protobuf:"bytes,5,rep,name=partial_matching_images,json=partialMatchingImages,proto3" json:"partial_matching_images,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *WebDetection_WebPage) Reset()         { *m = WebDetection_WebPage{} }
func (m *WebDetection_WebPage) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebPage) ProtoMessage()    {}
func (*WebDetection_WebPage) Descriptor() ([]byte, []int) {
	return fileDescriptor_web_detection_1e42c8eaa5ae855c, []int{0, 2}
}
func (m *WebDetection_WebPage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebPage.Unmarshal(m, b)
}
func (m *WebDetection_WebPage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebPage.Marshal(b, m, deterministic)
}
func (dst *WebDetection_WebPage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebPage.Merge(dst, src)
}
func (m *WebDetection_WebPage) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebPage.Size(m)
}
func (m *WebDetection_WebPage) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebPage.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebPage proto.InternalMessageInfo

func (m *WebDetection_WebPage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *WebDetection_WebPage) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *WebDetection_WebPage) GetPageTitle() string {
	if m != nil {
		return m.PageTitle
	}
	return ""
}

func (m *WebDetection_WebPage) GetFullMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.FullMatchingImages
	}
	return nil
}

func (m *WebDetection_WebPage) GetPartialMatchingImages() []*WebDetection_WebImage {
	if m != nil {
		return m.PartialMatchingImages
	}
	return nil
}

// Label to provide extra metadata for the web detection.
type WebDetection_WebLabel struct {
	// Label for extra metadata.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// The BCP-47 language code for `label`, such as "en-US" or "sr-Latn".
	// For more information, see
	// http://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
	LanguageCode         string   `protobuf:"bytes,2,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebDetection_WebLabel) Reset()         { *m = WebDetection_WebLabel{} }
func (m *WebDetection_WebLabel) String() string { return proto.CompactTextString(m) }
func (*WebDetection_WebLabel) ProtoMessage()    {}
func (*WebDetection_WebLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_web_detection_1e42c8eaa5ae855c, []int{0, 3}
}
func (m *WebDetection_WebLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebDetection_WebLabel.Unmarshal(m, b)
}
func (m *WebDetection_WebLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebDetection_WebLabel.Marshal(b, m, deterministic)
}
func (dst *WebDetection_WebLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebDetection_WebLabel.Merge(dst, src)
}
func (m *WebDetection_WebLabel) XXX_Size() int {
	return xxx_messageInfo_WebDetection_WebLabel.Size(m)
}
func (m *WebDetection_WebLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_WebDetection_WebLabel.DiscardUnknown(m)
}

var xxx_messageInfo_WebDetection_WebLabel proto.InternalMessageInfo

func (m *WebDetection_WebLabel) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WebDetection_WebLabel) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func init() {
	proto.RegisterType((*WebDetection)(nil), "google.cloud.vision.v1p1beta1.WebDetection")
	proto.RegisterType((*WebDetection_WebEntity)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebEntity")
	proto.RegisterType((*WebDetection_WebImage)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebImage")
	proto.RegisterType((*WebDetection_WebPage)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebPage")
	proto.RegisterType((*WebDetection_WebLabel)(nil), "google.cloud.vision.v1p1beta1.WebDetection.WebLabel")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p1beta1/web_detection.proto", fileDescriptor_web_detection_1e42c8eaa5ae855c)
}

var fileDescriptor_web_detection_1e42c8eaa5ae855c = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x95, 0x76, 0x1b, 0x8b, 0x5b, 0x04, 0xb3, 0x86, 0x16, 0x05, 0x26, 0x15, 0xb8, 0xf4,
	0x94, 0xa8, 0x1b, 0x9c, 0xb8, 0x6d, 0x4c, 0x68, 0x12, 0x48, 0x55, 0x40, 0x1a, 0xe2, 0x92, 0x39,
	0x89, 0x97, 0xbe, 0x92, 0x1b, 0x47, 0xb1, 0xd3, 0xaa, 0x37, 0x4e, 0x7c, 0x14, 0x3e, 0x23, 0x47,
	0xf4, 0xda, 0xce, 0x54, 0x51, 0x36, 0x31, 0x86, 0xb8, 0xf9, 0x7d, 0xac, 0xe7, 0xf9, 0xd9, 0xaf,
	0xff, 0x90, 0x49, 0x29, 0x65, 0x29, 0x78, 0x9c, 0x0b, 0xd9, 0x16, 0xf1, 0x02, 0x14, 0xc8, 0x2a,
	0x5e, 0x4c, 0xea, 0x49, 0xc6, 0x35, 0x9b, 0xc4, 0x4b, 0x9e, 0xa5, 0x05, 0xd7, 0x3c, 0xd7, 0x20,
	0xab, 0xa8, 0x6e, 0xa4, 0x96, 0xf4, 0xd0, 0x5a, 0x22, 0x63, 0x89, 0xac, 0x25, 0xba, 0xb6, 0x84,
	0xcf, 0x5c, 0x22, 0xab, 0x21, 0x66, 0x55, 0x25, 0x35, 0x43, 0xaf, 0xb2, 0xe6, 0x17, 0xdf, 0x7c,
	0x32, 0xbc, 0xe0, 0xd9, 0xdb, 0x2e, 0x93, 0x7e, 0x26, 0x43, 0x84, 0xf0, 0x4a, 0x83, 0x06, 0xae,
	0x02, 0x6f, 0xd4, 0x1f, 0x0f, 0x8e, 0x5e, 0x47, 0xb7, 0x42, 0xa2, 0xf5, 0x08, 0x2c, 0xce, 0xd0,
	0xbe, 0x4a, 0x06, 0x4b, 0x37, 0x04, 0xae, 0xe8, 0x15, 0xd9, 0xbf, 0x6a, 0x85, 0x48, 0xe7, 0x4c,
	0xe7, 0x33, 0xa8, 0xca, 0x14, 0xe6, 0xac, 0xe4, 0x2a, 0xe8, 0x19, 0xc2, 0xab, 0x3b, 0x12, 0xce,
	0xd1, 0x9c, 0x50, 0x4c, 0xfc, 0xe0, 0x02, 0x8d, 0xa4, 0xa8, 0x20, 0x07, 0x35, 0x6b, 0x34, 0xb0,
	0x4d, 0x54, 0xff, 0x1e, 0xa8, 0x27, 0x2e, 0xf4, 0x17, 0x5a, 0x4d, 0xc2, 0x1a, 0x07, 0xe9, 0x12,
	0xf4, 0x6c, 0x03, 0xb8, 0x65, 0x80, 0xc7, 0x77, 0x04, 0x4e, 0x91, 0x77, 0x60, 0x62, 0x2f, 0x40,
	0xcf, 0x36, 0xf7, 0xb7, 0x00, 0xd5, 0x32, 0x21, 0x56, 0xa9, 0x82, 0x39, 0x08, 0xd6, 0x74, 0xb8,
	0x9d, 0xfb, 0xec, 0xaf, 0x0b, 0xfd, 0x68, 0x33, 0x1d, 0xed, 0x92, 0xec, 0x65, 0x5c, 0xe9, 0xb4,
	0x6c, 0xb9, 0x52, 0xa9, 0x60, 0x19, 0x17, 0x2a, 0xd8, 0xfd, 0x2b, 0xce, 0x7b, 0x34, 0x27, 0x8f,
	0x30, 0xee, 0x1d, 0xa6, 0x99, 0x5a, 0x85, 0x97, 0xc4, 0xbf, 0xbe, 0x31, 0xf4, 0x29, 0xf1, 0xcd,
	0xd5, 0x5b, 0xa5, 0x50, 0x04, 0xde, 0xc8, 0x1b, 0xfb, 0xc9, 0xae, 0x15, 0xce, 0x0b, 0xba, 0x4f,
	0xb6, 0x55, 0x2e, 0x1b, 0x1e, 0xf4, 0x46, 0xde, 0xb8, 0x97, 0xd8, 0x82, 0x8e, 0xc8, 0xa0, 0xe0,
	0x2a, 0x6f, 0xa0, 0x46, 0x50, 0xd0, 0x37, 0xa6, 0x75, 0x29, 0x3c, 0x22, 0xbb, 0xdd, 0x36, 0xe9,
	0x63, 0xd2, 0x6f, 0x1b, 0xe1, 0xa2, 0x71, 0xf8, 0xfb, 0xd4, 0xf0, 0x7b, 0x8f, 0x3c, 0x70, 0x47,
	0xf1, 0xa7, 0x1e, 0x7a, 0x48, 0x08, 0x1e, 0x5a, 0xaa, 0x41, 0x0b, 0xee, 0x16, 0xe2, 0xa3, 0xf2,
	0x09, 0x85, 0x1b, 0x1f, 0xc0, 0xd6, 0xff, 0x7b, 0x00, 0xdb, 0xff, 0xfc, 0x01, 0x84, 0x67, 0xa6,
	0xb9, 0xe6, 0x2c, 0xb1, 0x2d, 0xe6, 0x86, 0xb8, 0x56, 0xd9, 0x82, 0xbe, 0x24, 0x0f, 0x05, 0xab,
	0xca, 0x16, 0x5b, 0x93, 0xcb, 0xc2, 0x36, 0xcd, 0x4f, 0x86, 0x9d, 0x78, 0x2a, 0x0b, 0x7e, 0xf2,
	0xd5, 0x23, 0xcf, 0x73, 0x39, 0xbf, 0x7d, 0x65, 0x27, 0x7b, 0xeb, 0x4b, 0x9b, 0xe2, 0x0f, 0x36,
	0xf5, 0xbe, 0x9c, 0x3a, 0x4f, 0x29, 0x31, 0x31, 0x92, 0x4d, 0x19, 0x97, 0xbc, 0x32, 0xff, 0x5b,
	0x6c, 0xa7, 0x58, 0x0d, 0xea, 0x86, 0x2f, 0xf5, 0x8d, 0x15, 0x7e, 0x78, 0x5e, 0xb6, 0x63, 0x2c,
	0xc7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x19, 0xa7, 0x1d, 0x84, 0x05, 0x00, 0x00,
}
