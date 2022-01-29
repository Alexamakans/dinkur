// Dinkur the task time tracking utility.
// <https://github.com/dinkur/dinkur>
//
// Copyright (C) 2021 Kalle Fagerberg
// SPDX-FileCopyrightText: 2021 Kalle Fagerberg
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.2
// source: api/dinkurapi/v1/alerter.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StreamAlertRequest is an empty message and unused. It is here as a
// placeholder for potential future use.
type StreamAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamAlertRequest) Reset() {
	*x = StreamAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAlertRequest) ProtoMessage() {}

func (x *StreamAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAlertRequest.ProtoReflect.Descriptor instead.
func (*StreamAlertRequest) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{0}
}

// StreamAlertResponse is an alert event. An alert has been created, updated,
// or deleted.
type StreamAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alert is the created, updated, or deleted alert.
	Alert *Alert `protobuf:"bytes,1,opt,name=alert,proto3" json:"alert,omitempty"`
	// Event is the type of event.
	Event Event `protobuf:"varint,2,opt,name=event,proto3,enum=dinkurapi.v1.Event" json:"event,omitempty"`
}

func (x *StreamAlertResponse) Reset() {
	*x = StreamAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAlertResponse) ProtoMessage() {}

func (x *StreamAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAlertResponse.ProtoReflect.Descriptor instead.
func (*StreamAlertResponse) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{1}
}

func (x *StreamAlertResponse) GetAlert() *Alert {
	if x != nil {
		return x.Alert
	}
	return nil
}

func (x *StreamAlertResponse) GetEvent() Event {
	if x != nil {
		return x.Event
	}
	return Event_EVENT_UNSPECIFIED
}

// GetAlertListRequest is an empty message and unused. It is here as a
// placeholder for potential future use.
type GetAlertListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAlertListRequest) Reset() {
	*x = GetAlertListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertListRequest) ProtoMessage() {}

func (x *GetAlertListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertListRequest.ProtoReflect.Descriptor instead.
func (*GetAlertListRequest) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{2}
}

// GetAlertListResponse holds a list of alerts.
type GetAlertListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alerts is the list of alerts.
	Alerts []*Alert `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *GetAlertListResponse) Reset() {
	*x = GetAlertListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlertListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertListResponse) ProtoMessage() {}

func (x *GetAlertListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlertListResponse.ProtoReflect.Descriptor instead.
func (*GetAlertListResponse) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{3}
}

func (x *GetAlertListResponse) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

// DeleteAlertRequest holds the ID of the alert to delete.
type DeleteAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the alert to delete.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAlertRequest) Reset() {
	*x = DeleteAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlertRequest) ProtoMessage() {}

func (x *DeleteAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlertRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlertRequest) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAlertRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// DeleteAlertResponse holds the alert that was just deleted.
type DeleteAlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DeletedAlert is the alert that was just deleted.
	DeletedAlert *Alert `protobuf:"bytes,1,opt,name=deleted_alert,json=deletedAlert,proto3" json:"deleted_alert,omitempty"`
}

func (x *DeleteAlertResponse) Reset() {
	*x = DeleteAlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlertResponse) ProtoMessage() {}

func (x *DeleteAlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlertResponse.ProtoReflect.Descriptor instead.
func (*DeleteAlertResponse) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAlertResponse) GetDeletedAlert() *Alert {
	if x != nil {
		return x.DeletedAlert
	}
	return nil
}

// Alert is a sort of notification issued by the Dinkur daemon, and contains a
// union type of different alert types.
type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id is the unique identifier of this alert, and is used when deleting an
	// alert via the Alerter service.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Created is a timestamp of when the alert was initially issued.
	Created *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	// Updated is a timestamp of when the alert was most recently changed. This
	// has the same value as when the alert was created if it has never been
	// updated.
	Updated *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated,proto3" json:"updated,omitempty"`
	// Type holds the inner type of this alert.
	//
	// Types that are assignable to Type:
	//	*Alert_PlainMessage
	//	*Alert_Afk
	Type isAlert_Type `protobuf_oneof:"type"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{6}
}

func (x *Alert) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Alert) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Alert) GetUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.Updated
	}
	return nil
}

func (m *Alert) GetType() isAlert_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Alert) GetPlainMessage() *AlertPlainMessage {
	if x, ok := x.GetType().(*Alert_PlainMessage); ok {
		return x.PlainMessage
	}
	return nil
}

func (x *Alert) GetAfk() *AlertAfk {
	if x, ok := x.GetType().(*Alert_Afk); ok {
		return x.Afk
	}
	return nil
}

type isAlert_Type interface {
	isAlert_Type()
}

type Alert_PlainMessage struct {
	// PlainMessage means this alert is a plain string message alert.
	PlainMessage *AlertPlainMessage `protobuf:"bytes,4,opt,name=plain_message,json=plainMessage,proto3,oneof"`
}

type Alert_Afk struct {
	// AFK means this alert is an AFK alert, where the user has been away from
	// their keyboard for too long, and any clients can start idling. The alert
	// also contains information for when the user has returned from being AFK,
	// for which the Dinkur clients are expected to prompt the user for how to
	// save the away time.
	Afk *AlertAfk `protobuf:"bytes,5,opt,name=afk,proto3,oneof"`
}

func (*Alert_PlainMessage) isAlert_Type() {}

func (*Alert_Afk) isAlert_Type() {}

// AlertPlainMessage is an alert subtype that only holds a generic message.
// No user actions are expected. Dinkur clients should show this message and
// let the user dismiss it by deleting the alert.
type AlertPlainMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message is a free text message that can contain any text. No format is
	// assumed.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AlertPlainMessage) Reset() {
	*x = AlertPlainMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertPlainMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertPlainMessage) ProtoMessage() {}

func (x *AlertPlainMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertPlainMessage.ProtoReflect.Descriptor instead.
func (*AlertPlainMessage) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{7}
}

func (x *AlertPlainMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// AlertAfk is an alert subtype that is issued when the user has been away from
// their keyboard for too long. Dinkur clients can start idling while this
// alert exists. This alert is automatically deleted when the user is no longer
// AFK. The alert is only issued when the user has an active entry.
//
// Only one instance of this alert may exist at a single time.
type AlertAfk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ActiveEntry is the Dinkur entry that is currently active now that the user
	// is AFK. This field should always be set, as this alert is only issued when
	// the user has an active entry.
	ActiveEntry *Entry `protobuf:"bytes,1,opt,name=active_entry,json=activeEntry,proto3" json:"active_entry,omitempty"`
	// AfkSince is a timestamp of when the user went AFK.
	AfkSince *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=afk_since,json=afkSince,proto3" json:"afk_since,omitempty"`
	// BackSince is a timestamp of when the user returned from being AFK.
	// This value is null/nil when the user has not yet returned.
	BackSince *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=back_since,json=backSince,proto3" json:"back_since,omitempty"`
}

func (x *AlertAfk) Reset() {
	*x = AlertAfk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertAfk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertAfk) ProtoMessage() {}

func (x *AlertAfk) ProtoReflect() protoreflect.Message {
	mi := &file_api_dinkurapi_v1_alerter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertAfk.ProtoReflect.Descriptor instead.
func (*AlertAfk) Descriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_alerter_proto_rawDescGZIP(), []int{8}
}

func (x *AlertAfk) GetActiveEntry() *Entry {
	if x != nil {
		return x.ActiveEntry
	}
	return nil
}

func (x *AlertAfk) GetAfkSince() *timestamppb.Timestamp {
	if x != nil {
		return x.AfkSince
	}
	return nil
}

func (x *AlertAfk) GetBackSince() *timestamppb.Timestamp {
	if x != nil {
		return x.BackSince
	}
	return nil
}

var File_api_dinkurapi_v1_alerter_proto protoreflect.FileDescriptor

var file_api_dinkurapi_v1_alerter_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a,
	0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x6e, 0x6b,
	0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x05,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x43, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x4f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0d, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x22, 0xff, 0x01, 0x0a, 0x05, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x46, 0x0a, 0x0d, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61, 0x66, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x41, 0x66, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x66, 0x6b, 0x42, 0x06, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x50, 0x6c,
	0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x08, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x41, 0x66,
	0x6b, 0x12, 0x36, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x66, 0x6b,
	0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x61, 0x66, 0x6b, 0x53, 0x69, 0x6e,
	0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x69, 0x6e, 0x63, 0x65, 0x32, 0x8a, 0x02,
	0x0a, 0x07, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12, 0x54, 0x0a, 0x0b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x20, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75,
	0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x69, 0x6e,
	0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12,
	0x55, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x21, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x20, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x2f,
	0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75,
	0x72, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dinkurapi_v1_alerter_proto_rawDescOnce sync.Once
	file_api_dinkurapi_v1_alerter_proto_rawDescData = file_api_dinkurapi_v1_alerter_proto_rawDesc
)

func file_api_dinkurapi_v1_alerter_proto_rawDescGZIP() []byte {
	file_api_dinkurapi_v1_alerter_proto_rawDescOnce.Do(func() {
		file_api_dinkurapi_v1_alerter_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dinkurapi_v1_alerter_proto_rawDescData)
	})
	return file_api_dinkurapi_v1_alerter_proto_rawDescData
}

var file_api_dinkurapi_v1_alerter_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_dinkurapi_v1_alerter_proto_goTypes = []interface{}{
	(*StreamAlertRequest)(nil),    // 0: dinkurapi.v1.StreamAlertRequest
	(*StreamAlertResponse)(nil),   // 1: dinkurapi.v1.StreamAlertResponse
	(*GetAlertListRequest)(nil),   // 2: dinkurapi.v1.GetAlertListRequest
	(*GetAlertListResponse)(nil),  // 3: dinkurapi.v1.GetAlertListResponse
	(*DeleteAlertRequest)(nil),    // 4: dinkurapi.v1.DeleteAlertRequest
	(*DeleteAlertResponse)(nil),   // 5: dinkurapi.v1.DeleteAlertResponse
	(*Alert)(nil),                 // 6: dinkurapi.v1.Alert
	(*AlertPlainMessage)(nil),     // 7: dinkurapi.v1.AlertPlainMessage
	(*AlertAfk)(nil),              // 8: dinkurapi.v1.AlertAfk
	(Event)(0),                    // 9: dinkurapi.v1.Event
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
	(*Entry)(nil),                 // 11: dinkurapi.v1.Entry
}
var file_api_dinkurapi_v1_alerter_proto_depIdxs = []int32{
	6,  // 0: dinkurapi.v1.StreamAlertResponse.alert:type_name -> dinkurapi.v1.Alert
	9,  // 1: dinkurapi.v1.StreamAlertResponse.event:type_name -> dinkurapi.v1.Event
	6,  // 2: dinkurapi.v1.GetAlertListResponse.alerts:type_name -> dinkurapi.v1.Alert
	6,  // 3: dinkurapi.v1.DeleteAlertResponse.deleted_alert:type_name -> dinkurapi.v1.Alert
	10, // 4: dinkurapi.v1.Alert.created:type_name -> google.protobuf.Timestamp
	10, // 5: dinkurapi.v1.Alert.updated:type_name -> google.protobuf.Timestamp
	7,  // 6: dinkurapi.v1.Alert.plain_message:type_name -> dinkurapi.v1.AlertPlainMessage
	8,  // 7: dinkurapi.v1.Alert.afk:type_name -> dinkurapi.v1.AlertAfk
	11, // 8: dinkurapi.v1.AlertAfk.active_entry:type_name -> dinkurapi.v1.Entry
	10, // 9: dinkurapi.v1.AlertAfk.afk_since:type_name -> google.protobuf.Timestamp
	10, // 10: dinkurapi.v1.AlertAfk.back_since:type_name -> google.protobuf.Timestamp
	0,  // 11: dinkurapi.v1.Alerter.StreamAlert:input_type -> dinkurapi.v1.StreamAlertRequest
	2,  // 12: dinkurapi.v1.Alerter.GetAlertList:input_type -> dinkurapi.v1.GetAlertListRequest
	4,  // 13: dinkurapi.v1.Alerter.DeleteAlert:input_type -> dinkurapi.v1.DeleteAlertRequest
	1,  // 14: dinkurapi.v1.Alerter.StreamAlert:output_type -> dinkurapi.v1.StreamAlertResponse
	3,  // 15: dinkurapi.v1.Alerter.GetAlertList:output_type -> dinkurapi.v1.GetAlertListResponse
	5,  // 16: dinkurapi.v1.Alerter.DeleteAlert:output_type -> dinkurapi.v1.DeleteAlertResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_api_dinkurapi_v1_alerter_proto_init() }
func file_api_dinkurapi_v1_alerter_proto_init() {
	if File_api_dinkurapi_v1_alerter_proto != nil {
		return
	}
	file_api_dinkurapi_v1_entries_proto_init()
	file_api_dinkurapi_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_dinkurapi_v1_alerter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAlertRequest); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAlertResponse); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertListRequest); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlertListResponse); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlertRequest); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlertResponse); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alert); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertPlainMessage); i {
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
		file_api_dinkurapi_v1_alerter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertAfk); i {
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
	file_api_dinkurapi_v1_alerter_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Alert_PlainMessage)(nil),
		(*Alert_Afk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_dinkurapi_v1_alerter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_dinkurapi_v1_alerter_proto_goTypes,
		DependencyIndexes: file_api_dinkurapi_v1_alerter_proto_depIdxs,
		MessageInfos:      file_api_dinkurapi_v1_alerter_proto_msgTypes,
	}.Build()
	File_api_dinkurapi_v1_alerter_proto = out.File
	file_api_dinkurapi_v1_alerter_proto_rawDesc = nil
	file_api_dinkurapi_v1_alerter_proto_goTypes = nil
	file_api_dinkurapi_v1_alerter_proto_depIdxs = nil
}
