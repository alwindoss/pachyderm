// Code generated by protoc-gen-go.
// source: google/protobuf/any.proto
// DO NOT EDIT!

/*
Package google_protobuf is a generated protocol buffer package.

It is generated from these files:
	google/protobuf/any.proto
	google/protobuf/api.proto
	google/protobuf/duration.proto
	google/protobuf/empty.proto
	google/protobuf/field_mask.proto
	google/protobuf/source_context.proto
	google/protobuf/struct.proto
	google/protobuf/timestamp.proto
	google/protobuf/type.proto
	google/protobuf/wrappers.proto

It has these top-level messages:
	Any
	Api
	Method
	Mixin
	Duration
	Empty
	FieldMask
	SourceContext
	Struct
	Value
	ListValue
	Timestamp
	Type
	Field
	Enum
	EnumValue
	Option
	DoubleValue
	FloatValue
	Int64Value
	UInt64Value
	Int32Value
	UInt32Value
	BoolValue
	StringValue
	BytesValue
*/
package google_protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Any` contains an arbitrary serialized message along with a URL
// that describes the type of the serialized message.
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the the `@type`
// field. Example (for message [google.protobuf.Duration][google.protobuf.Duration]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
type Any struct {
	// A URL/resource name whose content describes the type of the
	// serialized message.
	//
	// For URLs which use the schema `http`, `https`, or no schema, the
	// following restrictions and interpretations apply:
	//
	// * If no schema is provided, `https` is assumed.
	// * The last segment of the URL's path must represent the fully
	//   qualified name of the type (as in `path/google.protobuf.Duration`).
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][google.protobuf.Type]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Schemas other than `http`, `https` (or the empty schema) might be
	// used with implementation specific semantics.
	//
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url" json:"type_url,omitempty"`
	// Must be valid serialized data of the above specified type.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}
