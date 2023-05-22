// Code generated by ogen, DO NOT EDIT.

package gen

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type EntitiesResponse struct {
	TotalCount int      `json:"totalCount"`
	Items      []Entity `json:"items"`
}

// GetTotalCount returns the value of TotalCount.
func (s *EntitiesResponse) GetTotalCount() int {
	return s.TotalCount
}

// GetItems returns the value of Items.
func (s *EntitiesResponse) GetItems() []Entity {
	return s.Items
}

// SetTotalCount sets the value of TotalCount.
func (s *EntitiesResponse) SetTotalCount(val int) {
	s.TotalCount = val
}

// SetItems sets the value of Items.
func (s *EntitiesResponse) SetItems(val []Entity) {
	s.Items = val
}

// Entity object. Any object can be here.
// Ref: #/components/schemas/Entity
type Entity struct{}

// Entity mapping.
// Ref: #/components/schemas/EntityMapping
type EntityMapping struct {
	Fields []EntityMappingFieldsItem `json:"fields"`
}

// GetFields returns the value of Fields.
func (s *EntityMapping) GetFields() []EntityMappingFieldsItem {
	return s.Fields
}

// SetFields sets the value of Fields.
func (s *EntityMapping) SetFields(val []EntityMappingFieldsItem) {
	s.Fields = val
}

type EntityMappingFieldsItem struct {
	Name     string                          `json:"name"`
	Datatype EntityMappingFieldsItemDatatype `json:"datatype"`
}

// GetName returns the value of Name.
func (s *EntityMappingFieldsItem) GetName() string {
	return s.Name
}

// GetDatatype returns the value of Datatype.
func (s *EntityMappingFieldsItem) GetDatatype() EntityMappingFieldsItemDatatype {
	return s.Datatype
}

// SetName sets the value of Name.
func (s *EntityMappingFieldsItem) SetName(val string) {
	s.Name = val
}

// SetDatatype sets the value of Datatype.
func (s *EntityMappingFieldsItem) SetDatatype(val EntityMappingFieldsItemDatatype) {
	s.Datatype = val
}

type EntityMappingFieldsItemDatatype string

const (
	EntityMappingFieldsItemDatatypeString  EntityMappingFieldsItemDatatype = "string"
	EntityMappingFieldsItemDatatypeInteger EntityMappingFieldsItemDatatype = "integer"
	EntityMappingFieldsItemDatatypeFloat   EntityMappingFieldsItemDatatype = "float"
)

// MarshalText implements encoding.TextMarshaler.
func (s EntityMappingFieldsItemDatatype) MarshalText() ([]byte, error) {
	switch s {
	case EntityMappingFieldsItemDatatypeString:
		return []byte(s), nil
	case EntityMappingFieldsItemDatatypeInteger:
		return []byte(s), nil
	case EntityMappingFieldsItemDatatypeFloat:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *EntityMappingFieldsItemDatatype) UnmarshalText(data []byte) error {
	switch EntityMappingFieldsItemDatatype(data) {
	case EntityMappingFieldsItemDatatypeString:
		*s = EntityMappingFieldsItemDatatypeString
		return nil
	case EntityMappingFieldsItemDatatypeInteger:
		*s = EntityMappingFieldsItemDatatypeInteger
		return nil
	case EntityMappingFieldsItemDatatypeFloat:
		*s = EntityMappingFieldsItemDatatypeFloat
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Error schema.
// Ref: #/components/schemas/Error
type Error struct {
	Code           int       `json:"code"`
	Message        string    `json:"message"`
	DisplayMessage OptString `json:"display_message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// GetDisplayMessage returns the value of DisplayMessage.
func (s *Error) GetDisplayMessage() OptString {
	return s.DisplayMessage
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// SetDisplayMessage sets the value of DisplayMessage.
func (s *Error) SetDisplayMessage(val OptString) {
	s.DisplayMessage = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/responses/NoContent
type NoContent struct{}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}
