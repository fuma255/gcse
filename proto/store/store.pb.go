// Code generated by protoc-gen-go.
// source: github.com/daviddengcn/gcse/proto/store/store.proto
// DO NOT EDIT!

/*
Package gcse_store is a generated protocol buffer package.

It is generated from these files:
	github.com/daviddengcn/gcse/proto/store/store.proto

It has these top-level messages:
	PackageInfo
*/
package gcse_store

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import gcse_spider "github.com/daviddengcn/gcse/proto/spider"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type PackageInfo struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Package     string                     `protobuf:"bytes,2,opt,name=package" json:"package,omitempty"`
	Author      string                     `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	Stars       int32                      `protobuf:"varint,4,opt,name=stars" json:"stars,omitempty"`
	Synopsis    string                     `protobuf:"bytes,5,opt,name=synopsis" json:"synopsis,omitempty"`
	Description string                     `protobuf:"bytes,6,opt,name=description" json:"description,omitempty"`
	ProjectUrl  string                     `protobuf:"bytes,7,opt,name=project_url,json=projectUrl" json:"project_url,omitempty"`
	ReadmeFn    string                     `protobuf:"bytes,8,opt,name=readme_fn,json=readmeFn" json:"readme_fn,omitempty"`
	ReadmeData  string                     `protobuf:"bytes,9,opt,name=readme_data,json=readmeData" json:"readme_data,omitempty"`
	Imports     []string                   `protobuf:"bytes,10,rep,name=imports" json:"imports,omitempty"`
	TestImports []string                   `protobuf:"bytes,11,rep,name=test_imports,json=testImports" json:"test_imports,omitempty"`
	Exported    []string                   `protobuf:"bytes,12,rep,name=exported" json:"exported,omitempty"`
	LastCrawled *google_protobuf.Timestamp `protobuf:"bytes,13,opt,name=last_crawled,json=lastCrawled" json:"last_crawled,omitempty"`
	// Available if the package is the repo's root.
	FolderInfo *gcse_spider.FolderInfo `protobuf:"bytes,14,opt,name=folder_info,json=folderInfo" json:"folder_info,omitempty"`
	// Available if the package is a repo's root.
	RepoInfo *gcse_spider.RepoInfo `protobuf:"bytes,15,opt,name=repo_info,json=repoInfo" json:"repo_info,omitempty"`
}

func (m *PackageInfo) Reset()                    { *m = PackageInfo{} }
func (m *PackageInfo) String() string            { return proto.CompactTextString(m) }
func (*PackageInfo) ProtoMessage()               {}
func (*PackageInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PackageInfo) GetLastCrawled() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastCrawled
	}
	return nil
}

func (m *PackageInfo) GetFolderInfo() *gcse_spider.FolderInfo {
	if m != nil {
		return m.FolderInfo
	}
	return nil
}

func (m *PackageInfo) GetRepoInfo() *gcse_spider.RepoInfo {
	if m != nil {
		return m.RepoInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*PackageInfo)(nil), "gcse.store.PackageInfo")
}

var fileDescriptor0 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x4f, 0x8f, 0x13, 0x21,
	0x14, 0xcf, 0xba, 0x6d, 0xb7, 0x7d, 0x54, 0x4d, 0x88, 0x7f, 0x48, 0x3d, 0x58, 0x3d, 0x79, 0x9a,
	0x49, 0x76, 0x63, 0x62, 0x62, 0x3c, 0x69, 0x36, 0xd9, 0x9b, 0x99, 0xe8, 0xb9, 0xa1, 0x03, 0xc3,
	0xa2, 0x33, 0x40, 0x80, 0xaa, 0xfb, 0x1d, 0xfc, 0xd0, 0xc2, 0x83, 0xd6, 0xf5, 0xe6, 0x65, 0xe6,
	0xbd, 0xdf, 0x1f, 0xe0, 0xf1, 0x03, 0xae, 0x94, 0x8e, 0xb7, 0x87, 0x7d, 0xd3, 0xdb, 0xa9, 0x15,
	0xfc, 0x87, 0x16, 0x42, 0x1a, 0xd5, 0x9b, 0x56, 0xf5, 0x41, 0xb6, 0xce, 0xdb, 0x68, 0xdb, 0x10,
	0xad, 0x97, 0xe5, 0xdb, 0x20, 0x42, 0x21, 0xb3, 0x0d, 0x22, 0x9b, 0xf7, 0xf7, 0x16, 0x50, 0x76,
	0xe4, 0x46, 0x15, 0xdb, 0xfe, 0x30, 0xb4, 0x2e, 0xde, 0x39, 0x19, 0xda, 0xa8, 0x27, 0x19, 0x22,
	0x9f, 0xdc, 0xdf, 0xaa, 0x2c, 0xb4, 0x79, 0xfb, 0x1f, 0xbb, 0x3b, 0x2d, 0xa4, 0xaf, 0xbf, 0x62,
	0x7b, 0xfd, 0x7b, 0x06, 0xe4, 0x33, 0xef, 0xbf, 0x73, 0x25, 0x6f, 0xcc, 0x60, 0x29, 0x85, 0x99,
	0xe1, 0x93, 0x64, 0x67, 0xdb, 0xb3, 0x37, 0xab, 0x0e, 0x6b, 0xca, 0xe0, 0xc2, 0x15, 0x09, 0x7b,
	0x80, 0xf0, 0xb1, 0xa5, 0xcf, 0x60, 0xc1, 0x0f, 0xf1, 0xd6, 0x7a, 0x76, 0x8e, 0x44, 0xed, 0xe8,
	0x13, 0x98, 0xa7, 0xb3, 0xf9, 0xc0, 0x66, 0x09, 0x9e, 0x77, 0xa5, 0xa1, 0x1b, 0x58, 0x86, 0x3b,
	0x63, 0x5d, 0xd0, 0x81, 0xcd, 0x51, 0x7f, 0xea, 0xe9, 0x16, 0x88, 0x90, 0xa1, 0xf7, 0xda, 0x45,
	0x6d, 0x0d, 0x5b, 0x20, 0x7d, 0x1f, 0xa2, 0x2f, 0x81, 0xa4, 0x23, 0x7f, 0x93, 0x7d, 0xdc, 0x1d,
	0xfc, 0xc8, 0x2e, 0x50, 0x01, 0x15, 0xfa, 0xea, 0x47, 0xfa, 0x02, 0x56, 0x5e, 0x72, 0x31, 0xc9,
	0xdd, 0x60, 0xd8, 0xb2, 0xac, 0x5f, 0x80, 0x6b, 0x74, 0x57, 0x52, 0xf0, 0xc8, 0xd9, 0xaa, 0xb8,
	0x0b, 0xf4, 0x29, 0x21, 0x79, 0x48, 0x3d, 0x39, 0xeb, 0x63, 0x60, 0xb0, 0x3d, 0xcf, 0x43, 0xd6,
	0x96, 0xbe, 0x82, 0x75, 0x4c, 0x57, 0xbd, 0x3b, 0xd2, 0x04, 0x69, 0x92, 0xb1, 0x9b, 0x2a, 0x49,
	0x93, 0xc9, 0x5f, 0xb9, 0x94, 0x82, 0xad, 0x91, 0x3e, 0xf5, 0xf4, 0x03, 0xac, 0x47, 0x9e, 0xec,
	0xbd, 0xe7, 0x3f, 0xc7, 0xc4, 0x3f, 0x4c, 0x5b, 0x93, 0xcb, 0x4d, 0xa3, 0xac, 0x55, 0x63, 0x7d,
	0x06, 0x29, 0xe1, 0xe6, 0xcb, 0x31, 0xd0, 0x8e, 0x64, 0xfd, 0xc7, 0x22, 0xa7, 0xef, 0x80, 0x0c,
	0x76, 0x4c, 0x81, 0xed, 0x74, 0xca, 0x87, 0x3d, 0x42, 0xf7, 0xf3, 0xa6, 0x3c, 0x9b, 0x92, 0xe4,
	0x35, 0xf2, 0x39, 0xbe, 0x0e, 0x86, 0x53, 0x4d, 0x2f, 0xf3, 0x7d, 0x38, 0x5b, 0x7c, 0x8f, 0xd1,
	0xf7, 0xf4, 0x1f, 0x5f, 0x97, 0x58, 0x74, 0x2d, 0x7d, 0xad, 0xf6, 0x0b, 0x3c, 0xce, 0xd5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x70, 0x1f, 0x1e, 0xcc, 0x02, 0x00, 0x00,
}
