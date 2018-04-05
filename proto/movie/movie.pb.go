// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/movie/movie.proto

/*
Package movie is a generated protocol buffer package.

It is generated from these files:
	proto/movie/movie.proto

It has these top-level messages:
	GetRequest
	ListRequest
	Movie
	Response
	Error
*/
package movie

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Page uint32 `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type Movie struct {
	Imdb   string `protobuf:"bytes,1,opt,name=imdb" json:"imdb,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Year   string `protobuf:"bytes,3,opt,name=year" json:"year,omitempty"`
	Poster string `protobuf:"bytes,4,opt,name=poster" json:"poster,omitempty"`
	Genre  string `protobuf:"bytes,5,opt,name=genre" json:"genre,omitempty"`
	Plot   string `protobuf:"bytes,6,opt,name=plot" json:"plot,omitempty"`
	Rating string `protobuf:"bytes,7,opt,name=rating" json:"rating,omitempty"`
}

func (m *Movie) Reset()                    { *m = Movie{} }
func (m *Movie) String() string            { return proto.CompactTextString(m) }
func (*Movie) ProtoMessage()               {}
func (*Movie) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Movie) GetImdb() string {
	if m != nil {
		return m.Imdb
	}
	return ""
}

func (m *Movie) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Movie) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func (m *Movie) GetPoster() string {
	if m != nil {
		return m.Poster
	}
	return ""
}

func (m *Movie) GetGenre() string {
	if m != nil {
		return m.Genre
	}
	return ""
}

func (m *Movie) GetPlot() string {
	if m != nil {
		return m.Plot
	}
	return ""
}

func (m *Movie) GetRating() string {
	if m != nil {
		return m.Rating
	}
	return ""
}

type Response struct {
	Movie  *Movie   `protobuf:"bytes,1,opt,name=movie" json:"movie,omitempty"`
	Movies []*Movie `protobuf:"bytes,2,rep,name=movies" json:"movies,omitempty"`
	Errors []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetMovie() *Movie {
	if m != nil {
		return m.Movie
	}
	return nil
}

func (m *Response) GetMovies() []*Movie {
	if m != nil {
		return m.Movies
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code        int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "movie.GetRequest")
	proto.RegisterType((*ListRequest)(nil), "movie.ListRequest")
	proto.RegisterType((*Movie)(nil), "movie.Movie")
	proto.RegisterType((*Response)(nil), "movie.Response")
	proto.RegisterType((*Error)(nil), "movie.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for MovieService service

type MovieServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*Response, error)
}

type movieServiceClient struct {
	c           client.Client
	serviceName string
}

func NewMovieServiceClient(serviceName string, c client.Client) MovieServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "movie"
	}
	return &movieServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *movieServiceClient) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "MovieService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "MovieService.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MovieService service

type MovieServiceHandler interface {
	Get(context.Context, *GetRequest, *Response) error
	List(context.Context, *ListRequest, *Response) error
}

func RegisterMovieServiceHandler(s server.Server, hdlr MovieServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&MovieService{hdlr}, opts...))
}

type MovieService struct {
	MovieServiceHandler
}

func (h *MovieService) Get(ctx context.Context, in *GetRequest, out *Response) error {
	return h.MovieServiceHandler.Get(ctx, in, out)
}

func (h *MovieService) List(ctx context.Context, in *ListRequest, out *Response) error {
	return h.MovieServiceHandler.List(ctx, in, out)
}

func init() { proto.RegisterFile("proto/movie/movie.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x10, 0xfd, 0xb6, 0xdb, 0xec, 0xa7, 0xd3, 0x8a, 0x38, 0x88, 0x06, 0x4f, 0x4b, 0xf0, 0x50, 0x10,
	0x2b, 0x54, 0x3c, 0x7a, 0x94, 0x5e, 0xf4, 0x12, 0x7f, 0xc1, 0xb6, 0x1d, 0x4a, 0xc4, 0x6e, 0xd6,
	0x24, 0x16, 0xf4, 0xbf, 0xf8, 0x5f, 0x25, 0x93, 0x54, 0x5b, 0x14, 0x2f, 0xcb, 0x9b, 0xf7, 0xde,
	0xbe, 0x99, 0x9d, 0x59, 0x38, 0xed, 0x9c, 0x0d, 0xf6, 0x6a, 0x65, 0xd7, 0x86, 0xd2, 0x73, 0xcc,
	0x0c, 0x0a, 0x2e, 0x54, 0x0d, 0x30, 0xa5, 0xa0, 0xe9, 0xe5, 0x95, 0x7c, 0x40, 0x84, 0x7e, 0xdb,
	0xac, 0x48, 0x16, 0x75, 0x31, 0xda, 0xd7, 0x8c, 0xd5, 0x0d, 0x0c, 0xee, 0x8d, 0xff, 0xcb, 0x12,
	0xb9, 0xae, 0x59, 0x92, 0xec, 0xd5, 0xc5, 0xe8, 0x40, 0x33, 0x56, 0x1f, 0x05, 0x88, 0x87, 0xd8,
	0x22, 0xaa, 0x66, 0xb5, 0x98, 0x6d, 0xde, 0x88, 0xf8, 0x2b, 0xa5, 0xb7, 0x9b, 0xf2, 0x46, 0x8d,
	0x93, 0x65, 0xe2, 0x22, 0xc6, 0x13, 0xa8, 0x3a, 0xeb, 0x03, 0x39, 0xd9, 0x67, 0x36, 0x57, 0x78,
	0x0c, 0x62, 0x49, 0xad, 0x23, 0x29, 0x98, 0x4e, 0x05, 0xcf, 0xf1, 0x6c, 0x83, 0xac, 0x52, 0x42,
	0xc4, 0x31, 0xc1, 0x35, 0xc1, 0xb4, 0x4b, 0xf9, 0x3f, 0x25, 0xa4, 0x4a, 0xbd, 0xc3, 0x9e, 0x26,
	0xdf, 0xd9, 0xd6, 0x13, 0x2a, 0x48, 0xdb, 0xe0, 0x11, 0x07, 0x93, 0xe1, 0x38, 0x2d, 0x8a, 0xc7,
	0xd7, 0x49, 0xc2, 0x73, 0xa8, 0x18, 0x78, 0xd9, 0xab, 0xcb, 0x1f, 0xa6, 0xac, 0x45, 0x17, 0x39,
	0x67, 0x9d, 0x97, 0xe5, 0x8e, 0xeb, 0x2e, 0x92, 0x3a, 0x6b, 0xea, 0x16, 0x04, 0x13, 0x71, 0xe0,
	0xb9, 0x5d, 0xa4, 0xbe, 0x42, 0x33, 0xc6, 0x1a, 0x06, 0x0b, 0xf2, 0x73, 0x67, 0xba, 0x60, 0x6c,
	0x9b, 0x37, 0xb4, 0x4d, 0x4d, 0x9e, 0x60, 0xc8, 0x5d, 0x1f, 0xc9, 0xad, 0xcd, 0x9c, 0xf0, 0x02,
	0xca, 0x29, 0x05, 0x3c, 0xca, 0xbd, 0xbe, 0xef, 0x79, 0x76, 0x98, 0xa9, 0xcd, 0x97, 0xaa, 0x7f,
	0x78, 0x09, 0xfd, 0x78, 0x4e, 0xc4, 0x2c, 0x6d, 0xdd, 0xf6, 0x17, 0xfb, 0xac, 0xe2, 0xbf, 0xe5,
	0xfa, 0x33, 0x00, 0x00, 0xff, 0xff, 0x48, 0x21, 0x18, 0xcc, 0x48, 0x02, 0x00, 0x00,
}