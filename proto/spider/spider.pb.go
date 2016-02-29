// Code generated by protoc-gen-go.
// source: github.com/daviddengcn/gcse/proto/spider/spider.proto
// DO NOT EDIT!

/*
Package gcse_spider is a generated protocol buffer package.

It is generated from these files:
	github.com/daviddengcn/gcse/proto/spider/spider.proto

It has these top-level messages:
	GoFileInfo
	RepoInfo
	FolderInfo
*/
package gcse_spider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type GoFileInfo_Status int32

const (
	GoFileInfo_Unknown      GoFileInfo_Status = 0
	GoFileInfo_ParseSuccess GoFileInfo_Status = 1
	GoFileInfo_ParseFailed  GoFileInfo_Status = 2
	GoFileInfo_ShouldIgnore GoFileInfo_Status = 3
)

var GoFileInfo_Status_name = map[int32]string{
	0: "Unknown",
	1: "ParseSuccess",
	2: "ParseFailed",
	3: "ShouldIgnore",
}
var GoFileInfo_Status_value = map[string]int32{
	"Unknown":      0,
	"ParseSuccess": 1,
	"ParseFailed":  2,
	"ShouldIgnore": 3,
}

func (x GoFileInfo_Status) String() string {
	return proto.EnumName(GoFileInfo_Status_name, int32(x))
}
func (GoFileInfo_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type GoFileInfo struct {
	Status      GoFileInfo_Status `protobuf:"varint,1,opt,name=status,enum=gcse.spider.GoFileInfo_Status" json:"status,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	IsTest      bool              `protobuf:"varint,4,opt,name=is_test,json=isTest" json:"is_test,omitempty"`
	Imports     []string          `protobuf:"bytes,5,rep,name=imports" json:"imports,omitempty"`
}

func (m *GoFileInfo) Reset()                    { *m = GoFileInfo{} }
func (m *GoFileInfo) String() string            { return proto.CompactTextString(m) }
func (*GoFileInfo) ProtoMessage()               {}
func (*GoFileInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RepoInfo struct {
	LastCrawled *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=last_crawled,json=lastCrawled" json:"last_crawled,omitempty"`
	Stars       int32                      `protobuf:"varint,2,opt,name=stars" json:"stars,omitempty"`
	Description string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Where this project was forked from, full path
	Source      string                     `protobuf:"bytes,5,opt,name=source" json:"source,omitempty"`
	LastUpdated *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated" json:"last_updated,omitempty"`
}

func (m *RepoInfo) Reset()                    { *m = RepoInfo{} }
func (m *RepoInfo) String() string            { return proto.CompactTextString(m) }
func (*RepoInfo) ProtoMessage()               {}
func (*RepoInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RepoInfo) GetLastCrawled() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastCrawled
	}
	return nil
}

func (m *RepoInfo) GetLastUpdated() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdated
	}
	return nil
}

// Information for non-repository folder.
type FolderInfo struct {
	// E.g. "spider"
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// E.g. "spider/github"
	Path    string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha" json:"sha,omitempty"`
	HtmlUrl string `protobuf:"bytes,4,opt,name=html_url,json=htmlUrl" json:"html_url,omitempty"`
	// The timestamp this folder-info is crawled
	CrawlingTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=crawling_time,json=crawlingTime" json:"crawling_time,omitempty"`
}

func (m *FolderInfo) Reset()                    { *m = FolderInfo{} }
func (m *FolderInfo) String() string            { return proto.CompactTextString(m) }
func (*FolderInfo) ProtoMessage()               {}
func (*FolderInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FolderInfo) GetCrawlingTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CrawlingTime
	}
	return nil
}

func init() {
	proto.RegisterType((*GoFileInfo)(nil), "gcse.spider.GoFileInfo")
	proto.RegisterType((*RepoInfo)(nil), "gcse.spider.RepoInfo")
	proto.RegisterType((*FolderInfo)(nil), "gcse.spider.FolderInfo")
	proto.RegisterEnum("gcse.spider.GoFileInfo_Status", GoFileInfo_Status_name, GoFileInfo_Status_value)
}

var fileDescriptor0 = []byte{
	// 444 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0xdb, 0x8f, 0x2d, 0x4e, 0x81, 0xc8, 0x42, 0x10, 0xf6, 0x80, 0xaa, 0x9e, 0x38, 0x39,
	0xd2, 0x22, 0xb8, 0x20, 0xc4, 0x01, 0xa9, 0x68, 0x39, 0xa1, 0x74, 0x7b, 0xae, 0xdc, 0xc4, 0xeb,
	0x58, 0x38, 0xb6, 0xe5, 0x0f, 0x56, 0xfc, 0x1d, 0x7e, 0x15, 0x7f, 0x85, 0x1b, 0x63, 0xa7, 0x61,
	0x7b, 0x42, 0x3d, 0x75, 0xe6, 0xcd, 0x9b, 0xf8, 0xcd, 0x7b, 0x45, 0xef, 0xb8, 0xf0, 0x5d, 0x38,
	0x90, 0x46, 0xf7, 0x55, 0x4b, 0x7f, 0x88, 0xb6, 0x65, 0x8a, 0x37, 0xaa, 0xe2, 0x8d, 0x63, 0x95,
	0xb1, 0xda, 0xeb, 0xca, 0x19, 0xd1, 0x32, 0x7b, 0xfc, 0x21, 0x09, 0xc3, 0x79, 0x9c, 0x93, 0x01,
	0xba, 0xfa, 0x70, 0xf2, 0x0d, 0xae, 0x25, 0x55, 0x7c, 0xd8, 0x3c, 0x84, 0xbb, 0xca, 0xf8, 0x9f,
	0x86, 0xb9, 0xca, 0x8b, 0x9e, 0x39, 0x4f, 0x7b, 0xf3, 0x50, 0x0d, 0x5f, 0x5a, 0xff, 0xc9, 0x10,
	0xfa, 0xa2, 0x37, 0x42, 0xb2, 0x1b, 0x75, 0xa7, 0xf1, 0x7b, 0x34, 0x87, 0xa9, 0x0f, 0xae, 0xcc,
	0x56, 0xd9, 0x9b, 0xa7, 0xd7, 0xaf, 0xc9, 0xc9, 0x4b, 0xe4, 0x81, 0x48, 0xb6, 0x89, 0x55, 0x1f,
	0xd9, 0x18, 0xa3, 0xa9, 0xa2, 0x3d, 0x2b, 0x2f, 0x60, 0xeb, 0x71, 0x9d, 0x6a, 0xbc, 0x42, 0x79,
	0xcb, 0x5c, 0x63, 0x85, 0xf1, 0x42, 0xab, 0x72, 0x92, 0x46, 0xa7, 0x10, 0x7e, 0x89, 0x2e, 0x85,
	0xdb, 0x7b, 0x10, 0x54, 0x4e, 0x61, 0xba, 0xa8, 0xe7, 0xc2, 0xdd, 0x42, 0x87, 0x4b, 0x18, 0xf4,
	0x46, 0x5b, 0xef, 0xca, 0xd9, 0x6a, 0x02, 0x6b, 0x63, 0xbb, 0xfe, 0x8a, 0xe6, 0xc3, 0xd3, 0x38,
	0x47, 0x97, 0x3b, 0xf5, 0x5d, 0xe9, 0x7b, 0x55, 0x3c, 0xc2, 0x05, 0x5a, 0x7e, 0xa3, 0xd6, 0xb1,
	0x6d, 0x68, 0x1a, 0xe6, 0x5c, 0x91, 0xe1, 0x67, 0x28, 0x4f, 0xc8, 0x86, 0x82, 0xe4, 0xb6, 0xb8,
	0x88, 0x94, 0x6d, 0xa7, 0x83, 0x6c, 0x6f, 0xb8, 0xd2, 0x96, 0x15, 0x93, 0xf5, 0xef, 0x0c, 0x2d,
	0x6a, 0x66, 0x74, 0xba, 0xfc, 0x23, 0x5a, 0x4a, 0xea, 0xfc, 0xbe, 0xb1, 0xf4, 0x1e, 0x16, 0xd2,
	0xfd, 0xf9, 0xf5, 0x15, 0xe1, 0x5a, 0x73, 0xc9, 0xc8, 0xe8, 0x28, 0xb9, 0x1d, 0x0d, 0xac, 0xf3,
	0xc8, 0xff, 0x3c, 0xd0, 0xf1, 0x73, 0x34, 0x03, 0xd4, 0xba, 0xe4, 0xc0, 0xac, 0x1e, 0x9a, 0x33,
	0x2c, 0x78, 0x01, 0x86, 0xeb, 0x60, 0x1b, 0x06, 0x87, 0xc6, 0xe1, 0xb1, 0xfb, 0x27, 0x27, 0x98,
	0x96, 0x7a, 0x90, 0x33, 0x3d, 0x4f, 0xce, 0x6e, 0xa0, 0xaf, 0x7f, 0x41, 0xac, 0x1b, 0x2d, 0x21,
	0xb4, 0x74, 0xdc, 0x18, 0x4f, 0x76, 0x12, 0x0f, 0x60, 0x86, 0xfa, 0x6e, 0x8c, 0x2c, 0xd6, 0xe0,
	0xd1, 0xc4, 0x75, 0xf4, 0xa8, 0x33, 0x96, 0xf8, 0x15, 0x5a, 0x74, 0xbe, 0x97, 0xfb, 0x60, 0x65,
	0xd2, 0x00, 0x51, 0xc4, 0x7e, 0x67, 0x25, 0xfe, 0x84, 0x9e, 0x24, 0xb3, 0x84, 0xe2, 0xfb, 0xf8,
	0xb7, 0x4a, 0x17, 0xfc, 0x5f, 0xe3, 0x72, 0x5c, 0x88, 0xd0, 0x61, 0x9e, 0x18, 0x6f, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0x6e, 0xd3, 0x65, 0x05, 0x03, 0x00, 0x00,
}
