// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openconfig.proto

/*
Package openconfig is a generated protocol buffer package.

It is generated from these files:
	openconfig.proto

It has these top-level messages:
	ExecRequest
	ExecReply
	RegisterRequest
	RegisterReply
	RegisterModuleRequest
	RegisterModuleReply
	ExecModuleRequest
	ExecModuleReply
	SubscribeRequest
	ConfigRequest
	ConfigReply
	ShowRequest
	ShowReply
*/
package openconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Command message type.
type ExecType int32

const (
	ExecType_EXEC                    ExecType = 0
	ExecType_COMPLETE                ExecType = 1
	ExecType_COMPLETE_TRAILING_SPACE ExecType = 2
	ExecType_COMPLETE_FIRST_COMMANDS ExecType = 3
	ExecType_COMPLETE_DYNAMIC        ExecType = 4
)

var ExecType_name = map[int32]string{
	0: "EXEC",
	1: "COMPLETE",
	2: "COMPLETE_TRAILING_SPACE",
	3: "COMPLETE_FIRST_COMMANDS",
	4: "COMPLETE_DYNAMIC",
}
var ExecType_value = map[string]int32{
	"EXEC":                    0,
	"COMPLETE":                1,
	"COMPLETE_TRAILING_SPACE": 2,
	"COMPLETE_FIRST_COMMANDS": 3,
	"COMPLETE_DYNAMIC":        4,
}

func (x ExecType) String() string {
	return proto.EnumName(ExecType_name, int32(x))
}
func (ExecType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Command reply format type.
type ExecConentType int32

const (
	ExecConentType_CONTENT_TEXT ExecConentType = 0
	ExecConentType_CONTENT_JSON ExecConentType = 1
)

var ExecConentType_name = map[int32]string{
	0: "CONTENT_TEXT",
	1: "CONTENT_JSON",
}
var ExecConentType_value = map[string]int32{
	"CONTENT_TEXT": 0,
	"CONTENT_JSON": 1,
}

func (x ExecConentType) String() string {
	return proto.EnumName(ExecConentType_name, int32(x))
}
func (ExecConentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Command exec code.
type ExecCode int32

const (
	ExecCode_SUCCESS       ExecCode = 0
	ExecCode_NOMATCH       ExecCode = 1
	ExecCode_INCOMPLETE    ExecCode = 2
	ExecCode_AMBIGUOUS     ExecCode = 3
	ExecCode_SHOW          ExecCode = 4
	ExecCode_REDIRECT      ExecCode = 5
	ExecCode_REDIRECT_SHOW ExecCode = 6
)

var ExecCode_name = map[int32]string{
	0: "SUCCESS",
	1: "NOMATCH",
	2: "INCOMPLETE",
	3: "AMBIGUOUS",
	4: "SHOW",
	5: "REDIRECT",
	6: "REDIRECT_SHOW",
}
var ExecCode_value = map[string]int32{
	"SUCCESS":       0,
	"NOMATCH":       1,
	"INCOMPLETE":    2,
	"AMBIGUOUS":     3,
	"SHOW":          4,
	"REDIRECT":      5,
	"REDIRECT_SHOW": 6,
}

func (x ExecCode) String() string {
	return proto.EnumName(ExecCode_name, int32(x))
}
func (ExecCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Configuration service
type ConfigType int32

const (
	ConfigType_SET               ConfigType = 0
	ConfigType_DELETE            ConfigType = 1
	ConfigType_VALIDATE_START    ConfigType = 2
	ConfigType_VALIDATE_END      ConfigType = 3
	ConfigType_VALIDATE_SUCCESS  ConfigType = 4
	ConfigType_VALIDATE_FAILED   ConfigType = 5
	ConfigType_COMMIT_START      ConfigType = 6
	ConfigType_COMMIT_END        ConfigType = 7
	ConfigType_SUBSCRIBE         ConfigType = 8
	ConfigType_SUBSCRIBE_MULTI   ConfigType = 9
	ConfigType_SUBSCRIBE_REQUEST ConfigType = 10
	ConfigType_JSON_CONFIG       ConfigType = 11
)

var ConfigType_name = map[int32]string{
	0:  "SET",
	1:  "DELETE",
	2:  "VALIDATE_START",
	3:  "VALIDATE_END",
	4:  "VALIDATE_SUCCESS",
	5:  "VALIDATE_FAILED",
	6:  "COMMIT_START",
	7:  "COMMIT_END",
	8:  "SUBSCRIBE",
	9:  "SUBSCRIBE_MULTI",
	10: "SUBSCRIBE_REQUEST",
	11: "JSON_CONFIG",
}
var ConfigType_value = map[string]int32{
	"SET":               0,
	"DELETE":            1,
	"VALIDATE_START":    2,
	"VALIDATE_END":      3,
	"VALIDATE_SUCCESS":  4,
	"VALIDATE_FAILED":   5,
	"COMMIT_START":      6,
	"COMMIT_END":        7,
	"SUBSCRIBE":         8,
	"SUBSCRIBE_MULTI":   9,
	"SUBSCRIBE_REQUEST": 10,
	"JSON_CONFIG":       11,
}

func (x ConfigType) String() string {
	return proto.EnumName(ConfigType_name, int32(x))
}
func (ConfigType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Subscribe type.
type SubscribeType int32

const (
	SubscribeType_COMMAND SubscribeType = 0
	SubscribeType_JSON    SubscribeType = 1
)

var SubscribeType_name = map[int32]string{
	0: "COMMAND",
	1: "JSON",
}
var SubscribeType_value = map[string]int32{
	"COMMAND": 0,
	"JSON":    1,
}

func (x SubscribeType) String() string {
	return proto.EnumName(SubscribeType_name, int32(x))
}
func (SubscribeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// The request message containing user input string.
type ExecRequest struct {
	Type      ExecType       `protobuf:"varint,1,opt,name=type,enum=openconfig.ExecType" json:"type,omitempty"`
	Mode      string         `protobuf:"bytes,2,opt,name=mode" json:"mode,omitempty"`
	Privilege uint32         `protobuf:"varint,3,opt,name=privilege" json:"privilege,omitempty"`
	Line      string         `protobuf:"bytes,4,opt,name=line" json:"line,omitempty"`
	Commands  []string       `protobuf:"bytes,5,rep,name=commands" json:"commands,omitempty"`
	Args      []string       `protobuf:"bytes,6,rep,name=args" json:"args,omitempty"`
	Content   ExecConentType `protobuf:"varint,7,opt,name=content,enum=openconfig.ExecConentType" json:"content,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExecRequest) GetType() ExecType {
	if m != nil {
		return m.Type
	}
	return ExecType_EXEC
}

func (m *ExecRequest) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *ExecRequest) GetPrivilege() uint32 {
	if m != nil {
		return m.Privilege
	}
	return 0
}

func (m *ExecRequest) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *ExecRequest) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *ExecRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ExecRequest) GetContent() ExecConentType {
	if m != nil {
		return m.Content
	}
	return ExecConentType_CONTENT_TEXT
}

// The response message containing the completion with help.
type ExecReply struct {
	Code       ExecCode `protobuf:"varint,1,opt,name=code,enum=openconfig.ExecCode" json:"code,omitempty"`
	Lines      string   `protobuf:"bytes,2,opt,name=lines" json:"lines,omitempty"`
	Port       uint32   `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Candidates []string `protobuf:"bytes,4,rep,name=candidates" json:"candidates,omitempty"`
}

func (m *ExecReply) Reset()                    { *m = ExecReply{} }
func (m *ExecReply) String() string            { return proto.CompactTextString(m) }
func (*ExecReply) ProtoMessage()               {}
func (*ExecReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExecReply) GetCode() ExecCode {
	if m != nil {
		return m.Code
	}
	return ExecCode_SUCCESS
}

func (m *ExecReply) GetLines() string {
	if m != nil {
		return m.Lines
	}
	return ""
}

func (m *ExecReply) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ExecReply) GetCandidates() []string {
	if m != nil {
		return m.Candidates
	}
	return nil
}

// The request message for command definition.
type RegisterRequest struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Module    string   `protobuf:"bytes,2,opt,name=module" json:"module,omitempty"`
	Mode      string   `protobuf:"bytes,3,opt,name=mode" json:"mode,omitempty"`
	Line      string   `protobuf:"bytes,4,opt,name=line" json:"line,omitempty"`
	Privilege uint32   `protobuf:"varint,5,opt,name=privilege" json:"privilege,omitempty"`
	Helps     []string `protobuf:"bytes,6,rep,name=helps" json:"helps,omitempty"`
	Code      ExecCode `protobuf:"varint,7,opt,name=code,enum=openconfig.ExecCode" json:"code,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RegisterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *RegisterRequest) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *RegisterRequest) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *RegisterRequest) GetPrivilege() uint32 {
	if m != nil {
		return m.Privilege
	}
	return 0
}

func (m *RegisterRequest) GetHelps() []string {
	if m != nil {
		return m.Helps
	}
	return nil
}

func (m *RegisterRequest) GetCode() ExecCode {
	if m != nil {
		return m.Code
	}
	return ExecCode_SUCCESS
}

// The response message for callback ID.
type RegisterReply struct {
	Callbackid int32 `protobuf:"varint,1,opt,name=callbackid" json:"callbackid,omitempty"`
}

func (m *RegisterReply) Reset()                    { *m = RegisterReply{} }
func (m *RegisterReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()               {}
func (*RegisterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterReply) GetCallbackid() int32 {
	if m != nil {
		return m.Callbackid
	}
	return 0
}

// The request message mode information.
type RegisterModuleRequest struct {
	Module string `protobuf:"bytes,1,opt,name=module" json:"module,omitempty"`
	Port   string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *RegisterModuleRequest) Reset()                    { *m = RegisterModuleRequest{} }
func (m *RegisterModuleRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterModuleRequest) ProtoMessage()               {}
func (*RegisterModuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterModuleRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *RegisterModuleRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

// The response message containing the callbackID for the mode change.
type RegisterModuleReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *RegisterModuleReply) Reset()                    { *m = RegisterModuleReply{} }
func (m *RegisterModuleReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterModuleReply) ProtoMessage()               {}
func (*RegisterModuleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RegisterModuleReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

// The request message mode information.
type ExecModuleRequest struct {
	Callbackid int32    `protobuf:"varint,1,opt,name=callbackid" json:"callbackid,omitempty"`
	Args       []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *ExecModuleRequest) Reset()                    { *m = ExecModuleRequest{} }
func (m *ExecModuleRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecModuleRequest) ProtoMessage()               {}
func (*ExecModuleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ExecModuleRequest) GetCallbackid() int32 {
	if m != nil {
		return m.Callbackid
	}
	return 0
}

func (m *ExecModuleRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// The response message containing the callbackID for the mode change.
type ExecModuleReply struct {
	Result int32  `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Line   string `protobuf:"bytes,2,opt,name=line" json:"line,omitempty"`
}

func (m *ExecModuleReply) Reset()                    { *m = ExecModuleReply{} }
func (m *ExecModuleReply) String() string            { return proto.CompactTextString(m) }
func (*ExecModuleReply) ProtoMessage()               {}
func (*ExecModuleReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ExecModuleReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *ExecModuleReply) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

type SubscribeRequest struct {
	Type SubscribeType `protobuf:"varint,1,opt,name=type,enum=openconfig.SubscribeType" json:"type,omitempty"`
	Path string        `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SubscribeRequest) GetType() SubscribeType {
	if m != nil {
		return m.Type
	}
	return SubscribeType_COMMAND
}

func (m *SubscribeRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ConfigRequest struct {
	Type      ConfigType          `protobuf:"varint,1,opt,name=type,enum=openconfig.ConfigType" json:"type,omitempty"`
	Module    string              `protobuf:"bytes,2,opt,name=module" json:"module,omitempty"`
	Port      uint32              `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Subtype   SubscribeType       `protobuf:"varint,4,opt,name=subtype,enum=openconfig.SubscribeType" json:"subtype,omitempty"`
	Path      []string            `protobuf:"bytes,5,rep,name=path" json:"path,omitempty"`
	Subscribe []*SubscribeRequest `protobuf:"bytes,6,rep,name=subscribe" json:"subscribe,omitempty"`
}

func (m *ConfigRequest) Reset()                    { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()               {}
func (*ConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ConfigRequest) GetType() ConfigType {
	if m != nil {
		return m.Type
	}
	return ConfigType_SET
}

func (m *ConfigRequest) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ConfigRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ConfigRequest) GetSubtype() SubscribeType {
	if m != nil {
		return m.Subtype
	}
	return SubscribeType_COMMAND
}

func (m *ConfigRequest) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *ConfigRequest) GetSubscribe() []*SubscribeRequest {
	if m != nil {
		return m.Subscribe
	}
	return nil
}

type ConfigReply struct {
	Result int32      `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Type   ConfigType `protobuf:"varint,2,opt,name=type,enum=openconfig.ConfigType" json:"type,omitempty"`
	Path   []string   `protobuf:"bytes,3,rep,name=path" json:"path,omitempty"`
	Json   string     `protobuf:"bytes,4,opt,name=json" json:"json,omitempty"`
}

func (m *ConfigReply) Reset()                    { *m = ConfigReply{} }
func (m *ConfigReply) String() string            { return proto.CompactTextString(m) }
func (*ConfigReply) ProtoMessage()               {}
func (*ConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ConfigReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *ConfigReply) GetType() ConfigType {
	if m != nil {
		return m.Type
	}
	return ConfigType_SET
}

func (m *ConfigReply) GetPath() []string {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *ConfigReply) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

// Show service.
type ShowRequest struct {
	Line string `protobuf:"bytes,1,opt,name=line" json:"line,omitempty"`
	Json bool   `protobuf:"varint,2,opt,name=json" json:"json,omitempty"`
}

func (m *ShowRequest) Reset()                    { *m = ShowRequest{} }
func (m *ShowRequest) String() string            { return proto.CompactTextString(m) }
func (*ShowRequest) ProtoMessage()               {}
func (*ShowRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ShowRequest) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func (m *ShowRequest) GetJson() bool {
	if m != nil {
		return m.Json
	}
	return false
}

type ShowReply struct {
	Str string `protobuf:"bytes,1,opt,name=str" json:"str,omitempty"`
}

func (m *ShowReply) Reset()                    { *m = ShowReply{} }
func (m *ShowReply) String() string            { return proto.CompactTextString(m) }
func (*ShowReply) ProtoMessage()               {}
func (*ShowReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ShowReply) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecRequest)(nil), "openconfig.ExecRequest")
	proto.RegisterType((*ExecReply)(nil), "openconfig.ExecReply")
	proto.RegisterType((*RegisterRequest)(nil), "openconfig.RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "openconfig.RegisterReply")
	proto.RegisterType((*RegisterModuleRequest)(nil), "openconfig.RegisterModuleRequest")
	proto.RegisterType((*RegisterModuleReply)(nil), "openconfig.RegisterModuleReply")
	proto.RegisterType((*ExecModuleRequest)(nil), "openconfig.ExecModuleRequest")
	proto.RegisterType((*ExecModuleReply)(nil), "openconfig.ExecModuleReply")
	proto.RegisterType((*SubscribeRequest)(nil), "openconfig.SubscribeRequest")
	proto.RegisterType((*ConfigRequest)(nil), "openconfig.ConfigRequest")
	proto.RegisterType((*ConfigReply)(nil), "openconfig.ConfigReply")
	proto.RegisterType((*ShowRequest)(nil), "openconfig.ShowRequest")
	proto.RegisterType((*ShowReply)(nil), "openconfig.ShowReply")
	proto.RegisterEnum("openconfig.ExecType", ExecType_name, ExecType_value)
	proto.RegisterEnum("openconfig.ExecConentType", ExecConentType_name, ExecConentType_value)
	proto.RegisterEnum("openconfig.ExecCode", ExecCode_name, ExecCode_value)
	proto.RegisterEnum("openconfig.ConfigType", ConfigType_name, ConfigType_value)
	proto.RegisterEnum("openconfig.SubscribeType", SubscribeType_name, SubscribeType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Exec service

type ExecClient interface {
	DoExec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecReply, error)
}

type execClient struct {
	cc *grpc.ClientConn
}

func NewExecClient(cc *grpc.ClientConn) ExecClient {
	return &execClient{cc}
}

func (c *execClient) DoExec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecReply, error) {
	out := new(ExecReply)
	err := grpc.Invoke(ctx, "/openconfig.Exec/DoExec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Exec service

type ExecServer interface {
	DoExec(context.Context, *ExecRequest) (*ExecReply, error)
}

func RegisterExecServer(s *grpc.Server, srv ExecServer) {
	s.RegisterService(&_Exec_serviceDesc, srv)
}

func _Exec_DoExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecServer).DoExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openconfig.Exec/DoExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecServer).DoExec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Exec_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openconfig.Exec",
	HandlerType: (*ExecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoExec",
			Handler:    _Exec_DoExec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openconfig.proto",
}

// Client API for Register service

type RegisterClient interface {
	DoRegister(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	DoRegisterModule(ctx context.Context, in *RegisterModuleRequest, opts ...grpc.CallOption) (*RegisterModuleReply, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) DoRegister(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := grpc.Invoke(ctx, "/openconfig.Register/DoRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registerClient) DoRegisterModule(ctx context.Context, in *RegisterModuleRequest, opts ...grpc.CallOption) (*RegisterModuleReply, error) {
	out := new(RegisterModuleReply)
	err := grpc.Invoke(ctx, "/openconfig.Register/DoRegisterModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Register service

type RegisterServer interface {
	DoRegister(context.Context, *RegisterRequest) (*RegisterReply, error)
	DoRegisterModule(context.Context, *RegisterModuleRequest) (*RegisterModuleReply, error)
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_DoRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openconfig.Register/DoRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoRegister(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Register_DoRegisterModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).DoRegisterModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openconfig.Register/DoRegisterModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).DoRegisterModule(ctx, req.(*RegisterModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openconfig.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoRegister",
			Handler:    _Register_DoRegister_Handler,
		},
		{
			MethodName: "DoRegisterModule",
			Handler:    _Register_DoRegisterModule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openconfig.proto",
}

// Client API for ExecModule service

type ExecModuleClient interface {
	DoExecModule(ctx context.Context, in *ExecModuleRequest, opts ...grpc.CallOption) (*ExecModuleReply, error)
}

type execModuleClient struct {
	cc *grpc.ClientConn
}

func NewExecModuleClient(cc *grpc.ClientConn) ExecModuleClient {
	return &execModuleClient{cc}
}

func (c *execModuleClient) DoExecModule(ctx context.Context, in *ExecModuleRequest, opts ...grpc.CallOption) (*ExecModuleReply, error) {
	out := new(ExecModuleReply)
	err := grpc.Invoke(ctx, "/openconfig.ExecModule/DoExecModule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExecModule service

type ExecModuleServer interface {
	DoExecModule(context.Context, *ExecModuleRequest) (*ExecModuleReply, error)
}

func RegisterExecModuleServer(s *grpc.Server, srv ExecModuleServer) {
	s.RegisterService(&_ExecModule_serviceDesc, srv)
}

func _ExecModule_DoExecModule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecModuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecModuleServer).DoExecModule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openconfig.ExecModule/DoExecModule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecModuleServer).DoExecModule(ctx, req.(*ExecModuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ExecModule_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openconfig.ExecModule",
	HandlerType: (*ExecModuleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoExecModule",
			Handler:    _ExecModule_DoExecModule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "openconfig.proto",
}

// Client API for Config service

type ConfigClient interface {
	DoConfig(ctx context.Context, opts ...grpc.CallOption) (Config_DoConfigClient, error)
}

type configClient struct {
	cc *grpc.ClientConn
}

func NewConfigClient(cc *grpc.ClientConn) ConfigClient {
	return &configClient{cc}
}

func (c *configClient) DoConfig(ctx context.Context, opts ...grpc.CallOption) (Config_DoConfigClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Config_serviceDesc.Streams[0], c.cc, "/openconfig.Config/DoConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &configDoConfigClient{stream}
	return x, nil
}

type Config_DoConfigClient interface {
	Send(*ConfigRequest) error
	Recv() (*ConfigReply, error)
	grpc.ClientStream
}

type configDoConfigClient struct {
	grpc.ClientStream
}

func (x *configDoConfigClient) Send(m *ConfigRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *configDoConfigClient) Recv() (*ConfigReply, error) {
	m := new(ConfigReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Config service

type ConfigServer interface {
	DoConfig(Config_DoConfigServer) error
}

func RegisterConfigServer(s *grpc.Server, srv ConfigServer) {
	s.RegisterService(&_Config_serviceDesc, srv)
}

func _Config_DoConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConfigServer).DoConfig(&configDoConfigServer{stream})
}

type Config_DoConfigServer interface {
	Send(*ConfigReply) error
	Recv() (*ConfigRequest, error)
	grpc.ServerStream
}

type configDoConfigServer struct {
	grpc.ServerStream
}

func (x *configDoConfigServer) Send(m *ConfigReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *configDoConfigServer) Recv() (*ConfigRequest, error) {
	m := new(ConfigRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Config_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openconfig.Config",
	HandlerType: (*ConfigServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoConfig",
			Handler:       _Config_DoConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "openconfig.proto",
}

// Client API for Show service

type ShowClient interface {
	Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (Show_ShowClient, error)
}

type showClient struct {
	cc *grpc.ClientConn
}

func NewShowClient(cc *grpc.ClientConn) ShowClient {
	return &showClient{cc}
}

func (c *showClient) Show(ctx context.Context, in *ShowRequest, opts ...grpc.CallOption) (Show_ShowClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Show_serviceDesc.Streams[0], c.cc, "/openconfig.Show/Show", opts...)
	if err != nil {
		return nil, err
	}
	x := &showShowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Show_ShowClient interface {
	Recv() (*ShowReply, error)
	grpc.ClientStream
}

type showShowClient struct {
	grpc.ClientStream
}

func (x *showShowClient) Recv() (*ShowReply, error) {
	m := new(ShowReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Show service

type ShowServer interface {
	Show(*ShowRequest, Show_ShowServer) error
}

func RegisterShowServer(s *grpc.Server, srv ShowServer) {
	s.RegisterService(&_Show_serviceDesc, srv)
}

func _Show_Show_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShowRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ShowServer).Show(m, &showShowServer{stream})
}

type Show_ShowServer interface {
	Send(*ShowReply) error
	grpc.ServerStream
}

type showShowServer struct {
	grpc.ServerStream
}

func (x *showShowServer) Send(m *ShowReply) error {
	return x.ServerStream.SendMsg(m)
}

var _Show_serviceDesc = grpc.ServiceDesc{
	ServiceName: "openconfig.Show",
	HandlerType: (*ShowServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Show",
			Handler:       _Show_Show_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "openconfig.proto",
}

func init() { proto.RegisterFile("openconfig.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1025 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xcd, 0x72, 0xe3, 0x44,
	0x10, 0xb6, 0x6c, 0xf9, 0xaf, 0xbd, 0x4e, 0x26, 0xb3, 0xf9, 0x71, 0x9c, 0x5d, 0x08, 0x3a, 0x50,
	0x2e, 0x57, 0xed, 0xb2, 0x95, 0x5d, 0x2e, 0xa9, 0xe2, 0xa0, 0x48, 0x93, 0xac, 0x28, 0x4b, 0x5e,
	0x24, 0x19, 0x02, 0x17, 0x97, 0x7f, 0x44, 0x62, 0xb0, 0x25, 0x21, 0xc9, 0x80, 0x4f, 0xbc, 0x10,
	0x6f, 0xc1, 0x8b, 0x50, 0xc5, 0x95, 0x87, 0xa0, 0x66, 0xa4, 0xd1, 0x8f, 0xe3, 0x4d, 0x38, 0xa5,
	0xbb, 0xa7, 0xbb, 0xf5, 0x7d, 0x5f, 0xcf, 0x74, 0x0c, 0xc8, 0xf3, 0x1d, 0x77, 0xe6, 0xb9, 0x3f,
	0x2e, 0xee, 0x5e, 0xfb, 0x81, 0x17, 0x79, 0x18, 0xb2, 0x88, 0xf4, 0x8f, 0x00, 0x2d, 0xf2, 0xbb,
	0x33, 0x33, 0x9d, 0x5f, 0xd6, 0x4e, 0x18, 0xe1, 0x1e, 0x88, 0xd1, 0xc6, 0x77, 0x3a, 0xc2, 0xb9,
	0xd0, 0xdb, 0xbb, 0x38, 0x7c, 0x9d, 0x2b, 0xa6, 0x69, 0xf6, 0xc6, 0x77, 0x4c, 0x96, 0x81, 0x31,
	0x88, 0x2b, 0x6f, 0xee, 0x74, 0xca, 0xe7, 0x42, 0xaf, 0x69, 0x32, 0x1b, 0xbf, 0x80, 0xa6, 0x1f,
	0x2c, 0x7e, 0x5d, 0x2c, 0x9d, 0x3b, 0xa7, 0x53, 0x39, 0x17, 0x7a, 0x6d, 0x33, 0x0b, 0xd0, 0x8a,
	0xe5, 0xc2, 0x75, 0x3a, 0x62, 0x5c, 0x41, 0x6d, 0xdc, 0x85, 0xc6, 0xcc, 0x5b, 0xad, 0x26, 0xee,
	0x3c, 0xec, 0x54, 0xcf, 0x2b, 0xbd, 0xa6, 0x99, 0xfa, 0x34, 0x7f, 0x12, 0xdc, 0x85, 0x9d, 0x1a,
	0x8b, 0x33, 0x1b, 0xbf, 0x83, 0xfa, 0xcc, 0x73, 0x23, 0xc7, 0x8d, 0x3a, 0x75, 0x06, 0xb1, 0xbb,
	0x0d, 0x51, 0xf1, 0x5c, 0xc7, 0x8d, 0x18, 0x50, 0x9e, 0x2a, 0xfd, 0x01, 0xcd, 0x98, 0xa4, 0xbf,
	0xdc, 0x50, 0x8a, 0x33, 0x0a, 0xfc, 0x23, 0x14, 0x15, 0x6f, 0xee, 0x98, 0x2c, 0x03, 0x1f, 0x42,
	0x95, 0x82, 0x0c, 0x13, 0x8e, 0xb1, 0x43, 0x61, 0xf9, 0x5e, 0x10, 0x25, 0xfc, 0x98, 0x8d, 0x3f,
	0x01, 0x98, 0x4d, 0xdc, 0xf9, 0x62, 0x3e, 0x89, 0x9c, 0xb0, 0x23, 0x32, 0xc0, 0xb9, 0x88, 0xf4,
	0x97, 0x00, 0xfb, 0xa6, 0x73, 0xb7, 0x08, 0x23, 0x27, 0xe0, 0x52, 0x63, 0x10, 0xdd, 0xc9, 0x2a,
	0xc6, 0xd1, 0x34, 0x99, 0x8d, 0x8f, 0xa1, 0xb6, 0xf2, 0xe6, 0xeb, 0x25, 0x97, 0x35, 0xf1, 0x52,
	0xb1, 0x2b, 0x39, 0xb1, 0x77, 0xc9, 0x59, 0x18, 0x40, 0x75, 0x7b, 0x00, 0x87, 0x50, 0xbd, 0x77,
	0x96, 0x3e, 0x57, 0x34, 0x76, 0x52, 0x3d, 0xea, 0x4f, 0xe9, 0x21, 0x7d, 0x01, 0xed, 0x8c, 0x04,
	0x95, 0x92, 0xd1, 0x5e, 0x2e, 0xa7, 0x93, 0xd9, 0xcf, 0x8b, 0x39, 0x23, 0x52, 0x35, 0x73, 0x11,
	0x49, 0x81, 0x23, 0x5e, 0xa0, 0x33, 0x22, 0x9c, 0x7b, 0xc6, 0x53, 0xd8, 0xe6, 0xc9, 0xb4, 0x4d,
	0x2e, 0x15, 0xb5, 0xa5, 0x57, 0xf0, 0x7c, 0xbb, 0x09, 0xfd, 0xf6, 0x31, 0xd4, 0x02, 0x27, 0x5c,
	0x2f, 0xa3, 0xe4, 0xbb, 0x89, 0x27, 0xdd, 0xc0, 0x01, 0x85, 0x5d, 0xfc, 0xde, 0x13, 0x40, 0xd3,
	0xab, 0x56, 0xce, 0xae, 0x9a, 0xf4, 0x15, 0xec, 0xe7, 0x1b, 0x3d, 0xf2, 0xcd, 0x74, 0x14, 0xe5,
	0x6c, 0x14, 0xd2, 0x08, 0x90, 0xb5, 0x9e, 0x86, 0xb3, 0x60, 0x31, 0x4d, 0x61, 0xbc, 0x2a, 0xbc,
	0xae, 0xd3, 0xbc, 0xd4, 0x69, 0x6e, 0xf1, 0x89, 0xf9, 0x93, 0xe8, 0x3e, 0x55, 0x63, 0x12, 0xdd,
	0x4b, 0xff, 0x0a, 0xd0, 0x56, 0x58, 0x09, 0x6f, 0xda, 0x2f, 0x34, 0x3d, 0xce, 0x37, 0x8d, 0x13,
	0x73, 0x1d, 0x1f, 0xb9, 0x5f, 0x0f, 0xee, 0xf4, 0x5b, 0xa8, 0x87, 0xeb, 0x29, 0x6b, 0x2d, 0x3e,
	0x85, 0x97, 0x67, 0xa6, 0x90, 0xe3, 0xb7, 0xcc, 0x6c, 0x7c, 0x09, 0xcd, 0x90, 0x67, 0xb3, 0xab,
	0xd7, 0xba, 0x78, 0xb1, 0xb3, 0x55, 0xc2, 0xc8, 0xcc, 0xd2, 0xa5, 0x0d, 0xb4, 0x38, 0xdb, 0xc7,
	0x06, 0xc0, 0x35, 0x28, 0xff, 0x0f, 0x0d, 0x38, 0xc4, 0x4a, 0x0e, 0x22, 0x06, 0xf1, 0xa7, 0xd0,
	0x73, 0xf9, 0x5b, 0xa2, 0xb6, 0xf4, 0x25, 0xb4, 0xac, 0x7b, 0xef, 0xb7, 0xdc, 0x73, 0x65, 0x33,
	0x16, 0x72, 0xcf, 0x8d, 0x97, 0xd1, 0xcf, 0x36, 0x92, 0xb2, 0x97, 0xd0, 0x8c, 0xcb, 0x28, 0x5e,
	0x04, 0x95, 0x30, 0x0a, 0x92, 0x1a, 0x6a, 0xf6, 0x03, 0x68, 0xf0, 0x45, 0x8a, 0x1b, 0x20, 0x92,
	0x5b, 0xa2, 0xa0, 0x12, 0x7e, 0x06, 0x0d, 0x65, 0xa8, 0x7f, 0x18, 0x10, 0x9b, 0x20, 0x01, 0x9f,
	0xc1, 0x09, 0xf7, 0xc6, 0xb6, 0x29, 0x6b, 0x03, 0xcd, 0xb8, 0x19, 0x5b, 0x1f, 0x64, 0x85, 0xa0,
	0x72, 0xe1, 0xf0, 0x5a, 0x33, 0x2d, 0x7b, 0xac, 0x0c, 0x75, 0x5d, 0x36, 0x54, 0x0b, 0x55, 0xf0,
	0x21, 0xa0, 0xf4, 0x50, 0xfd, 0xde, 0x90, 0x75, 0x4d, 0x41, 0x62, 0xff, 0x1d, 0xec, 0x15, 0x37,
	0x23, 0x46, 0xf0, 0x4c, 0x19, 0x1a, 0x36, 0x31, 0xec, 0xb1, 0x4d, 0x6e, 0x6d, 0x54, 0xca, 0x47,
	0xbe, 0xb6, 0x86, 0x06, 0x12, 0xfa, 0x6e, 0x8c, 0x94, 0xbe, 0x7f, 0xdc, 0x82, 0xba, 0x35, 0x52,
	0x14, 0x62, 0x59, 0xa8, 0x44, 0x1d, 0x63, 0xa8, 0xcb, 0xb6, 0xf2, 0x1e, 0x09, 0x78, 0x0f, 0x40,
	0x33, 0x52, 0xec, 0x65, 0xdc, 0x86, 0xa6, 0xac, 0x5f, 0x69, 0x37, 0xa3, 0xe1, 0x88, 0x02, 0x6a,
	0x80, 0x68, 0xbd, 0x1f, 0x7e, 0x87, 0x44, 0x4a, 0xd1, 0x24, 0xaa, 0x66, 0x12, 0xc5, 0x46, 0x55,
	0x7c, 0x00, 0x6d, 0xee, 0x8d, 0x59, 0x42, 0xad, 0xff, 0xb7, 0x00, 0x90, 0x0d, 0x0b, 0xd7, 0xa1,
	0x62, 0x11, 0x8a, 0x0c, 0xa0, 0xa6, 0x92, 0x44, 0x19, 0x0c, 0x7b, 0xdf, 0xca, 0x03, 0x4d, 0x95,
	0x6d, 0x32, 0xb6, 0x6c, 0xd9, 0xb4, 0x51, 0x99, 0x22, 0x4f, 0x63, 0xc4, 0x50, 0x63, 0x15, 0xb2,
	0xac, 0x04, 0xb6, 0x88, 0x9f, 0xc3, 0x7e, 0x1a, 0xbd, 0x96, 0xb5, 0x01, 0x51, 0x51, 0x35, 0xa6,
	0xad, 0xeb, 0x9a, 0x9d, 0xb4, 0xab, 0x51, 0x42, 0x49, 0x84, 0x36, 0xab, 0x53, 0x42, 0xd6, 0xe8,
	0xca, 0x52, 0x4c, 0xed, 0x8a, 0xa0, 0x06, 0xed, 0x92, 0xba, 0x63, 0x7d, 0x34, 0xb0, 0x35, 0xd4,
	0xc4, 0x47, 0x70, 0x90, 0x05, 0x4d, 0xf2, 0xcd, 0x88, 0x58, 0x36, 0x02, 0xbc, 0x0f, 0x2d, 0xaa,
	0xe5, 0x58, 0x19, 0x1a, 0xd7, 0xda, 0x0d, 0x6a, 0xf5, 0x3f, 0x87, 0x76, 0xe1, 0xdd, 0x50, 0x29,
	0x93, 0xe9, 0xa1, 0x12, 0xd5, 0x2a, 0x96, 0xfe, 0xe2, 0x0a, 0x44, 0x2a, 0x3d, 0xbe, 0x84, 0x9a,
	0xea, 0x31, 0xeb, 0x64, 0x7b, 0x2d, 0x27, 0xd7, 0xb2, 0x7b, 0xf4, 0xf0, 0xc0, 0x5f, 0x6e, 0xa4,
	0xd2, 0xc5, 0x9f, 0x02, 0x34, 0xf8, 0xde, 0xc4, 0xd7, 0x00, 0xaa, 0x97, 0x7a, 0x67, 0xf9, 0x9a,
	0xad, 0x7f, 0x4b, 0xdd, 0xd3, 0xdd, 0x87, 0xac, 0x29, 0xbe, 0x05, 0x94, 0xf5, 0x89, 0x37, 0x23,
	0xfe, 0x6c, 0x57, 0x41, 0x61, 0xfd, 0x76, 0x3f, 0x7d, 0x2c, 0x25, 0x86, 0xfb, 0x03, 0x40, 0xb6,
	0x6d, 0xf1, 0x00, 0x9e, 0xc5, 0xc4, 0x13, 0xff, 0xe5, 0x36, 0xcb, 0x62, 0xff, 0xb3, 0x8f, 0x1d,
	0xc7, 0xbd, 0x0d, 0xa8, 0xc5, 0x17, 0x0b, 0xab, 0xd0, 0x50, 0xbd, 0xc4, 0x3e, 0x7d, 0xb8, 0x25,
	0x78, 0xbf, 0x93, 0x5d, 0x47, 0xac, 0x57, 0x4f, 0x78, 0xc3, 0xc6, 0x43, 0x9f, 0x38, 0xbe, 0x4c,
	0xfe, 0x16, 0xd2, 0x73, 0x3b, 0xa3, 0x38, 0x9c, 0x74, 0x2b, 0x48, 0xa5, 0x37, 0xc2, 0xb4, 0xc6,
	0x7e, 0x8b, 0xbd, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x66, 0x51, 0x0c, 0x7a, 0x9f, 0x09, 0x00,
	0x00,
}
