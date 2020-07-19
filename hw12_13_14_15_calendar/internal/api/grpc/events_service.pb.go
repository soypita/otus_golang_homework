// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events_service.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Event struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Header               string               `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	OwnerId              string               `protobuf:"bytes,6,opt,name=ownerId,proto3" json:"ownerId,omitempty"`
	NotifyBefore         *duration.Duration   `protobuf:"bytes,7,opt,name=notifyBefore,proto3" json:"notifyBefore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Event) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Event) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

func (m *Event) GetNotifyBefore() *duration.Duration {
	if m != nil {
		return m.NotifyBefore
	}
	return nil
}

type CreateEventRequest struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventRequest) Reset()         { *m = CreateEventRequest{} }
func (m *CreateEventRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEventRequest) ProtoMessage()    {}
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{1}
}

func (m *CreateEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventRequest.Unmarshal(m, b)
}
func (m *CreateEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventRequest.Marshal(b, m, deterministic)
}
func (m *CreateEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventRequest.Merge(m, src)
}
func (m *CreateEventRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEventRequest.Size(m)
}
func (m *CreateEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventRequest proto.InternalMessageInfo

func (m *CreateEventRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type CreateEventResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEventResponse) Reset()         { *m = CreateEventResponse{} }
func (m *CreateEventResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEventResponse) ProtoMessage()    {}
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{2}
}

func (m *CreateEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEventResponse.Unmarshal(m, b)
}
func (m *CreateEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEventResponse.Marshal(b, m, deterministic)
}
func (m *CreateEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEventResponse.Merge(m, src)
}
func (m *CreateEventResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEventResponse.Size(m)
}
func (m *CreateEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEventResponse proto.InternalMessageInfo

func (m *CreateEventResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EventUpdateRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                *Event   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventUpdateRequest) Reset()         { *m = EventUpdateRequest{} }
func (m *EventUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*EventUpdateRequest) ProtoMessage()    {}
func (*EventUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{3}
}

func (m *EventUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventUpdateRequest.Unmarshal(m, b)
}
func (m *EventUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventUpdateRequest.Marshal(b, m, deterministic)
}
func (m *EventUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUpdateRequest.Merge(m, src)
}
func (m *EventUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_EventUpdateRequest.Size(m)
}
func (m *EventUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventUpdateRequest proto.InternalMessageInfo

func (m *EventUpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventUpdateRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type DeleteEventRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventRequest) Reset()         { *m = DeleteEventRequest{} }
func (m *DeleteEventRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEventRequest) ProtoMessage()    {}
func (*DeleteEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{4}
}

func (m *DeleteEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventRequest.Unmarshal(m, b)
}
func (m *DeleteEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventRequest.Merge(m, src)
}
func (m *DeleteEventRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEventRequest.Size(m)
}
func (m *DeleteEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventRequest proto.InternalMessageInfo

func (m *DeleteEventRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetEventByIDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventByIDRequest) Reset()         { *m = GetEventByIDRequest{} }
func (m *GetEventByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventByIDRequest) ProtoMessage()    {}
func (*GetEventByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{5}
}

func (m *GetEventByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventByIDRequest.Unmarshal(m, b)
}
func (m *GetEventByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetEventByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventByIDRequest.Merge(m, src)
}
func (m *GetEventByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventByIDRequest.Size(m)
}
func (m *GetEventByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventByIDRequest proto.InternalMessageInfo

func (m *GetEventByIDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetEventByIDResponse struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventByIDResponse) Reset()         { *m = GetEventByIDResponse{} }
func (m *GetEventByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventByIDResponse) ProtoMessage()    {}
func (*GetEventByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{6}
}

func (m *GetEventByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventByIDResponse.Unmarshal(m, b)
}
func (m *GetEventByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetEventByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventByIDResponse.Merge(m, src)
}
func (m *GetEventByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventByIDResponse.Size(m)
}
func (m *GetEventByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventByIDResponse proto.InternalMessageInfo

func (m *GetEventByIDResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type GetAllEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllEventsResponse) Reset()         { *m = GetAllEventsResponse{} }
func (m *GetAllEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllEventsResponse) ProtoMessage()    {}
func (*GetAllEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{7}
}

func (m *GetAllEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllEventsResponse.Unmarshal(m, b)
}
func (m *GetAllEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetAllEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllEventsResponse.Merge(m, src)
}
func (m *GetAllEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllEventsResponse.Size(m)
}
func (m *GetAllEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllEventsResponse proto.InternalMessageInfo

func (m *GetAllEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type FindDayEventsRequest struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FindDayEventsRequest) Reset()         { *m = FindDayEventsRequest{} }
func (m *FindDayEventsRequest) String() string { return proto.CompactTextString(m) }
func (*FindDayEventsRequest) ProtoMessage()    {}
func (*FindDayEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{8}
}

func (m *FindDayEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindDayEventsRequest.Unmarshal(m, b)
}
func (m *FindDayEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindDayEventsRequest.Marshal(b, m, deterministic)
}
func (m *FindDayEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindDayEventsRequest.Merge(m, src)
}
func (m *FindDayEventsRequest) XXX_Size() int {
	return xxx_messageInfo_FindDayEventsRequest.Size(m)
}
func (m *FindDayEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindDayEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindDayEventsRequest proto.InternalMessageInfo

func (m *FindDayEventsRequest) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type FindDayEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindDayEventsResponse) Reset()         { *m = FindDayEventsResponse{} }
func (m *FindDayEventsResponse) String() string { return proto.CompactTextString(m) }
func (*FindDayEventsResponse) ProtoMessage()    {}
func (*FindDayEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{9}
}

func (m *FindDayEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindDayEventsResponse.Unmarshal(m, b)
}
func (m *FindDayEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindDayEventsResponse.Marshal(b, m, deterministic)
}
func (m *FindDayEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindDayEventsResponse.Merge(m, src)
}
func (m *FindDayEventsResponse) XXX_Size() int {
	return xxx_messageInfo_FindDayEventsResponse.Size(m)
}
func (m *FindDayEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindDayEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindDayEventsResponse proto.InternalMessageInfo

func (m *FindDayEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type FindWeekEventsRequest struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FindWeekEventsRequest) Reset()         { *m = FindWeekEventsRequest{} }
func (m *FindWeekEventsRequest) String() string { return proto.CompactTextString(m) }
func (*FindWeekEventsRequest) ProtoMessage()    {}
func (*FindWeekEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{10}
}

func (m *FindWeekEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWeekEventsRequest.Unmarshal(m, b)
}
func (m *FindWeekEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWeekEventsRequest.Marshal(b, m, deterministic)
}
func (m *FindWeekEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWeekEventsRequest.Merge(m, src)
}
func (m *FindWeekEventsRequest) XXX_Size() int {
	return xxx_messageInfo_FindWeekEventsRequest.Size(m)
}
func (m *FindWeekEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWeekEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindWeekEventsRequest proto.InternalMessageInfo

func (m *FindWeekEventsRequest) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type FindWeekEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWeekEventsResponse) Reset()         { *m = FindWeekEventsResponse{} }
func (m *FindWeekEventsResponse) String() string { return proto.CompactTextString(m) }
func (*FindWeekEventsResponse) ProtoMessage()    {}
func (*FindWeekEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{11}
}

func (m *FindWeekEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWeekEventsResponse.Unmarshal(m, b)
}
func (m *FindWeekEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWeekEventsResponse.Marshal(b, m, deterministic)
}
func (m *FindWeekEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWeekEventsResponse.Merge(m, src)
}
func (m *FindWeekEventsResponse) XXX_Size() int {
	return xxx_messageInfo_FindWeekEventsResponse.Size(m)
}
func (m *FindWeekEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWeekEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindWeekEventsResponse proto.InternalMessageInfo

func (m *FindWeekEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type FindMonthEventsRequest struct {
	Date                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FindMonthEventsRequest) Reset()         { *m = FindMonthEventsRequest{} }
func (m *FindMonthEventsRequest) String() string { return proto.CompactTextString(m) }
func (*FindMonthEventsRequest) ProtoMessage()    {}
func (*FindMonthEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{12}
}

func (m *FindMonthEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMonthEventsRequest.Unmarshal(m, b)
}
func (m *FindMonthEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMonthEventsRequest.Marshal(b, m, deterministic)
}
func (m *FindMonthEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMonthEventsRequest.Merge(m, src)
}
func (m *FindMonthEventsRequest) XXX_Size() int {
	return xxx_messageInfo_FindMonthEventsRequest.Size(m)
}
func (m *FindMonthEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMonthEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMonthEventsRequest proto.InternalMessageInfo

func (m *FindMonthEventsRequest) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

type FindMonthEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMonthEventsResponse) Reset()         { *m = FindMonthEventsResponse{} }
func (m *FindMonthEventsResponse) String() string { return proto.CompactTextString(m) }
func (*FindMonthEventsResponse) ProtoMessage()    {}
func (*FindMonthEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_192ada64c58763bd, []int{13}
}

func (m *FindMonthEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMonthEventsResponse.Unmarshal(m, b)
}
func (m *FindMonthEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMonthEventsResponse.Marshal(b, m, deterministic)
}
func (m *FindMonthEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMonthEventsResponse.Merge(m, src)
}
func (m *FindMonthEventsResponse) XXX_Size() int {
	return xxx_messageInfo_FindMonthEventsResponse.Size(m)
}
func (m *FindMonthEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMonthEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMonthEventsResponse proto.InternalMessageInfo

func (m *FindMonthEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*CreateEventRequest)(nil), "CreateEventRequest")
	proto.RegisterType((*CreateEventResponse)(nil), "CreateEventResponse")
	proto.RegisterType((*EventUpdateRequest)(nil), "EventUpdateRequest")
	proto.RegisterType((*DeleteEventRequest)(nil), "DeleteEventRequest")
	proto.RegisterType((*GetEventByIDRequest)(nil), "GetEventByIDRequest")
	proto.RegisterType((*GetEventByIDResponse)(nil), "GetEventByIDResponse")
	proto.RegisterType((*GetAllEventsResponse)(nil), "GetAllEventsResponse")
	proto.RegisterType((*FindDayEventsRequest)(nil), "FindDayEventsRequest")
	proto.RegisterType((*FindDayEventsResponse)(nil), "FindDayEventsResponse")
	proto.RegisterType((*FindWeekEventsRequest)(nil), "FindWeekEventsRequest")
	proto.RegisterType((*FindWeekEventsResponse)(nil), "FindWeekEventsResponse")
	proto.RegisterType((*FindMonthEventsRequest)(nil), "FindMonthEventsRequest")
	proto.RegisterType((*FindMonthEventsResponse)(nil), "FindMonthEventsResponse")
}

func init() {
	proto.RegisterFile("events_service.proto", fileDescriptor_192ada64c58763bd)
}

var fileDescriptor_192ada64c58763bd = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x95, 0x34, 0x71, 0xc3, 0xb8, 0x14, 0x69, 0xe3, 0xa4, 0xc6, 0xa0, 0x12, 0x59, 0x54,
	0xea, 0x69, 0x2b, 0x85, 0x7f, 0x85, 0xaa, 0x48, 0x4d, 0xdd, 0x96, 0x1e, 0xb8, 0x44, 0x20, 0x24,
	0x2e, 0xc8, 0x8d, 0x27, 0xa9, 0x85, 0xe3, 0x35, 0xeb, 0x4d, 0x51, 0x8e, 0x7c, 0x04, 0xbe, 0x31,
	0xca, 0xae, 0x1d, 0xf9, 0x2f, 0x8d, 0xd4, 0xe3, 0xee, 0xbc, 0x37, 0xf3, 0x3c, 0xfe, 0xd9, 0x60,
	0xe0, 0x1d, 0x86, 0x22, 0xfe, 0x11, 0x23, 0xbf, 0xf3, 0x27, 0x48, 0x23, 0xce, 0x04, 0xb3, 0x5e,
	0xcc, 0x18, 0x9b, 0x05, 0x78, 0x24, 0x4f, 0x37, 0x8b, 0xe9, 0x91, 0xf0, 0xe7, 0x18, 0x0b, 0x77,
	0x1e, 0x25, 0x82, 0xfd, 0xa2, 0xc0, 0x5b, 0x70, 0x57, 0xf8, 0x2c, 0x4c, 0xea, 0xcf, 0x8a, 0x75,
	0x9c, 0x47, 0x62, 0xa9, 0x8a, 0xf6, 0xdf, 0x26, 0xb4, 0x2f, 0x56, 0x63, 0xc9, 0x2e, 0x34, 0x7d,
	0xcf, 0x6c, 0x0c, 0x1a, 0x87, 0x8f, 0xc6, 0x4d, 0xdf, 0x23, 0x7d, 0xd0, 0x6e, 0xd1, 0xf5, 0x90,
	0x9b, 0x4d, 0x79, 0x97, 0x9c, 0x08, 0x85, 0x96, 0xe7, 0x0a, 0x34, 0xb7, 0x06, 0x8d, 0x43, 0x7d,
	0x68, 0x51, 0xd5, 0x9d, 0xa6, 0xdd, 0xe9, 0x97, 0x34, 0xde, 0x58, 0xea, 0xc8, 0x1b, 0xe8, 0xa4,
	0x81, 0xcc, 0x96, 0xf4, 0x3c, 0x2d, 0x79, 0x9c, 0x44, 0x30, 0x5e, 0x4b, 0xc9, 0x00, 0x74, 0x0f,
	0xe3, 0x09, 0xf7, 0x23, 0xe9, 0x6c, 0xcb, 0x0c, 0xd9, 0x2b, 0x62, 0xc2, 0x36, 0xfb, 0x1d, 0x22,
	0xbf, 0xf6, 0x4c, 0x4d, 0x56, 0xd3, 0x23, 0x39, 0x85, 0x9d, 0x90, 0x09, 0x7f, 0xba, 0x1c, 0xe1,
	0x94, 0x71, 0x34, 0xb7, 0xef, 0x1b, 0x9b, 0x93, 0xdb, 0x43, 0x20, 0xe7, 0x1c, 0x5d, 0x81, 0x72,
	0x31, 0x63, 0xfc, 0xb5, 0xc0, 0x58, 0x90, 0xe7, 0xd0, 0x96, 0xef, 0x47, 0xae, 0x48, 0x1f, 0x6a,
	0x54, 0x55, 0xd5, 0xa5, 0x7d, 0x00, 0xdd, 0x9c, 0x27, 0x8e, 0x58, 0x18, 0x63, 0x71, 0xa9, 0xf6,
	0x08, 0x88, 0x14, 0x7c, 0x8d, 0x56, 0xbb, 0x49, 0x5b, 0x17, 0x57, 0xbf, 0x1e, 0xd5, 0xac, 0x1a,
	0xf5, 0x12, 0x88, 0x83, 0x01, 0x16, 0xe2, 0x15, 0x27, 0x1d, 0x40, 0xf7, 0x0a, 0x85, 0x94, 0x8c,
	0x96, 0xd7, 0x4e, 0x9d, 0xec, 0x35, 0x18, 0x79, 0x59, 0x12, 0xfc, 0xff, 0x4f, 0xfb, 0x56, 0xba,
	0xce, 0x82, 0x40, 0xde, 0xc6, 0x6b, 0xd7, 0x3e, 0x68, 0x8a, 0x61, 0xb3, 0x31, 0xd8, 0xca, 0xd8,
	0x92, 0x5b, 0xfb, 0x12, 0x8c, 0x4b, 0x3f, 0xf4, 0x1c, 0x77, 0x99, 0x1a, 0x55, 0xaa, 0x94, 0xa9,
	0xc6, 0x66, 0x4c, 0xd9, 0xef, 0xa0, 0x57, 0xe8, 0xb3, 0x61, 0x80, 0x2b, 0x65, 0xfc, 0x86, 0xf8,
	0xf3, 0x61, 0x09, 0x8e, 0xa1, 0x5f, 0x6c, 0xb4, 0x61, 0x84, 0x4f, 0xca, 0xf9, 0x99, 0x85, 0xe2,
	0xf6, 0x61, 0x19, 0xde, 0xc3, 0x5e, 0xa9, 0xd3, 0x66, 0x21, 0x86, 0x7f, 0x5a, 0xd0, 0x39, 0x77,
	0x03, 0x0c, 0x3d, 0x97, 0x93, 0x63, 0xd0, 0x33, 0xec, 0x92, 0x2e, 0x2d, 0xd3, 0x6f, 0x19, 0xb4,
	0x0a, 0xef, 0x0f, 0xa0, 0x2b, 0x92, 0x53, 0x67, 0x19, 0x6e, 0xab, 0x5f, 0x7a, 0x8e, 0x8b, 0xd5,
	0xff, 0x67, 0xe5, 0xcd, 0x60, 0x4c, 0xba, 0xb4, 0x0c, 0x75, 0xad, 0xf7, 0x14, 0x76, 0xb2, 0xfc,
	0x91, 0x1a, 0x9d, 0xd5, 0xa3, 0x95, 0x98, 0x9e, 0x48, 0xfb, 0x1a, 0x7a, 0x62, 0xd0, 0x8a, 0x4f,
	0x45, 0x99, 0xcb, 0x5f, 0xc6, 0x47, 0x78, 0x9c, 0x63, 0x8f, 0xf4, 0x68, 0x15, 0xd3, 0x56, 0x9f,
	0x56, 0x23, 0x7a, 0x06, 0xbb, 0x79, 0x72, 0x88, 0x52, 0x96, 0x98, 0xb4, 0xf6, 0x68, 0x0d, 0x62,
	0x0e, 0x3c, 0x29, 0xbc, 0x78, 0xa2, 0xb4, 0x65, 0xa8, 0x2c, 0x93, 0xd6, 0x30, 0x32, 0xea, 0x7c,
	0xd7, 0xe8, 0xc9, 0x8c, 0x47, 0x93, 0x1b, 0x4d, 0xae, 0xed, 0xd5, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xf4, 0x8c, 0x1f, 0xbe, 0x81, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	UpdateEvent(ctx context.Context, in *EventUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetAllEvents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllEventsResponse, error)
	GetEventByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*GetEventByIDResponse, error)
	FindDayEvents(ctx context.Context, in *FindDayEventsRequest, opts ...grpc.CallOption) (*FindDayEventsResponse, error)
	FindWeekEvents(ctx context.Context, in *FindWeekEventsRequest, opts ...grpc.CallOption) (*FindWeekEventsResponse, error)
	FindMonthEvents(ctx context.Context, in *FindMonthEventsRequest, opts ...grpc.CallOption) (*FindMonthEventsResponse, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/Calendar/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) UpdateEvent(ctx context.Context, in *EventUpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Calendar/UpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) DeleteEvent(ctx context.Context, in *DeleteEventRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Calendar/DeleteEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetAllEvents(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllEventsResponse, error) {
	out := new(GetAllEventsResponse)
	err := c.cc.Invoke(ctx, "/Calendar/GetAllEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) GetEventByID(ctx context.Context, in *GetEventByIDRequest, opts ...grpc.CallOption) (*GetEventByIDResponse, error) {
	out := new(GetEventByIDResponse)
	err := c.cc.Invoke(ctx, "/Calendar/GetEventByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) FindDayEvents(ctx context.Context, in *FindDayEventsRequest, opts ...grpc.CallOption) (*FindDayEventsResponse, error) {
	out := new(FindDayEventsResponse)
	err := c.cc.Invoke(ctx, "/Calendar/FindDayEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) FindWeekEvents(ctx context.Context, in *FindWeekEventsRequest, opts ...grpc.CallOption) (*FindWeekEventsResponse, error) {
	out := new(FindWeekEventsResponse)
	err := c.cc.Invoke(ctx, "/Calendar/FindWeekEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) FindMonthEvents(ctx context.Context, in *FindMonthEventsRequest, opts ...grpc.CallOption) (*FindMonthEventsResponse, error) {
	out := new(FindMonthEventsResponse)
	err := c.cc.Invoke(ctx, "/Calendar/FindMonthEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	UpdateEvent(context.Context, *EventUpdateRequest) (*empty.Empty, error)
	DeleteEvent(context.Context, *DeleteEventRequest) (*empty.Empty, error)
	GetAllEvents(context.Context, *empty.Empty) (*GetAllEventsResponse, error)
	GetEventByID(context.Context, *GetEventByIDRequest) (*GetEventByIDResponse, error)
	FindDayEvents(context.Context, *FindDayEventsRequest) (*FindDayEventsResponse, error)
	FindWeekEvents(context.Context, *FindWeekEventsRequest) (*FindWeekEventsResponse, error)
	FindMonthEvents(context.Context, *FindMonthEventsRequest) (*FindMonthEventsResponse, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) CreateEvent(ctx context.Context, req *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedCalendarServer) UpdateEvent(ctx context.Context, req *EventUpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateEvent not implemented")
}
func (*UnimplementedCalendarServer) DeleteEvent(ctx context.Context, req *DeleteEventRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEvent not implemented")
}
func (*UnimplementedCalendarServer) GetAllEvents(ctx context.Context, req *empty.Empty) (*GetAllEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEvents not implemented")
}
func (*UnimplementedCalendarServer) GetEventByID(ctx context.Context, req *GetEventByIDRequest) (*GetEventByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventByID not implemented")
}
func (*UnimplementedCalendarServer) FindDayEvents(ctx context.Context, req *FindDayEventsRequest) (*FindDayEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindDayEvents not implemented")
}
func (*UnimplementedCalendarServer) FindWeekEvents(ctx context.Context, req *FindWeekEventsRequest) (*FindWeekEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindWeekEvents not implemented")
}
func (*UnimplementedCalendarServer) FindMonthEvents(ctx context.Context, req *FindMonthEventsRequest) (*FindMonthEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMonthEvents not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_UpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).UpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/UpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).UpdateEvent(ctx, req.(*EventUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_DeleteEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).DeleteEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/DeleteEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).DeleteEvent(ctx, req.(*DeleteEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetAllEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetAllEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/GetAllEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetAllEvents(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_GetEventByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEventByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/GetEventByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEventByID(ctx, req.(*GetEventByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_FindDayEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindDayEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).FindDayEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/FindDayEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).FindDayEvents(ctx, req.(*FindDayEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_FindWeekEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindWeekEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).FindWeekEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/FindWeekEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).FindWeekEvents(ctx, req.(*FindWeekEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_FindMonthEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindMonthEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).FindMonthEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/FindMonthEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).FindMonthEvents(ctx, req.(*FindMonthEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _Calendar_CreateEvent_Handler,
		},
		{
			MethodName: "UpdateEvent",
			Handler:    _Calendar_UpdateEvent_Handler,
		},
		{
			MethodName: "DeleteEvent",
			Handler:    _Calendar_DeleteEvent_Handler,
		},
		{
			MethodName: "GetAllEvents",
			Handler:    _Calendar_GetAllEvents_Handler,
		},
		{
			MethodName: "GetEventByID",
			Handler:    _Calendar_GetEventByID_Handler,
		},
		{
			MethodName: "FindDayEvents",
			Handler:    _Calendar_FindDayEvents_Handler,
		},
		{
			MethodName: "FindWeekEvents",
			Handler:    _Calendar_FindWeekEvents_Handler,
		},
		{
			MethodName: "FindMonthEvents",
			Handler:    _Calendar_FindMonthEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events_service.proto",
}
