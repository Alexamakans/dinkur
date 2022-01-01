// Dinkur the task time tracking utility.
// <https://github.com/dinkur/dinkur>
//
// Copyright (C) 2021 Kalle Fagerberg
// SPDX-FileCopyrightText: 2021 Kalle Fagerberg
// SPDX-License-Identifier: GPL-3.0-or-later
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more
// details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: api/dinkurapi/v1/event.proto

package v1

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

type Event int32

const (
	Event_UNKNOWN Event = 0
	Event_CREATED Event = 1
	Event_UPDATED Event = 2
	Event_DELETED Event = 3
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATED",
		2: "UPDATED",
		3: "DELETED",
	}
	Event_value = map[string]int32{
		"UNKNOWN": 0,
		"CREATED": 1,
		"UPDATED": 2,
		"DELETED": 3,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_api_dinkurapi_v1_event_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_api_dinkurapi_v1_event_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_api_dinkurapi_v1_event_proto_rawDescGZIP(), []int{0}
}

var File_api_dinkurapi_v1_event_proto protoreflect.FileDescriptor

var file_api_dinkurapi_v1_event_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2a, 0x3b, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72, 0x2f, 0x64,
	0x69, 0x6e, 0x6b, 0x75, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69, 0x6e, 0x6b, 0x75, 0x72,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_dinkurapi_v1_event_proto_rawDescOnce sync.Once
	file_api_dinkurapi_v1_event_proto_rawDescData = file_api_dinkurapi_v1_event_proto_rawDesc
)

func file_api_dinkurapi_v1_event_proto_rawDescGZIP() []byte {
	file_api_dinkurapi_v1_event_proto_rawDescOnce.Do(func() {
		file_api_dinkurapi_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_dinkurapi_v1_event_proto_rawDescData)
	})
	return file_api_dinkurapi_v1_event_proto_rawDescData
}

var file_api_dinkurapi_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_dinkurapi_v1_event_proto_goTypes = []interface{}{
	(Event)(0), // 0: dinkurapi.v1.Event
}
var file_api_dinkurapi_v1_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_dinkurapi_v1_event_proto_init() }
func file_api_dinkurapi_v1_event_proto_init() {
	if File_api_dinkurapi_v1_event_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_dinkurapi_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_dinkurapi_v1_event_proto_goTypes,
		DependencyIndexes: file_api_dinkurapi_v1_event_proto_depIdxs,
		EnumInfos:         file_api_dinkurapi_v1_event_proto_enumTypes,
	}.Build()
	File_api_dinkurapi_v1_event_proto = out.File
	file_api_dinkurapi_v1_event_proto_rawDesc = nil
	file_api_dinkurapi_v1_event_proto_goTypes = nil
	file_api_dinkurapi_v1_event_proto_depIdxs = nil
}
