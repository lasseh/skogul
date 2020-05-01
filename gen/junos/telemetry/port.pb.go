// Code generated by protoc-gen-go. DO NOT EDIT.
// source: port.proto

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

//
// Top-level message
//
type Port struct {
	InterfaceStats       []*InterfaceInfos `protobuf:"bytes,1,rep,name=interface_stats,json=interfaceStats" json:"interface_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{0}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetInterfaceStats() []*InterfaceInfos {
	if m != nil {
		return m.InterfaceStats
	}
	return nil
}

//
// Interface information
//
type InterfaceInfos struct {
	// Interface name, e.g., xe-0/0/0
	IfName *string `protobuf:"bytes,1,req,name=if_name,json=ifName" json:"if_name,omitempty"`
	// Time when interface is created
	InitTime *uint64 `protobuf:"varint,2,opt,name=init_time,json=initTime" json:"init_time,omitempty"`
	// Global Index
	SnmpIfIndex *uint32 `protobuf:"varint,3,opt,name=snmp_if_index,json=snmpIfIndex" json:"snmp_if_index,omitempty"`
	// Name of parent for AE interface, if applicable
	ParentAeName *string `protobuf:"bytes,4,opt,name=parent_ae_name,json=parentAeName" json:"parent_ae_name,omitempty"`
	// Egress queue information
	EgressQueueInfo []*QueueStats `protobuf:"bytes,5,rep,name=egress_queue_info,json=egressQueueInfo" json:"egress_queue_info,omitempty"`
	// Ingress queue information
	IngressQueueInfo []*QueueStats `protobuf:"bytes,6,rep,name=ingress_queue_info,json=ingressQueueInfo" json:"ingress_queue_info,omitempty"`
	// Inbound traffic statistics
	IngressStats *InterfaceStats `protobuf:"bytes,7,opt,name=ingress_stats,json=ingressStats" json:"ingress_stats,omitempty"`
	// Outbound traffic statistics
	EgressStats *InterfaceStats `protobuf:"bytes,8,opt,name=egress_stats,json=egressStats" json:"egress_stats,omitempty"`
	// Inbound traffic errors
	IngressErrors *IngressInterfaceErrors `protobuf:"bytes,9,opt,name=ingress_errors,json=ingressErrors" json:"ingress_errors,omitempty"`
	// Interface administration status
	IfAdministrationStatus *string `protobuf:"bytes,10,opt,name=if_administration_status,json=ifAdministrationStatus" json:"if_administration_status,omitempty"`
	// Interface operational status
	IfOperationalStatus *string `protobuf:"bytes,11,opt,name=if_operational_status,json=ifOperationalStatus" json:"if_operational_status,omitempty"`
	// Interface description
	IfDescription *string `protobuf:"bytes,12,opt,name=if_description,json=ifDescription" json:"if_description,omitempty"`
	// Counter: number of carrier transitions on this interface
	IfTransitions *uint64 `protobuf:"varint,13,opt,name=if_transitions,json=ifTransitions" json:"if_transitions,omitempty"`
	// This corresponds to the ifLastChange object in the standard interface MIB
	IfLastChange *uint32 `protobuf:"varint,14,opt,name=ifLastChange" json:"ifLastChange,omitempty"`
	// This corresponds to the ifHighSpeed object in the standard interface MIB
	IfHighSpeed *uint32 `protobuf:"varint,15,opt,name=ifHighSpeed" json:"ifHighSpeed,omitempty"`
	// Outbound traffic errors
	EgressErrors         *EgressInterfaceErrors `protobuf:"bytes,16,opt,name=egress_errors,json=egressErrors" json:"egress_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *InterfaceInfos) Reset()         { *m = InterfaceInfos{} }
func (m *InterfaceInfos) String() string { return proto.CompactTextString(m) }
func (*InterfaceInfos) ProtoMessage()    {}
func (*InterfaceInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{1}
}

func (m *InterfaceInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceInfos.Unmarshal(m, b)
}
func (m *InterfaceInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceInfos.Marshal(b, m, deterministic)
}
func (m *InterfaceInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceInfos.Merge(m, src)
}
func (m *InterfaceInfos) XXX_Size() int {
	return xxx_messageInfo_InterfaceInfos.Size(m)
}
func (m *InterfaceInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceInfos.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceInfos proto.InternalMessageInfo

func (m *InterfaceInfos) GetIfName() string {
	if m != nil && m.IfName != nil {
		return *m.IfName
	}
	return ""
}

func (m *InterfaceInfos) GetInitTime() uint64 {
	if m != nil && m.InitTime != nil {
		return *m.InitTime
	}
	return 0
}

func (m *InterfaceInfos) GetSnmpIfIndex() uint32 {
	if m != nil && m.SnmpIfIndex != nil {
		return *m.SnmpIfIndex
	}
	return 0
}

func (m *InterfaceInfos) GetParentAeName() string {
	if m != nil && m.ParentAeName != nil {
		return *m.ParentAeName
	}
	return ""
}

func (m *InterfaceInfos) GetEgressQueueInfo() []*QueueStats {
	if m != nil {
		return m.EgressQueueInfo
	}
	return nil
}

func (m *InterfaceInfos) GetIngressQueueInfo() []*QueueStats {
	if m != nil {
		return m.IngressQueueInfo
	}
	return nil
}

func (m *InterfaceInfos) GetIngressStats() *InterfaceStats {
	if m != nil {
		return m.IngressStats
	}
	return nil
}

func (m *InterfaceInfos) GetEgressStats() *InterfaceStats {
	if m != nil {
		return m.EgressStats
	}
	return nil
}

func (m *InterfaceInfos) GetIngressErrors() *IngressInterfaceErrors {
	if m != nil {
		return m.IngressErrors
	}
	return nil
}

func (m *InterfaceInfos) GetIfAdministrationStatus() string {
	if m != nil && m.IfAdministrationStatus != nil {
		return *m.IfAdministrationStatus
	}
	return ""
}

func (m *InterfaceInfos) GetIfOperationalStatus() string {
	if m != nil && m.IfOperationalStatus != nil {
		return *m.IfOperationalStatus
	}
	return ""
}

func (m *InterfaceInfos) GetIfDescription() string {
	if m != nil && m.IfDescription != nil {
		return *m.IfDescription
	}
	return ""
}

func (m *InterfaceInfos) GetIfTransitions() uint64 {
	if m != nil && m.IfTransitions != nil {
		return *m.IfTransitions
	}
	return 0
}

func (m *InterfaceInfos) GetIfLastChange() uint32 {
	if m != nil && m.IfLastChange != nil {
		return *m.IfLastChange
	}
	return 0
}

func (m *InterfaceInfos) GetIfHighSpeed() uint32 {
	if m != nil && m.IfHighSpeed != nil {
		return *m.IfHighSpeed
	}
	return 0
}

func (m *InterfaceInfos) GetEgressErrors() *EgressInterfaceErrors {
	if m != nil {
		return m.EgressErrors
	}
	return nil
}

//
// Interface queue statistics
//
type QueueStats struct {
	// Queue number
	QueueNumber *uint32 `protobuf:"varint,1,opt,name=queue_number,json=queueNumber" json:"queue_number,omitempty"`
	// The total number of packets that have been added to this queue
	Packets *uint64 `protobuf:"varint,2,opt,name=packets" json:"packets,omitempty"`
	// The total number of bytes that have been added to this queue
	Bytes *uint64 `protobuf:"varint,3,opt,name=bytes" json:"bytes,omitempty"`
	// The total number of tail dropped packets
	TailDropPackets *uint64 `protobuf:"varint,4,opt,name=tail_drop_packets,json=tailDropPackets" json:"tail_drop_packets,omitempty"`
	// The total number of rate-limited packets
	RlDropPackets *uint64 `protobuf:"varint,5,opt,name=rl_drop_packets,json=rlDropPackets" json:"rl_drop_packets,omitempty"`
	// The total number of rate-limited bytes
	RlDropBytes *uint64 `protobuf:"varint,6,opt,name=rl_drop_bytes,json=rlDropBytes" json:"rl_drop_bytes,omitempty"`
	// The total number of red-dropped packets
	RedDropPackets *uint64 `protobuf:"varint,7,opt,name=red_drop_packets,json=redDropPackets" json:"red_drop_packets,omitempty"`
	// The total number of red-dropped bytes
	RedDropBytes *uint64 `protobuf:"varint,8,opt,name=red_drop_bytes,json=redDropBytes" json:"red_drop_bytes,omitempty"`
	// Average queue depth, in packets
	AvgBufferOccupancy *uint64 `protobuf:"varint,9,opt,name=avg_buffer_occupancy,json=avgBufferOccupancy" json:"avg_buffer_occupancy,omitempty"`
	// Current queue depth, in packets
	CurBufferOccupancy *uint64 `protobuf:"varint,10,opt,name=cur_buffer_occupancy,json=curBufferOccupancy" json:"cur_buffer_occupancy,omitempty"`
	// The max measured queue depth, in packets, across all measurements since boot
	PeakBufferOccupancy *uint64 `protobuf:"varint,11,opt,name=peak_buffer_occupancy,json=peakBufferOccupancy" json:"peak_buffer_occupancy,omitempty"`
	// Allocated buffer size
	AllocatedBufferSize  *uint64  `protobuf:"varint,12,opt,name=allocated_buffer_size,json=allocatedBufferSize" json:"allocated_buffer_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueueStats) Reset()         { *m = QueueStats{} }
func (m *QueueStats) String() string { return proto.CompactTextString(m) }
func (*QueueStats) ProtoMessage()    {}
func (*QueueStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{2}
}

func (m *QueueStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueueStats.Unmarshal(m, b)
}
func (m *QueueStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueueStats.Marshal(b, m, deterministic)
}
func (m *QueueStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueueStats.Merge(m, src)
}
func (m *QueueStats) XXX_Size() int {
	return xxx_messageInfo_QueueStats.Size(m)
}
func (m *QueueStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueueStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueueStats proto.InternalMessageInfo

func (m *QueueStats) GetQueueNumber() uint32 {
	if m != nil && m.QueueNumber != nil {
		return *m.QueueNumber
	}
	return 0
}

func (m *QueueStats) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *QueueStats) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *QueueStats) GetTailDropPackets() uint64 {
	if m != nil && m.TailDropPackets != nil {
		return *m.TailDropPackets
	}
	return 0
}

func (m *QueueStats) GetRlDropPackets() uint64 {
	if m != nil && m.RlDropPackets != nil {
		return *m.RlDropPackets
	}
	return 0
}

func (m *QueueStats) GetRlDropBytes() uint64 {
	if m != nil && m.RlDropBytes != nil {
		return *m.RlDropBytes
	}
	return 0
}

func (m *QueueStats) GetRedDropPackets() uint64 {
	if m != nil && m.RedDropPackets != nil {
		return *m.RedDropPackets
	}
	return 0
}

func (m *QueueStats) GetRedDropBytes() uint64 {
	if m != nil && m.RedDropBytes != nil {
		return *m.RedDropBytes
	}
	return 0
}

func (m *QueueStats) GetAvgBufferOccupancy() uint64 {
	if m != nil && m.AvgBufferOccupancy != nil {
		return *m.AvgBufferOccupancy
	}
	return 0
}

func (m *QueueStats) GetCurBufferOccupancy() uint64 {
	if m != nil && m.CurBufferOccupancy != nil {
		return *m.CurBufferOccupancy
	}
	return 0
}

func (m *QueueStats) GetPeakBufferOccupancy() uint64 {
	if m != nil && m.PeakBufferOccupancy != nil {
		return *m.PeakBufferOccupancy
	}
	return 0
}

func (m *QueueStats) GetAllocatedBufferSize() uint64 {
	if m != nil && m.AllocatedBufferSize != nil {
		return *m.AllocatedBufferSize
	}
	return 0
}

//
// Interface statistics
//
type InterfaceStats struct {
	// The total number of packets sent/received by this interface
	IfPkts *uint64 `protobuf:"varint,1,opt,name=if_pkts,json=ifPkts" json:"if_pkts,omitempty"`
	// The total number of bytes sent/received by this interface
	IfOctets *uint64 `protobuf:"varint,2,opt,name=if_octets,json=ifOctets" json:"if_octets,omitempty"`
	// The rate at which packets are sent/received by this interface (in packets/sec)
	If_1SecPkts *uint64 `protobuf:"varint,3,opt,name=if_1sec_pkts,json=if1secPkts" json:"if_1sec_pkts,omitempty"`
	// The rate at which bytes are sent/received by this interface
	If_1SecOctets *uint64 `protobuf:"varint,4,opt,name=if_1sec_octets,json=if1secOctets" json:"if_1sec_octets,omitempty"`
	// Total number of unicast packets sent/received by this interface
	IfUcPkts *uint64 `protobuf:"varint,5,opt,name=if_uc_pkts,json=ifUcPkts" json:"if_uc_pkts,omitempty"`
	// Total number of multicast packets sent/received by this interface
	IfMcPkts *uint64 `protobuf:"varint,6,opt,name=if_mc_pkts,json=ifMcPkts" json:"if_mc_pkts,omitempty"`
	// Total number of broadcast packets sent/received by this interface
	IfBcPkts *uint64 `protobuf:"varint,7,opt,name=if_bc_pkts,json=ifBcPkts" json:"if_bc_pkts,omitempty"`
	// Counter: total no of error packets sent/rcvd by this interface
	IfError *uint64 `protobuf:"varint,8,opt,name=if_error,json=ifError" json:"if_error,omitempty"`
	// Counter: total no of PAUSE packets sent/rcvd by this interface
	IfPausePkts *uint64 `protobuf:"varint,9,opt,name=if_pause_pkts,json=ifPausePkts" json:"if_pause_pkts,omitempty"`
	// Counter: total no of UNKNOWN proto packets sent/rcvd by this interface
	IfUnknownProtoPkts   *uint64  `protobuf:"varint,10,opt,name=if_unknown_proto_pkts,json=ifUnknownProtoPkts" json:"if_unknown_proto_pkts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceStats) Reset()         { *m = InterfaceStats{} }
func (m *InterfaceStats) String() string { return proto.CompactTextString(m) }
func (*InterfaceStats) ProtoMessage()    {}
func (*InterfaceStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{3}
}

func (m *InterfaceStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceStats.Unmarshal(m, b)
}
func (m *InterfaceStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceStats.Marshal(b, m, deterministic)
}
func (m *InterfaceStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceStats.Merge(m, src)
}
func (m *InterfaceStats) XXX_Size() int {
	return xxx_messageInfo_InterfaceStats.Size(m)
}
func (m *InterfaceStats) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceStats.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceStats proto.InternalMessageInfo

func (m *InterfaceStats) GetIfPkts() uint64 {
	if m != nil && m.IfPkts != nil {
		return *m.IfPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfOctets() uint64 {
	if m != nil && m.IfOctets != nil {
		return *m.IfOctets
	}
	return 0
}

func (m *InterfaceStats) GetIf_1SecPkts() uint64 {
	if m != nil && m.If_1SecPkts != nil {
		return *m.If_1SecPkts
	}
	return 0
}

func (m *InterfaceStats) GetIf_1SecOctets() uint64 {
	if m != nil && m.If_1SecOctets != nil {
		return *m.If_1SecOctets
	}
	return 0
}

func (m *InterfaceStats) GetIfUcPkts() uint64 {
	if m != nil && m.IfUcPkts != nil {
		return *m.IfUcPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfMcPkts() uint64 {
	if m != nil && m.IfMcPkts != nil {
		return *m.IfMcPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfBcPkts() uint64 {
	if m != nil && m.IfBcPkts != nil {
		return *m.IfBcPkts
	}
	return 0
}

func (m *InterfaceStats) GetIfError() uint64 {
	if m != nil && m.IfError != nil {
		return *m.IfError
	}
	return 0
}

func (m *InterfaceStats) GetIfPausePkts() uint64 {
	if m != nil && m.IfPausePkts != nil {
		return *m.IfPausePkts
	}
	return 0
}

func (m *InterfaceStats) GetIfUnknownProtoPkts() uint64 {
	if m != nil && m.IfUnknownProtoPkts != nil {
		return *m.IfUnknownProtoPkts
	}
	return 0
}

//
// Inbound traffic error statistics
//
type IngressInterfaceErrors struct {
	// The number of packets that contained errors
	IfErrors *uint64 `protobuf:"varint,1,opt,name=if_errors,json=ifErrors" json:"if_errors,omitempty"`
	// The number of packets dropped by the input queue of the I/O Manager ASIC
	IfInQdrops *uint64 `protobuf:"varint,2,opt,name=if_in_qdrops,json=ifInQdrops" json:"if_in_qdrops,omitempty"`
	// The number of packets which were misaligned
	IfInFrameErrors *uint64 `protobuf:"varint,3,opt,name=if_in_frame_errors,json=ifInFrameErrors" json:"if_in_frame_errors,omitempty"`
	// The number of non-error packets which were chosen to be discarded
	IfDiscards *uint64 `protobuf:"varint,4,opt,name=if_discards,json=ifDiscards" json:"if_discards,omitempty"`
	// The number of runt packets
	IfInRunts *uint64 `protobuf:"varint,5,opt,name=if_in_runts,json=ifInRunts" json:"if_in_runts,omitempty"`
	// The number of packets that fail Layer 3 sanity checks of the header
	IfInL3Incompletes *uint64 `protobuf:"varint,6,opt,name=if_in_l3_incompletes,json=ifInL3Incompletes" json:"if_in_l3_incompletes,omitempty"`
	// The number of packets for which the software could not find a valid logical interface
	IfInL2ChanErrors *uint64 `protobuf:"varint,7,opt,name=if_in_l2chan_errors,json=ifInL2chanErrors" json:"if_in_l2chan_errors,omitempty"`
	// The number of malform or short packets
	IfInL2MismatchTimeouts *uint64 `protobuf:"varint,8,opt,name=if_in_l2_mismatch_timeouts,json=ifInL2MismatchTimeouts" json:"if_in_l2_mismatch_timeouts,omitempty"`
	// The number of FIFO errors
	IfInFifoErrors *uint64 `protobuf:"varint,9,opt,name=if_in_fifo_errors,json=ifInFifoErrors" json:"if_in_fifo_errors,omitempty"`
	// The number of resource errors
	IfInResourceErrors   *uint64  `protobuf:"varint,10,opt,name=if_in_resource_errors,json=ifInResourceErrors" json:"if_in_resource_errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngressInterfaceErrors) Reset()         { *m = IngressInterfaceErrors{} }
func (m *IngressInterfaceErrors) String() string { return proto.CompactTextString(m) }
func (*IngressInterfaceErrors) ProtoMessage()    {}
func (*IngressInterfaceErrors) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{4}
}

func (m *IngressInterfaceErrors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressInterfaceErrors.Unmarshal(m, b)
}
func (m *IngressInterfaceErrors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressInterfaceErrors.Marshal(b, m, deterministic)
}
func (m *IngressInterfaceErrors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressInterfaceErrors.Merge(m, src)
}
func (m *IngressInterfaceErrors) XXX_Size() int {
	return xxx_messageInfo_IngressInterfaceErrors.Size(m)
}
func (m *IngressInterfaceErrors) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressInterfaceErrors.DiscardUnknown(m)
}

var xxx_messageInfo_IngressInterfaceErrors proto.InternalMessageInfo

func (m *IngressInterfaceErrors) GetIfErrors() uint64 {
	if m != nil && m.IfErrors != nil {
		return *m.IfErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInQdrops() uint64 {
	if m != nil && m.IfInQdrops != nil {
		return *m.IfInQdrops
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInFrameErrors() uint64 {
	if m != nil && m.IfInFrameErrors != nil {
		return *m.IfInFrameErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfDiscards() uint64 {
	if m != nil && m.IfDiscards != nil {
		return *m.IfDiscards
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInRunts() uint64 {
	if m != nil && m.IfInRunts != nil {
		return *m.IfInRunts
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInL3Incompletes() uint64 {
	if m != nil && m.IfInL3Incompletes != nil {
		return *m.IfInL3Incompletes
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInL2ChanErrors() uint64 {
	if m != nil && m.IfInL2ChanErrors != nil {
		return *m.IfInL2ChanErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInL2MismatchTimeouts() uint64 {
	if m != nil && m.IfInL2MismatchTimeouts != nil {
		return *m.IfInL2MismatchTimeouts
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInFifoErrors() uint64 {
	if m != nil && m.IfInFifoErrors != nil {
		return *m.IfInFifoErrors
	}
	return 0
}

func (m *IngressInterfaceErrors) GetIfInResourceErrors() uint64 {
	if m != nil && m.IfInResourceErrors != nil {
		return *m.IfInResourceErrors
	}
	return 0
}

//
// Outbound traffic error statistics
//
type EgressInterfaceErrors struct {
	// The number of packets that contained errors
	IfErrors *uint64 `protobuf:"varint,1,opt,name=if_errors,json=ifErrors" json:"if_errors,omitempty"`
	// The number of non-error packets which were chosen to be discarded
	IfDiscards           *uint64  `protobuf:"varint,2,opt,name=if_discards,json=ifDiscards" json:"if_discards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EgressInterfaceErrors) Reset()         { *m = EgressInterfaceErrors{} }
func (m *EgressInterfaceErrors) String() string { return proto.CompactTextString(m) }
func (*EgressInterfaceErrors) ProtoMessage()    {}
func (*EgressInterfaceErrors) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{5}
}

func (m *EgressInterfaceErrors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EgressInterfaceErrors.Unmarshal(m, b)
}
func (m *EgressInterfaceErrors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EgressInterfaceErrors.Marshal(b, m, deterministic)
}
func (m *EgressInterfaceErrors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EgressInterfaceErrors.Merge(m, src)
}
func (m *EgressInterfaceErrors) XXX_Size() int {
	return xxx_messageInfo_EgressInterfaceErrors.Size(m)
}
func (m *EgressInterfaceErrors) XXX_DiscardUnknown() {
	xxx_messageInfo_EgressInterfaceErrors.DiscardUnknown(m)
}

var xxx_messageInfo_EgressInterfaceErrors proto.InternalMessageInfo

func (m *EgressInterfaceErrors) GetIfErrors() uint64 {
	if m != nil && m.IfErrors != nil {
		return *m.IfErrors
	}
	return 0
}

func (m *EgressInterfaceErrors) GetIfDiscards() uint64 {
	if m != nil && m.IfDiscards != nil {
		return *m.IfDiscards
	}
	return 0
}

var E_JnprInterfaceExt = &proto.ExtensionDesc{
	ExtendedType:  (*JuniperNetworksSensors)(nil),
	ExtensionType: (*Port)(nil),
	Field:         3,
	Name:          "jnpr_interface_ext",
	Tag:           "bytes,3,opt,name=jnpr_interface_ext",
	Filename:      "port.proto",
}

func init() {
	proto.RegisterType((*Port)(nil), "Port")
	proto.RegisterType((*InterfaceInfos)(nil), "InterfaceInfos")
	proto.RegisterType((*QueueStats)(nil), "QueueStats")
	proto.RegisterType((*InterfaceStats)(nil), "InterfaceStats")
	proto.RegisterType((*IngressInterfaceErrors)(nil), "IngressInterfaceErrors")
	proto.RegisterType((*EgressInterfaceErrors)(nil), "EgressInterfaceErrors")
	proto.RegisterExtension(E_JnprInterfaceExt)
}

func init() { proto.RegisterFile("port.proto", fileDescriptor_729c3d36e9010a8e) }

var fileDescriptor_729c3d36e9010a8e = []byte{
	// 1094 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0xe1, 0xc4, 0x8e, 0xed, 0xe3, 0xaf, 0x84, 0x69, 0x52, 0xa1, 0x05, 0x36, 0xc3, 0x5b,
	0x37, 0x0f, 0xdd, 0xbc, 0xd5, 0x2d, 0xd6, 0x76, 0x03, 0x86, 0x36, 0x6b, 0x86, 0x79, 0x48, 0x93,
	0xd4, 0x49, 0xaf, 0x09, 0x46, 0x26, 0x13, 0xce, 0x16, 0xa5, 0x92, 0x54, 0x9b, 0xf4, 0x72, 0xc0,
	0xde, 0x63, 0x6f, 0xb3, 0x87, 0xda, 0xcd, 0xc0, 0x0f, 0xd9, 0x92, 0x9c, 0xdd, 0xec, 0x52, 0xe7,
	0xfc, 0xfe, 0x87, 0xd4, 0xe1, 0x9f, 0x47, 0x02, 0x48, 0x62, 0xa9, 0x47, 0x89, 0x8c, 0x75, 0x7c,
	0x6f, 0x57, 0xd3, 0x05, 0x8d, 0xa8, 0x96, 0x37, 0x58, 0xc7, 0x89, 0x0b, 0x0e, 0x5e, 0x40, 0xf5,
	0x34, 0x96, 0x1a, 0x3d, 0x83, 0x1e, 0x17, 0x9a, 0x4a, 0x46, 0x42, 0x8a, 0x95, 0x26, 0x5a, 0x05,
	0x95, 0xfe, 0xe6, 0xb0, 0x35, 0xee, 0x8d, 0x26, 0x59, 0x7c, 0x22, 0x58, 0xac, 0xa6, 0xdd, 0x25,
	0x77, 0x66, 0xb0, 0xc1, 0x3f, 0x35, 0xe8, 0x16, 0x11, 0xf4, 0x09, 0xd4, 0x39, 0xc3, 0x82, 0x44,
	0x34, 0xa8, 0xf4, 0x37, 0x86, 0xcd, 0x83, 0xda, 0x1f, 0x2f, 0x36, 0x1a, 0x95, 0xe9, 0x16, 0x67,
	0xc7, 0x24, 0xa2, 0xe8, 0x3e, 0x34, 0xb9, 0xe0, 0x1a, 0x6b, 0x1e, 0xd1, 0x60, 0xa3, 0x5f, 0x19,
	0x56, 0xa7, 0x0d, 0x13, 0x38, 0xe7, 0x11, 0x45, 0x03, 0xe8, 0x28, 0x11, 0x25, 0x98, 0x33, 0xcc,
	0xc5, 0x8c, 0x5e, 0x07, 0x9b, 0xfd, 0xca, 0xb0, 0x33, 0x6d, 0x99, 0xe0, 0x84, 0x4d, 0x4c, 0x08,
	0x7d, 0x0e, 0xdd, 0x84, 0x48, 0x2a, 0x34, 0x26, 0xd4, 0xad, 0x53, 0xed, 0x57, 0x86, 0xcd, 0x69,
	0xdb, 0x45, 0x5f, 0x52, 0xbb, 0xcc, 0x53, 0xd8, 0xa1, 0x97, 0x92, 0x2a, 0x85, 0xdf, 0xa5, 0x34,
	0xa5, 0x98, 0x0b, 0x16, 0x07, 0x35, 0xfb, 0x56, 0xad, 0xd1, 0x1b, 0x13, 0xb2, 0x6f, 0x30, 0xed,
	0x39, 0xca, 0x46, 0xcc, 0x0b, 0xa0, 0xe7, 0x80, 0xb8, 0x58, 0x53, 0x6e, 0xad, 0x2b, 0xb7, 0x3d,
	0xb6, 0x92, 0x3e, 0x81, 0x4e, 0x26, 0x75, 0x5d, 0xac, 0xf7, 0x2b, 0xc5, 0x2e, 0x3a, 0x65, 0xdb,
	0x53, 0xf6, 0x09, 0x8d, 0xa1, 0x4d, 0xf3, 0xa2, 0xc6, 0xed, 0xa2, 0x16, 0xcd, 0x69, 0x7e, 0x82,
	0x6e, 0xb6, 0x12, 0x95, 0x32, 0x96, 0x2a, 0x68, 0x5a, 0xd5, 0xdd, 0xd1, 0xc4, 0x85, 0x97, 0xe2,
	0x43, 0x9b, 0x9e, 0x66, 0x1b, 0x73, 0x8f, 0xe8, 0x19, 0x04, 0x9c, 0x61, 0x32, 0x8b, 0xb8, 0xe0,
	0x4a, 0x4b, 0xa2, 0x79, 0x2c, 0xec, 0xf2, 0xa9, 0x0a, 0xc0, 0x76, 0x73, 0x9f, 0xb3, 0x97, 0x85,
	0xf4, 0x99, 0xcd, 0xa2, 0x31, 0xec, 0x71, 0x86, 0xe3, 0x84, 0xba, 0x28, 0x59, 0x64, 0xb2, 0x96,
	0x95, 0xed, 0x72, 0x76, 0xb2, 0xca, 0x79, 0xcd, 0x03, 0xe8, 0x72, 0x86, 0x67, 0x54, 0x85, 0x92,
	0x27, 0x26, 0x13, 0xb4, 0x2d, 0xdc, 0xe1, 0xec, 0xd5, 0x2a, 0x88, 0xbe, 0xb6, 0x98, 0x96, 0x44,
	0x28, 0x6e, 0x02, 0x2a, 0xe8, 0x18, 0x7b, 0x58, 0x03, 0x05, 0x15, 0x43, 0x9f, 0xaf, 0x72, 0x68,
	0x00, 0x6d, 0xce, 0x8e, 0x88, 0xd2, 0x3f, 0x5f, 0x11, 0x71, 0x49, 0x83, 0xae, 0x75, 0x4a, 0x21,
	0x86, 0xfa, 0xd0, 0xe2, 0xec, 0x57, 0x7e, 0x79, 0x75, 0x96, 0x50, 0x3a, 0x0b, 0x7a, 0xce, 0x4c,
	0xb9, 0x10, 0xfa, 0x11, 0x3a, 0xb4, 0xd0, 0xc7, 0x6d, 0xdb, 0xc7, 0xfd, 0xd1, 0xe1, 0xad, 0x6d,
	0xf4, 0x27, 0xe5, 0x9e, 0x06, 0x7f, 0x57, 0x01, 0x56, 0x86, 0x40, 0x43, 0x68, 0x3b, 0xc7, 0x88,
	0x34, 0xba, 0xa0, 0x32, 0xa8, 0x98, 0xe5, 0x32, 0xfb, 0xb7, 0x6c, 0xea, 0xd8, 0x66, 0xd0, 0xa7,
	0x50, 0x4f, 0x48, 0x38, 0xa7, 0x5a, 0xb9, 0x1b, 0x90, 0xbd, 0x62, 0x16, 0x45, 0xf7, 0xa1, 0x76,
	0x71, 0xa3, 0xa9, 0xb2, 0xfe, 0x5f, 0xa6, 0x5d, 0x0c, 0x3d, 0x82, 0x1d, 0x4d, 0xf8, 0x02, 0xcf,
	0x64, 0x9c, 0xe0, 0xac, 0x4e, 0x35, 0x0f, 0xf6, 0x4c, 0xfe, 0x95, 0x8c, 0x93, 0x53, 0x5f, 0xef,
	0x1b, 0xe8, 0xc9, 0x92, 0xa0, 0x56, 0xe8, 0xad, 0x2c, 0xe0, 0x5f, 0x41, 0x27, 0xc3, 0xdd, 0x36,
	0xb6, 0xf2, 0x70, 0xcb, 0xc1, 0x07, 0x76, 0x33, 0xdf, 0xc2, 0xb6, 0xa4, 0xb3, 0x62, 0xe9, 0x7a,
	0x9e, 0xee, 0x4a, 0x3a, 0xcb, 0xd7, 0x7e, 0x08, 0xdd, 0xa5, 0xc0, 0x15, 0x6f, 0xe4, 0xf1, 0xb6,
	0xc7, 0x5d, 0xf5, 0xa7, 0x70, 0x87, 0xbc, 0xbf, 0xc4, 0x17, 0x29, 0x63, 0x54, 0xe2, 0x38, 0x0c,
	0xd3, 0x84, 0x88, 0xf0, 0xc6, 0xba, 0xdd, 0x49, 0xfa, 0x95, 0x29, 0x22, 0xef, 0x2f, 0x0f, 0x2c,
	0x71, 0x92, 0x01, 0x46, 0x18, 0xa6, 0x72, 0x5d, 0x08, 0x05, 0x61, 0x98, 0xca, 0xb2, 0xf0, 0x39,
	0xec, 0x25, 0x94, 0xcc, 0xd7, 0x95, 0xad, 0xbc, 0x72, 0xd7, 0x30, 0xb7, 0x48, 0xc9, 0x62, 0x11,
	0x87, 0x44, 0xd3, 0x59, 0xa6, 0x57, 0xfc, 0x23, 0xb5, 0x6e, 0x5f, 0x49, 0x97, 0x8c, 0xd3, 0x9f,
	0xf1, 0x8f, 0x74, 0xf0, 0xd7, 0x66, 0x6e, 0x8e, 0x3a, 0x37, 0xb9, 0x39, 0x9a, 0xcc, 0xed, 0x30,
	0xce, 0x35, 0x68, 0x8b, 0xb3, 0xd3, 0xb9, 0x36, 0xfe, 0x6f, 0x9a, 0x8b, 0x18, 0xea, 0x35, 0x17,
	0x35, 0x38, 0x3b, 0xb1, 0x61, 0xf4, 0xa5, 0xb9, 0x23, 0xf8, 0x91, 0xa2, 0xa1, 0x2b, 0xb4, 0x99,
	0xdf, 0x08, 0x70, 0x66, 0x32, 0xb6, 0xd8, 0x43, 0x7b, 0xf5, 0x2c, 0xe8, 0x2b, 0x56, 0xf3, 0x68,
	0xdb, 0xa1, 0xbe, 0xea, 0x67, 0x00, 0x9c, 0xe1, 0xd4, 0xd7, 0xac, 0x95, 0x96, 0x7e, 0xeb, 0x2a,
	0x3a, 0x28, 0xf2, 0xd0, 0x56, 0x09, 0x7a, 0x9d, 0x87, 0x2e, 0x3c, 0x54, 0x2f, 0x41, 0x07, 0x0e,
	0xea, 0x43, 0x83, 0x33, 0x77, 0x3d, 0x8b, 0x56, 0xa9, 0x73, 0x66, 0x2f, 0xa2, 0xb1, 0xab, 0x69,
	0x15, 0x49, 0x15, 0x75, 0x95, 0x9a, 0x05, 0xbb, 0x72, 0x76, 0x6a, 0x52, 0xb6, 0xd8, 0x33, 0x3b,
	0xbe, 0x52, 0x31, 0x17, 0xf1, 0x07, 0x81, 0xed, 0x67, 0xd0, 0x49, 0x20, 0x2f, 0x41, 0x9c, 0xbd,
	0x75, 0xc8, 0xa9, 0x21, 0x8c, 0x72, 0xf0, 0x67, 0x15, 0xf6, 0x6f, 0x1f, 0xae, 0xfe, 0x28, 0xfc,
	0x00, 0xa9, 0x94, 0xde, 0xc2, 0x33, 0xee, 0x28, 0xb8, 0xc0, 0xef, 0x8c, 0xf3, 0x4b, 0x27, 0x06,
	0x9c, 0x4d, 0xc4, 0x1b, 0x9b, 0x40, 0x63, 0x40, 0x0e, 0x64, 0x92, 0x44, 0x34, 0xab, 0x5a, 0x98,
	0x03, 0x3d, 0x83, 0xff, 0x62, 0xd2, 0xbe, 0xf8, 0x17, 0x66, 0xce, 0xe1, 0x19, 0x57, 0x21, 0x91,
	0xb3, 0xd2, 0x2c, 0x00, 0xce, 0x5e, 0xf9, 0x04, 0x7a, 0x60, 0x39, 0x2e, 0xb0, 0x4c, 0x45, 0xf9,
	0xe8, 0x9a, 0xa6, 0xe8, 0xd4, 0xc4, 0xd1, 0xf7, 0x70, 0xc7, 0x61, 0x8b, 0xc7, 0x98, 0x8b, 0x30,
	0x8e, 0x92, 0x05, 0x5d, 0x9b, 0x02, 0x3b, 0x86, 0x3f, 0x7a, 0x3c, 0x59, 0xe5, 0xd1, 0x13, 0xd8,
	0xf5, 0xba, 0x71, 0x78, 0x45, 0x44, 0xb6, 0xf7, 0xc2, 0xb9, 0x6e, 0x5b, 0x99, 0xcd, 0xfb, 0xcd,
	0xbf, 0x84, 0x7b, 0x99, 0x0a, 0x47, 0x5c, 0x45, 0x44, 0x87, 0x57, 0xf6, 0xef, 0x20, 0x4e, 0x75,
	0x69, 0x38, 0xec, 0x3b, 0xf1, 0x6b, 0x4f, 0x9d, 0x7b, 0x08, 0x7d, 0x07, 0x3b, 0xbe, 0x67, 0x9c,
	0xc5, 0xf9, 0x2f, 0xe2, 0x6a, 0x0a, 0xd9, 0x96, 0x71, 0x16, 0x2f, 0x3f, 0x80, 0x7b, 0xbe, 0x13,
	0x54, 0xc5, 0xa9, 0x0c, 0x97, 0x8d, 0x2e, 0xfb, 0x60, 0x22, 0xa6, 0x9e, 0xf0, 0x43, 0x3f, 0x84,
	0xbd, 0xc3, 0xff, 0xed, 0x82, 0xd2, 0x41, 0x6d, 0xfc, 0xc7, 0x41, 0xfd, 0x70, 0x04, 0xe8, 0x77,
	0x91, 0x48, 0xbc, 0xfa, 0x2d, 0xa3, 0xd7, 0x1a, 0xdd, 0x1d, 0xfd, 0x96, 0x0a, 0x9e, 0x50, 0x79,
	0x4c, 0xf5, 0x87, 0x58, 0xce, 0xd5, 0x19, 0x15, 0x2a, 0x73, 0x47, 0x6b, 0x5c, 0x1b, 0x99, 0xbf,
	0xb9, 0xe9, 0xb6, 0x51, 0xae, 0x36, 0x77, 0xad, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x76,
	0x73, 0x5d, 0x09, 0x0a, 0x00, 0x00,
}
