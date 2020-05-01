// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agentd.proto

package telemetry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TelemetrySystem struct {
	Subscriptions        *TelemetrySystemSubscriptionsType `protobuf:"bytes,151,opt,name=subscriptions" json:"subscriptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TelemetrySystem) Reset()         { *m = TelemetrySystem{} }
func (m *TelemetrySystem) String() string { return proto.CompactTextString(m) }
func (*TelemetrySystem) ProtoMessage()    {}
func (*TelemetrySystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0}
}

func (m *TelemetrySystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystem.Unmarshal(m, b)
}
func (m *TelemetrySystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystem.Marshal(b, m, deterministic)
}
func (m *TelemetrySystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystem.Merge(m, src)
}
func (m *TelemetrySystem) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystem.Size(m)
}
func (m *TelemetrySystem) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystem.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystem proto.InternalMessageInfo

func (m *TelemetrySystem) GetSubscriptions() *TelemetrySystemSubscriptionsType {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

type TelemetrySystemSubscriptionsType struct {
	Dynamic              *TelemetrySystemSubscriptionsTypeDynamicType `protobuf:"bytes,151,opt,name=dynamic" json:"dynamic,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                     `json:"-"`
	XXX_unrecognized     []byte                                       `json:"-"`
	XXX_sizecache        int32                                        `json:"-"`
}

func (m *TelemetrySystemSubscriptionsType) Reset()         { *m = TelemetrySystemSubscriptionsType{} }
func (m *TelemetrySystemSubscriptionsType) String() string { return proto.CompactTextString(m) }
func (*TelemetrySystemSubscriptionsType) ProtoMessage()    {}
func (*TelemetrySystemSubscriptionsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0}
}

func (m *TelemetrySystemSubscriptionsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsType.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsType.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsType.Size(m)
}
func (m *TelemetrySystemSubscriptionsType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsType) GetDynamic() *TelemetrySystemSubscriptionsTypeDynamicType {
	if m != nil {
		return m.Dynamic
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicType struct {
	Subscription         []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList `protobuf:"bytes,151,rep,name=subscription" json:"subscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                       `json:"-"`
	XXX_unrecognized     []byte                                                         `json:"-"`
	XXX_sizecache        int32                                                          `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicType) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0, 0}
}

func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicType) GetSubscription() []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList {
	if m != nil {
		return m.Subscription
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList struct {
	SubscriptionId       *uint64                                                                     `protobuf:"varint,51,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	State                *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType       `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	SensorPaths          *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType `protobuf:"bytes,152,opt,name=sensor_paths,json=sensorPaths" json:"sensor_paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                    `json:"-"`
	XXX_unrecognized     []byte                                                                      `json:"-"`
	XXX_sizecache        int32                                                                       `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0, 0, 0}
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) GetSubscriptionId() uint64 {
	if m != nil && m.SubscriptionId != nil {
		return *m.SubscriptionId
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) GetState() *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList) GetSensorPaths() *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType {
	if m != nil {
		return m.SensorPaths
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType struct {
	SubscriptionId       *uint64  `protobuf:"varint,51,opt,name=subscription_id,json=subscriptionId" json:"subscription_id,omitempty"`
	SampleInterval       *uint64  `protobuf:"varint,52,opt,name=sample_interval,json=sampleInterval" json:"sample_interval,omitempty"`
	HeartbeatInterval    *uint64  `protobuf:"varint,53,opt,name=heartbeat_interval,json=heartbeatInterval" json:"heartbeat_interval,omitempty"`
	SuppressRedundant    *bool    `protobuf:"varint,54,opt,name=suppress_redundant,json=suppressRedundant" json:"suppress_redundant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0, 0, 0, 0}
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetSubscriptionId() uint64 {
	if m != nil && m.SubscriptionId != nil {
		return *m.SubscriptionId
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetSampleInterval() uint64 {
	if m != nil && m.SampleInterval != nil {
		return *m.SampleInterval
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetHeartbeatInterval() uint64 {
	if m != nil && m.HeartbeatInterval != nil {
		return *m.HeartbeatInterval
	}
	return 0
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType) GetSuppressRedundant() bool {
	if m != nil && m.SuppressRedundant != nil {
		return *m.SuppressRedundant
	}
	return false
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType struct {
	SensorPath           []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList `protobuf:"bytes,151,rep,name=sensor_path,json=sensorPath" json:"sensor_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                    `json:"-"`
	XXX_unrecognized     []byte                                                                                      `json:"-"`
	XXX_sizecache        int32                                                                                       `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) ProtoMessage() {}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0, 0, 0, 1}
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType) GetSensorPath() []*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList {
	if m != nil {
		return m.SensorPath
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList struct {
	Path                 *string                                                                                            `protobuf:"bytes,51,opt,name=path" json:"path,omitempty"`
	State                *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType `protobuf:"bytes,151,opt,name=state" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                                           `json:"-"`
	XXX_unrecognized     []byte                                                                                             `json:"-"`
	XXX_sizecache        int32                                                                                              `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) ProtoMessage() {
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0, 0, 0, 1, 0}
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList) GetState() *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType {
	if m != nil {
		return m.State
	}
	return nil
}

type TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType struct {
	Path                 *string  `protobuf:"bytes,51,opt,name=path" json:"path,omitempty"`
	ExcludeFilter        *string  `protobuf:"bytes,52,opt,name=exclude_filter,json=excludeFilter" json:"exclude_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) Reset() {
	*m = TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType{}
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) String() string {
	return proto.CompactTextString(m)
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) ProtoMessage() {
}
func (*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f929025cc6d20165, []int{0, 0, 0, 0, 1, 0, 0}
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Unmarshal(m, b)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Marshal(b, m, deterministic)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Merge(m, src)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_Size() int {
	return xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.Size(m)
}
func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) XXX_DiscardUnknown() {
	xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType.DiscardUnknown(m)
}

var xxx_messageInfo_TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType proto.InternalMessageInfo

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType) GetExcludeFilter() string {
	if m != nil && m.ExcludeFilter != nil {
		return *m.ExcludeFilter
	}
	return ""
}

var E_JnprTelemetrySystemExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*TelemetrySystem)(nil),
	Field:         31,
	Name:          "jnpr_telemetry_system_ext",
	Tag:           "bytes,31,opt,name=jnpr_telemetry_system_ext",
	Filename:      "agentd.proto",
}

func init() {
	proto.RegisterType((*TelemetrySystem)(nil), "telemetry_system")
	proto.RegisterType((*TelemetrySystemSubscriptionsType)(nil), "telemetry_system.subscriptions_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicType)(nil), "telemetry_system.subscriptions_type.dynamic_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionList)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListStateType)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.state_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsType)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.sensor_paths_type")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathList)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.sensor_paths_type.sensor_path_list")
	proto.RegisterType((*TelemetrySystemSubscriptionsTypeDynamicTypeSubscriptionListSensorPathsTypeSensorPathListStateType)(nil), "telemetry_system.subscriptions_type.dynamic_type.subscription_list.sensor_paths_type.sensor_path_list.state_type")
	proto.RegisterExtension(E_JnprTelemetrySystemExt)
}

func init() { proto.RegisterFile("agentd.proto", fileDescriptor_f929025cc6d20165) }

var fileDescriptor_f929025cc6d20165 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0x8a, 0xd4, 0x40,
	0x14, 0x85, 0xc9, 0xfc, 0x30, 0xdd, 0xb7, 0x7b, 0xc6, 0x49, 0x09, 0x9a, 0xc9, 0xc6, 0x46, 0x14,
	0x7b, 0x63, 0xc0, 0xf1, 0x67, 0xe1, 0x4a, 0x14, 0x95, 0x51, 0x19, 0x24, 0xa3, 0xeb, 0xa2, 0xa6,
	0x73, 0xc7, 0x8e, 0xa6, 0x2b, 0x65, 0xd5, 0x8d, 0x76, 0x96, 0x8a, 0x0b, 0x7d, 0x02, 0x5d, 0xfb,
	0x1c, 0xbe, 0x81, 0x4f, 0xe2, 0x5b, 0x48, 0x2a, 0x89, 0x49, 0x3a, 0x08, 0x0a, 0xa3, 0xdb, 0xef,
	0x9e, 0x3a, 0xf7, 0xd4, 0xe5, 0xc0, 0x58, 0xbc, 0x40, 0x49, 0x51, 0xa0, 0x74, 0x4a, 0xa9, 0x7f,
	0x96, 0x30, 0xc1, 0x05, 0x92, 0xce, 0x39, 0xa5, 0xaa, 0x84, 0x17, 0xbf, 0x0f, 0x61, 0xb7, 0xe1,
	0x26, 0x37, 0x84, 0x0b, 0xf6, 0x18, 0xb6, 0x4d, 0x76, 0x6c, 0x66, 0x3a, 0x56, 0x14, 0xa7, 0xd2,
	0x78, 0x9f, 0x9d, 0x89, 0x33, 0x1d, 0xed, 0x5f, 0x0a, 0x56, 0xa5, 0x41, 0x47, 0xc7, 0x29, 0x57,
	0x18, 0x76, 0xdf, 0xfa, 0x3f, 0x06, 0xc0, 0xfa, 0x2a, 0xf6, 0x04, 0xb6, 0xa2, 0x5c, 0x8a, 0x45,
	0x3c, 0xab, 0xdd, 0xaf, 0xfd, 0x89, 0x7b, 0x50, 0x3d, 0x2a, 0x57, 0xd5, 0x16, 0xfe, 0x87, 0x01,
	0x8c, 0xdb, 0x13, 0x36, 0x87, 0x71, 0xfb, 0x71, 0xb1, 0x63, 0x7d, 0x3a, 0xda, 0xbf, 0xf7, 0xd7,
	0x3b, 0x3a, 0x73, 0x9e, 0xc4, 0x86, 0xc2, 0x8e, 0xb3, 0xff, 0x69, 0x0b, 0xdc, 0x9e, 0x86, 0x05,
	0x70, 0xa6, 0x03, 0xe3, 0xc8, 0xbb, 0x3e, 0x71, 0xa6, 0x1b, 0x77, 0x37, 0xdf, 0xdf, 0x59, 0x1b,
	0x38, 0xe1, 0x4e, 0x7b, 0x7a, 0x10, 0x31, 0x84, 0x4d, 0x43, 0x82, 0xb0, 0x3e, 0xc6, 0xe1, 0x29,
	0x04, 0x0d, 0xac, 0x63, 0x79, 0xa9, 0xd2, 0x9d, 0xe5, 0x30, 0x36, 0x28, 0x4d, 0xaa, 0xb9, 0x12,
	0x34, 0x37, 0xde, 0x97, 0x72, 0xdb, 0xf3, 0x53, 0xd9, 0xd6, 0x32, 0x2e, 0x97, 0x8e, 0x4a, 0xf4,
	0xb4, 0x20, 0xfe, 0x37, 0x07, 0xa0, 0x09, 0xc4, 0xae, 0xfc, 0xe6, 0x40, 0xbd, 0xcb, 0x14, 0x42,
	0xb1, 0x50, 0x09, 0xf2, 0x58, 0x12, 0xea, 0x37, 0x22, 0xf1, 0x6e, 0x54, 0x42, 0x8b, 0x0f, 0x2a,
	0xca, 0xae, 0x02, 0x9b, 0xa3, 0xd0, 0x74, 0x8c, 0x82, 0x1a, 0xed, 0x4d, 0xab, 0x75, 0x7f, 0x4d,
	0xda, 0x72, 0x93, 0x29, 0xa5, 0xd1, 0x18, 0xae, 0x31, 0xca, 0x64, 0x24, 0x24, 0x79, 0xb7, 0x26,
	0xce, 0x74, 0x10, 0xba, 0xf5, 0x24, 0xac, 0x07, 0xfe, 0xd7, 0x75, 0x70, 0x7b, 0x3f, 0x64, 0x1f,
	0x1d, 0x18, 0xb5, 0x68, 0x5d, 0xb3, 0x93, 0x7f, 0x72, 0xcf, 0x36, 0x29, 0x9b, 0x08, 0xcd, 0x81,
	0xfd, 0x77, 0x6b, 0xb0, 0xbb, 0x2a, 0x60, 0x7b, 0xb0, 0x61, 0x73, 0x15, 0xa7, 0x1d, 0xd6, 0xdd,
	0xb3, 0xa8, 0x88, 0xde, 0xad, 0xdc, 0xeb, 0xff, 0x13, 0xba, 0xdf, 0x4a, 0xff, 0x61, 0xa7, 0x19,
	0xac, 0x9d, 0xb9, 0x0a, 0x7b, 0x19, 0x76, 0x70, 0x39, 0x4b, 0xb2, 0x08, 0xf9, 0x49, 0x9c, 0x10,
	0x6a, 0xdb, 0x81, 0x61, 0xb8, 0x5d, 0xd1, 0x07, 0x16, 0xde, 0x46, 0xd8, 0x7b, 0x29, 0x95, 0xe6,
	0xab, 0x1f, 0xe1, 0xb8, 0x24, 0x76, 0x3e, 0x78, 0x94, 0xc9, 0x58, 0xa1, 0x3e, 0x44, 0x7a, 0x9b,
	0xea, 0x57, 0xe6, 0xc8, 0x26, 0x33, 0xde, 0x05, 0xfb, 0x7d, 0xb7, 0xf7, 0xfd, 0xf0, 0x5c, 0x61,
	0xf6, 0xac, 0xa6, 0x47, 0x16, 0xde, 0x5f, 0xd2, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x8b,
	0x5b, 0x44, 0x58, 0x05, 0x00, 0x00,
}
